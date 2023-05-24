package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCustomFormatConditionReleaseGroupDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccCustomFormatConditionReleaseGroupDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_custom_format_condition_release_group.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_custom_format_condition_release_group.test", "name", "HDBits"),
					resource.TestCheckResourceAttr("readarr_custom_format.test", "specifications.0.value", ".*HDBits.*")),
			},
		},
	})
}

const testAccCustomFormatConditionReleaseGroupDataSourceConfig = `
data  "readarr_custom_format_condition_release_group" "test" {
	name = "HDBits"
	negate = false
	required = false
	value = ".*HDBits.*"
}

resource "readarr_custom_format" "test" {
	include_custom_format_when_renaming = false
	name = "TestWithDSReleaseGroup"
	
	specifications = [data.readarr_custom_format_condition_release_group.test]	
}`
