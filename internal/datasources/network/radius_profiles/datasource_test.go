package radius_profiles_test

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
var testAccRadiusProfilesDataSourceConfig string

func TestAccRadiusProfilesDataSource(t *testing.T) {
	t.Skip("not yet implemented")
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"unifi": providerserver.NewProtocol6WithError(provider.New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccRadiusProfilesDataSourceConfig,
				ExpectError: regexp.MustCompile(`connection refused`),
				PlanOnly:    true,
			},
		},
	})
}
