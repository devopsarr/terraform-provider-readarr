---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_import_list_readarr Resource - terraform-provider-readarr"
subcategory: "Import Lists"
description: |-
  Import List Readarr resource.
  For more information refer to Import List https://wiki.servarr.com/readarr/settings#import-lists and Readarr https://wiki.servarr.com/readarr/supported#readarrimport.
---

# readarr_import_list_readarr (Resource)

<!-- subcategory:Import Lists -->Import List Readarr resource.
For more information refer to [Import List](https://wiki.servarr.com/readarr/settings#import-lists) and [Readarr](https://wiki.servarr.com/readarr/supported#readarrimport).

## Example Usage

```terraform
resource "readarr_import_list_readarr" "example" {
  enable_automatic_add = false
  should_monitor       = "specificBook"
  should_search        = false
  root_folder_path     = "/books"
  quality_profile_id   = 1
  metadata_profile_id  = 1
  name                 = "Example"
  base_url             = "http://127.0.0.1:8787"
  api_key              = "APIKey"
  tags                 = [1, 2, 3]
  profile_ids          = [1, 2]
  tag_ids              = [1, 2, 3]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive) API key.
- `base_url` (String) Base URL.
- `name` (String) Import List name.

### Optional

- `enable_automatic_add` (Boolean) Enable automatic add flag.
- `list_order` (Number) List order.
- `metadata_profile_id` (Number) Metadata profile ID.
- `monitor_new_items` (String) Monitor new items.
- `profile_ids` (Set of Number) Profile IDs.
- `quality_profile_id` (Number) Quality profile ID.
- `root_folder_path` (String) Root folder path.
- `should_monitor` (String) Should monitor.
- `should_monitor_existing` (Boolean) Should monitor existing flag.
- `should_search` (Boolean) Should search flag.
- `tag_ids` (Set of Number) Tag IDs.
- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) Import List ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import readarr_import_list_readarr.example 1
```
