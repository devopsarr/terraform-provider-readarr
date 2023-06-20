package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationProwlResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationProwlResourceConfig("resourceProwlTest", "testAPIKey") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationProwlResourceConfig("resourceProwlTest", "testAPIKey"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_prowl.test", "api_key", "testAPIKey"),
					resource.TestCheckResourceAttrSet("readarr_notification_prowl.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationProwlResourceConfig("resourceProwlTest", "testAPIKey") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationProwlResourceConfig("resourceProwlTest", "testAPIKey123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_prowl.test", "api_key", "testAPIKey123"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_prowl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationProwlResourceConfig(name, key string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_prowl" "test" {
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
		
		priority = 2
		api_key = "%s"
	}`, name, key)
}
