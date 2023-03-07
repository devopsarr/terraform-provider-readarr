package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccQualityProfileResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccQualityProfileResourceError + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccQualityProfileResourceConfig("example-epub"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_quality_profile.test", "name", "example-epub"),
					resource.TestCheckResourceAttrSet("readarr_quality_profile.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccQualityProfileResourceError + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccQualityProfileResourceConfig("example-mobi"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_quality_profile.test", "name", "example-mobi"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_quality_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

const testAccQualityProfileResourceError = `
resource "readarr_quality_profile" "test" {
	name            = "Error"
	upgrade_allowed = true
	cutoff          = 2000
	quality_groups = []
}
`

func testAccQualityProfileResourceConfig(name string) string {
	return fmt.Sprintf(`
	data "readarr_quality" "epub" {
		name = "EPUB"
	}

	data "readarr_quality" "mobi" {
		name = "MOBI"
	}

	data "readarr_quality" "pdf" {
		name = "PDF"
	}

	resource "readarr_quality_profile" "test" {
		name            = "%s"
		upgrade_allowed = true
		cutoff          = 2000

		quality_groups = [
			{
				id   = 2000
				name = "lossless"
				qualities = [
					data.readarr_quality.mobi,
					data.readarr_quality.epub,
				]
			},
			{
				qualities = [data.readarr_quality.pdf]
			}
		]
	}
	`, name)
}
