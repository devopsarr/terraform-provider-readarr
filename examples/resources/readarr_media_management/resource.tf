resource "readarr_media_management" "example" {
  unmonitor_previous_books    = false
  hardlinks_copy              = true
  create_empty_author_folders = true
  delete_empty_folders        = true
  watch_ibrary_for_changes    = true
  import_extra_files          = true
  set_permissions             = true
  skip_free_space_check       = true
  minimum_free_space          = 100
  recycle_bin_days            = 7
  chmod_folder                = "755"
  chown_group                 = "arrs"
  download_propers_repacks    = "preferAndUpgrade"
  allow_fingerprinting        = "never"
  extra_file_extensions       = "info"
  file_date                   = "bookReleaseDate"
  recycle_bin_path            = "/bin"
  rescan_after_refresh        = "always"
}