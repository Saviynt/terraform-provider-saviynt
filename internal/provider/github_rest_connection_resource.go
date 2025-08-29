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
	"os"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &githubRestConnectionResource{}
var _ resource.ResourceWithImportState = &githubRestConnectionResource{}

// Initialize error codes for GitHub REST Connection operations
var githubRestErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeGithubREST)

type GithubRestConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                      types.String `tfsdk:"id"`
	ConnectionJSON          types.String `tfsdk:"connection_json"`
	ImportAccountEntJSON    types.String `tfsdk:"import_account_ent_json"`
	Access_Tokens           types.String `tfsdk:"access_tokens"`
	Organization_List       types.String `tfsdk:"organization_list"`
	Status_Threshold_Config types.String `tfsdk:"status_threshold_config"`
}

type githubRestConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	connectionFactory client.ConnectionFactoryInterface
}

func NewGithubRestConnectionResource() resource.Resource {
	return &githubRestConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewGithubRestConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &githubRestConnectionResource{
		connectionFactory: factory,
	}
}

func (r *githubRestConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
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
			WriteOnly:   true,
			Description: "Property for ConnectionJSON",
		},
		"import_account_ent_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ImportAccountEntJSON",
		},
		"access_tokens": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for ACCESS_TOKENS",
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

func (r *githubRestConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.GithubRestConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), GithubRestConnectorResourceSchema()),
	}
}

func (r *githubRestConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection resource configuration")

	if req.ProviderData == nil {
		opCtx.LogOperationEnd(ctx, "GitHub REST connection resource configuration completed - no provider data")
		return
	}

	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		errorCode := githubRestErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *saviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*saviyntProvider"})
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken

	opCtx.LogOperationEnd(ctx, "GitHub REST connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *githubRestConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *githubRestConnectionResource) SetToken(token string) {
	r.token = token
}

func (r *githubRestConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

func (r *githubRestConnectionResource) CreateGithubRestConnection(ctx context.Context, plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection creation")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	// Check if connection already exists
	existingResource, _, _ := connectionOps.GetConnectionDetails(ctx, connectionName)
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

	apiResp, _, err := connectionOps.CreateOrUpdateConnection(ctx, githubRestConnRequest)
	if err != nil {
		errorCode := githubRestErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "create", connectionName, err)
	}

	// Check for API business logic errors
	if apiResp != nil && *apiResp.ErrorCode != "0" {
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

func (r *githubRestConnectionResource) BuildGithubRestConnector(plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) openapi.GithubRESTConnector {
	githubRestConn := openapi.GithubRESTConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:  "GithubRest",
			ConnectionName:  plan.ConnectionName.ValueString(),
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		ConnectionJSON:          util.StringPointerOrEmpty(config.ConnectionJSON),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJSON),
		ACCESS_TOKENS:           util.StringPointerOrEmpty(config.Access_Tokens),
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

func (r *githubRestConnectionResource) UpdateModelFromCreateResponse(plan *GithubRestConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
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

func (r *githubRestConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

func (r *githubRestConnectionResource) ReadGithubRestConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection read operation")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	apiResp, _, err := connectionOps.GetConnectionDetails(ctx, connectionName)
	if err != nil {
		errorCode := githubRestErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "read", connectionName, err)
	}

	// Check for API business logic errors
	if apiResp != nil && apiResp.GithubRESTConnectionResponse != nil && *apiResp.GithubRESTConnectionResponse.Errorcode != 0 {
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

func (r *githubRestConnectionResource) UpdateModelFromReadResponse(state *GithubRestConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
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

	apiMessage := util.SafeDeref(apiResp.GithubRESTConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.GithubRESTConnectionResponse.Errorcode)
}

func (r *githubRestConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	_, err := r.UpdateGithubRestConnection(ctx, &plan, &config)
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

func (r *githubRestConnectionResource) UpdateGithubRestConnection(ctx context.Context, plan *GithubRestConnectorResourceModel, config *GithubRestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection update")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	githubRestConn := r.BuildGithubRestConnector(plan, config)

	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		githubRestConn.BaseConnector.VaultConnection = &emptyStr
		githubRestConn.BaseConnector.VaultConfiguration = &emptyStr
		githubRestConn.BaseConnector.Saveinvault = &emptyStr
	}

	githubRestConnRequest := openapi.CreateOrUpdateRequest{
		GithubRESTConnector: &githubRestConn,
	}

	apiResp, _, err := connectionOps.CreateOrUpdateConnection(ctx, githubRestConnRequest)
	if err != nil {
		errorCode := githubRestErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update GitHub REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "update", connectionName, err)
	}

	// Check for API business logic errors
	if apiResp != nil && *apiResp.ErrorCode != "0" {
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

func (r *githubRestConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *githubRestConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
