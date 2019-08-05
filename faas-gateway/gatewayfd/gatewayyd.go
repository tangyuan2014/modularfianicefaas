package gatewayfd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

//const ping = "http://localhost:8080/ping"
//const operation = "http://localhost:8080/operations"
const default_function = "http://loclahost:8080/default"

var routerMap map[string]ContainerData

type ContainerData struct {
	lables string
	status string
	ip     string
	port   string
}

func handleRequestAndRedirect(writer http.ResponseWriter, request *http.Request) {
	var urll string
	switch getFunction(request) {
	case "ping":
		urll = "http://" + routerMap["ping"].ip + routerMap["ping"].port
	case "operation":
		urll = "http://" + routerMap["ping"].ip + routerMap["operation"].port
	default:
		urll = default_function
	}
	targetUrl, _ := url.Parse(urll)
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	//request.URL.Host = targetUrl.Host
	//request.URL.Scheme = targetUrl.Scheme
	request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
	//request.Host = targetUrl.Host
	proxy.ServeHTTP(writer, request)
}

func pollingStatusOfContainers() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(10).Seconds().Do(GetContainerStatus)
}

func GetContainerStatus() {
	var routerTable map[string]ContainerData
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "faas.name")
	filterArgs.Add("status","running")
	ctx := context.Background();
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) //TODO
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
		status := container.State
		ipAddress := container.NetworkSettings.Networks["bridge"].IPAddress
		port := strconv.FormatUint(uint64(container.Ports[0].PrivatePort), 10)
		container1 := ContainerData{labels, status, ipAddress, port}
		routerTable[labels] = container1
	}
	routerMap = routerTable
}

//func parseRequestBody(request http.Request){

func getFunction(request *http.Request) string {
	return request.URL.Path[len("/gateway/"):]
}

func defaultFunction(writer http.ResponseWriter,request *http.Request){
	fmt.Println("No service, please correct function name")
}

func main() {
	http.HandleFunc("/gateway/", handleRequestAndRedirect)
	http.HandleFunc("/default",defaultFunction)
}
