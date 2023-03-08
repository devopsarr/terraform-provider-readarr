resource "readarr_import_list_goodreads_owned_books" "example" {
  enable_automatic_add = false
  should_monitor       = "specificBook"
  should_search        = false
  root_folder_path     = "/books"
  quality_profile_id   = 1
  metadata_profile_id  = 1
  name                 = "Example"
  access_token         = "ExampleAccessToken"
  access_token_secret  = "ExampleAccessTokenSecret"
  user_id              = "163730408"
  username             = "Test User"
  tags                 = [1, 2, 3]
}
