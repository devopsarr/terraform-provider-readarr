package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListExclusionResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListExclusionResourceConfig("test", "b1a9c0e9-d987-4042-ae91-78d6a3267d69") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccImportListExclusionResourceConfig("test", "b1a9c0e9-d987-4042-ae91-78d6a3267d69"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_exclusion.test", "foreign_id", "b1a9c0e9-d987-4042-ae91-78d6a3267d69"),
					resource.TestCheckResourceAttrSet("readarr_import_list_exclusion.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListExclusionResourceConfig("test", "b1a9c0e9-d987-4042-ae91-78d6a3267d69") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListExclusionResourceConfig("test", "46a098f3-272d-4bec-9746-67e8ab48ed40"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_exclusion.test", "foreign_id", "46a098f3-272d-4bec-9746-67e8ab48ed40"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list_exclusion.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListExclusionResourceConfig(name, ID string) string {
	return fmt.Sprintf(`
		resource "readarr_import_list_exclusion" "%s" {
			author_name = "Agatha Christie"
			foreign_id = "%s"
		}
	`, name, ID)
}
