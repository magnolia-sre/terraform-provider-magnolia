package magnolia

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	subscriptionRestClient "github.com/magnolia-sre/terraform-provider-magnolia/internal/subscription-service-client"
)

type MagnoliaClient struct {
	conn  *subscriptionRestClient.APIClient
	token string
}

func Client() (*MagnoliaClient, diag.Diagnostics) {
	var client MagnoliaClient

	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	accessToken, err := ioutil.ReadFile(dirname + "/.mgnl/access_token")
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("[Err] No Token for Magnolia"))
	}

	client.token = string(accessToken)

	cfg := subscriptionRestClient.NewConfiguration()
	cfg.Servers = subscriptionRestClient.ServerConfigurations{
		{
			URL:         "https://subscription-service.beta.de.magnolia-cloud.com",
			Description: "Staging environment",
		},
	}

	magnoliaClient := subscriptionRestClient.NewAPIClient(cfg)

	// DELETEME: Just for example
	_, _, err = magnoliaClient.UserApi.ListUsersOfSubscription(context.WithValue(context.TODO(), subscriptionRestClient.ContextAccessToken, client.token), "mabdtq1l6bx4ic94").Execute()

	if err != nil {
		return nil, diag.FromErr(err)
	}

	client.conn = magnoliaClient
	return &client, nil
}
