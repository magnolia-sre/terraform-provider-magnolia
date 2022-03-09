package magnolia

import (
	"testing"
)

func TestConfig_Client(t *testing.T) {
	_, diagnostics := client()

	if diagnostics.HasError() {
		t.Errorf("Failed to create client")
	}
}
