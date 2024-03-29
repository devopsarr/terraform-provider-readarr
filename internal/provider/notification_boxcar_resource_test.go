package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationBoxcarResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationBoxcarResourceConfig("resourceBoxcarTest", "token123") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationBoxcarResourceConfig("resourceBoxcarTest", "token123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_boxcar.test", "token", "token123"),
					resource.TestCheckResourceAttrSet("readarr_notification_boxcar.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationBoxcarResourceConfig("resourceBoxcarTest", "token123") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationBoxcarResourceConfig("resourceBoxcarTest", "token234"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_boxcar.test", "token", "token234"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_boxcar.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationBoxcarResourceConfig(name, token string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_boxcar" "test" {
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
	  
		token = "%s"
	}`, name, token)
}
