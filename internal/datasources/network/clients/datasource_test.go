package clients_test

import (
	_ "embed"
	"regexp"
	"testing"

	"github.com/alienchow/unitf/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//go:embed testdata/datasource.tf
var testAccClientsDataSourceConfig string

func TestAccClientsDataSource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"unifi": providerserver.NewProtocol6WithError(provider.New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccClientsDataSourceConfig,
				ExpectError: regexp.MustCompile(`connection refused`),
				PlanOnly:    true,
			},
		},
	})
}
