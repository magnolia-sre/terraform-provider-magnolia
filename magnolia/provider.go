package magnolia

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAGNOLIA_TOKEN", nil),
				Description: "Magnolia TOKEN  API",
			},
		},
		ResourcesMap:   map[string]*schema.Resource{
			"magnolia_git_config" : resourceGitConfig(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}

	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return Client()
	}
	return provider
}

/*
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
}*/
