---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_kavita Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Kavita resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Kavita https://wiki.servarr.com/readarr/supported#kavita.
---

# readarr_notification_kavita (Resource)

<!-- subcategory:Notifications -->Notification Kavita resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Kavita](https://wiki.servarr.com/readarr/supported#kavita).

## Example Usage

```terraform
resource "readarr_notification_kavita" "example" {
  on_book_retag                   = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_release_import               = false

  name = "Example"

  api_key = "APIKey"
  host    = "kavita.local"
  port    = 4040
  notify  = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive) API key.
- `host` (String) Host.
- `name` (String) Notification name.
- `notify` (Boolean) Update library.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_book_retag` (Boolean) On book retag flag.
- `on_release_import` (Boolean) On release import flag.
- `on_upgrade` (Boolean) On upgrade flag.
- `port` (Number) Port.

### Optional

- `tags` (Set of Number) List of associated tags.
- `use_ssl` (Boolean) Use SSL flag.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_kavita.example 1
```