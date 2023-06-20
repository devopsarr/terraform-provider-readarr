package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNotificationKavitaResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationKavitaResourceConfig("resourceKavitaTest", "true") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationKavitaResourceConfig("resourceKavitaTest", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_kavita.test", "notify", "true"),
					resource.TestCheckResourceAttrSet("readarr_notification_kavita.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationKavitaResourceConfig("resourceKavitaTest", "true") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationKavitaResourceConfig("resourceKavitaTest", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_kavita.test", "notify", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_kavita.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationKavitaResourceConfig(name, notify string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_kavita" "test" {
		on_book_retag                     = false
		on_upgrade                        = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_release_import                 = false
	  
		name                    = "%s"
	  
		api_key = "APIKey"
		host = "kavita.local"
		port = 4040
		notify = %s
	}`, name, notify)
}
