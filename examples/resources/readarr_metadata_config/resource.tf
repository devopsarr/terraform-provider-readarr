resource "readarr_metadata_config" "example" {
  write_audio_tags = "no"
  write_book_tags  = "newFiles"
  update_covers    = true
  embed_metadata   = true
  scrub_audio_tags = false
}