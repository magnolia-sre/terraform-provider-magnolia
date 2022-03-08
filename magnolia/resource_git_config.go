package magnolia

import (
	"context"
	_ "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	subscriptionRestClient "github.com/magnolia-sre/terraform-provider-magnolia/internal/subscription-service-client"
)

func resourceMagnoliaGitConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceMagnoliaGitConfigCreate,
		ReadContext:   resourceMagnoliaGitConfigRead,
		UpdateContext: resourceMagnoliaGitConfigUpdate,
		DeleteContext: resourceMagnoliaGitConfigDelete,

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the magnolia subscription",
			},
			"git_provider": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Git provider can be GITHUB, BITBUCKET or BITBUCKET_CLOUD ",
			},
			"git_clone_url": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "User SSH URl to be clone",
			},
			"public_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Description: "Magnolia Public Key to interact whit",
			},
			"git_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Description: "Secret generated to be injected in the user git repository",
			},
			"webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Description: "Webhook URL generated to be injected in the user git repository",
			},
			"webhook_content_type": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "Webhook Content-Type when the git repository is GITHUB",
			},
		},
	}
}

func resourceMagnoliaGitConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	d.SetId(d.Get("subscription_id").(string))

	return resourceMagnoliaGitConfigRead(ctx, d, meta)
}

func resourceMagnoliaGitConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	token := meta.(MagnoliaClient).token
	conn := meta.(MagnoliaClient).conn

	var diags diag.Diagnostics

	getRequest := conn.SubscriptionGitApi.GetGitConfiguration(context.WithValue(context.TODO(),
		subscriptionRestClient.ContextAccessToken, token), d.Get("subscription_id").(string))

	response, _, err := getRequest.Execute()
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

	return diags
}

func resourceMagnoliaGitConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	d.SetId(d.Get("subscription_id").(string))

	return resourceMagnoliaGitConfigRead(ctx, d, meta)
}

func resourceMagnoliaGitConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
