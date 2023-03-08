package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationGoodreadsBookshelvesResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNotificationGoodreadsBookshelvesResourceConfig("resourceGoodreadsBookshelvesTest", "testAccessToken") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNotificationGoodreadsBookshelvesResourceConfig("resourceGoodreadsBookshelvesTest", "testAccessToken"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_goodreads_bookshelves.test", "access_token", "testAccessToken"),
					resource.TestCheckResourceAttrSet("readarr_notification_goodreads_bookshelves.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNotificationGoodreadsBookshelvesResourceConfig("resourceGoodreadsBookshelvesTest", "testAccessToken") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNotificationGoodreadsBookshelvesResourceConfig("resourceGoodreadsBookshelvesTest", "testAccessToken123"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_goodreads_bookshelves.test", "access_token", "testAccessToken123"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_goodreads_bookshelves.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationGoodreadsBookshelvesResourceConfig(name, token string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_goodreads_bookshelves" "test" {
		on_upgrade                        = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		access_token = "%s"
		access_token_secret = "testAccessTokenSecret"
		user_id = "163730408"
		username = "Test User"
		add_ids = ["currently-reading","read","to-read"]
		remove_ids = ["test"]
	}`, name, token)
}
