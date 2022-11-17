resource "readarr_download_client_config" "example" {
  remove_completed_downloads         = false
  remove_failed_downloads            = false
  enable_completed_download_handling = true
  auto_redownload_failed             = false
}