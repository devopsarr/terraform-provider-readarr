package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCustomFormatConditionSizeDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccCustomFormatConditionSizeDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_custom_format_condition_size.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_custom_format_condition_size.test", "name", "Test"),
					resource.TestCheckResourceAttr("readarr_custom_format.test", "specifications.0.min", "5"),
					resource.TestCheckResourceAttr("readarr_custom_format.test", "specifications.0.max", "50")),
			},
		},
	})
}

const testAccCustomFormatConditionSizeDataSourceConfig = `
data  "readarr_custom_format_condition_size" "test" {
	name = "Test"
	negate = false
	required = false
	min = 5
	max = 50
}

resource "readarr_custom_format" "test" {
	include_custom_format_when_renaming = false
	name = "TestWithDSSize"
	
	specifications = [data.readarr_custom_format_condition_size.test]	
}`
