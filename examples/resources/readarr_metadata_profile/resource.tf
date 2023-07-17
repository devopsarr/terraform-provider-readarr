resource "readarr_metadata_profile" "example" {
  name                  = "Example"
  allowed_languages     = "eng,ita"
  ignored               = ["alpha", "beta"]
  min_popularity        = 3.5
  min_pages             = 10
  skip_missing_date     = false
  skip_missing_isbn     = true
  skip_parts_and_sets   = false
  skip_series_secondary = false
}