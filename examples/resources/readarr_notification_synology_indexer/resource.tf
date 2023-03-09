resource "readarr_notification_synology_indexer" "example" {
  on_upgrade                      = false
  on_rename                       = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_book_retag                   = true
  on_author_delete                = false
  on_release_import               = false

  name = "Example"

  update_library = true
}