resource "readarr_notification" "example" {
  on_grab                         = false
  on_download_failure             = true
  on_upgrade                      = false
  on_rename                       = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_book_retag                   = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  implementation  = "CustomScript"
  config_contract = "CustomScriptSettings"

  path = "/scripts/readarr.sh"
}