package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationCustomScriptResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccNotificationCustomScriptResourceConfig("resourceScriptTest", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_custom_script.test", "on_upgrade", "false"),
					resource.TestCheckResourceAttrSet("readarr_notification_custom_script.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccNotificationCustomScriptResourceConfig("resourceScriptTest", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("readarr_notification_custom_script.test", "on_upgrade", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "readarr_notification_custom_script.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNotificationCustomScriptResourceConfig(name, upgrade string) string {
	return fmt.Sprintf(`
	resource "readarr_notification_custom_script" "test" {
		on_grab                            = false
		on_download_failure                = true
		on_upgrade                         = %s
		on_rename                          = false
		on_import_failure                  = false
		on_book_delete                    = false
		on_book_file_delete               = false
		on_book_file_delete_for_upgrade   = true
		on_health_issue                   = false
		on_application_update             = false
		on_book_retag 					  = false
		on_author_delete                  = false
		on_release_import                 = false
	  
		include_health_warnings = false
		name                    = "%s"
	  
		path = "/scripts/test.sh"
	}`, upgrade, name)
}
