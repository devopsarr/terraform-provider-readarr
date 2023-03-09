resource "readarr_tag" "example" {
  monitored          = true
  author_name        = "Leo Tolstoy"
  path               = "/books/leotolstoy"
  quality_profile_id = 1
  foreign_author_id  = "128382"
}
