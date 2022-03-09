package magnolia

import (
	"context"

	_ "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	subscriptionRestClient "github.com/magnolia-sre/terraform-provider-magnolia/internal/subscription-service-client"
)

func resourceSubscription() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Magnolia Subscription",

		CreateContext: resourceSubscriptionCreate,
		ReadContext:   resourceSubscriptionRead,
		UpdateContext: resourceSubscriptionUpdate,
		DeleteContext: resourceSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Description: "email",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "password",
				Sensitive:   true,
				Required:    true,
			},
		},
	}
}

func resourceSubscriptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	token := meta.(MagnoliaClient).token
	conn := meta.(MagnoliaClient).conn

	req := conn.AdminSubscriptionApi.CreateSubscription(context.WithValue(context.TODO(),
		subscriptionRestClient.ContextAccessToken, token))
	planId := "970d51a7-aa97-4639-85b8-83b6a2d7b9f5" // :shrug
	req.CreatedSubscription(subscriptionRestClient.CreatedSubscription{
		FirstName: "Hackathon",
		LastName:  "SRE",
		Email:     d.Get("email").(string),
		Function:  "Developer",
		Company:   "Magnolia",
		Password:  d.Get("password").(string),
		Country:   "Vietnam",
		PlanId:    &planId,
	})

	response, _, err := req.Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("id", response.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", response.Alias); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSubscriptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceSubscriptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
