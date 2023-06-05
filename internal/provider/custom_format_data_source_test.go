package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCustomFormatDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized
			{
				Config:      testAccCustomFormatDataSourceConfig("\"Error\"") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Not found testing
			{
				Config:      testAccCustomFormatDataSourceConfig("\"Error\""),
				ExpectError: regexp.MustCompile("Unable to find custom_format"),
			},
			// Read testing
			{
				Config: testAccCustomFormatResourceConfig("dataTest", "false") + testAccCustomFormatDataSourceConfig("readarr_custom_format.test.name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_custom_format.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_custom_format.test", "include_custom_format_when_renaming", "false")),
			},
		},
	})
}

func testAccCustomFormatDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	data "readarr_custom_format" "test" {
		name = %s
	}
	`, name)
}