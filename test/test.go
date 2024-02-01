package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	// 查找某个确定的docker镜像
	filters := filters.NewArgs()
	filters.Add("reference", "registry.cnbita.com:5000/training/pytorch-1.6-cuda-11.0-py3:v1.0")

	options := types.ImageListOptions{
		Filters: filters,
	}
	images, err := cli.ImageList(ctx, options)
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Printf("%s\n", image.ID)
	}
}
