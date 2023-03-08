---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_notifiarr Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Notifiarr resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Notifiarr https://wiki.servarr.com/readarr/supported#notifiarr.
---

# readarr_notification_notifiarr (Resource)

<!-- subcategory:Notifications -->Notification Notifiarr resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Notifiarr](https://wiki.servarr.com/readarr/supported#notifiarr).

## Example Usage

```terraform
resource "readarr_notification_notifiarr" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  api_key = "Token"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive) API Key.
- `include_health_warnings` (Boolean) Include health warnings.
- `name` (String) Notification name.
- `on_author_delete` (Boolean) On author deleted flag.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_grab` (Boolean) On grab flag.
- `on_health_issue` (Boolean) On health issue flag.
- `on_release_import` (Boolean) On release import flag.
- `on_upgrade` (Boolean) On upgrade flag.

### Optional

- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_notifiarr.example 1
```