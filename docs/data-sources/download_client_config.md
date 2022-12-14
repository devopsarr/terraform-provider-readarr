---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "readarr_download_client_config Data Source - terraform-provider-readarr"
subcategory: "Download Clients"
description: |-
  Download Client Config ../resources/download_client_config.
---

# readarr_download_client_config (Data Source)

<!-- subcategory:Download Clients -->[Download Client Config](../resources/download_client_config).

## Example Usage

```terraform
data "readarr_download_client_config" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `auto_redownload_failed` (Boolean) Auto Redownload Failed flag.
- `download_client_working_folders` (String) Download Client Working Folders.
- `enable_completed_download_handling` (Boolean) Enable Completed Download Handling flag.
- `id` (Number) Download Client Config ID.
- `remove_completed_downloads` (Boolean) Remove completed downloads flag.
- `remove_failed_downloads` (Boolean) Remove failed downloads flag.

