resource "readarr_notification_discord" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_rename                       = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_book_retag                   = true
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  web_hook_url  = "http://discord-web-hook.com"
  username      = "User"
  avatar        = "https://i.imgur.com/oBPXx0D.png"
  grab_fields   = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
  import_fields = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
}