---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_sendgrid Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification Sendgrid resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and Sendgrid https://wiki.servarr.com/readarr/supported#sendgrid.
---

# readarr_notification_sendgrid (Resource)

<!-- subcategory:Notifications -->Notification Sendgrid resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Sendgrid](https://wiki.servarr.com/readarr/supported#sendgrid).

## Example Usage

```terraform
resource "readarr_notification_sendgrid" "example" {
  on_grab                         = false
  on_download_failure             = false
  on_upgrade                      = false
  on_import_failure               = false
  on_book_delete                  = false
  on_book_file_delete             = true
  on_book_file_delete_for_upgrade = false
  on_health_issue                 = false
  on_author_delete                = true
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  api_key    = "APIkey"
  from       = "from_sendgrid@example.com"
  recipients = ["user1@example.com", "user2@example.com"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `from` (String) From.
- `name` (String) Notification name.
- `recipients` (Set of String) Recipients.

### Optional

- `api_key` (String, Sensitive) API key.
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
- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_sendgrid.example 1
```
