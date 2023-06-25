package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListGoodreadsBookshelfResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListGoodreadsBookshelfResourceConfig("resourceGoodreadsBookshelfTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListGoodreadsBookshelfResourceConfig("resourceGoodreadsBookshelfTest", "entireAuthor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_bookshelf.test", "should_monitor", "entireAuthor"),
					resource.TestCheckResourceAttrSet("readarr_import_list_goodreads_bookshelf.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListGoodreadsBookshelfResourceConfig("resourceGoodreadsBookshelfTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListGoodreadsBookshelfResourceConfig("resourceGoodreadsBookshelfTest", "specificBook"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_bookshelf.test", "should_monitor", "specificBook"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list_goodreads_bookshelf.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListGoodreadsBookshelfResourceConfig(name, folder string) string {
	return fmt.Sprintf(`
	resource "readarr_import_list_goodreads_bookshelf" "test" {
		enable_automatic_add = false
		should_monitor = "%s"
		monitor_new_items = "none"
		should_search = false
		root_folder_path = "/config"
		quality_profile_id = 1
		metadata_profile_id = 1
		name = "%s"
		access_token = "testAccessToken"
		access_token_secret = "testAccessTokenSecret"
		user_id = "163730408"
		username = "Test User"
		bookshelf_ids = ["currently-reading","read","to-read"]
	}`, folder, name)
}
