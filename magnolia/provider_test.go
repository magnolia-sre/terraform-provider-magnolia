package magnolia

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){
		"magnolia": func() (*schema.Provider, error) { return testAccProvider, nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

//Use later for comming tests
//PreCheck: func() { testAccPreCheck(t) },
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MAGNOLIA_TOKEN"); v == "" {
		t.Fatal("MAGNOLIA_TOKEN must be set for acceptance tests")
	}
}
