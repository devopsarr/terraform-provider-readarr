package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAuthorResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccAuthorResourceConfig("Error", "test", "656983") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				PreConfig: rootFolderDSInit,
				Config:    testAccAuthorResourceConfig("J.R.R. Tolkien", "test", "656983"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_author.test", "path", "/config/test"),
					resource.TestCheckResourceAttrSet("readarr_author.test", "id"),
					resource.TestCheckResourceAttr("readarr_author.test", "author_name", "J.R.R. Tolkien"),
					resource.TestCheckResourceAttr("readarr_author.test", "status", "continuing"),
					resource.TestCheckResourceAttr("readarr_author.test", "monitored", "false"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccAuthorResourceConfig("Error", "test", "656983") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccAuthorResourceConfig("J.R.R. Tolkien", "test123", "656983"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_author.test", "path", "/config/test123"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_author.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAuthorResourceConfig(title, path, foreignID string) string {
	return fmt.Sprintf(`
		resource "readarr_author" "test" {
			monitored = false
			author_name = "%s"
			path = "/config/%s"
			quality_profile_id = 1
			foreign_author_id = "%s"
		}
	`, title, path, foreignID)
}
