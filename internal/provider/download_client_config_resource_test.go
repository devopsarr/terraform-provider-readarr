package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDownloadClientConfigResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccDownloadClientConfigResourceConfig("true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_download_client_config.test", "auto_redownload_failed", "true"),
					resource.TestCheckResourceAttrSet("readarr_download_client_config.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccDownloadClientConfigResourceConfig("false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_download_client_config.test", "auto_redownload_failed", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_download_client_config.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccDownloadClientConfigResourceConfig(redownload string) string {
	return fmt.Sprintf(`
	resource "readarr_download_client_config" "test" {
		remove_completed_downloads = false
		remove_failed_downloads = false
		enable_completed_download_handling = true
		auto_redownload_failed = %s
	}`, redownload)
}
