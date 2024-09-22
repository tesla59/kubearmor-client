package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/kubearmor/kubearmor-client/recommend"
	"github.com/kubearmor/kubearmor-client/recommend/common"
	"strings"
)

type Client struct {
	Client *client.Client
}

func NewClient() (*Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: dockerClient,
	}, nil

}

func (c *Client) ListDeployments(o common.Options) ([]recommend.Deployment, error) {
	var result []recommend.Deployment
	containers, err := c.Client.ContainerList(context.Background(), container.ListOptions{
		Filters: filters.NewArgs(),
	})
	if err != nil {
		return nil, err
	}
	for _, ctr := range containers {
		result = append(result, recommend.Deployment{
			Name:   strings.TrimPrefix(ctr.Names[0], "/"),
			Images: []string{ctr.Image},
			Labels: ctr.Labels,
		})
	}
	return result, nil
}
