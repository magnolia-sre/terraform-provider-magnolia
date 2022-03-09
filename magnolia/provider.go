package magnolia

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"magnolia_subscription": resourceSubscription(),
			"magnolia_git_config":   resourceMagnoliaGitConfig(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}

	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return client()
	}
	return provider
}
