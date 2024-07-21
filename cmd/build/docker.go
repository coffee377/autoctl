package build

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"os"
	"strings"
)

type DockerCli struct {
	cli *client.Client
}

func NewDockerCli() DockerCli {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return DockerCli{cli: cli}
}

func (d DockerCli) EnsureImage(ctx context.Context) error {
	// 获取所有镜像信息
	images, err := d.cli.ImageList(ctx, types.ImageListOptions{ContainerCount: true})
	for _, image := range images {
		if len(image.RepoTags) == 0 {
			continue
		}
		for _, tag := range image.RepoTags {
			if strings.Contains(tag, "node:18.18.2") {
				fmt.Printf("%v \n", tag)
			}
		}
	}
	reader, err := d.cli.ImagePull(ctx, "alpine", types.ImagePullOptions{})
	//d.cli.ImageBuild()
	io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}
	return nil
}

func (d DockerCli) RunImage(ctx context.Context) error {
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf(resp.ID)
	return nil
}
