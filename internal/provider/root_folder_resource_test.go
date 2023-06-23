package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRootFolderResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccRootFolderResourceConfig("/error", "Error") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccRootFolderResourceConfig("/config/asp", "ResourceTest"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_root_folder.test", "path", "/config/asp"),
					resource.TestCheckResourceAttrSet("readarr_root_folder.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccRootFolderResourceConfig("/error", "Error") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccRootFolderResourceConfig("/config/logs", "ResourceTest"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_root_folder.test", "path", "/config/logs"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_root_folder.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRootFolderResourceConfig(path, name string) string {
	return fmt.Sprintf(`
		resource "readarr_root_folder" "test" {
  			path = "%s"
			name = "%s"
			default_metadata_profile_id = 1
			default_quality_profile_id = 1
			default_monitor_option = "all"
			default_monitor_new_item_option = "all"
			default_tags = []

			output_profile = "default"

			is_calibre_library = false
		}
	`, path, name)
}
