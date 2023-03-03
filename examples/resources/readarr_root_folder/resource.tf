resource "readarr_root_folder" "example" {
  path                        = "/example"
  name                        = "Example"
  default_metadata_profile_id = 1
  default_quality_profile_id  = 1
  default_monitor_option      = "all"
  is_calibre_library          = false
  # keep "default" if not used
  output_profile = "default"
}

# with calibre
resource "readarr_root_folder" "calibre_example" {
  path                        = "/calibre"
  name                        = "Calibre"
  default_metadata_profile_id = 1
  default_quality_profile_id  = 1
  default_monitor_option      = "all"
  is_calibre_library          = true

  # calibre server values
  use_ssl        = false
  host           = "calibre-host"
  port           = 8081
  username       = "User"
  password       = "Pass"
  output_profile = "default"
  library        = "Calibre_Library"
}
