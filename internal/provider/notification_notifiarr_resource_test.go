package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationNotifiarrResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationNotifiarrResourceConfig("resourceNotifiarrTest", "testAPIKey") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationNotifiarrResourceConfig("resourceNotifiarrTest", "testAPIKey"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_notifiarr.test", "api_key", "testAPIKey"),
					resource.TestCheckResourceAttrSet("readarr_notification_notifiarr.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationNotifiarrResourceConfig("resourceNotifiarrTest", "testAPIKey") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationNotifiarrResourceConfig("resourceNotifiarrTest", "testAPIKey123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_notifiarr.test", "api_key", "testAPIKey123"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_notifiarr.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationNotifiarrResourceConfig(name, key string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_notifiarr" "test" {
		on_grab                           = false
		on_upgrade                        = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_health_issue                   = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		api_key = "%s"
	}`, name, key)
}
