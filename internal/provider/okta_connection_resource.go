// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_okta_connection_resource manages Okta connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new okta connector using the supplied configuration.
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
	"terraform-provider-Saviynt/util/errorsutil"

	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OktaConnectionResource{}
var _ resource.ResourceWithImportState = &OktaConnectionResource{}

// Initialize error codes for Okta Connection operations
var oktaErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeOkta)

type OktaConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                            types.String `tfsdk:"id"`
	ImportUrl                     types.String `tfsdk:"import_url"`
	AuthToken                     types.String `tfsdk:"auth_token"`
	AccountFieldMappings          types.String `tfsdk:"account_field_mappings"`
	UserFieldMappings             types.String `tfsdk:"user_field_mappings"`
	EntitlementTypesMappings      types.String `tfsdk:"entitlement_types_mappings"`
	ImportInactiveApps            types.String `tfsdk:"import_inactive_apps"`
	OktaApplicationSecuritySystem types.String `tfsdk:"okta_application_securitysystem"`
	OktaGroupsFilter              types.String `tfsdk:"okta_groups_filter"`
	AppAccountFieldMappings       types.String `tfsdk:"app_account_field_mappings"`
	StatusThresholdConfig         types.String `tfsdk:"status_threshold_config"`
	AuditFilter                   types.String `tfsdk:"audit_filter"`
	ModifyUserDataJson            types.String `tfsdk:"modify_user_data_json"`
	ActivateEndpoint              types.String `tfsdk:"activate_endpoint"`
	ConfigJson                    types.String `tfsdk:"config_json"`
	PamConfig                     types.String `tfsdk:"pam_config"`
}

type OktaConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewOktaConnectionResource() resource.Resource {
	return &OktaConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewOktaConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &OktaConnectionResource{
		connectionFactory: factory,
	}
}

func (r *OktaConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_okta_connection_resource"
}

func OktaConnectorSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"import_url": schema.StringAttribute{
			Required:    true,
			Description: "Base URL for Okta API calls.",
		},
		"auth_token": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "API token for Okta authentication.",
		},
		"account_field_mappings": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Maps Okta user fields to Saviynt account fields.",
		},
		"user_field_mappings": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " Maps Okta user fields to Saviynt user fields.",
		},
		"entitlement_types_mappings": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " Maps Okta entitlements to Saviynt entitlement types.",
		},
		"import_inactive_apps": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " Controls import of inactive/disabled Okta applications.",
		},
		"okta_application_securitysystem": schema.StringAttribute{
			Required:    true,
			Description: "Saviynt security system name for Okta applications.",
		},
		"okta_groups_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Filter criteria for selective group import from Okta.",
		},
		"app_account_field_mappings": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Maps Okta application user fields to Saviynt account field.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON config for status mapping, thresholds, and bulk operation safety controls.",
		},
		"audit_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Filter for importing specific audit events.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " JSON configuration for user data modification operations during provisioning.",
		},
		"activate_endpoint": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Auto-enables disabled endpoints based on application status.",
		},
		"config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "General connector configuration including timeouts, retries, and connector-specific settings.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Privileged Access Management configuration for PAM operations and bootstrap processes.",
		},
	}
}

func (r *OktaConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.OktaConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), OktaConnectorSchema()),
	}
}

func (r *OktaConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Okta connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := oktaErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "Okta connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *OktaConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *OktaConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *OktaConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *OktaConnectionResource) BuildOktaConnector(plan *OktaConnectorResourceModel, config *OktaConnectorResourceModel) openapi.OktaConnector {
	oktaConn := openapi.OktaConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "Okta",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required field
		IMPORTURL:                       plan.ImportUrl.ValueString(),
		AUTHTOKEN:                       config.AuthToken.ValueString(),
		OKTA_APPLICATION_SECURITYSYSTEM: plan.OktaApplicationSecuritySystem.ValueString(),
		//optional field
		ACCOUNTFIELDMAPPINGS:     util.StringPointerOrEmpty(plan.AccountFieldMappings),
		USERFIELDMAPPINGS:        util.StringPointerOrEmpty(plan.UserFieldMappings),
		ENTITLEMENTTYPESMAPPINGS: util.StringPointerOrEmpty(plan.EntitlementTypesMappings),
		IMPORT_INACTIVE_APPS:     util.StringPointerOrEmpty(plan.ImportInactiveApps),
		OKTA_GROUPS_FILTER:       util.StringPointerOrEmpty(plan.OktaGroupsFilter),
		APPACCOUNTFIELDMAPPINGS:  util.StringPointerOrEmpty(plan.AppAccountFieldMappings),
		STATUS_THRESHOLD_CONFIG:  util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		AUDIT_FILTER:             util.StringPointerOrEmpty(plan.AuditFilter),
		MODIFYUSERDATAJSON:       util.StringPointerOrEmpty(plan.ModifyUserDataJson),
		ACTIVATE_ENDPOINT:        util.StringPointerOrEmpty(plan.ActivateEndpoint),
		ConfigJSON:               util.StringPointerOrEmpty(plan.ConfigJson),
		PAM_CONFIG:               util.StringPointerOrEmpty(plan.PamConfig),
	}

	if plan.VaultConnection.ValueString() != "" {
		oktaConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		oktaConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		oktaConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return oktaConn
}

func (r *OktaConnectionResource) UpdateModelFromCreateResponse(plan *OktaConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ImportUrl = util.SafeStringDatasource(plan.ImportUrl.ValueStringPointer())
	plan.OktaApplicationSecuritySystem = util.SafeStringDatasource(plan.OktaApplicationSecuritySystem.ValueStringPointer())
	plan.AccountFieldMappings = util.SafeStringDatasource(plan.AccountFieldMappings.ValueStringPointer())
	plan.UserFieldMappings = util.SafeStringDatasource(plan.UserFieldMappings.ValueStringPointer())
	plan.EntitlementTypesMappings = util.SafeStringDatasource(plan.EntitlementTypesMappings.ValueStringPointer())
	plan.ImportInactiveApps = util.SafeStringDatasource(plan.ImportInactiveApps.ValueStringPointer())
	plan.OktaGroupsFilter = util.SafeStringDatasource(plan.OktaGroupsFilter.ValueStringPointer())
	plan.AppAccountFieldMappings = util.SafeStringDatasource(plan.AppAccountFieldMappings.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.AuditFilter = util.SafeStringDatasource(plan.AuditFilter.ValueStringPointer())
	plan.ModifyUserDataJson = util.SafeStringDatasource(plan.ModifyUserDataJson.ValueStringPointer())
	plan.ActivateEndpoint = util.SafeStringDatasource(plan.ActivateEndpoint.ValueStringPointer())
	plan.ConfigJson = util.SafeStringDatasource(plan.ConfigJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *OktaConnectionResource) CreateOktaConnection(ctx context.Context, plan *OktaConnectorResourceModel, config *OktaConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Okta connection creation")

	// Check if connection already exists (idempotency check) with retry logic
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
		errorCode := oktaErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.OktaConnectionResponse != nil &&
		existingResource.OktaConnectionResponse.Errorcode != nil &&
		*existingResource.OktaConnectionResponse.Errorcode == 0 {

		errorCode := oktaErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "create", connectionName, nil)
	}

	// Build Okta connection create request
	tflog.Debug(ctx, "Building Okta connection create request")

	oktaConn := r.BuildOktaConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		OktaConnector: &oktaConn,
	}

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_okta_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})
	if err != nil {
		errorCode := oktaErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create Okta connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := oktaErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Okta connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": func() interface{} {
					if apiResp != nil && apiResp.ErrorCode != nil {
						return *apiResp.ErrorCode
					}
					return "unknown"
				}(),
				"message": func() interface{} {
					if apiResp != nil && apiResp.Msg != nil {
						return errorsutil.SanitizeMessage(apiResp.Msg)
					}
					return "API call failed"
				}(),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Okta connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp != nil && apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *OktaConnectionResource) ReadOktaConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Okta connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_okta_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := oktaErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Okta connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "read", connectionName, err)
	}

	if apiResp != nil && apiResp.OktaConnectionResponse != nil && apiResp.OktaConnectionResponse.Errorcode != nil && *apiResp.OktaConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.OktaConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.OktaConnectionResponse.Msg))
		errorCode := oktaErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Okta connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.OktaConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.OktaConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Okta connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.OktaConnectionResponse != nil && apiResp.OktaConnectionResponse.Connectionkey != nil {
				return *apiResp.OktaConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *OktaConnectionResource) UpdateOktaConnection(ctx context.Context, plan *OktaConnectorResourceModel, config *OktaConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Okta connection update")

	// Build Okta connection update request
	tflog.Debug(logCtx, "Building Okta connection update request")

	oktaConn := r.BuildOktaConnector(plan, config)
	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		oktaConn.BaseConnector.VaultConnection = &emptyStr
		oktaConn.BaseConnector.VaultConfiguration = &emptyStr
		oktaConn.BaseConnector.Saveinvault = &emptyStr
	}

	updateReq := openapi.CreateOrUpdateRequest{
		OktaConnector: &oktaConn,
	}

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_okta_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := oktaErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update Okta connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := oktaErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "Okta connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Okta connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *OktaConnectionResource) UpdateModelFromReadResponse(state *OktaConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.OktaConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.OktaConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Emailtemplate)
	state.ImportUrl = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.IMPORTURL)
	state.OktaApplicationSecuritySystem = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.OKTA_APPLICATION_SECURITYSYSTEM)
	state.AccountFieldMappings = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.ACCOUNTFIELDMAPPINGS)
	state.UserFieldMappings = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.USERFIELDMAPPINGS)
	state.EntitlementTypesMappings = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.ENTITLEMENTTYPESMAPPINGS)
	state.ImportInactiveApps = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.IMPORT_INACTIVE_APPS)
	state.OktaGroupsFilter = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.OKTA_GROUPS_FILTER)
	state.AppAccountFieldMappings = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.APPACCOUNTFIELDMAPPINGS)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.AuditFilter = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.AUDIT_FILTER)
	state.ModifyUserDataJson = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.ActivateEndpoint = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.ACTIVATE_ENDPOINT)
	state.ConfigJson = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.ConfigJSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionattributes.PAM_CONFIG)

	apiMessage := util.SafeDeref(apiResp.OktaConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.OktaConnectionResponse.Errorcode)
}

func (r *OktaConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config OktaConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaErrorCodes.PlanExtraction()
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
		errorCode := oktaErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateOktaConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Okta connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"Okta Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "Okta connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *OktaConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OktaConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadOktaConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Okta connection read failed", "", err)
		resp.Diagnostics.AddError(
			"Okta Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Okta connection resource read completed successfully")
}

func (r *OktaConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config OktaConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaErrorCodes.StateExtraction()
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
		errorCode := oktaErrorCodes.PlanExtraction()
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
		errorCode := oktaErrorCodes.ConfigExtraction()
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
		errorCode := oktaErrorCodes.NameImmutable()
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
	_, err := r.UpdateOktaConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Okta connection update failed", "", err)
		resp.Diagnostics.AddError(
			"Okta Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadOktaConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated Okta connection", "", err)
		resp.Diagnostics.AddError(
			"Okta Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Okta connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *OktaConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
func (r *OktaConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an Okta connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "Okta connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
