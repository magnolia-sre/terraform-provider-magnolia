package magnolia

import (
	"context"
	_ "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	subscriptionRestClient "github.com/magnolia-sre/terraform-provider-magnolia/internal/subscription-service-client"
)

func resourceGitConfig() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceGitConfigCreate,
		ReadContext:   resourceGitConfigRead,
		UpdateContext: resourceGitConfigUpdate,
		DeleteContext: resourceGitConfigDelete,

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "subscription_id",
			},
			"git_provider": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Git provider",
			},
			"git_clone_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "git_clone_url",
			},
		},
	}
}

func resourceGitConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	token := meta.(MagnoliaClient).token
	conn := meta.(MagnoliaClient).conn

	updateRequest := conn.SubscriptionGitApi.UpdateGitConfiguration(context.WithValue(context.TODO(),
		subscriptionRestClient.ContextAccessToken, token), d.Get("subscription_id").(string))
	updateRequest.UpdateGitConfigurationRequest(subscriptionRestClient.UpdateGitConfigurationRequest{
		GitCloneUrl: d.Get("git_clone_url").(string),
		GitProvider: d.Get("git_provider").(subscriptionRestClient.GitProvider),
	})

	response, _, err := updateRequest.Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("git_provider", response.GitProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("git_clone_url", response.GitCloneUrl); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("public_key", response.PublicKey); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("git_secret", response.GitSecret); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("webhook_url", response.WebhookUrl); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGitConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceGitConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceGitConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
