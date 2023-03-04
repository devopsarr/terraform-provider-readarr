package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTagDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized
			{
				Config:      testAccTagDataSourceConfig("\"error\"") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Not found testing
			{
				Config:      testAccTagDataSourceConfig("\"error\""),
				ExpectError: regexp.MustCompile("Unable to find tag"),
			},
			// Read testing
			{
				Config: testAccTagResourceConfig("test", "tag_datasource") + testAccTagDataSourceConfig("readarr_tag.test.label"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_tag.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_tag.test", "label", "tag_datasource"),
				),
			},
		},
	})
}

func testAccTagDataSourceConfig(label string) string {
	return fmt.Sprintf(`
	data "readarr_tag" "test" {
		label = %s
	}
	`, label)
}
