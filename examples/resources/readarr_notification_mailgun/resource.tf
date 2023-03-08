resource "readarr_notification_mailgun" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  api_key    = "APIkey"
  from       = "from_mailgun@example.com"
  recipients = ["user1@example.com", "user2@example.com"]
}