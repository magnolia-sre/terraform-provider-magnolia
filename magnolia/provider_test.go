package magnolia

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProvider(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Provider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_providerConfigure(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 diag.Diagnostics
	}{
		{name: "test", args: {
			ctx: context.TODO(),
			d: &{
				
			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := providerConfigure(tt.args.ctx, tt.args.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("providerConfigure() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("providerConfigure() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
