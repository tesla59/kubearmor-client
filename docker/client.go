package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/kubearmor/kubearmor-client/recommend/common"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Client struct {
	*client.Client
}

func ConnectDockerClient() (*Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Client{cli}, nil
}

func (c *Client) ListObjects(o common.Options) ([]common.Object, error) {
	log.Info("Listing objects in docker")
	var result []common.Object
	containers, err := c.Client.ContainerList(context.Background(), container.ListOptions{
		Filters: filters.NewArgs(),
	})
	log.Info("containers: ", containers)
	if err != nil {
		return nil, err
	}
	for _, ctr := range containers {
		result = append(result, common.Object{
			Name:   strings.TrimPrefix(ctr.Names[0], "/"),
			Images: []string{ctr.Image},
			Labels: ctr.Labels,
		})
	}
	return result, nil
}
