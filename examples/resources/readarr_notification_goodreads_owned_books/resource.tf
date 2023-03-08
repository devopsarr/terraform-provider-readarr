resource "readarr_notification_gooodreads_owned_books" "example" {
  on_upgrade        = false
  on_release_import = false

  name = "Example"

  access_token        = "exampleAccessToken"
  access_token_secret = "exampleAccessTokenSecret"
  user_id             = "163730408"
  username            = "Example User"
  condition           = 20
  description         = "with issues"
  location            = "Dubai"
}