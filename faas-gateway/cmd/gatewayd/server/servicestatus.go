package server

import (
	"log"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var ServiceMap = make(map[string]ContainerData)

type ContainerData struct {
	Labels    string
	Status    string
	IPAddress string
	Port      string
}

func GetContainerStatus() {
	var serviceTable = make(map[string]ContainerData)
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "faas.name")
	filterArgs.Add("status", "running")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("Failed to generate docker client with error: " + err.Error())
		return
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterArgs})

	if err != nil {
		log.Println("Failed to get container data with error: " + err.Error())
		return
	}
	if containers != nil {
		for _, container := range containers {
			labelFaasName := container.Labels["faas.name"]
			status := container.State
			netWorks := container.NetworkSettings.Networks
			hostConfig := container.HostConfig.NetworkMode
			IPAddress := netWorks[hostConfig].IPAddress
			port := strconv.FormatUint(uint64(container.Ports[0].PrivatePort), 10)
			container1 := ContainerData{labelFaasName, status, IPAddress, port}
			serviceTable[labelFaasName] = container1
		}
	}
	ServiceMap = serviceTable
}
