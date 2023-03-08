resource "readarr_import_list_goodreads_list" "example" {
  enable_automatic_add = false
  should_monitor       = "specificBook"
  should_search        = false
  root_folder_path     = "/config"
  quality_profile_id   = 1
  metadata_profile_id  = 1
  list_id              = 8544254
  tags                 = [1, 2, 3]
}
