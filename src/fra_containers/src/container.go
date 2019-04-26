package containers

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func GetContainersList(options types.ContainerListOptions) ([]types.Container, error) {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	container, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	return container, err
}

// 创建容器
func CreateContainer(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	createbody, err := cli.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, containerName)
	if err != nil {
		panic(err)
	}

	return createbody, err
}

// 移除容器
func MoveContainer(containerID string, options types.ContainerRemoveOptions) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	result := true
	if err := cli.ContainerRemove(context.Background(), containerID, options); err != nil {
		panic(err)
		result = false
	}
	return result
}

// 启动容器
func StartContainer() {

}

// 获取容器详情
func GetContainerDetail() {

}

// 扫描容器 查看是否有状态异常的容器
func ScanContainer() {

}
