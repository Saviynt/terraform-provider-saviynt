// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_github_rest_connection_resource manages GithubRest connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new GithubRest connector using the supplied configuration.
//   - Read: fetches the current connector state from Saviynt to keep Terraform's state in sync.
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

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GithubRestConnectionResource{}
var _ resource.ResourceWithImportState = &GithubRestConnectionResource{}

// Initialize error codes for GitHub REST Connection operations
var githubRestErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeGithubREST)

type GithubRestConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                      types.String `tfsdk:"id"`
	ConnectionJSON          types.String `tfsdk:"connection_json"`
	ConnectionJSONWO        types.String `tfsdk:"connection_json_wo"`
	ImportAccountEntJSON    types.String `tfsdk:"import_account_ent_json"`
	Access_Tokens           types.String `tfsdk:"access_tokens"`
	Access_TokensWO         types.String `tfsdk:"access_tokens_wo"`
	Organization_List       types.String `tfsdk:"organization_list"`
	Status_Threshold_Config types.String `tfsdk:"status_threshold_config"`
}

type GithubRestConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	saviyntVersion    string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewGithubRestConnectionResource() resource.Resource {
	return &GithubRestConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewGithubRestConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &GithubRestConnectionResource{
		connectionFactory: factory,
	}
}

func (r *GithubRestConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_github_rest_connection_resource"
}

func GithubRestConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_json": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Property for ConnectionJSON. For setting connection_json either this field or connection_json_wo need to be set",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json_wo")),
			},
		},
		"connection_json_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for ConnectionJSON (write-only). For setting connection_json either this field or connection_json_wo need to be set",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json")),
			},
		},
		"import_account_ent_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ImportAccountEntJSON",
		},
		"access_tokens": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Property for ACCESS_TOKENS. For setting access_tokens either this field or access_tokens_wo need to be set",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("access_tokens_wo")),
			},
		},
		"access_tokens_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for ACCESS_TOKENS (write-only). For setting access_tokens either this field or access_tokens_wo need to be set",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("access_tokens")),
			},
		},
		"organization_list": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ORGANIZATION_LIST",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for STATUS_THRESHOLD_CONFIG",
		},
	}
}

func (r *GithubRestConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.GithubRestConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), GithubRestConnectorResourceSchema()),
	}
}

func (r *GithubRestConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection resource configuration")

	if req.ProviderData == nil {
		opCtx.LogOperationEnd(ctx, "GitHub REST connection resource configuration completed - no provider data")
		return
	}

	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := githubRestErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.saviyntVersion = prov.saviyntVersion
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "GitHub REST connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *GithubRestConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *GithubRestConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *GithubRestConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *GithubRestConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config GithubRestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.PlanExtraction()
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
		errorCode := githubRestErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Validate version-specific attributes for GitHub REST connector
	util.ValidateAttributeCompatibility(r.saviyntVersion, "GitHubREST", "status_threshold_config", plan.Status_Threshold_Config.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateGithubRestConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "GitHub REST connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"GitHub REST Connection Creation Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "GitHub REST connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *GithubRestConnectionResource) CreateGithubRestConnection(ctx context.Context, plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection creation")

	// Check if connection already exists with retry logic
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
		errorCode := githubRestErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.GithubRESTConnectionResponse != nil &&
		existingResource.GithubRESTConnectionResponse.Errorcode != nil &&
		*existingResource.GithubRESTConnectionResponse.Errorcode == 0 {

		errorCode := githubRestErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "create", connectionName, nil)
	}

	githubRestConn := r.BuildGithubRestConnector(plan, config)
	githubRestConnRequest := openapi.CreateOrUpdateRequest{
		GithubRESTConnector: &githubRestConn,
	}

	// Execute create operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_githubrest_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, githubRestConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := githubRestErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "create", connectionName, err)
	}

	// Check for API business logic errors
	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := githubRestErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "GitHub REST connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "GitHub REST connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp != nil && apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return nil
		}()})

	return apiResp, nil
}

func (r *GithubRestConnectionResource) BuildGithubRestConnector(plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) openapi.GithubRESTConnector {
	var connectionJson string
	if !config.ConnectionJSON.IsNull() && !config.ConnectionJSON.IsUnknown() {
		connectionJson = config.ConnectionJSON.ValueString()
	} else if !config.ConnectionJSONWO.IsNull() && !config.ConnectionJSONWO.IsUnknown() {
		connectionJson = config.ConnectionJSONWO.ValueString()
	}

	var accessTokens string
	if !config.Access_Tokens.IsNull() && !config.Access_Tokens.IsUnknown() {
		accessTokens = config.Access_Tokens.ValueString()
	} else if !config.Access_TokensWO.IsNull() && !config.Access_TokensWO.IsUnknown() {
		accessTokens = config.Access_TokensWO.ValueString()
	}

	githubRestConn := openapi.GithubRESTConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:        "GithubRest",
			ConnectionName:        plan.ConnectionName.ValueString(),
			ConnectionDescription: util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:       util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:         util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		ConnectionJSON:          util.StringPointerOrEmpty(types.StringValue(connectionJson)),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJSON),
		ACCESS_TOKENS:           util.StringPointerOrEmpty(types.StringValue(accessTokens)),
		ORGANIZATION_LIST:       util.StringPointerOrEmpty(plan.Organization_List),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.Status_Threshold_Config),
	}

	if plan.VaultConnection.ValueString() != "" {
		githubRestConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		githubRestConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		githubRestConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return githubRestConn
}

func (r *GithubRestConnectionResource) UpdateModelFromCreateResponse(plan *GithubRestConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	if apiResp != nil && apiResp.ConnectionKey != nil {
		plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
		plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	}

	// Update all optional fields to maintain state
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ImportAccountEntJSON = util.SafeStringDatasource(plan.ImportAccountEntJSON.ValueStringPointer())
	plan.Organization_List = util.SafeStringDatasource(plan.Organization_List.ValueStringPointer())
	plan.Status_Threshold_Config = util.SafeStringDatasource(plan.Status_Threshold_Config.ValueStringPointer())

	if apiResp != nil {
		plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
		plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	}
}

func (r *GithubRestConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GithubRestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadGithubRestConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "GitHub REST connection read failed", "", err)
		resp.Diagnostics.AddError(
			"GitHub REST Connection Read Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromReadResponse(&state, apiResp)

	apiMessage := util.SafeDeref(apiResp.GithubRESTConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.GithubRESTConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "GitHub REST connection resource read completed successfully")
}

func (r *GithubRestConnectionResource) ReadGithubRestConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_githubrest_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := githubRestErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateGithubRestConnectionResponse(apiResp); err != nil {
		errorCode := githubRestErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for GithubREST datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	// Check for API business logic errors
	if apiResp != nil && apiResp.GithubRESTConnectionResponse != nil && apiResp.GithubRESTConnectionResponse.Errorcode != nil && *apiResp.GithubRESTConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.GithubRESTConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.GithubRESTConnectionResponse.Msg))
		errorCode := githubRestErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "GitHub REST connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.GithubRESTConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.GithubRESTConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "GitHub REST connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp != nil && apiResp.GithubRESTConnectionResponse != nil && apiResp.GithubRESTConnectionResponse.Connectionkey != nil {
				return *apiResp.GithubRESTConnectionResponse.Connectionkey
			}
			return nil
		}()})

	return apiResp, nil
}

func (r *GithubRestConnectionResource) UpdateModelFromReadResponse(state *GithubRestConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.GithubRESTConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.GithubRESTConnectionResponse.Connectionkey))

	// Update all fields from API response
	state.ConnectionName = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Emailtemplate)
	state.ImportAccountEntJSON = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	state.Organization_List = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ORGANIZATION_LIST)
	state.Status_Threshold_Config = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
}

func (r *GithubRestConnectionResource) ValidateGithubRestConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.GithubRESTConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - GithubREST connection response is nil")
	}
	return nil
}

func (r *GithubRestConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config GithubRestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.StateExtraction()
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
		errorCode := githubRestErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request for connection '%s'", errorCode, state.ConnectionName.ValueString()),
		)
		return
	}

	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.ConfigExtraction()
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
		errorCode := githubRestErrorCodes.NameImmutable()
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

	// Validate version-specific attributes for GitHub REST connector
	util.ValidateAttributeCompatibility(r.saviyntVersion, "GitHubREST", "status_threshold_config", plan.Status_Threshold_Config.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}
	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	updateResp, err := r.UpdateGithubRestConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "GitHub REST connection update failed", "", err)
		resp.Diagnostics.AddError(
			"GitHub REST Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadGithubRestConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated GitHub REST connection", "", err)
		resp.Diagnostics.AddError(
			"GitHub REST Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromReadResponse(&plan, getResp)

	apiMessage := util.SafeDeref(updateResp.Msg)
	plan.Msg = types.StringValue(apiMessage)
	plan.ErrorCode = types.StringValue(*updateResp.ErrorCode)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "GitHub REST connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *GithubRestConnectionResource) UpdateGithubRestConnection(ctx context.Context, plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection update")

	githubRestConn := r.BuildGithubRestConnector(plan, config)

	githubRestConnRequest := openapi.CreateOrUpdateRequest{
		GithubRESTConnector: &githubRestConn,
	}

	// Execute update operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_githubrest_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, githubRestConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := githubRestErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "update", connectionName, err)
	}

	// Check for API business logic errors
	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := githubRestErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "GitHub REST connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "GitHub REST connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp != nil && apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return nil
		}()})

	return apiResp, nil
}

func (r *GithubRestConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *GithubRestConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an GithubRest connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GithubRest connection resource import")

	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "AD connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})

}
