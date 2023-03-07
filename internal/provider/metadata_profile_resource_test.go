package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMetadataProfileResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccMetadataProfileResourceConfig("error", "eng") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccMetadataProfileResourceConfig("remotemapResourceTest", "eng"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_metadata_profile.test", "allowed_languages", "eng"),
					resource.TestCheckResourceAttrSet("readarr_metadata_profile.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccMetadataProfileResourceConfig("error", "eng") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccMetadataProfileResourceConfig("profileResourceTest", "ita"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_metadata_profile.test", "allowed_languages", "ita"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_metadata_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccMetadataProfileResourceConfig(name, language string) string {
	return fmt.Sprintf(`
		resource "readarr_metadata_profile" "test" {
  			name = "%s"
			allowed_languages = "%s"
			min_popularity = 3.5
			min_pages = 10
			skip_missing_date = false
			skip_missing_isbn = true
			skip_parts_and_sets = false
			skip_series_secondary = false
		}
	`, name, language)
}
