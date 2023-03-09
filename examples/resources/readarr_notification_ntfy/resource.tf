resource "readarr_notification_ntfy" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_author_delete                = false
  on_release_import               = true

  include_health_warnings = false
  name                    = "Example"

  priority   = 1
  server_url = "https://ntfy.sh"
  username   = "User"
  password   = "Pass"
  topics     = ["Topic1234", "Topic4321"]
  field_tags = ["warning", "skull"]
}