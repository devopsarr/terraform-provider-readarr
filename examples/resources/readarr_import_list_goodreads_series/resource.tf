resource "readarr_import_list_goodreads_series" "example" {
  enable_automatic_add = false
  should_monitor       = "specificBook"
  should_search        = false
  root_folder_path     = "/books"
  quality_profile_id   = 1
  metadata_profile_id  = 1
  series_id            = 45175
  tags                 = [1, 2, 3]
}
