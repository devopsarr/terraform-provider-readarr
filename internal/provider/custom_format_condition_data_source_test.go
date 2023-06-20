package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomFormatConditionDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccCustomFormatConditionDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_custom_format_condition.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_custom_format_condition.test", "name", "Preferred Words"),
					resource.TestCheckResourceAttr("readarr_custom_format.test", "specifications.0.implementation", "ReleaseTitleSpecification")),
			},
		},
	})
}

const testAccCustomFormatConditionDataSourceConfig = `
data  "readarr_custom_format_condition" "test" {
	name = "Preferred Words"
	implementation = "ReleaseTitleSpecification"
	negate = false
	required = false
	value = "\\b(SPARKS|Framestor)\\b"
}

data  "readarr_custom_format_condition" "test1" {
	name = "Size"
	implementation = "SizeSpecification"
	negate = false
	required = false
	min = 0
	max = 100
}

resource "readarr_custom_format" "test" {
	include_custom_format_when_renaming = false
	name = "TestWithDS"
	
	specifications = [data.readarr_custom_format_condition.test,data.readarr_custom_format_condition.test1]	
}`
