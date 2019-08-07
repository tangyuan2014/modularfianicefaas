package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

//const ping = "http://localhost:8080/ping"
//const operation = "http://localhost:8080/operations"
const default_function = "http://loclahost:8080"


   var routerMap=make(map[string]ContainerData)
   var hostProxy=make(map[string]*httputil.ReverseProxy)


type ContainerData struct {
	labels string
	status string
	ip     string
	port   string
}

func handleRequestAndRedirect(writer http.ResponseWriter, request *http.Request) {
	var host string
	functionName:=getFunction(request)
	if function,ok:=routerMap[functionName];ok{
		host = "http://" + function.ip +":"+ function.port
	}else{
		host=default_function
	}
	//switch getFunction(request) {
	//case "ping":
	//	host = "http://" + routerMap["ping"].ip + routerMap["ping"].port
	//case "operation":
	//	host = "http://" + routerMap["ping"].ip + routerMap["operation"].port
	//default:
	//	host = default_function
	//}
	if fn, ok :=hostProxy[host]; ok{
		fn.ServeHTTP(writer,request)
		return
	}
	proxy:=ServeHttp(host,writer,request)
	hostProxy[host]=proxy
}

func ServeHttp(target string, writer http.ResponseWriter, request *http.Request) *httputil.ReverseProxy {
	log.Println(target)
	targetUrl, err := url.Parse(target)
	if err!=nil{
		log.Println("url fail")
		return nil
	}
	proxy:=httputil.NewSingleHostReverseProxy(targetUrl)
	request.URL.Host = targetUrl.Host
	request.URL.Scheme = targetUrl.Scheme
	request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
	request.Host = targetUrl.Host
	proxy.ServeHTTP(writer,request)
	return proxy
}

func pollingStatusOfContainers() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(10).Seconds().Do(GetContainerStatus)
	<- scheduler.Start()
}

func GetContainerStatus() {
	var routerTable map[string]ContainerData
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "faas.name")
	filterArgs.Add("status", "running")
	ctx := context.Background();
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) //TODO define host version
	if err != nil {
		panic(err.Error())
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterArgs,})
	if err != nil {
		panic(err.Error())
	}
	if containers == nil {
		panic("No containers")
	}
	for _, container := range containers {
		labels := container.Labels["faas.name"]
		log.Println(labels)
		status := container.State
		log.Println(status)
		ipAddress := container.NetworkSettings.Networks["bridge"].IPAddress
		log.Println(ipAddress)
		port := strconv.FormatUint(uint64(container.Ports[0].PrivatePort), 10)
		log.Println(port)
		container1 := ContainerData{labels, status, ipAddress, port}
		routerTable[labels] = container1
	}
	routerMap = routerTable
}

//func parseRequestBody(request http.Request){

func getFunction(request *http.Request) string {
	return request.URL.Path[len("/gateway/"):]
}

func defaultFunction(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("No service, please correct function name"))
}

func main() {
	go pollingStatusOfContainers()
	http.HandleFunc("/gateway/", handleRequestAndRedirect)
	http.HandleFunc("/default", defaultFunction)
	http.ListenAndServe(":80", nil)
}

