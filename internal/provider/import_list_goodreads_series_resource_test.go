package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImportListGoodreadsSeriesResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccImportListGoodreadsSeriesResourceConfig("resourceGoodreadsSeriesTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccImportListGoodreadsSeriesResourceConfig("resourceGoodreadsSeriesTest", "entireAuthor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_series.test", "should_monitor", "entireAuthor"),
					resource.TestCheckResourceAttrSet("readarr_import_list_goodreads_series.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccImportListGoodreadsSeriesResourceConfig("resourceGoodreadsSeriesTest", "entireAuthor") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccImportListGoodreadsSeriesResourceConfig("resourceGoodreadsSeriesTest", "specificBook"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_import_list_goodreads_series.test", "should_monitor", "specificBook"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_import_list_goodreads_series.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccImportListGoodreadsSeriesResourceConfig(name, folder string) string {
	return fmt.Sprintf(`
	resource "readarr_import_list_goodreads_series" "test" {
		enable_automatic_add = false
		should_monitor = "%s"
		should_search = false
		root_folder_path = "/config"
		quality_profile_id = 1
		metadata_profile_id = 1
		name = "%s"
		series_id = 45175
	}`, folder, name)
}
