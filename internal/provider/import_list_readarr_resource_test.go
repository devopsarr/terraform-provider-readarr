package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListReadarrResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListReadarrResourceConfig("resourceReadarrTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListReadarrResourceConfig("resourceReadarrTest", "entireAuthor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_readarr.test", "should_monitor", "entireAuthor"),
					resource.TestCheckResourceAttrSet("readarr_import_list_readarr.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListReadarrResourceConfig("resourceReadarrTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListReadarrResourceConfig("resourceReadarrTest", "specificBook"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_readarr.test", "should_monitor", "specificBook"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list_readarr.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListReadarrResourceConfig(name, folder string) string {
	return fmt.Sprintf(`
	resource "readarr_import_list_readarr" "test" {
		enable_automatic_add = false
		should_monitor = "%s"
		should_search = false
		root_folder_path = "/config"
		quality_profile_id = 1
		metadata_profile_id = 1
		name = "%s"
		base_url = "http://127.0.0.1:8787"
		api_key = "testAPIKey"
	}`, folder, name)
}
