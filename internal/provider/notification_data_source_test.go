package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccNotificationDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.readarr_notification.test", "id"),
					resource.TestCheckResourceAttr("data.readarr_notification.test", "path", "/scripts/test.sh")),
			},
		},
	})
}

const testAccNotificationDataSourceConfig = `
resource "readarr_notification" "test" {
	on_grab                            = false
	on_download_failure                = true
	on_upgrade                         = false
	on_rename                          = false
	on_import_failure                  = false
	on_book_delete                    = false
	on_book_file_delete               = false
	on_book_file_delete_for_upgrade   = true
	on_health_issue                   = false
	on_book_retag 					  = false
	on_author_delete                  = false
	on_release_import                 = false
  
	include_health_warnings = false
	name                    = "notificationData"
  
	implementation  = "CustomScript"
	config_contract = "CustomScriptSettings"
  
	path = "/scripts/test.sh"
}

data "readarr_notification" "test" {
	name = readarr_notification.test.name
}
`
