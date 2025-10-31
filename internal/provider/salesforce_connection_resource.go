// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_salesforce_connection_resource manages Salesforce connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Salesforce connector using the supplied configuration.
//   - Read: fetches the current connector state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing connector.
//   - Import: brings an existing connector under Terraform management by its name.
package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SalesforceConnectionResource{}
var _ resource.ResourceWithImportState = &SalesforceConnectionResource{}

// Initialize error codes for Salesforce Connection operations
var salesforceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSalesforce)

type SalesforceConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                     types.String `tfsdk:"id"`
	ClientId               types.String `tfsdk:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	ClientSecretWO         types.String `tfsdk:"client_secret_wo"`
	RefreshToken           types.String `tfsdk:"refresh_token"`
	RefreshTokenWO         types.String `tfsdk:"refresh_token_wo"`
	RedirectUri            types.String `tfsdk:"redirect_uri"`
	InstanceUrl            types.String `tfsdk:"instance_url"`
	ObjectToBeImported     types.String `tfsdk:"object_to_be_imported"`
	FeatureLicenseJson     types.String `tfsdk:"feature_license_json"`
	CustomCreateaccountUrl types.String `tfsdk:"custom_createaccount_url"`
	Createaccountjson      types.String `tfsdk:"createaccountjson"`
	AccountFilterQuery     types.String `tfsdk:"account_filter_query"`
	AccountFieldQuery      types.String `tfsdk:"account_field_query"`
	FieldMappingJson       types.String `tfsdk:"field_mapping_json"`
	Modifyaccountjson      types.String `tfsdk:"modifyaccountjson"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	Customconfigjson       types.String `tfsdk:"customconfigjson"`
	PamConfig              types.String `tfsdk:"pam_config"`
}

type SalesforceConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewSalesforceConnectionResource() resource.Resource {
	return &SalesforceConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewSalesforceConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &SalesforceConnectionResource{
		connectionFactory: factory,
	}
}

func (r *SalesforceConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_salesforce_connection_resource"
}

func SalesforceConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"client_id": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The OAuth client ID for Salesforce.",
		},
		"client_secret": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "The OAuth client secret for Salesforce. Either this field or the client_secret_wo field must be provided to configure the client_secret attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("client_secret_wo")),
			},
		},
		"client_secret_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "The OAuth client secret for Salesforce (write-only). Either this field or the client_secret field must be provided to configure the client_secret attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("client_secret")),
			},
		},
		"refresh_token": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "The OAuth refresh token used to get access tokens from Salesforce. Either this field or the refresh_token_wo field must be provided to configure the refresh_token attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("refresh_token_wo")),
			},
		},
		"refresh_token_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "The OAuth refresh token used to get access tokens from Salesforce (write-only). Either this field or the refresh_token field must be provided to configure the refresh_token attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("refresh_token")),
			},
		},
		"redirect_uri": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The redirect URI used in OAuth flows. Example: https://@INSTANCE_NAME@.salesforce.com/services/oauth2/success",
		},
		"instance_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Salesforce instance base URL. Example: https://@INSTANCE_NAME@.salesforce.com",
		},
		"object_to_be_imported": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: `Comma-separated list of Salesforce objects to import. Example: "Profile,Role,Group,PermissionSet"`,
		},
		"feature_license_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON mapping of feature licenses to permission fields in Salesforce.",
		},
		"custom_createaccount_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom URL used when creating a Salesforce account.",
		},
		"createaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template used for account creation in Salesforce.",
		},
		"account_filter_query": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Query used to filter Salesforce accounts.",
		},
		"account_field_query": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Fields to retrieve for Salesforce accounts. Example: Id, Username, LastName, FirstName, etc.",
		},
		"field_mapping_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON mapping of local fields to Salesforce fields with data types.",
		},
		"modifyaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template used for modifying Salesforce accounts.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to define active/inactive thresholds and lock statuses.",
		},
		"customconfigjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration options for Salesforce connector.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Privileged Access Management (PAM) configuration in JSON format.",
		},
	}
}

func (r *SalesforceConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SalesforceConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), SalesforceConnectorResourceSchema()),
	}
}

func (r *SalesforceConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Salesforce connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := salesforceErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "Salesforce connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *SalesforceConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *SalesforceConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *SalesforceConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *SalesforceConnectionResource) CreateSalesforceConnection(ctx context.Context, plan *SalesforceConnectorResourceModel, config *SalesforceConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Salesforce connection creation")

	// Check if Salesforce connection already exists (idempotency check) with retry logic
	tflog.Debug(logCtx, "Checking if connection already exists")
	var existingResource *openapi.GetConnectionDetailsResponse
	var finalHttpResp *http.Response

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_connection_details_idempotency", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		finalHttpResp = httpResp // Update on every call including retries
		return err
	})

	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		errorCode := salesforceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.SalesforceConnectionResponse != nil &&
		existingResource.SalesforceConnectionResponse.Errorcode != nil &&
		*existingResource.SalesforceConnectionResponse.Errorcode == 0 {

		errorCode := salesforceErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "create", connectionName, nil)
	}

	// Build Salesforce connection create request
	tflog.Debug(ctx, "Building Salesforce connection create request")

	salesforceConn := r.BuildSalesforceConnector(plan, config)
	salesforceConnRequest := openapi.CreateOrUpdateRequest{
		SalesforceConnector: &salesforceConn,
	}

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_salesforce_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, salesforceConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := salesforceErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create Salesforce connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := salesforceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Salesforce connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Salesforce connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SalesforceConnectionResource) BuildSalesforceConnector(plan *SalesforceConnectorResourceModel, config *SalesforceConnectorResourceModel) openapi.SalesforceConnector {
	var clientSecret string
	if !config.ClientSecret.IsNull() && !config.ClientSecret.IsUnknown() {
		clientSecret = config.ClientSecret.ValueString()
	} else if !config.ClientSecretWO.IsNull() && !config.ClientSecretWO.IsUnknown() {
		clientSecret = config.ClientSecretWO.ValueString()
	}

	var refreshToken string
	if !config.RefreshToken.IsNull() && !config.RefreshToken.IsUnknown() {
		refreshToken = config.RefreshToken.ValueString()
	} else if !config.RefreshTokenWO.IsNull() && !config.RefreshTokenWO.IsUnknown() {
		refreshToken = config.RefreshTokenWO.ValueString()
	}

	salesforceConn := openapi.SalesforceConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:  "SalesForce",
			ConnectionName:  plan.ConnectionName.ValueString(),
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		CLIENT_ID:                util.StringPointerOrEmpty(plan.ClientId),
		CLIENT_SECRET:            util.StringPointerOrEmpty(types.StringValue(clientSecret)),
		REFRESH_TOKEN:            util.StringPointerOrEmpty(types.StringValue(refreshToken)),
		REDIRECT_URI:             util.StringPointerOrEmpty(plan.RedirectUri),
		INSTANCE_URL:             util.StringPointerOrEmpty(plan.InstanceUrl),
		OBJECT_TO_BE_IMPORTED:    util.StringPointerOrEmpty(plan.ObjectToBeImported),
		FEATURE_LICENSE_JSON:     util.StringPointerOrEmpty(plan.FeatureLicenseJson),
		CUSTOM_CREATEACCOUNT_URL: util.StringPointerOrEmpty(plan.CustomCreateaccountUrl),
		CREATEACCOUNTJSON:        util.StringPointerOrEmpty(plan.Createaccountjson),
		ACCOUNT_FILTER_QUERY:     util.StringPointerOrEmpty(plan.AccountFilterQuery),
		ACCOUNT_FIELD_QUERY:      util.StringPointerOrEmpty(plan.AccountFieldQuery),
		FIELD_MAPPING_JSON:       util.StringPointerOrEmpty(plan.FieldMappingJson),
		MODIFYACCOUNTJSON:        util.StringPointerOrEmpty(plan.Modifyaccountjson),
		STATUS_THRESHOLD_CONFIG:  util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CUSTOMCONFIGJSON:         util.StringPointerOrEmpty(plan.Customconfigjson),
		PAM_CONFIG:               util.StringPointerOrEmpty(plan.PamConfig),
	}

	if plan.VaultConnection.ValueString() != "" {
		salesforceConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		salesforceConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		salesforceConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return salesforceConn
}

func (r *SalesforceConnectionResource) UpdateModelFromCreateResponse(plan *SalesforceConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))

	// Update all optional fields to maintain state
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ClientId = util.SafeStringDatasource(plan.ClientId.ValueStringPointer())
	plan.RedirectUri = util.SafeStringDatasource(plan.RedirectUri.ValueStringPointer())
	plan.InstanceUrl = util.SafeStringDatasource(plan.InstanceUrl.ValueStringPointer())
	plan.ObjectToBeImported = util.SafeStringDatasource(plan.ObjectToBeImported.ValueStringPointer())
	plan.FeatureLicenseJson = util.SafeStringDatasource(plan.FeatureLicenseJson.ValueStringPointer())
	plan.CustomCreateaccountUrl = util.SafeStringDatasource(plan.CustomCreateaccountUrl.ValueStringPointer())
	plan.Createaccountjson = util.SafeStringDatasource(plan.Createaccountjson.ValueStringPointer())
	plan.AccountFilterQuery = util.SafeStringDatasource(plan.AccountFilterQuery.ValueStringPointer())
	plan.AccountFieldQuery = util.SafeStringDatasource(plan.AccountFieldQuery.ValueStringPointer())
	plan.FieldMappingJson = util.SafeStringDatasource(plan.FieldMappingJson.ValueStringPointer())
	plan.Modifyaccountjson = util.SafeStringDatasource(plan.Modifyaccountjson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.Customconfigjson = util.SafeStringDatasource(plan.Customconfigjson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())

	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *SalesforceConnectionResource) ReadSalesforceConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Salesforce connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_salesforce_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})
	if err != nil {
		errorCode := salesforceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Salesforce connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateSalesforceConnectionResponse(apiResp); err != nil {
		errorCode := salesforceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Salesforce datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.SalesforceConnectionResponse != nil && apiResp.SalesforceConnectionResponse.Errorcode != nil && *apiResp.SalesforceConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.SalesforceConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.SalesforceConnectionResponse.Msg))
		errorCode := salesforceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Salesforce connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.SalesforceConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.SalesforceConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Salesforce connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.SalesforceConnectionResponse != nil && apiResp.SalesforceConnectionResponse.Connectionkey != nil {
				return *apiResp.SalesforceConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SalesforceConnectionResource) UpdateModelFromReadResponse(state *SalesforceConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.SalesforceConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.SalesforceConnectionResponse.Connectionkey))

	// Update all fields from API response
	state.ConnectionName = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Emailtemplate)
	state.ObjectToBeImported = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.OBJECT_TO_BE_IMPORTED)
	state.FeatureLicenseJson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FEATURE_LICENSE_JSON)
	state.Createaccountjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.RedirectUri = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.REDIRECT_URI)
	state.Modifyaccountjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.MODIFYACCOUNTJSON)
	state.ClientId = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CLIENT_ID)
	state.PamConfig = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.Customconfigjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	state.FieldMappingJson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FIELD_MAPPING_JSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.AccountFieldQuery = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FIELD_QUERY)
	state.CustomCreateaccountUrl = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOM_CREATEACCOUNT_URL)
	state.AccountFilterQuery = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FILTER_QUERY)
	state.InstanceUrl = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.INSTANCE_URL)
}

func (r *SalesforceConnectionResource) ValidateSalesforceConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SalesforceConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Salesforce connection response is nil")
	}
	return nil
}

func (r *SalesforceConnectionResource) UpdateSalesforceConnection(ctx context.Context, plan *SalesforceConnectorResourceModel, config *SalesforceConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Salesforce connection update")

	// Build Salesforce connection update request
	tflog.Debug(logCtx, "Building Salesforce connection update request")

	salesforceConn := r.BuildSalesforceConnector(plan, config)

	salesforceConnRequest := openapi.CreateOrUpdateRequest{
		SalesforceConnector: &salesforceConn,
	}

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_salesforce_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, salesforceConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := salesforceErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update Salesforce connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := salesforceErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "Salesforce connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Salesforce connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SalesforceConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config SalesforceConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateSalesforceConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Salesforce connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"Salesforce Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "Salesforce connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *SalesforceConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SalesforceConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	connectionName := state.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.ReadSalesforceConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Salesforce connection read failed", "", err)
		resp.Diagnostics.AddError(
			"Salesforce Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	apiMessage := util.SafeDeref(apiResp.SalesforceConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.SalesforceConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Salesforce connection resource read completed successfully")
}

func (r *SalesforceConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config SalesforceConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request for connection '%s'", errorCode, state.ConnectionName.ValueString()),
		)
		return
	}

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, plan.ConnectionName.ValueString()),
		)
		return
	}

	// Validate that connection name cannot be updated
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		errorCode := salesforceErrorCodes.NameImmutable()
		opCtx.LogOperationError(ctx, "Connection name cannot be updated", errorCode,
			fmt.Errorf("attempted to change connection name from '%s' to '%s'", state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
			map[string]interface{}{
				"old_name": state.ConnectionName.ValueString(),
				"new_name": plan.ConnectionName.ValueString(),
			})
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorCode),
			fmt.Sprintf("[%s] Cannot change connection name from '%s' to '%s'", errorCode, state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	updateResp, err := r.UpdateSalesforceConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Salesforce connection update failed", "", err)
		resp.Diagnostics.AddError(
			"Salesforce Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Refresh state after update
	getResp, err := r.ReadSalesforceConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated Salesforce connection", "", err)
		resp.Diagnostics.AddError(
			"Salesforce Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	apiMessage := util.SafeDeref(updateResp.Msg)
	plan.Msg = types.StringValue(apiMessage)
	plan.ErrorCode = types.StringValue(*updateResp.ErrorCode)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Salesforce connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}
func (r *SalesforceConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Resource deletion is not supported by this provider. Please remove the resource manually if required, or contact your administrator.",
	)
}

func (r *SalesforceConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing a Salesforce connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "Salesforce connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
