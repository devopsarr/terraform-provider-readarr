package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListResourceConfig("importListResourceTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListResourceConfig("importListResourceTest", "entireAuthor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list.test", "should_monitor", "entireAuthor"),
					resource.TestCheckResourceAttrSet("readarr_import_list.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListResourceConfig("importListResourceTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListResourceConfig("importListResourceTest", "specificBook"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list.test", "should_monitor", "specificBook"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListResourceConfig(name, monitor string) string {
	return fmt.Sprintf(`
	resource "readarr_import_list" "test" {
		enable_automatic_add = false
		should_monitor = "%s"
		monitor_new_items = "none"
		should_search = false
		list_type = "program"
		root_folder_path = "/config"
		quality_profile_id = 1
		metadata_profile_id = 1
		name = "%s"
		implementation = "ReadarrImport"
    	config_contract = "ReadarrSettings"
		base_url = "http://127.0.0.1:8787"
		api_key = "testAPIKey"
		tag_ids = [1,2]
		profile_ids = [1]
		tags = []
	}`, monitor, name)
}
