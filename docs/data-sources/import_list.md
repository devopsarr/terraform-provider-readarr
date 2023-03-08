---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_import_list Data Source - terraform-provider-readarr"
subcategory: "Import Lists"
description: |-
  Single Import List ../resources/import_list.
---

# readarr_import_list (Data Source)

<!-- subcategory:Import Lists -->Single [Import List](../resources/import_list).

## Example Usage

```terraform
data "readarr_import_list" "example" {
  name = "Example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Import List name.

### Read-Only

- `access_token` (String, Sensitive) Access token.
- `access_token_secret` (String, Sensitive) Access token secret.
- `api_key` (String, Sensitive) API key.
- `base_url` (String) Base URL.
- `bookshelf_ids` (Set of String) Bookshelf IDs.
- `config_contract` (String) ImportList configuration template.
- `enable_automatic_add` (Boolean) Enable automatic add flag.
- `id` (Number) Import List ID.
- `implementation` (String) ImportList implementation name.
- `list_id` (Number) List ID.
- `list_order` (Number) List order.
- `list_type` (String) List type.
- `metadata_profile_id` (Number) Metadata profile ID.
- `profile_ids` (Set of Number) Profile IDs.
- `quality_profile_id` (Number) Quality profile ID.
- `request_token_secret` (String, Sensitive) Request token secret.
- `root_folder_path` (String) Root folder path.
- `series_id` (Number) Series ID.
- `should_monitor` (String) Should monitor.
- `should_search` (Boolean) Should search flag.
- `tag_ids` (Set of Number) Tag IDs.
- `tags` (Set of Number) List of associated tags.
- `user_id` (String) User ID.
- `username` (String) Username.

