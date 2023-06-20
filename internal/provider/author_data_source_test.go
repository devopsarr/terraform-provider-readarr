package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAuthorDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized
			{
				Config:      testAccAuthorDataSourceConfig("\"999\"") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Not found testing
			{
				Config:      testAccAuthorDataSourceConfig("\"999\""),
				ExpectError: regexp.MustCompile("Unable to find author"),
			},
			// Read testing
			{
				Config: testAccAuthorResourceConfig("Agatha Christie", "agathachristie", "123715") + testAccAuthorDataSourceConfig("readarr_author.test.foreign_author_id"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_author.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_author.test", "author_name", "Agatha Christie"),
				),
			},
		},
	})
}

func testAccAuthorDataSourceConfig(id string) string {
	return fmt.Sprintf(`
	data "readarr_author" "test" {
		foreign_author_id = %s
	}
	`, id)
}
