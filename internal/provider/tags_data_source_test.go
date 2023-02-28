package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTagsDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a resource to have a value to check
			{
				Config: testAccTagResourceConfig("test-1", "fiction") + testAccTagResourceConfig("test-2", "travel"),
			},
			// Read testing
			{
				Config: testAccTagsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs("data.readarr_tags.test", "tags.*", map[string]string{"label": "fiction"}),
				),
			},
		},
	})
}

const testAccTagsDataSourceConfig = `
data "readarr_tags" "test" {
}
`
