package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDownloadClientFloodResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccDownloadClientFloodResourceConfig("resourceFloodTest", "flood") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccDownloadClientFloodResourceConfig("resourceFloodTest", "flood"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_download_client_flood.test", "host", "flood"),
					resource.TestCheckResourceAttr("readarr_download_client_flood.test", "url_base", "/flood/"),
					resource.TestCheckResourceAttrSet("readarr_download_client_flood.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccDownloadClientFloodResourceConfig("resourceFloodTest", "flood") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccDownloadClientFloodResourceConfig("resourceFloodTest", "flood-host"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_download_client_flood.test", "host", "flood-host"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_download_client_flood.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccDownloadClientFloodResourceConfig(name, host string) string {
	return fmt.Sprintf(`
	resource "readarr_download_client_flood" "test" {
		enable = false
		priority = 1
		name = "%s"
		host = "%s"
		url_base = "/flood/"
		port = 9091
		add_paused = true
		additional_tags = [0,1]
		field_tags = ["readarr"]
	}`, name, host)
}
