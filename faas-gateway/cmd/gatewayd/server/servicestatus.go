package server

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
	"strconv"
)
var ServiceMap = make(map[string]ContainerData)
type ContainerData struct {
	Labels    string
	Status    string
	IpAddress string
	Port      string
}

func GetContainerStatus()  {
	var serviceTable = make(map[string]ContainerData)
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "faas.name")
	filterArgs.Add("status", "running")
	ctx := context.Background();
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) //TODO define host version
	if err != nil {
		log.Println(err.Error())
		return
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterArgs,})
	if err != nil {
		log.Println(err.Error())
		return
	}
	if containers != nil {
		for _, container := range containers {
			labelFaasName := container.Labels["faas.name"]
			status := container.State
			netWorks := container.NetworkSettings.Networks
			hostConfig := container.HostConfig.NetworkMode
			ipAddress := netWorks[hostConfig].IPAddress
			port := strconv.FormatUint(uint64(container.Ports[0].PrivatePort), 10)
			container1 := ContainerData{labelFaasName, status, ipAddress, port}
			serviceTable[labelFaasName] = container1
		}
	}
	ServiceMap=serviceTable
}
