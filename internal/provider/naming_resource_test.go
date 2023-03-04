package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNamingResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccNamingResourceConfig("{Author Name}") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccNamingResourceConfig("{Author Name}"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_naming.test", "author_folder_format", "{Author Name}"),
					resource.TestCheckResourceAttrSet("readarr_naming.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccNamingResourceConfig("{Author Name}") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccNamingResourceConfig("{Author_Name}"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_naming.test", "author_folder_format", "{Author_Name}"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_naming.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNamingResourceConfig(author string) string {
	return fmt.Sprintf(`
	resource "readarr_naming" "test" {
		rename_books               = true
		replace_illegal_characters = true
		author_folder_format       = "%s"
		standard_book_format       = "{Book Title}/{Author Name} - {Book Title}{ (PartNumber)}"
	}`, author)
}
