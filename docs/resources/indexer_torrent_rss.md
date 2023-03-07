---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_indexer_torrent_rss Resource - terraform-provider-readarr"
subcategory: "Indexers"
description: |-
  Indexer Torrent RSS resource.
  For more information refer to Indexer https://wiki.servarr.com/readarr/settings#indexers and Torrent RSS https://wiki.servarr.com/readarr/supported#torrentrssindexer.
---

# readarr_indexer_torrent_rss (Resource)

<!-- subcategory:Indexers -->Indexer Torrent RSS resource.
For more information refer to [Indexer](https://wiki.servarr.com/readarr/settings#indexers) and [Torrent RSS](https://wiki.servarr.com/readarr/supported#torrentrssindexer).

## Example Usage

```terraform
resource "readarr_indexer_torrent_rss" "example" {
  enable_automatic_search = true
  name                    = "Example"
  base_url                = "https://rss.io"
  allow_zero_size         = true
  minimum_seeders         = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `base_url` (String) Base URL.
- `name` (String) IndexerTorrentRss name.

### Optional

- `allow_zero_size` (Boolean) Allow zero size files.
- `author_seed_time` (Number) Author seed time.
- `cookie` (String) Cookie.
- `early_release_limit` (Number) Early release limit.
- `enable_rss` (Boolean) Enable RSS flag.
- `minimum_seeders` (Number) Minimum seeders.
- `priority` (Number) Priority.
- `seed_ratio` (Number) Seed ratio.
- `seed_time` (Number) Seed time.
- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) IndexerTorrentRss ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_indexer_torrent_rss.example 1
```