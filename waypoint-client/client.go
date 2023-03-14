package client

import (
	"context"

	client "github.com/hashicorp/waypoint-client/client"
	"github.com/hashicorp/waypoint/pkg/client/gen/client/waypoint"
)

type Client struct {
	c waypoint.ClientService
}

func New(isHCP bool) (*Client, error) {
	var (
		clientConfig *client.Config
		err          error
	)

	// TODO(briancain): Abstract these out to to be configured
	if isHCP {
		clientConfig, err = client.NewWithHCP("hcp_org_id",
			"hcp_project_id",
			"client_id",
			"client_secret", context.Background())
		if err != nil {
			return nil, err
		}
	} else {
		clientConfig, err = client.NewWithSelfManaged("waypoint-user-token", "localhost:9702", true)
		if err != nil {
			return nil, err
		}
	}

	client, err := client.New(*clientConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		c: client,
	}, nil
}
