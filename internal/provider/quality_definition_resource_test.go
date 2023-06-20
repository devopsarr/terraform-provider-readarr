package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccQualityDefinitionResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccQualityDefinitionResourceConfig("example-EPUB") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccQualityDefinitionResourceConfig("example-EPUB"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_quality_definition.test", "title", "example-EPUB"),
					resource.TestCheckResourceAttrSet("readarr_quality_definition.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccQualityDefinitionResourceConfig("example-EPUB") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccQualityDefinitionResourceConfig("example-MOBI"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_quality_definition.test", "title", "example-MOBI"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_quality_definition.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccQualityDefinitionResourceConfig(name string) string {
	return fmt.Sprintf(`
	resource "readarr_quality_definition" "test" {
		id = 1
		title    = "%s"
		min_size = 35.0
		max_size = 400
	}
	`, name)
}
