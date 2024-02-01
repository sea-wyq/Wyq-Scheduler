package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// 提供查询镜像是否存在功能
func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reference := r.URL.Query().Get("reference")

		filters := filters.NewArgs()
		filters.Add("reference", reference)
		options := types.ImageListOptions{
			Filters: filters,
		}
		images, err := cli.ImageList(context.Background(), options)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, image := range images {
			fmt.Fprintf(w, "%s\n", image.ID)
		}
	})

	http.ListenAndServe(":8088", nil)
}
