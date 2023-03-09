resource "readarr_notification_slack" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_rename                       = true
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_book_retag                   = true
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  web_hook_url = "http://my.slack.com/test"
  username     = "user"
  channel      = "example-channel"
}