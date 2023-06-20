package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationGotifyResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0) + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_gotify.test", "priority", "0"),
					resource.TestCheckResourceAttrSet("readarr_notification_gotify.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationGotifyResourceConfig("resourceGotifyTest", 0) + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationGotifyResourceConfig("resourceGotifyTest", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_gotify.test", "priority", "2"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_gotify.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationGotifyResourceConfig(name string, priority int) string {
	return fmt.Sprintf(`
	resource "readarr_notification_gotify" "test" {
		on_grab                           = false
		on_download_failure               = false
		on_upgrade                        = false
		on_import_failure                 = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_health_issue                   = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		server = "http://gotify-server.net"
		app_token = "Token"
		priority = %d
	}`, name, priority)
}
