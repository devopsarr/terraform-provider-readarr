package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIndexerTorrentRssResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccIndexerTorrentRssResourceConfig("rssResourceTest", "https://rss.org") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccIndexerTorrentRssResourceConfig("rssResourceTest", "https://rss.org"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_indexer_torrent_rss.test", "base_url", "https://rss.org"),
					resource.TestCheckResourceAttrSet("readarr_indexer_torrent_rss.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccIndexerTorrentRssResourceConfig("rssResourceTest", "https://rss.org") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccIndexerTorrentRssResourceConfig("rssResourceTest", "https://rss.net"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_indexer_torrent_rss.test", "base_url", "https://rss.net"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_indexer_torrent_rss.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIndexerTorrentRssResourceConfig(name, url string) string {
	return fmt.Sprintf(`
	resource "readarr_indexer_torrent_rss" "test" {
		enable_rss = false
		name = "%s"
		base_url = "%s"
		allow_zero_size = true
		minimum_seeders = 1
	}`, name, url)
}
