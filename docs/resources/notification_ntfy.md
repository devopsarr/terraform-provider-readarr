---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_ntfy Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Ntfy resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Ntfy https://wiki.servarr.com/readarr/supported#ntfy.
---

# readarr_notification_ntfy (Resource)

<!-- subcategory:Notifications -->Notification Ntfy resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Ntfy](https://wiki.servarr.com/readarr/supported#ntfy).

## Example Usage

```terraform
resource "readarr_notification_ntfy" "example" {
  on_grab                         = false
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = true
  on_health_issue                 = false
  on_author_delete                = false
  on_release_import               = true

  include_health_warnings = false
  name                    = "Example"

  priority   = 1
  server_url = "https://ntfy.sh"
  username   = "User"
  password   = "Pass"
  topics     = ["Topic1234", "Topic4321"]
  field_tags = ["warning", "skull"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Notification name.
- `topics` (Set of String) Topics.

### Optional

- `click_url` (String) Click URL.
- `field_tags` (Set of String) Tags and emojis.
- `include_health_warnings` (Boolean) Include health warnings.
- `on_application_update` (Boolean) On application update flag.
- `on_author_delete` (Boolean) On author deleted flag.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_grab` (Boolean) On grab flag.
- `on_health_issue` (Boolean) On health issue flag.
- `on_release_import` (Boolean) On release import flag.
- `on_upgrade` (Boolean) On upgrade flag.
- `password` (String, Sensitive) Password.
- `priority` (Number) Priority. `1` Min, `2` Low, `3` Default, `4` High, `5` Max.
- `server_url` (String) Server URL.
- `tags` (Set of Number) List of associated tags.
- `username` (String) Username.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_ntfy.example 1
```
