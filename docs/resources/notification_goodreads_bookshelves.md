---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_notification_goodreads_bookshelves Resource - terraform-provider-readarr"
subcategory: "Notifications"
description: |-
  Notification GoodreadsBookshelves resource.
  For more information refer to Notification https://wiki.servarr.com/readarr/settings#connect and GoodreadsBookshelves https://wiki.servarr.com/readarr/supported#goodreadsbookshelf.
---

# readarr_notification_goodreads_bookshelves (Resource)

<!-- subcategory:Notifications -->Notification GoodreadsBookshelves resource.
For more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [GoodreadsBookshelves](https://wiki.servarr.com/readarr/supported#goodreadsbookshelf).

## Example Usage

```terraform
resource "readarr_notification_gooodreads_bookshelves" "example" {
  on_upgrade                      = false
  on_book_delete                  = false
  on_book_file_delete             = false
  on_book_file_delete_for_upgrade = false
  on_author_delete                = false
  on_release_import               = false

  include_health_warnings = false
  name                    = "Example"

  access_token        = "exampleAccessToken"
  access_token_secret = "exampleAccessTokenSecret"
  user_id             = "163730408"
  username            = "Example User"
  add_ids             = ["currently-reading", "read", "to-read"]
  remove_ids          = ["test"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_token` (String, Sensitive) Access token.
- `access_token_secret` (String, Sensitive) Access token secret.
- `include_health_warnings` (Boolean) Include health warnings.
- `name` (String) Notification name.
- `on_author_delete` (Boolean) On author deleted flag.
- `on_book_delete` (Boolean) On book delete flag.
- `on_book_file_delete` (Boolean) On book file delete flag.
- `on_book_file_delete_for_upgrade` (Boolean) On book file delete for upgrade flag.
- `on_release_import` (Boolean) On release import flag.
- `on_upgrade` (Boolean) On upgrade flag.
- `username` (String) Username.

### Optional

- `add_ids` (Set of String) Add IDs.
- `remove_ids` (Set of String) Remove IDs.
- `request_token_secret` (String, Sensitive) Request token secret.
- `tags` (Set of Number) List of associated tags.
- `user_id` (String) User ID.

### Read-Only

- `id` (Number) Notification ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_notification_gooodreads_bookshelves.example 1
```