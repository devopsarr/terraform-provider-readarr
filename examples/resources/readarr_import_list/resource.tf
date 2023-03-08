resource "readarr_import_list" "example" {
  enable_automatic_add = false
  should_monitor       = "entireAuthor"
  should_search        = false
  list_type            = "program"
  monitor_new_items    = "all"
  root_folder_path     = readarr_root_folder.example.path
  quality_profile_id   = readarr_quality_profile.example.id
  metadata_profile_id  = readarr_metadata_profile.example.id
  name                 = "Example"
  implementation       = "ReadarrImport"
  config_contract      = "ReadarrSettings"
  tags                 = [1, 2]

  tag_ids     = [1, 2]
  profile_ids = [1]
  base_url    = "http://127.0.0.1:8686"
  api_key     = "APIKey"
}