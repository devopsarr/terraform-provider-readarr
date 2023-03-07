resource "readarr_indexer_gazelle" "example" {
  enable_automatic_search = true
  name                    = "Example"
  base_url                = "https://orpheus.network"
  username                = "User"
  password                = "Pass"
  minimum_seeders         = 1
}
