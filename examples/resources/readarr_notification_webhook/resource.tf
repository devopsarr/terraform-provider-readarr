resource "readarr_notification_webhook" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_rename                       = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_book_retag                   = true
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  url      = "https://example.webhook.com/example"
  method   = 1
  username = "exampleUser"
  password = "examplePass"
}