package images

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImagesList() []types.ImageSummary{

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	return images
}

