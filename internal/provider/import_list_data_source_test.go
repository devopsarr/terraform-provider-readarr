package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized
			{
				Config:      testAccImportListDataSourceConfig("\"Error\"") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Not found testing
			{
				Config:      testAccImportListDataSourceConfig("\"Error\""),
				ExpectError: regexp.MustCompile("Unable to find import_list"),
			},
			// Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListResourceConfig("importListDataTest", "none") + testAccImportListDataSourceConfig("readarr_import_list.test.name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_import_list.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_import_list.test", "should_monitor", "none")),
			},
		},
	})
}

func testAccImportListDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	data "readarr_import_list" "test" {
		name = %s
	}
	`, name)
}
