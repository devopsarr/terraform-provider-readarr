package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationGoodreadsOwnedBooksResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "testAccessToken") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "testAccessToken"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_goodreads_owned_books.test", "access_token", "testAccessToken"),
					resource.TestCheckResourceAttrSet("readarr_notification_goodreads_owned_books.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "testAccessToken") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationGoodreadsOwnedBooksResourceConfig("resourceGoodreadsOwnedBooksTest", "testAccessToken123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_goodreads_owned_books.test", "access_token", "testAccessToken123"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_goodreads_owned_books.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationGoodreadsOwnedBooksResourceConfig(name, token string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_goodreads_owned_books" "test" {
		on_upgrade                        = false
		on_release_import                 = false
	  
		name                    = "%s"
	  
		access_token = "%s"
		access_token_secret = "testAccessTokenSecret"
		user_id = "163730408"
		username = "Test User"
		condition = 20
		description = "with issues"
		location = "Dubai"
	}`, name, token)
}
