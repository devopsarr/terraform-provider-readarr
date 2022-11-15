package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/devopsarr/terraform-provider-sonarr/tools"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golift.io/starr/readarr"
)

const (
	notificationCustomScriptResourceName   = "notification_custom_script"
	NotificationCustomScriptImplementation = "CustomScript"
	NotificationCustomScriptConfigContrat  = "CustomScriptSettings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NotificationCustomScriptResource{}
var _ resource.ResourceWithImportState = &NotificationCustomScriptResource{}

func NewNotificationCustomScriptResource() resource.Resource {
	return &NotificationCustomScriptResource{}
}

// NotificationCustomScriptResource defines the notification implementation.
type NotificationCustomScriptResource struct {
	client *readarr.Readarr
}

// NotificationCustomScript describes the notification data model.
type NotificationCustomScript struct {
	Tags                       types.Set    `tfsdk:"tags"`
	Arguments                  types.String `tfsdk:"arguments"`
	Path                       types.String `tfsdk:"path"`
	Name                       types.String `tfsdk:"name"`
	ID                         types.Int64  `tfsdk:"id"`
	OnGrab                     types.Bool   `tfsdk:"on_grab"`
	IncludeHealthWarnings      types.Bool   `tfsdk:"include_health_warnings"`
	OnApplicationUpdate        types.Bool   `tfsdk:"on_application_update"`
	OnHealthIssue              types.Bool   `tfsdk:"on_health_issue"`
	OnRename                   types.Bool   `tfsdk:"on_rename"`
	OnUpgrade                  types.Bool   `tfsdk:"on_upgrade"`
	OnReleaseImport            types.Bool   `tfsdk:"on_release_import"`
	OnAuthorDelete             types.Bool   `tfsdk:"on_author_delete"`
	OnBookDelete               types.Bool   `tfsdk:"on_book_delete"`
	OnBookFileDelete           types.Bool   `tfsdk:"on_book_file_delete"`
	OnBookFileDeleteForUpgrade types.Bool   `tfsdk:"on_book_file_delete_for_upgrade"`
	OnDownloadFailure          types.Bool   `tfsdk:"on_download_failure"`
	OnImportFailure            types.Bool   `tfsdk:"on_import_failure"`
	OnBookRetag                types.Bool   `tfsdk:"on_book_retag"`
}

func (n NotificationCustomScript) toNotification() *Notification {
	return &Notification{
		Tags:                       n.Tags,
		Path:                       n.Path,
		Arguments:                  n.Arguments,
		Name:                       n.Name,
		ID:                         n.ID,
		OnGrab:                     n.OnGrab,
		OnReleaseImport:            n.OnReleaseImport,
		OnAuthorDelete:             n.OnAuthorDelete,
		IncludeHealthWarnings:      n.IncludeHealthWarnings,
		OnApplicationUpdate:        n.OnApplicationUpdate,
		OnHealthIssue:              n.OnHealthIssue,
		OnBookDelete:               n.OnBookDelete,
		OnBookFileDelete:           n.OnBookFileDelete,
		OnRename:                   n.OnRename,
		OnUpgrade:                  n.OnUpgrade,
		OnBookFileDeleteForUpgrade: n.OnBookFileDeleteForUpgrade,
		OnDownloadFailure:          n.OnDownloadFailure,
		OnImportFailure:            n.OnImportFailure,
		OnBookRetag:                n.OnBookRetag,
	}
}

func (n *NotificationCustomScript) fromNotification(notification *Notification) {
	n.Tags = notification.Tags
	n.Path = notification.Path
	n.Arguments = notification.Arguments
	n.Name = notification.Name
	n.ID = notification.ID
	n.OnGrab = notification.OnGrab
	n.OnBookFileDeleteForUpgrade = notification.OnBookFileDeleteForUpgrade
	n.OnBookFileDelete = notification.OnBookFileDelete
	n.IncludeHealthWarnings = notification.IncludeHealthWarnings
	n.OnApplicationUpdate = notification.OnApplicationUpdate
	n.OnHealthIssue = notification.OnHealthIssue
	n.OnAuthorDelete = notification.OnAuthorDelete
	n.OnBookDelete = notification.OnBookDelete
	n.OnRename = notification.OnRename
	n.OnUpgrade = notification.OnUpgrade
	n.OnDownloadFailure = notification.OnDownloadFailure
	n.OnBookRetag = notification.OnBookRetag
	n.OnImportFailure = notification.OnImportFailure
	n.OnReleaseImport = notification.OnReleaseImport
}

func (r *NotificationCustomScriptResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + notificationCustomScriptResourceName
}

func (r *NotificationCustomScriptResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "<!-- subcategory:Notifications -->Notification Custom Script resource.\nFor more information refer to [Notification](https://wiki.servarr.com/readarr/settings#connect) and [Custom Script](https://wiki.servarr.com/readarr/supported#customscript).",
		Attributes: map[string]tfsdk.Attribute{
			"on_grab": {
				MarkdownDescription: "On grab flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_download_failure": {
				MarkdownDescription: "On download failure flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_upgrade": {
				MarkdownDescription: "On upgrade flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_rename": {
				MarkdownDescription: "On rename flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_author_delete": {
				MarkdownDescription: "On author deleted flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_book_delete": {
				MarkdownDescription: "On book delete flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_book_file_delete": {
				MarkdownDescription: "On book file delete flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_book_file_delete_for_upgrade": {
				MarkdownDescription: "On book file delete for upgrade flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_health_issue": {
				MarkdownDescription: "On health issue flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_import_failure": {
				MarkdownDescription: "On import failure flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_book_retag": {
				MarkdownDescription: "On book retag flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_release_import": {
				MarkdownDescription: "On release import flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"on_application_update": {
				MarkdownDescription: "On application update flag.",
				Required:            true,
				Type:                types.BoolType,
			},
			"include_health_warnings": {
				MarkdownDescription: "Include health warnings.",
				Required:            true,
				Type:                types.BoolType,
			},
			"name": {
				MarkdownDescription: "NotificationCustomScript name.",
				Required:            true,
				Type:                types.StringType,
			},
			"tags": {
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				Type: types.SetType{
					ElemType: types.Int64Type,
				},
			},
			"id": {
				MarkdownDescription: "Notification ID.",
				Computed:            true,
				Type:                types.Int64Type,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			// Field values
			"arguments": {
				MarkdownDescription: "Arguments.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"path": {
				MarkdownDescription: "Path.",
				Required:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (r *NotificationCustomScriptResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*readarr.Readarr)
	if !ok {
		resp.Diagnostics.AddError(
			tools.UnexpectedResourceConfigureType,
			fmt.Sprintf("Expected *readarr.Readarr, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NotificationCustomScriptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var notification *NotificationCustomScript

	resp.Diagnostics.Append(req.Plan.Get(ctx, &notification)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new NotificationCustomScript
	request := notification.read(ctx)

	response, err := r.client.AddNotificationContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to create %s, got error: %s", notificationCustomScriptResourceName, err))

		return
	}

	tflog.Trace(ctx, "created "+notificationCustomScriptResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	notification.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &notification)...)
}

func (r *NotificationCustomScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var notification *NotificationCustomScript

	resp.Diagnostics.Append(req.State.Get(ctx, &notification)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get NotificationCustomScript current value
	response, err := r.client.GetNotificationContext(ctx, int(notification.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", notificationCustomScriptResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+notificationCustomScriptResourceName+": "+strconv.Itoa(int(response.ID)))
	// Map response body to resource schema attribute
	notification.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &notification)...)
}

func (r *NotificationCustomScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var notification *NotificationCustomScript

	resp.Diagnostics.Append(req.Plan.Get(ctx, &notification)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update NotificationCustomScript
	request := notification.read(ctx)

	response, err := r.client.UpdateNotificationContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to update %s, got error: %s", notificationCustomScriptResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+notificationCustomScriptResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	notification.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &notification)...)
}

func (r *NotificationCustomScriptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var notification *NotificationCustomScript

	resp.Diagnostics.Append(req.State.Get(ctx, &notification)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete NotificationCustomScript current value
	err := r.client.DeleteNotificationContext(ctx, notification.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", notificationCustomScriptResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+notificationCustomScriptResourceName+": "+strconv.Itoa(int(notification.ID.ValueInt64())))
	resp.State.RemoveResource(ctx)
}

func (r *NotificationCustomScriptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			tools.UnexpectedImportIdentifier,
			fmt.Sprintf("Expected import identifier with format: ID. Got: %q", req.ID),
		)

		return
	}

	tflog.Trace(ctx, "imported "+notificationCustomScriptResourceName+": "+strconv.Itoa(id))
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (n *NotificationCustomScript) write(ctx context.Context, notification *readarr.NotificationOutput) {
	genericNotification := Notification{
		OnGrab:                     types.BoolValue(notification.OnGrab),
		OnReleaseImport:            types.BoolValue(notification.OnReleaseImport),
		OnUpgrade:                  types.BoolValue(notification.OnUpgrade),
		OnRename:                   types.BoolValue(notification.OnRename),
		OnAuthorDelete:             types.BoolValue(notification.OnAuthorDelete),
		OnBookDelete:               types.BoolValue(notification.OnBookDelete),
		OnBookFileDelete:           types.BoolValue(notification.OnBookFileDelete),
		OnBookFileDeleteForUpgrade: types.BoolValue(notification.OnBookFileDeleteForUpgrade),
		OnHealthIssue:              types.BoolValue(notification.OnHealthIssue),
		OnDownloadFailure:          types.BoolValue(notification.OnDownloadFailure),
		OnImportFailure:            types.BoolValue(notification.OnImportFailure),
		OnBookRetag:                types.BoolValue(notification.OnBookRetag),
		OnApplicationUpdate:        types.BoolValue(notification.OnApplicationUpdate),
		IncludeHealthWarnings:      types.BoolValue(notification.IncludeHealthWarnings),
		ID:                         types.Int64Value(notification.ID),
		Name:                       types.StringValue(notification.Name),
		Tags:                       types.SetValueMust(types.Int64Type, nil),
	}
	tfsdk.ValueFrom(ctx, notification.Tags, genericNotification.Tags.Type(ctx), &genericNotification.Tags)
	genericNotification.writeFields(ctx, notification.Fields)
	n.fromNotification(&genericNotification)
}

func (n *NotificationCustomScript) read(ctx context.Context) *readarr.NotificationInput {
	var tags []int

	tfsdk.ValueAs(ctx, n.Tags, &tags)

	return &readarr.NotificationInput{
		OnGrab:                     n.OnGrab.ValueBool(),
		OnReleaseImport:            n.OnReleaseImport.ValueBool(),
		OnUpgrade:                  n.OnUpgrade.ValueBool(),
		OnRename:                   n.OnRename.ValueBool(),
		OnAuthorDelete:             n.OnAuthorDelete.ValueBool(),
		OnBookDelete:               n.OnBookDelete.ValueBool(),
		OnBookFileDelete:           n.OnBookFileDelete.ValueBool(),
		OnBookFileDeleteForUpgrade: n.OnBookFileDeleteForUpgrade.ValueBool(),
		OnHealthIssue:              n.OnHealthIssue.ValueBool(),
		OnDownloadFailure:          n.OnDownloadFailure.ValueBool(),
		OnImportFailure:            n.OnImportFailure.ValueBool(),
		OnBookRetag:                n.OnBookRetag.ValueBool(),
		OnApplicationUpdate:        n.OnApplicationUpdate.ValueBool(),
		IncludeHealthWarnings:      n.IncludeHealthWarnings.ValueBool(),
		ID:                         n.ID.ValueInt64(),
		Name:                       n.Name.ValueString(),
		Implementation:             NotificationCustomScriptImplementation,
		ConfigContract:             NotificationCustomScriptConfigContrat,
		Tags:                       tags,
		Fields:                     n.toNotification().readFields(ctx),
	}
}
