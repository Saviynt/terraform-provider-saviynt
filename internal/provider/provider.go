// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// provider.go defines the Saviynt Terraform provider using the Terraform Plugin Framework.
// It handles authentication, schema configuration, and registration of both standard
// and ephemeral resources and data sources for managing entities in the Saviynt
// Security Manager.

package provider

import (
	"context"
	"log"
	"strings"
	"sync"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
)

// Ensure SaviyntProvider satisfies Terraform's provider interfaces.

var _ provider.Provider = &SaviyntProvider{}
var _ provider.ProviderWithEphemeralResources = &SaviyntProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SaviyntProvider{
			version: version,
		}
	}
}

// SaviyntProvider defines the provider implementation.
type SaviyntProvider struct {
	version      string
	client       *s.Client // your Go client SDK instance
	accessToken  string
	refreshToken string
	tokenMutex   sync.RWMutex // Protects token refresh operations
}

// SaviyntProviderModel describes the provider data model.
type SaviyntProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
}

func (p *SaviyntProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "saviynt"
	resp.Version = p.version
}

func (p *SaviyntProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ProviderDescription,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				Required:    true,
				Description: "URL of Saviynt server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username for authentication.",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "Password for user authentication.",
			},
		},
	}
}

// Configure prepares a Saviynt API client for data sources and resources.
func (p *SaviyntProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SaviyntProviderModel

	configDiagnostics := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(configDiagnostics...)

	if resp.Diagnostics.HasError() {
		log.Println("Diagnostics contain errors, returning early.")
		return
	}

	if config.ServerURL.IsUnknown() || config.ServerURL.IsNull() ||
		config.Username.IsUnknown() || config.Username.IsNull() ||
		config.Password.IsUnknown() || config.Password.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"server_url, username, and password must be set.",
		)
		return
	}
	ctx = context.Background()

	serverURL := config.ServerURL.ValueString()

	serverURL = strings.TrimPrefix(strings.TrimPrefix(serverURL, "https://"), "http://")

	client, err := s.NewClient(ctx, s.Credentials{
		ServerURL: "https://" + serverURL,
		Username:  config.Username.ValueString(),
		Password:  config.Password.ValueString(),
	})
	if err != nil {
		log.Printf("Failed to create Saviynt client: %v", err)
		resp.Diagnostics.AddError(
			"Failed to create Saviynt client",
			"Could not initialize Saviynt API client: "+err.Error(),
		)
		return
	}

	token := client.Token()
	if token == nil {
		log.Printf("Token error: Failed to fetch access token.")
		resp.Diagnostics.AddError("Token Error", "Failed to fetch access token.")
		return
	}

	// Store the token details in the provider struct.
	p.client = client
	p.accessToken = token.AccessToken
	p.refreshToken = token.RefreshToken
	//Storing in Resource and Datasource
	resp.ResourceData = p
	resp.DataSourceData = p

}

// DataSources defines the data sources implemented in the provider.
func (p *SaviyntProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSecuritySystemsDataSource,
		NewEndpointsDataSource,
		NewConnectionsDataSource,
		NewADConnectionsDataSource,
		NewRESTConnectionsDataSource,
		NewADSIConnectionsDataSource,
		NewDBConnectionsDataSource,
		NewWorkdayConnectionsDataSource,
		NewSalesforceConnectionsDataSource,
		NewEntraIDConnectionsDataSource,
		NewSAPConnectionsDataSource,
		NewUnixConnectionsDataSource,
		NewGithubRestConnectionsDataSource,
		NewDynamicAttributeDataSource,
		NewOktaConnectionsDataSource,
		NewRolesDataSource,
		NewEntitlementTypeDataSource,
		NewEntitlementDataSource,
		NewPrivilegeDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *SaviyntProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSecuritySystemResource,
		NewADConnectionResource,
		NewRestConnectionResource,
		NewDBConnectionResource,
		NewADSIConnectionResource,
		NewWorkdayConnectionResource,
		NewEntraIdConnectionResource,
		NewSalesforceConnectionResource,
		NewSapConnectionResource,
		NewUnixConnectionResource,
		NewGithubRestConnectionResource,
		NewEndpointResource,
		NewRolesResource,
		NewDynamicAttributeResource,
		NewOktaConnectionResource,
		NewEntitlementTypeResource,
		NewEntitlementResource,
		NewPrivilegeResource,
		NewApplicationDataImportJobResource,
		NewAccountsImportFullJobResource,
		NewWSRetryJobResource,
		NewWSRetryBlockingJobResource,
		NewEcmJobResource,
		NewEcmSapUserJobResource,
		NewUserImportJobResource,
		NewAccountsImportIncrementalJobResource,
		NewSchemaAccountJobResource,
		NewSchemaRoleJobResource,
		NewSchemaUserJobResource,
	}
}
func (p *SaviyntProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewFileCredentialsResource,
		NewEnvCredentialsResource,
	}
}
