package magnolia

import (
	"testing"
)

func TestConfig_Client(t *testing.T) {
	c := Config{}
	_, diagnostics := c.Client()

	if diagnostics.HasError() {
		t.Errorf("Failed to create client")
	}
}