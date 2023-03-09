package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationSubsonicResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationSubsonicResourceConfig("resourceSubsonicTest", "key") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationSubsonicResourceConfig("resourceSubsonicTest", "key"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_subsonic.test", "password", "key"),
					resource.TestCheckResourceAttrSet("readarr_notification_subsonic.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationSubsonicResourceConfig("resourceSubsonicTest", "key") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationSubsonicResourceConfig("resourceSubsonicTest", "key1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_subsonic.test", "password", "key1"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_subsonic.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationSubsonicResourceConfig(name, password string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_subsonic" "test" {
		on_grab                           = false
		on_upgrade                        = false
		on_rename                         = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_health_issue                   = false
		on_book_retag 					  = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		host = "http://subsonic.com"
		port = 8080
		username = "User"
		password = "%s"
		notify = true
	}`, name, password)
}
