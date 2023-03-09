resource "readarr_notification_kavita" "example" {
  on_book_retag                   = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_release_import               = false

  name = "Example"

  api_key = "APIKey"
  host    = "kavita.local"
  port    = 4040
  notify  = true
}