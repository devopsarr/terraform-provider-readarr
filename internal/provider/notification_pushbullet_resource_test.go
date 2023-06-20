package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationPushbulletResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationPushbulletResourceConfig("resourcePushbulletTest", "key1") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationPushbulletResourceConfig("resourcePushbulletTest", "key1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_pushbullet.test", "api_key", "key1"),
					resource.TestCheckResourceAttrSet("readarr_notification_pushbullet.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationPushbulletResourceConfig("resourcePushbulletTest", "key1") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationPushbulletResourceConfig("resourcePushbulletTest", "key2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_pushbullet.test", "api_key", "key2"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_pushbullet.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationPushbulletResourceConfig(name, key string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_pushbullet" "test" {
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
	  
		api_key = "%s"
		device_ids = ["test"]
	}`, name, key)
}
