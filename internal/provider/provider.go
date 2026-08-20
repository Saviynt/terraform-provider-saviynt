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
	version        string
	client         *s.Client // your Go client SDK instance
	accessToken    string
	refreshToken   string
	saviyntVersion string
	tokenMutex     sync.RWMutex // Protects token refresh operations
}

// SaviyntProviderModel describes the provider data model.
type SaviyntProviderModel struct {
	ServerURL        types.String `tfsdk:"server_url"`
	Username         types.String `tfsdk:"username"`
	Password         types.String `tfsdk:"password"`
	SubjectToken     types.String `tfsdk:"subject_token"`
	SubjectTokenType types.String `tfsdk:"subject_token_type"`
	GrantType        types.String `tfsdk:"grant_type"`
	Scope            types.String `tfsdk:"scope"`
	AccessToken      types.String `tfsdk:"access_token"`
	RefreshToken     types.String `tfsdk:"refresh_token"`
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
				Optional:    true,
				Description: "Username for authentication. Used with password as the fallback auth method.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for user authentication. Used with username as the fallback auth method.",
			},
			"subject_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Entra ID (or other IdP) access token for OAuth2 Token Exchange authentication. Highest priority auth method. Requires scope to also be set.",
			},
			"subject_token_type": schema.StringAttribute{
				Optional:    true,
				Description: "Token type URI for the subject_token. Defaults to urn:ietf:params:oauth:token-type:access_token.",
			},
			"grant_type": schema.StringAttribute{
				Optional:    true,
				Description: "OAuth2 grant type for token exchange. Defaults to urn:ietf:params:oauth:grant-type:token-exchange.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Description: "Saviynt ExternalConnection name used as the scope in token exchange authentication.",
			},
			"access_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Saviynt Bearer access token. Used directly without re-authentication. Second priority auth method.",
			},
			"refresh_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Saviynt refresh token. Optional — used with access_token to enable automatic token refresh when the access token expires.",
			},
		},
	}
}

// Configure prepares a Saviynt API client for data sources and resources.
// Auth priority: (1) token exchange via subject_token+scope, (2) direct access_token, (3) username+password.
func (p *SaviyntProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SaviyntProviderModel

	configDiagnostics := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(configDiagnostics...)

	if resp.Diagnostics.HasError() {
		log.Println("Diagnostics contain errors, returning early.")
		return
	}

	if config.ServerURL.IsUnknown() || config.ServerURL.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"server_url must be set.",
		)
		return
	}

	ctx = context.Background()
	serverURL := "https://" + strings.TrimPrefix(strings.TrimPrefix(config.ServerURL.ValueString(), "https://"), "http://")

	isSet := func(v types.String) bool {
		return !v.IsNull() && !v.IsUnknown() && v.ValueString() != ""
	}

	var client *s.Client

	switch {
	// Priority 1: Token Exchange (e.g. Entra ID subject_token → Saviynt token)
	case isSet(config.SubjectToken):
		if !isSet(config.Scope) {
			resp.Diagnostics.AddError("Missing Configuration", "scope must be set when using subject_token authentication.")
			return
		}
		subjectTokenType := "urn:ietf:params:oauth:token-type:access_token"
		if isSet(config.SubjectTokenType) {
			subjectTokenType = config.SubjectTokenType.ValueString()
		}
		grantType := "urn:ietf:params:oauth:grant-type:token-exchange"
		if isSet(config.GrantType) {
			grantType = config.GrantType.ValueString()
		}
		var err error
		client, err = s.NewClientTokenExchange(ctx, serverURL, config.SubjectToken.ValueString(), subjectTokenType, grantType, config.Scope.ValueString())
		if err != nil {
			log.Printf("Token exchange failed: %v", err)
			resp.Diagnostics.AddError("Token Exchange Failed", "Could not exchange token with Saviynt: "+err.Error())
			return
		}
		log.Printf("[DEBUG] Authenticated via OAuth2 token exchange (scope: %s)", config.Scope.ValueString())

	// Priority 2: Direct Saviynt access token
	case isSet(config.AccessToken):
		client = s.NewClientAccessToken(ctx, serverURL, config.AccessToken.ValueString(), config.RefreshToken.ValueString())
		log.Printf("[DEBUG] Authenticated via direct Bearer access token")

	// Priority 3: Username + Password
	case isSet(config.Username) && isSet(config.Password):
		var err error
		client, err = s.NewClient(ctx, s.Credentials{
			ServerURL: serverURL,
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
		log.Printf("[DEBUG] Authenticated via username/password")

	case isSet(config.Username) && !isSet(config.Password):
		resp.Diagnostics.AddError("Missing Configuration", "username was provided without password. Both are required for credential authentication.")
		return

	case !isSet(config.Username) && isSet(config.Password):
		resp.Diagnostics.AddError("Missing Configuration", "password was provided without username. Both are required for credential authentication.")
		return

	default:
		resp.Diagnostics.AddError(
			"Missing Authentication Configuration",
			"One of the following auth methods must be provided: "+
				"(1) subject_token + scope for token exchange, "+
				"(2) access_token for direct token auth, "+
				"(3) username + password for credential auth.",
		)
		return
	}

	token := client.Token()
	if token == nil {
		log.Printf("Token error: Failed to fetch access token.")
		resp.Diagnostics.AddError("Token Error", "Failed to fetch access token.")
		return
	}

	saviyntVersion, _, saviyntVersionerr := client.Utility.GetEcmVersion(ctx).Execute()
	if saviyntVersionerr != nil {
		log.Printf("Version error: Failed to fetch Saviynt version: %v", saviyntVersionerr)
		resp.Diagnostics.AddWarning("Version Warning", "Failed to fetch Saviynt version, continuing without version info.")
	}

	// Store the token details in the provider struct.
	p.client = client
	p.accessToken = token.AccessToken
	p.refreshToken = token.RefreshToken

	// Store Saviynt version if available
	if saviyntVersion != nil && saviyntVersion.Version != nil {
		if versionStr, ok := saviyntVersion.Version.(string); ok {
			p.saviyntVersion = versionStr
		}
	}

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
		NewWorkdaySOAPConnectionsDataSource,
		NewSFTPConnectionsDataSource,
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
		NewWorkdaySOAPConnectionResource,
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
		NewImportTransportPackageResource,
		NewExportTransportPackageResource,
		NewFileUploadResource,
		NewSFTPConnectionResource,
		NewJobControlResource,
		NewFileTransferJobResource,
	}
}

func (p *SaviyntProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewFileCredentialsResource,
		NewEnvCredentialsResource,
	}
}
