package magnolia

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	subscriptionRestClient "github.com/magnolia-sre/terraform-provider-magnolia/internal/subscription-service-client"
)

const APIKeyEnvVar = "MAGNOLIA_TOKEN"

type Config struct {
	Token string
}

type MagnoliaClient struct {
	conn *subscriptionRestClient.APIClient
}

func (c *Config) Client() (*MagnoliaClient, diag.Diagnostics) {
	var client MagnoliaClient

	// FIXME: make credential file configuration
	accessToken, err := ioutil.ReadFile("/Users/chanh.hua/.mgnl/access_token")
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("[Err] No Token for Magnolia"))
	}

	c.Token = string(accessToken)

	cfg := subscriptionRestClient.NewConfiguration()
	cfg.Servers = subscriptionRestClient.ServerConfigurations{
		{
			URL:         "https://subscription-service.beta.de.magnolia-cloud.com",
			Description: "Staging environment",
		},
	}

	magnoliaClient, err := subscriptionRestClient.NewAPIClient(cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// DELETEME: Just for example
	_, _, err = magnoliaClient.UserApi.ListUsersOfSubscription(context.WithValue(context.TODO(), subscriptionRestClient.ContextAccessToken, c.Token), "mabdtq1l6bx4ic94").Execute()

	if err != nil {
		return nil, diag.FromErr(err)
	}

	client.conn = magnoliaClient
	return &client, nil
}
