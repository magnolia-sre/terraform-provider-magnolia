package magnolia

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMagnoliaSubscription(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testResource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("magnolia_subscription.this", "email", "chanh.hua+terraform@magnolia-cms.com"),
				),
			},
		},
	})
}

const testResource = `
resource "magnolia_subscription" "this" {
  email = "chanh.hua+terraform@magnolia-cms.com"
	password = "Abcd@1234"
}
`
