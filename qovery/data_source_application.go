package qovery

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/qovery/terraform-provider-qovery/client"
)

// Ensure provider defined types fully satisfy terraform framework interfaces.
var _ datasource.DataSourceWithConfigure = &applicationDataSource{}

type applicationDataSource struct {
	client *client.Client
}

func newApplicationDataSource() datasource.DataSource {
	return &applicationDataSource{}
}

func (d applicationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

func (d *applicationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*qProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *qProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = provider.client
}

func (d applicationDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Use this data source to retrieve information about an existing application.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "Id of the application.",
				Type:        types.StringType,
				Required:    true,
			},
			"environment_id": {
				Description: "Id of the environment.",
				Type:        types.StringType,
				Computed:    true,
			},
			"name": {
				Description: "Name of the application.",
				Type:        types.StringType,
				Computed:    true,
			},
			"git_repository": {
				Description: "Git repository of the application.",
				Computed:    true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"url": {
						Description: "URL of the git repository.",
						Type:        types.StringType,
						Computed:    true,
					},
					"branch": {
						Description: "Branch of the git repository.",
						Type:        types.StringType,
						Computed:    true,
					},
					"root_path": {
						Description: "Root path of the application.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"build_mode": {
				Description: "Build Mode of the application.",
				Type:        types.StringType,
				Computed:    true,
			},
			"dockerfile_path": {
				Description: "Dockerfile Path of the application.",
				Type:        types.StringType,
				Computed:    true,
			},
			"buildpack_language": {
				Description: "Buildpack Language framework.",
				Type:        types.StringType,
				Computed:    true,
			},
			"cpu": {
				Description: "CPU of the application in millicores (m) [1000m = 1 CPU].",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"memory": {
				Description: "RAM of the application in MB [1024MB = 1GB].",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"min_running_instances": {
				Description: "Minimum number of instances running for the application.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"max_running_instances": {
				Description: "Maximum number of instances running for the application.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"auto_preview": {
				Description: "Specify if the environment preview option is activated or not for this application.",
				Type:        types.BoolType,
				Computed:    true,
			},
			"entrypoint": {
				Description: "Entrypoint of the application.",
				Type:        types.StringType,
				Computed:    true,
			},
			"arguments": {
				Description: "List of arguments of this container.",
				Computed:    true,
				Type: types.SetType{
					ElemType: types.StringType,
				},
			},
			"storage": {
				Description: "List of storages linked to this application.",
				Computed:    true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the storage.",
						Type:        types.StringType,
						Computed:    true,
					},
					"type": {
						Description: "Type of the storage for the application.",
						Type:        types.StringType,
						Computed:    true,
					},
					"size": {
						Description: "Size of the storage for the application in GB [1024MB = 1GB].",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"mount_point": {
						Description: "Mount point of the storage for the application.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"ports": {
				Description: "List of storages linked to this application.",
				Computed:    true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the port.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						Description: "Name of the port.",
						Type:        types.StringType,
						Computed:    true,
					},
					"internal_port": {
						Description: "Internal port of the application.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"external_port": {
						Description: "External port of the application.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"publicly_accessible": {
						Description: "Specify if the port is exposed to the world or not for this application.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"protocol": {
						Description: "Protocol used for the port of the application.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"built_in_environment_variables": {
				Description: "List of built-in environment variables linked to this application.",
				Computed:    true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
					"key": {
						Description: "Key of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						Description: "Value of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"environment_variables": {
				Description: "List of environment variables linked to this application.",
				Computed:    true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
					"key": {
						Description: "Key of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						Description: "Value of the environment variable.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"secrets": {
				Description: "List of secrets linked to this application.",
				Optional:    true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the secret.",
						Type:        types.StringType,
						Computed:    true,
					},
					"key": {
						Description: "Key of the secret.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						Description: "Value of the secret [NOTE: will always be empty].",
						Type:        types.StringType,
						Computed:    true,
						Sensitive:   true,
					},
				}),
			},
			"custom_domains": {
				Description: "List of custom domains linked to this application.",
				Computed:    true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Id of the custom domain.",
						Type:        types.StringType,
						Computed:    true,
					},
					"domain": {
						Description: "Your custom domain.",
						Type:        types.StringType,
						Computed:    true,
					},
					"validation_domain": {
						Description: "URL provided by Qovery. You must create a CNAME on your DNS provider using that URL.",
						Type:        types.StringType,
						Computed:    true,
					},
					"status": {
						Description: "Status of the custom domain.",
						Type:        types.StringType,
						Computed:    true,
					},
				}),
			},
			"external_host": {
				Description: "The application external FQDN host [NOTE: only if your application is using a publicly accessible port].",
				Type:        types.StringType,
				Computed:    true,
			},
			"internal_host": {
				Description: "The application internal host.",
				Type:        types.StringType,
				Computed:    true,
			},
			"state": {
				Description: "State of the application.",
				Type:        types.StringType,
				Computed:    true,
			},
		},
	}, nil
}

// Read qovery application data source
func (d applicationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var data Application
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get application from API
	application, apiErr := d.client.GetApplication(ctx, data.Id.Value)
	if apiErr != nil {
		resp.Diagnostics.AddError(apiErr.Summary(), apiErr.Detail())
		return
	}

	state := convertResponseToApplication(data, application)
	tflog.Trace(ctx, "read application", map[string]interface{}{"application_id": state.Id.Value})

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
