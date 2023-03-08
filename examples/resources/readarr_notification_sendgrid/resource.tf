resource "readarr_notification_sendgrid" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  api_key    = "APIkey"
  from       = "from_sendgrid@example.com"
  recipients = ["user1@example.com", "user2@example.com"]
}