resource "readarr_notification_gotify" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  server    = "http://gotify-server.net"
  app_token = "Token"
  priority  = 5
}