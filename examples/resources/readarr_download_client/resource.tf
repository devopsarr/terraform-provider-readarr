resource "readarr_download_client" "example" {
  enable                     = true
  remove_completed_downloads = false
  remove_failed_downloads    = false
  priority                   = 1
  name                       = "Example"
  implementation             = "Transmission"
  protocol                   = "torrent"
  config_contract            = "TransmissionSettings"
  host                       = "transmission"
  url_base                   = "/transmission/"
  port                       = 9091
}