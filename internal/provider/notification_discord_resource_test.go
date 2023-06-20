package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationDiscordResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationDiscordResourceConfig("resourceDiscordTest", "dog-picture") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationDiscordResourceConfig("resourceDiscordTest", "dog-picture"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_discord.test", "avatar", "dog-picture"),
					resource.TestCheckResourceAttrSet("readarr_notification_discord.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationDiscordResourceConfig("resourceDiscordTest", "dog-picture") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationDiscordResourceConfig("resourceDiscordTest", "cat-picture"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_discord.test", "avatar", "cat-picture"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_discord.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationDiscordResourceConfig(name, avatar string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_discord" "test" {
		on_grab                           = false
		on_download_failure               = false
		on_upgrade                        = false
		on_rename                         = false
		on_import_failure                 = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_health_issue                   = false
		on_book_retag 					  = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		web_hook_url  = "http://discord-web-hook.com"
		username      = "User"
		avatar        = "%s"
	}`, name, avatar)
}
