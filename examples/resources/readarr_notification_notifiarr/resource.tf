resource "readarr_notification_notifiarr" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  api_key = "Token"
}