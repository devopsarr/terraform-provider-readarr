resource "readarr_notification_twitter" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = true
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  access_token        = "Token"
  access_token_secret = "TokenSecret"
  consumer_key        = "Key"
  consumer_secret     = "Secret"
  mention             = "someone"
}