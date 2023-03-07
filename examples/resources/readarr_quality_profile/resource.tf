resource "readarr_quality_profile" "example" {
  name            = "example-native"
  upgrade_allowed = true
  cutoff          = 1100

  quality_groups = [
    {
      id   = 1100
      name = "native"
      qualities = [
        {
          id   = 3
          name = "EPUB"
        },
        {
          id   = 2
          name = "MOBI"
        }
      ]
    }
  ]
}