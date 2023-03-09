resource "readarr_notification_subsonic" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_rename                       = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_book_retag                   = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  host     = "http://subsonic.com"
  port     = 8080
  username = "User"
  password = "MyPass"
  notify   = true
}