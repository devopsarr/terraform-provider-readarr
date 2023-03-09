resource "readarr_notification_gooodreads_bookshelves" "example" {
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  access_token        = "exampleAccessToken"
  access_token_secret = "exampleAccessTokenSecret"
  user_id             = "163730408"
  username            = "Example User"
  add_ids             = ["currently-reading", "read", "to-read"]
  remove_ids          = ["test"]
}