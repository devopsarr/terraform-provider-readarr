package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccImportListGoodreadsOwnedBooksResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "entireAuthor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_owned_books.test", "should_monitor", "entireAuthor"),
					resource.TestCheckResourceAttrSet("readarr_import_list_goodreads_owned_books.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "specificBook"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_owned_books.test", "should_monitor", "specificBook"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list_goodreads_owned_books.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListGoodreadsOwnedBooksResourceConfig(name, folder string) string {
	return fmt.Sprintf(`
	resource "readarr_import_list_goodreads_owned_books" "test" {
		enable_automatic_add = false
		should_monitor = "%s"
		should_search = false
		root_folder_path = "/config"
		quality_profile_id = 1
		metadata_profile_id = 1
		name = "%s"
		access_token = "testAccessToken"
		access_token_secret = "testAccessTokenSecret"
		user_id = "163730408"
		username = "Test User"
	}`, folder, name)
}
