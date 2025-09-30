// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &RestConnectionResource{}
var _ resource.ResourceWithImportState = &RestConnectionResource{}

// Initialize error codes for REST Connection operations
var restErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeREST)

type RestConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                    types.String `tfsdk:"id"`
	ConnectionJSON        types.String `tfsdk:"connection_json"`
	ImportUserJson        types.String `tfsdk:"import_user_json"`
	ImportAccountEntJson  types.String `tfsdk:"import_account_ent_json"`
	StatusThresholdConfig types.String `tfsdk:"status_threshold_config"`
	CreateAccountJson     types.String `tfsdk:"create_account_json"`
	UpdateAccountJson     types.String `tfsdk:"update_account_json"`
	EnableAccountJson     types.String `tfsdk:"enable_account_json"`
	DisableAccountJson    types.String `tfsdk:"disable_account_json"`
	AddAccessJson         types.String `tfsdk:"add_access_json"`
	RemoveAccessJson      types.String `tfsdk:"remove_access_json"`
	UpdateUserJson        types.String `tfsdk:"update_user_json"`
	ChangePassJson        types.String `tfsdk:"change_pass_json"`
	RemoveAccountJson     types.String `tfsdk:"remove_account_json"`
	TicketStatusJson      types.String `tfsdk:"ticket_status_json"`
	CreateTicketJson      types.String `tfsdk:"create_ticket_json"`
	EndpointsFilter       types.String `tfsdk:"endpoints_filter"`
	PasswdPolicyJson      types.String `tfsdk:"passwd_policy_json"`
	ConfigJSON            types.String `tfsdk:"config_json"`
	AddFFIDAccessJson     types.String `tfsdk:"add_ffid_access_json"`
	RemoveFFIDAccessJson  types.String `tfsdk:"remove_ffid_access_json"`
	ModifyUserdataJson    types.String `tfsdk:"modify_user_data_json"`
	SendOtpJson           types.String `tfsdk:"send_otp_json"`
	ValidateOtpJson       types.String `tfsdk:"validate_otp_json"`
	PamConfig             types.String `tfsdk:"pam_config"`
	// TER-176
	ApplicationDiscoveryJson types.String `tfsdk:"application_discovery_json"`
	CreateEntitlementJson    types.String `tfsdk:"create_entitlement_json"`
	DeleteEntitlementJson    types.String `tfsdk:"delete_entitlement_json"`
	UpdateEntitlementJson    types.String `tfsdk:"update_entitlement_json"`
}

type RestConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	connectionFactory client.ConnectionFactoryInterface
}

func NewRestConnectionResource() resource.Resource {
	return &RestConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewRestConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &RestConnectionResource{
		connectionFactory: factory,
	}
}

func (r *RestConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_rest_connection_resource"
}

func RestConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Dynamic JSON configuration for the connection. Must be a valid JSON object string.",
		},
		"import_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for importing users.",
		},
		"import_account_ent_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for importing accounts and entitlements.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for status thresholds.",
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create an account.",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update an account.",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to enable an account.",
		},
		"disable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to disable an account.",
		},
		"add_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to add access.",
		},
		"remove_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to remove access.",
		},
		"update_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update a user.",
		},
		"change_pass_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "JSON to change a user's password.",
		},
		"remove_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to remove an account.",
		},
		"ticket_status_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to check ticket status.",
		},
		"create_ticket_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create a ticket.",
		},
		"endpoints_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Filter criteria for endpoints.",
		},
		"passwd_policy_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON defining the password policy.",
		},
		"config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "General configuration JSON for the REST connector.",
		},
		"add_ffid_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to add FFID access.",
		},
		"remove_ffid_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to remove FFID access.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for modifying user data.",
		},
		"send_otp_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to send OTP.",
		},
		"validate_otp_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to validate OTP.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "PAM configuration JSON.",
		},
		//TER-176
		"application_discovery_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The ApplicationDiscoveryJSON attribute is specifically implemented for ServiceNow application discovery, allowing automated discovery and import of applications from ServiceNow instances.",
		},
		"create_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The three entitlement JSON attributes (Create, Update, Delete) are part of a comprehensive entitlement management system for REST connectors, with supporting constants and service classes.",
		},
		"delete_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The three entitlement JSON attributes (Create, Update, Delete) are part of a comprehensive entitlement management system for REST connectors, with supporting constants and service classes.",
		},
		"update_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The three entitlement JSON attributes (Create, Update, Delete) are part of a comprehensive entitlement management system for REST connectors, with supporting constants and service classes.",
		},
	}
}

func (r *RestConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.RestConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), RestConnectorResourceSchema()),
	}
}

func (r *RestConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "REST connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		errorCode := restErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *saviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*saviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *saviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken

	opCtx.LogOperationEnd(ctx, "REST connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *RestConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *RestConnectionResource) SetToken(token string) {
	r.token = token
}

func (r *RestConnectionResource) CreateRESTConnection(ctx context.Context, plan *RestConnectorResourceModel, config *RestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting REST connection creation")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	// Check if connection already exists (idempotency check)
	tflog.Debug(logCtx, "Checking if connection already exists")

	// Use original context for API calls to maintain test compatibility
	existingResource, _, _ := connectionOps.GetConnectionDetails(ctx, connectionName)
	if existingResource != nil &&
		existingResource.RESTConnectionResponse != nil &&
		existingResource.RESTConnectionResponse.Errorcode != nil &&
		*existingResource.RESTConnectionResponse.Errorcode == 0 {

		errorCode := restErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "create", connectionName, nil)
	}

	// Build REST connection create request
	tflog.Debug(ctx, "Building REST connection create request")

	restConn := r.BuildRESTConnector(plan, config)
	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Execute create operation through interface
	tflog.Debug(ctx, "Executing create operation")

	apiResp, _, err := connectionOps.CreateOrUpdateConnection(ctx, restConnRequest)
	if err != nil {
		errorCode := restErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := restErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "REST connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "REST connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *RestConnectionResource) BuildRESTConnector(plan *RestConnectorResourceModel, config *RestConnectorResourceModel) openapi.RESTConnector {
	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "REST",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//optional fields
		ConnectionJSON:          util.StringPointerOrEmpty(config.ConnectionJSON),
		ImportUserJSON:          util.StringPointerOrEmpty(plan.ImportUserJson),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJson),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CreateAccountJSON:       util.StringPointerOrEmpty(plan.CreateAccountJson),
		UpdateAccountJSON:       util.StringPointerOrEmpty(plan.UpdateAccountJson),
		EnableAccountJSON:       util.StringPointerOrEmpty(plan.EnableAccountJson),
		DisableAccountJSON:      util.StringPointerOrEmpty(plan.DisableAccountJson),
		AddAccessJSON:           util.StringPointerOrEmpty(plan.AddAccessJson),
		RemoveAccessJSON:        util.StringPointerOrEmpty(plan.RemoveAccessJson),
		UpdateUserJSON:          util.StringPointerOrEmpty(plan.UpdateUserJson),
		ChangePassJSON:          util.StringPointerOrEmpty(config.ChangePassJson),
		RemoveAccountJSON:       util.StringPointerOrEmpty(plan.RemoveAccountJson),
		TicketStatusJSON:        util.StringPointerOrEmpty(plan.TicketStatusJson),
		CreateTicketJSON:        util.StringPointerOrEmpty(plan.CreateTicketJson),
		ENDPOINTS_FILTER:        util.StringPointerOrEmpty(plan.EndpointsFilter),
		PasswdPolicyJSON:        util.StringPointerOrEmpty(plan.PasswdPolicyJson),
		ConfigJSON:              util.StringPointerOrEmpty(plan.ConfigJSON),
		AddFFIDAccessJSON:       util.StringPointerOrEmpty(plan.AddFFIDAccessJson),
		RemoveFFIDAccessJSON:    util.StringPointerOrEmpty(plan.RemoveFFIDAccessJson),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		SendOtpJSON:             util.StringPointerOrEmpty(plan.SendOtpJson),
		ValidateOtpJSON:         util.StringPointerOrEmpty(plan.ValidateOtpJson),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.PamConfig),
		//TER-176
		ApplicationDiscoveryJSON: util.StringPointerOrEmpty(plan.ApplicationDiscoveryJson),
		CreateEntitlementJSON:    util.StringPointerOrEmpty(plan.CreateEntitlementJson),
		DeleteEntitlementJSON:    util.StringPointerOrEmpty(plan.DeleteEntitlementJson),
		UpdateEntitlementJSON:    util.StringPointerOrEmpty(plan.UpdateEntitlementJson),
	}

	if plan.VaultConnection.ValueString() != "" {
		restConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		restConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		restConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return restConn
}

func (r *RestConnectionResource) UpdateModelFromCreateResponse(plan *RestConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))

	// Update all optional fields to maintain state
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ImportUserJson = util.SafeStringDatasource(plan.ImportUserJson.ValueStringPointer())
	plan.ImportAccountEntJson = util.SafeStringDatasource(plan.ImportAccountEntJson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.AddAccessJson = util.SafeStringDatasource(plan.AddAccessJson.ValueStringPointer())
	plan.RemoveAccessJson = util.SafeStringDatasource(plan.RemoveAccessJson.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.ChangePassJson = util.SafeStringDatasource(plan.ChangePassJson.ValueStringPointer())
	plan.RemoveAccountJson = util.SafeStringDatasource(plan.RemoveAccountJson.ValueStringPointer())
	plan.TicketStatusJson = util.SafeStringDatasource(plan.TicketStatusJson.ValueStringPointer())
	plan.CreateTicketJson = util.SafeStringDatasource(plan.CreateTicketJson.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.PasswdPolicyJson = util.SafeStringDatasource(plan.PasswdPolicyJson.ValueStringPointer())
	plan.ConfigJSON = util.SafeStringDatasource(plan.ConfigJSON.ValueStringPointer())
	plan.AddFFIDAccessJson = util.SafeStringDatasource(plan.AddFFIDAccessJson.ValueStringPointer())
	plan.RemoveFFIDAccessJson = util.SafeStringDatasource(plan.RemoveFFIDAccessJson.ValueStringPointer())
	plan.ModifyUserdataJson = util.SafeStringDatasource(plan.ModifyUserdataJson.ValueStringPointer())
	plan.SendOtpJson = util.SafeStringDatasource(plan.SendOtpJson.ValueStringPointer())
	plan.ValidateOtpJson = util.SafeStringDatasource(plan.ValidateOtpJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	//TER-176
	plan.ApplicationDiscoveryJson = util.SafeStringDatasource(plan.ApplicationDiscoveryJson.ValueStringPointer())
	plan.CreateEntitlementJson = util.SafeStringDatasource(plan.CreateEntitlementJson.ValueStringPointer())
	plan.DeleteEntitlementJson = util.SafeStringDatasource(plan.DeleteEntitlementJson.ValueStringPointer())
	plan.UpdateEntitlementJson = util.SafeStringDatasource(plan.UpdateEntitlementJson.ValueStringPointer())

	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *RestConnectionResource) ReadRESTConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting REST connection read operation")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	// Execute read operation through interface - use original context for API calls
	apiResp, _, err := connectionOps.GetConnectionDetails(ctx, connectionName)
	if err != nil {
		errorCode := restErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "read", connectionName, err)
	}

	if apiResp != nil && apiResp.RESTConnectionResponse != nil && apiResp.RESTConnectionResponse.Errorcode != nil && *apiResp.RESTConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.RESTConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.RESTConnectionResponse.Msg))
		errorCode := restErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "REST connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.RESTConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.RESTConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "REST connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.RESTConnectionResponse != nil && apiResp.RESTConnectionResponse.Connectionkey != nil {
				return *apiResp.RESTConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *RestConnectionResource) UpdateModelFromReadResponse(state *RestConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.RESTConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.RESTConnectionResponse.Connectionkey))

	// Update all fields from API response
	state.ConnectionName = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Emailtemplate)
	state.ImportUserJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportUserJSON)
	state.ImportAccountEntJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateAccountJSON)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateAccountJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.EnableAccountJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.DisableAccountJSON)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddAccessJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccessJSON)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateUserJSON)
	state.ChangePassJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ChangePassJSON)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccountJSON)
	state.TicketStatusJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.TicketStatusJSON)
	state.CreateTicketJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateTicketJSON)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.PasswdPolicyJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PasswdPolicyJSON)
	state.ConfigJSON = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ConfigJSON)
	state.AddFFIDAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddFFIDAccessJSON)
	state.RemoveFFIDAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveFFIDAccessJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.SendOtpJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.SendOtpJSON)
	state.ValidateOtpJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ValidateOtpJSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PAM_CONFIG)
	//TER-176
	state.ApplicationDiscoveryJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ApplicationDiscoveryJSON)
	state.CreateEntitlementJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateEntitlementJSON)
	state.DeleteEntitlementJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.DeleteEntitlementJSON)
	state.UpdateEntitlementJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateEntitlementJSON)

	apiMessage := util.SafeDeref(apiResp.RESTConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.RESTConnectionResponse.Errorcode)
}

func (r *RestConnectionResource) UpdateRESTConnection(ctx context.Context, plan *RestConnectorResourceModel, config *RestConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting REST connection update")

	// Use the factory to create connection operations
	connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), r.token)

	// Build REST connection update request
	tflog.Debug(logCtx, "Building REST connection update request")

	restConn := r.BuildRESTConnector(plan, config)
	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		restConn.BaseConnector.VaultConnection = &emptyStr
		restConn.BaseConnector.VaultConfiguration = &emptyStr
		restConn.BaseConnector.Saveinvault = &emptyStr
	}

	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Execute update operation through interface
	tflog.Debug(logCtx, "Executing update operation")

	// Use original context for API calls to maintain test compatibility
	apiResp, _, err := connectionOps.CreateOrUpdateConnection(ctx, restConnRequest)
	if err != nil {
		errorCode := restErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update REST connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := restErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "REST connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "REST connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *RestConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config RestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.PlanExtraction()
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

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateRESTConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "REST connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"REST Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "REST connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *RestConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadRESTConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "REST connection read failed", "", err)
		resp.Diagnostics.AddError(
			"REST Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "REST connection resource read completed successfully")
}

func (r *RestConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config RestConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.StateExtraction()
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
		errorCode := restErrorCodes.PlanExtraction()
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
		errorCode := restErrorCodes.ConfigExtraction()
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
		errorCode := restErrorCodes.NameImmutable()
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
	_, err := r.UpdateRESTConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "REST connection update failed", "", err)
		resp.Diagnostics.AddError(
			"REST Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadRESTConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated REST connection", "", err)
		resp.Diagnostics.AddError(
			"REST Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := restErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "REST connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *RestConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *RestConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing a REST connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "REST connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
