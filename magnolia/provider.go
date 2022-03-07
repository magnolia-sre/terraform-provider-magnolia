package magnolia

import (
	"context"
	"fmt"
	subscriptionRestClient "terraform-provider-magnolia/internal/subscription-service-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAGNOLIA_TOKEN", nil),
			},
		},
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	fmt.Println("Get token..." + token)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if token == "" {
		fmt.Errorf("Error: Missing Magnolia token")
	}

	ctx, client, err := okta.NewClient(
		ctx,
		okta.WithOrgUrl("https://magnolia-cloud.oktapreview.com"),
		okta.WithToken(token),
	)

	ssClient := subscriptionRestClient.NewAPIClient(nil)
	ssClient.SubscriptionApi.FindSubscriptionById(ctx, "mockId")

	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	return client, diags
}
