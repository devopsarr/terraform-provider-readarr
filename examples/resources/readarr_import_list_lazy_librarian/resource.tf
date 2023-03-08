resource "readarr_import_list_lazy_librarian" "example" {
  enable_automatic_add = false
  should_monitor       = "specificBook"
  should_search        = false
  root_folder_path     = "/books"
  quality_profile_id   = 1
  metadata_profile_id  = 1
  name                 = "Example"
  base_url             = "http://127.0.0.1:5299"
  api_key              = "APIKey"
  tags                 = [1, 2, 3]
}
