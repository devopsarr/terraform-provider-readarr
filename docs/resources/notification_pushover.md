---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_pushover Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Pushover resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Pushover https://wiki.servarr.com/readarr/supported#pushover.
---

# readarr_notification_pushover (Resource)

<!-- subcategory:Notifications -->Notification Pushover resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Pushover](https://wiki.servarr.com/readarr/supported#pushover).

## Example Usage

```terraform
resource "readarr_notification_join" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false


  include_health_warnings = false
  name                    = "Example"

  api_key  = "Key"
  priority = 2
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive) API key.
- `name` (String) Notification name.

### Optional

- `devices` (Set of String) List of devices.
- `expire` (Number) Expire.
- `include_health_warnings` (Boolean) Include health warnings.
- `on_application_update` (Boolean) On application update flag.
- `on_author_delete` (Boolean) On author deleted flag.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_download_failure` (Boolean) On download failure flag.
- `on_grab` (Boolean) On grab flag.
- `on_health_issue` (Boolean) On health issue flag.
- `on_import_failure` (Boolean) On import failure flag.
- `on_release_import` (Boolean) On release import flag.
- `on_upgrade` (Boolean) On upgrade flag.
- `priority` (Number) Priority. `-2` Silent, `-1` Quiet, `0` Normal, `1` High, `2` Emergency, `8` High.
- `retry` (Number) Retry.
- `sound` (String) Sound.
- `tags` (Set of Number) List of associated tags.
- `user_key` (String, Sensitive) User key.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_pushover.example 1
```
