---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_indexer_gazelle Resource - terraform-provider-readarr"
subcategory: "Indexers"
description: |-
  Indexer Gazelle resource.
  For more information refer to Indexer https://wiki.servarr.com/readarr/settings#indexers and Gazelle https://wiki.servarr.com/readarr/supported#gazelle.
---

# readarr_indexer_gazelle (Resource)

<!-- subcategory:Indexers -->Indexer Gazelle resource.
For more information refer to [Indexer](https://wiki.servarr.com/readarr/settings#indexers) and [Gazelle](https://wiki.servarr.com/readarr/supported#gazelle).

## Example Usage

```terraform
resource "readarr_indexer_gazelle" "example" {
  enable_automatic_search = true
  name                    = "Example"
  base_url                = "https://orpheus.network"
  username                = "User"
  password                = "Pass"
  minimum_seeders         = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) IndexerGazelle name.
- `password` (String, Sensitive) Password.
- `username` (String) Username.

### Optional

- `author_seed_time` (Number) Author seed time.
- `base_url` (String) Base URL.
- `early_release_limit` (Number) Early release limit.
- `enable_automatic_search` (Boolean) Enable automatic search flag.
- `enable_interactive_search` (Boolean) Enable interactive search flag.
- `enable_rss` (Boolean) Enable RSS flag.
- `minimum_seeders` (Number) Minimum seeders.
- `priority` (Number) Priority.
- `seed_ratio` (Number) Seed ratio.
- `seed_time` (Number) Seed time.
- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) IndexerGazelle ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_indexer_gazelle.example 1
```
