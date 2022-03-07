package magnolia

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	subscriptionRestClient "terraform-provider-magnolia/internal/subscription-service-client"
)

type Config struct {
	Token string
}

type MagnoliaClient struct {
	conn *subscriptionRestClient.APIClient
}

func (c *Config) Client() (*MagnoliaClient, diag.Diagnostics) {
	var client MagnoliaClient

	if c.Token == "" {
		return nil, diag.FromErr(fmt.Errorf("[Err] No Token for Magnolia"))
	}

	//subscriptionRestClient.ContextAccessToken

	cfg := subscriptionRestClient.NewConfiguration()

	magnoliaClient, err := subscriptionRestClient.NewAPIClient(cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	client.conn = magnoliaClient
	return &client, nil
}
