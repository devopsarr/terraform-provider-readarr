resource "readarr_notification_email" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_author_delete                = false
  on_release_import               = false

  server       = "http://email-server.net"
  port         = 587
  from         = "from_email@example.com"
  to           = ["user1@example.com", "user2@example.com"]
  attach_files = true
}