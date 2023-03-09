package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAuthorsDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized
			{
				Config:      testAccAuthorResourceConfig("Error", "error", "error") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Read testing
			{
				Config: testAccAuthorResourceConfig("J.K. Rowling", "jkrowling", "1077326") + testAccAuthorsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs("data.readarr_authors.test", "authors.*", map[string]string{"author_name": "J.K. Rowling"}),
				),
			},
		},
	})
}

const testAccAuthorsDataSourceConfig = `
data "readarr_authors" "test" {
	depends_on = [readarr_author.test]
}
`
