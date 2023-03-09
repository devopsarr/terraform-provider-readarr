package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationTwitterResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationTwitterResourceConfig("resourceTwitterTest", "me") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationTwitterResourceConfig("resourceTwitterTest", "me"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_twitter.test", "mention", "me"),
					resource.TestCheckResourceAttrSet("readarr_notification_twitter.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationTwitterResourceConfig("resourceTwitterTest", "me") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationTwitterResourceConfig("resourceTwitterTest", "myself"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_twitter.test", "mention", "myself"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_twitter.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationTwitterResourceConfig(name, mention string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_twitter" "test" {
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
	  
		access_token = "Token"
		access_token_secret = "TokenSecret"
		consumer_key = "Key"
		consumer_secret = "Secret"
		mention = "%s"
	}`, name, mention)
}
