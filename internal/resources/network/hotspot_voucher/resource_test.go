package network_test

import (
	"embed"
	"testing"

	"github.com/alienchow/unitf/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//go:embed testdata/*
var testDataFS embed.FS

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"unifi": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func getTestData(t *testing.T, filename string) string {
	b, err := testDataFS.ReadFile("testdata/" + filename)
	if err != nil {
		t.Fatalf("Failed to load testdata %s: %s", filename, err)
	}
	return string(b)
}

func TestAccHotspotVoucherResource_SchemaValidation(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             getTestData(t, "voucher_valid.tf"),
				ExpectError:        nil,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
