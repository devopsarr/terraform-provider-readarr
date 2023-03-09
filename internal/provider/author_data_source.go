package provider

import (
	"context"
	"fmt"

	"github.com/devopsarr/readarr-go/readarr"
	"github.com/devopsarr/terraform-provider-readarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const authorDataSourceName = "author"

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &AuthorDataSource{}

func NewAuthorDataSource() datasource.DataSource {
	return &AuthorDataSource{}
}

// AuthorDataSource defines the author implementation.
type AuthorDataSource struct {
	client *readarr.APIClient
}

func (d *AuthorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + authorDataSourceName
}

func (d *AuthorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "<!-- subcategory:Authors -->Single [Author](../resources/author).",
		Attributes: map[string]schema.Attribute{
			"monitored": schema.BoolAttribute{
				MarkdownDescription: "Monitored flag.",
				Computed:            true,
			},
			"quality_profile_id": schema.Int64Attribute{
				MarkdownDescription: "Quality profile ID.",
				Computed:            true,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Author ID.",
				Computed:            true,
			},
			"author_name": schema.StringAttribute{
				MarkdownDescription: "Author name.",
				Computed:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "Full author path.",
				Computed:            true,
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

func (d *AuthorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if client := helpers.DataSourceConfigure(ctx, req, resp); client != nil {
		d.client = client
	}
}

func (d *AuthorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var author *Author

	resp.Diagnostics.Append(req.Config.Get(ctx, &author)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get authors current value
	response, _, err := d.client.AuthorApi.ListAuthor(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, authorDataSourceName, err))

		return
	}

	value, err := findAuthor(author.ForeignAuthorID.ValueString(), response)
	if err != nil {
		resp.Diagnostics.AddError(helpers.DataSourceError, fmt.Sprintf("Unable to find %s, got error: %s", authorDataSourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+authorDataSourceName)
	author.write(ctx, value)
	// Map response body to resource schema attribute
	resp.Diagnostics.Append(resp.State.Set(ctx, &author)...)
}

func findAuthor(ID string, authors []*readarr.AuthorResource) (*readarr.AuthorResource, error) {
	for _, t := range authors {
		if t.GetForeignAuthorId() == ID {
			return t, nil
		}
	}

	return nil, helpers.ErrDataNotFoundError(authorDataSourceName, "TMDB ID", ID)
}
