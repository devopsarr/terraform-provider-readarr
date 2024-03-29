---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_synology_indexer Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Synology resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Synology https://wiki.servarr.com/readarr/supported#synologyindexer.
---

# readarr_notification_synology_indexer (Resource)

<!-- subcategory:Notifications -->Notification Synology resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Synology](https://wiki.servarr.com/readarr/supported#synologyindexer).

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Notification name.

### Optional

- `on_author_delete` (Boolean) On author deleted flag.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_book_retag` (Boolean) On book retag flag.
- `on_release_import` (Boolean) On release import flag.
- `on_rename` (Boolean) On rename flag.
- `on_upgrade` (Boolean) On upgrade flag.
- `tags` (Set of Number) List of associated tags.
- `update_library` (Boolean) Update library flag.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_synology_indexer.example 1
```
