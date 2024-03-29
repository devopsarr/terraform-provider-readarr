resource "readarr_media_management" "example" {
  rename_books               = true
  replace_illegal_characters = true
  colon_replacement_format   = 0
  author_folder_format       = "{Author Name}"
  standard_book_format       = "{Book Title}/{Author Name} - {Book Title}{ (PartNumber)}"
}