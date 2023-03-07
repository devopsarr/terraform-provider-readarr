resource "readarr_release_profile" "example" {
  enabled                         = true
  include_preferred_when_renaming = true
  indexer_id                      = 0
  required                        = "epub,isdn"
  ignored                         = "pdf"
  preferred = [
    {
      term  = "higher"
      score = 100
    },
    {
      term  = "lower"
      score = -100
    },
  ]
}