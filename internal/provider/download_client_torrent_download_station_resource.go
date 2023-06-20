package provider

import (
	"context"
	"strconv"

	"github.com/devopsarr/readarr-go/readarr"
	"github.com/devopsarr/terraform-provider-readarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	downloadClientTorrentDownloadStationResourceName   = "download_client_torrent_download_station"
	downloadClientTorrentDownloadStationImplementation = "TorrentDownloadStation"
	downloadClientTorrentDownloadStationConfigContract = "DownloadStationSettings"
	downloadClientTorrentDownloadStationProtocol       = "torrent"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &DownloadClientTorrentDownloadStationResource{}
	_ resource.ResourceWithImportState = &DownloadClientTorrentDownloadStationResource{}
)

func NewDownloadClientTorrentDownloadStationResource() resource.Resource {
	return &DownloadClientTorrentDownloadStationResource{}
}

// DownloadClientTorrentDownloadStationResource defines the download client implementation.
type DownloadClientTorrentDownloadStationResource struct {
	client *readarr.APIClient
}

// DownloadClientTorrentDownloadStation describes the download client data model.
type DownloadClientTorrentDownloadStation struct {
	Tags          types.Set    `tfsdk:"tags"`
	Name          types.String `tfsdk:"name"`
	Host          types.String `tfsdk:"host"`
	Username      types.String `tfsdk:"username"`
	Password      types.String `tfsdk:"password"`
	MusicCategory types.String `tfsdk:"book_category"`
	TVDirectory   types.String `tfsdk:"book_directory"`
	Priority      types.Int64  `tfsdk:"priority"`
	Port          types.Int64  `tfsdk:"port"`
	ID            types.Int64  `tfsdk:"id"`
	UseSsl        types.Bool   `tfsdk:"use_ssl"`
	Enable        types.Bool   `tfsdk:"enable"`
}

func (d DownloadClientTorrentDownloadStation) toDownloadClient() *DownloadClient {
	return &DownloadClient{
		Tags:           d.Tags,
		Name:           d.Name,
		Host:           d.Host,
		Username:       d.Username,
		Password:       d.Password,
		MusicCategory:  d.MusicCategory,
		TVDirectory:    d.TVDirectory,
		Priority:       d.Priority,
		Port:           d.Port,
		ID:             d.ID,
		UseSsl:         d.UseSsl,
		Enable:         d.Enable,
		Implementation: types.StringValue(downloadClientTorrentDownloadStationImplementation),
		ConfigContract: types.StringValue(downloadClientTorrentDownloadStationConfigContract),
		Protocol:       types.StringValue(downloadClientTorrentDownloadStationProtocol),
	}
}

func (d *DownloadClientTorrentDownloadStation) fromDownloadClient(client *DownloadClient) {
	d.Tags = client.Tags
	d.Name = client.Name
	d.Host = client.Host
	d.Username = client.Username
	d.Password = client.Password
	d.MusicCategory = client.MusicCategory
	d.TVDirectory = client.TVDirectory
	d.Priority = client.Priority
	d.Port = client.Port
	d.ID = client.ID
	d.UseSsl = client.UseSsl
	d.Enable = client.Enable
}

func (r *DownloadClientTorrentDownloadStationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + downloadClientTorrentDownloadStationResourceName
}

func (r *DownloadClientTorrentDownloadStationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<!-- subcategory:Download Clients -->Download Client TorrentDownloadStation resource.\nFor more information refer to [Download Client](https://wiki.servarr.com/readarr/settings#download-clients) and [TorrentDownloadStation](https://wiki.servarr.com/readarr/supported#torrentdownloadstation).",
		Attributes: map[string]schema.Attribute{
			"enable": schema.BoolAttribute{
				MarkdownDescription: "Enable flag.",
				Optional:            true,
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority.",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Download Client name.",
				Required:            true,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.Int64Type,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Download Client ID.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			// Field values
			"use_ssl": schema.BoolAttribute{
				MarkdownDescription: "Use SSL flag.",
				Optional:            true,
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port.",
				Optional:            true,
				Computed:            true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "host.",
				Optional:            true,
				Computed:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username.",
				Optional:            true,
				Computed:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password.",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"book_category": schema.StringAttribute{
				MarkdownDescription: "Book category.",
				Optional:            true,
				Computed:            true,
			},
			"book_directory": schema.StringAttribute{
				MarkdownDescription: "Book directory.",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *DownloadClientTorrentDownloadStationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if client := helpers.ResourceConfigure(ctx, req, resp); client != nil {
		r.client = client
	}
}

func (r *DownloadClientTorrentDownloadStationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var client *DownloadClientTorrentDownloadStation

	resp.Diagnostics.Append(req.Plan.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new DownloadClientTorrentDownloadStation
	request := client.read(ctx)

	response, _, err := r.client.DownloadClientApi.CreateDownloadClient(ctx).DownloadClientResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Create, downloadClientTorrentDownloadStationResourceName, err))

		return
	}

	tflog.Trace(ctx, "created "+downloadClientTorrentDownloadStationResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTorrentDownloadStationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var client DownloadClientTorrentDownloadStation

	resp.Diagnostics.Append(req.State.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get DownloadClientTorrentDownloadStation current value
	response, _, err := r.client.DownloadClientApi.GetDownloadClientById(ctx, int32(client.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, downloadClientTorrentDownloadStationResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+downloadClientTorrentDownloadStationResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Map response body to resource schema attribute
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTorrentDownloadStationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var client *DownloadClientTorrentDownloadStation

	resp.Diagnostics.Append(req.Plan.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update DownloadClientTorrentDownloadStation
	request := client.read(ctx)

	response, _, err := r.client.DownloadClientApi.UpdateDownloadClient(ctx, strconv.Itoa(int(request.GetId()))).DownloadClientResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Update, downloadClientTorrentDownloadStationResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+downloadClientTorrentDownloadStationResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTorrentDownloadStationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var client *DownloadClientTorrentDownloadStation

	resp.Diagnostics.Append(req.State.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete DownloadClientTorrentDownloadStation current value
	_, err := r.client.DownloadClientApi.DeleteDownloadClient(ctx, int32(client.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Delete, downloadClientTorrentDownloadStationResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+downloadClientTorrentDownloadStationResourceName+": "+strconv.Itoa(int(client.ID.ValueInt64())))
	resp.State.RemoveResource(ctx)
}

func (r *DownloadClientTorrentDownloadStationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	helpers.ImportStatePassthroughIntID(ctx, path.Root("id"), req, resp)
	tflog.Trace(ctx, "imported "+downloadClientTorrentDownloadStationResourceName+": "+req.ID)
}

func (d *DownloadClientTorrentDownloadStation) write(ctx context.Context, downloadClient *readarr.DownloadClientResource) {
	genericDownloadClient := d.toDownloadClient()
	genericDownloadClient.write(ctx, downloadClient)
	d.fromDownloadClient(genericDownloadClient)
}

func (d *DownloadClientTorrentDownloadStation) read(ctx context.Context) *readarr.DownloadClientResource {
	return d.toDownloadClient().read(ctx)
}
