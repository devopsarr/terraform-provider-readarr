package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/devopsarr/readarr-go/readarr"
	"github.com/devopsarr/terraform-provider-readarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const authorResourceName = "author"

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &AuthorResource{}
	_ resource.ResourceWithImportState = &AuthorResource{}
)

func NewAuthorResource() resource.Resource {
	return &AuthorResource{}
}

// AuthorResource defines the author implementation.
type AuthorResource struct {
	client *readarr.APIClient
}

// Author describes the author data model.
type Author struct {
	Genres           types.Set    `tfsdk:"genres"`
	Tags             types.Set    `tfsdk:"tags"`
	AuthorName       types.String `tfsdk:"author_name"`
	ForeignAuthorID  types.String `tfsdk:"foreign_author_id"`
	Status           types.String `tfsdk:"status"`
	Path             types.String `tfsdk:"path"`
	Overview         types.String `tfsdk:"overview"`
	ID               types.Int64  `tfsdk:"id"`
	QualityProfileID types.Int64  `tfsdk:"quality_profile_id"`
	Monitored        types.Bool   `tfsdk:"monitored"`

	// TODO: future Implementation
	// Links          types.Set    `tfsdk:"links"`
	// SortName       types.String `tfsdk:"sortName"`
	// Ended          types.Bool   `tfsdk:"ended"`
	// RootFolderPath types.String `tfsdk:"root_folder_path"`
	// FolderName     types.String `tfsdk:"folderName"`
	// CleanName      types.String `tfsdk:"cleanName"`
	// Added          types.String `tfsdk:"added"`
	// Ratings        types.Object `tfsdk:"ratings"`
}

func (r *AuthorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + authorResourceName
}

func (r *AuthorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<!-- subcategory:Authors -->Author resource.\nFor more information refer to [Authors](https://wiki.servarr.com/readarr/library#authors) documentation.",
		Attributes: map[string]schema.Attribute{
			"monitored": schema.BoolAttribute{
				MarkdownDescription: "Monitored flag.",
				Required:            true,
			},
			"quality_profile_id": schema.Int64Attribute{
				MarkdownDescription: "Quality profile ID.",
				Required:            true,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Author ID.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"author_name": schema.StringAttribute{
				MarkdownDescription: "Author name.",
				Required:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "Full author path.",
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Author status.",
				Computed:            true,
			},
			"overview": schema.StringAttribute{
				MarkdownDescription: "Overview.",
				Computed:            true,
			},
			"foreign_author_id": schema.StringAttribute{
				MarkdownDescription: "Foreign author ID.",
				Required:            true,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.Int64Type,
			},
			"genres": schema.SetAttribute{
				MarkdownDescription: "List genres.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (r *AuthorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if client := helpers.ResourceConfigure(ctx, req, resp); client != nil {
		r.client = client
	}
}

func (r *AuthorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var author *Author

	resp.Diagnostics.Append(req.Plan.Get(ctx, &author)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new Author
	request := author.read(ctx)
	// TODO: can parametrize AddAuthorOptions
	options := readarr.NewAddAuthorOptions()
	options.SetMonitor(readarr.MONITORTYPES_ALL)
	options.SetSearchForMissingBooks(true)

	response, _, err := r.client.AuthorApi.CreateAuthor(ctx).AuthorResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Create, authorResourceName, err))

		return
	}

	tflog.Trace(ctx, "created author: "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct
	author.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &author)...)
}

func (r *AuthorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var author *Author

	resp.Diagnostics.Append(req.State.Get(ctx, &author)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get author current value
	response, _, err := r.client.AuthorApi.GetAuthorById(ctx, int32(author.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, authorResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+authorResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Map response body to resource schema attribute
	author.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &author)...)
}

func (r *AuthorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var author *Author

	resp.Diagnostics.Append(req.Plan.Get(ctx, &author)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update Author
	request := author.read(ctx)

	response, _, err := r.client.AuthorApi.UpdateAuthor(ctx, fmt.Sprint(request.GetId())).AuthorResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Update, authorResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+authorResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct
	author.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &author)...)
}

func (r *AuthorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var author *Author

	resp.Diagnostics.Append(req.State.Get(ctx, &author)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete author current value
	_, err := r.client.AuthorApi.DeleteAuthor(ctx, int32(author.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, authorResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+authorResourceName+": "+strconv.Itoa(int(author.ID.ValueInt64())))
	resp.State.RemoveResource(ctx)
}

func (r *AuthorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	helpers.ImportStatePassthroughIntID(ctx, path.Root("id"), req, resp)
	tflog.Trace(ctx, "imported "+authorResourceName+": "+req.ID)
}

func (m *Author) write(ctx context.Context, author *readarr.AuthorResource) {
	m.Monitored = types.BoolValue(author.GetMonitored())
	m.ID = types.Int64Value(int64(author.GetId()))
	m.AuthorName = types.StringValue(author.GetAuthorName())
	m.Path = types.StringValue(author.GetPath())
	m.QualityProfileID = types.Int64Value(int64(author.GetQualityProfileId()))
	m.ForeignAuthorID = types.StringValue(author.GetForeignAuthorId())
	m.Tags = types.SetValueMust(types.Int64Type, nil)
	tfsdk.ValueFrom(ctx, author.Tags, m.Tags.Type(ctx), &m.Tags)
	// Read only values
	m.Status = types.StringValue(string(author.GetStatus()))
	m.Overview = types.StringValue(author.GetOverview())
	m.Genres = types.SetValueMust(types.StringType, nil)
	tfsdk.ValueFrom(ctx, author.Genres, m.Genres.Type(ctx), &m.Genres)
}

func (m *Author) read(ctx context.Context) *readarr.AuthorResource {
	tags := make([]*int32, len(m.Tags.Elements()))
	tfsdk.ValueAs(ctx, m.Tags, &tags)

	author := readarr.NewAuthorResource()
	author.SetMonitored(m.Monitored.ValueBool())
	author.SetAuthorName(m.AuthorName.ValueString())
	author.SetPath(m.Path.ValueString())
	author.SetQualityProfileId(int32(m.QualityProfileID.ValueInt64()))
	author.SetForeignAuthorId(m.ForeignAuthorID.ValueString())
	author.SetId(int32(m.ID.ValueInt64()))
	author.SetTags(tags)
	// Fix unused but required profile
	author.SetMetadataProfileId(1)

	return author
}
