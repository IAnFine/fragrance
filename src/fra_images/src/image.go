package images

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
)

// 镜像列表
func GetImagesList(options types.ImageListOptions) []types.ImageSummary {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	return images
}

// 创建镜像
func CreateImage(parentReference string, options types.ImageCreateOptions) io.ReadCloser {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	readCloser, err := cli.ImageCreate(context.Background(), parentReference, options)
	if err != nil {
		panic(err)
	}
	return readCloser
}

// 移除镜像
func RemoveImage(imageID string, options types.ImageRemoveOptions) []types.ImageDeleteResponseItem {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	items, err := cli.ImageRemove(context.Background(), imageID, options)
	if err != nil {
		panic(err)
	}
	return items
}

// 推送镜像
func PullImage(refStr string, options types.ImagePullOptions) io.ReadCloser {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	readCloser, err := cli.ImagePull(context.Background(), refStr, options)
	if err != nil {
		panic(err)
	}
	return readCloser
}
