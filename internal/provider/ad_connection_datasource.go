// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_ad_connection_datasource retrieves ad connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing ad connections by name.
package provider

import (
	"context"
	"fmt"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &AdConnectionsDataSource{}

// Initialize error codes for AD Connection datasource operations
var adDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeAD)

// ADConnectionsDataSource defines the data source
type AdConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type ADConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *ADConnectionAttributes `tfsdk:"connection_attributes"`
}

type ADConnectionAttributes struct {
	URL                        types.String             `tfsdk:"url"`
	ConnectionType             types.String             `tfsdk:"connection_type"`
	LastImportTime             types.String             `tfsdk:"last_import_time"`
	CreateAccountJSON          types.String             `tfsdk:"create_account_json"`
	DisableAccountJSON         types.String             `tfsdk:"disable_account_json"`
	GroupSearchBaseDN          types.String             `tfsdk:"group_search_base_dn"`
	PasswordNoOfSplChars       types.String             `tfsdk:"password_no_of_spl_chars"`
	PasswordNoOfDigits         types.String             `tfsdk:"password_no_of_digits"`
	StatusKeyJSON              types.String             `tfsdk:"status_key_json"`
	SearchFilter               types.String             `tfsdk:"search_filter"`
	ConfigJSON                 types.String             `tfsdk:"config_json"`
	RemoveAccountAction        types.String             `tfsdk:"remove_account_action"`
	AccountAttribute           types.String             `tfsdk:"account_attribute"`
	AccountNameRule            types.String             `tfsdk:"account_name_rule"`
	AdvSearch                  types.String             `tfsdk:"adv_search"`
	LDAPOrAD                   types.String             `tfsdk:"ldap_or_ad"`
	EntitlementAttribute       types.String             `tfsdk:"entitlement_attribute"`
	SetRandomPassword          types.String             `tfsdk:"set_random_password"`
	PasswordMinLength          types.String             `tfsdk:"password_min_length"`
	PasswordMaxLength          types.String             `tfsdk:"password_max_length"`
	PasswordNoOfCapsAlpha      types.String             `tfsdk:"password_no_of_caps_alpha"`
	SetDefaultPageSize         types.String             `tfsdk:"set_default_page_size"`
	IsTimeoutSupported         types.Bool               `tfsdk:"is_timeout_supported"`
	ReuseInactiveAccount       types.String             `tfsdk:"reuse_inactive_account"`
	ImportJSON                 types.String             `tfsdk:"import_json"`
	CreateUpdateMappings       types.String             `tfsdk:"create_update_mappings"`
	AdvanceFilterJSON          types.String             `tfsdk:"advance_filter_json"`
	OrgImportJSON              types.String             `tfsdk:"org_import_json"`
	PAMConfig                  types.String             `tfsdk:"pam_config"`
	PageSize                   types.String             `tfsdk:"page_size"`
	Base                       types.String             `tfsdk:"base"`
	DCLocator                  types.String             `tfsdk:"dc_locator"`
	StatusThresholdConfig      types.String             `tfsdk:"status_threshold_config"`
	ResetAndChangePasswordJSON types.String             `tfsdk:"reset_and_change_password_json"`
	SupportEmptyString         types.String             `tfsdk:"support_empty_string"`
	ReadOperationalAttributes  types.String             `tfsdk:"read_operational_attributes"`
	EnableAccountJSON          types.String             `tfsdk:"enable_account_json"`
	UserAttribute              types.String             `tfsdk:"user_attribute"`
	DefaultUserRole            types.String             `tfsdk:"default_user_role"`
	EndpointsFilter            types.String             `tfsdk:"endpoints_filter"`
	UpdateAccountJSON          types.String             `tfsdk:"update_account_json"`
	ReuseAccountJSON           types.String             `tfsdk:"reuse_account_json"`
	EnforceTreeDeletion        types.String             `tfsdk:"enforce_tree_deletion"`
	Filter                     types.String             `tfsdk:"filter"`
	ObjectFilter               types.String             `tfsdk:"object_filter"`
	UpdateUserJSON             types.String             `tfsdk:"update_user_json"`
	SaveConnection             types.String             `tfsdk:"save_connection"`
	SystemName                 types.String             `tfsdk:"system_name"`
	GroupImportMapping         types.String             `tfsdk:"group_import_mapping"`
	UnlockAccountJSON          types.String             `tfsdk:"unlock_account_json"`
	EnableGroupManagement      types.String             `tfsdk:"enable_group_management"`
	ModifyUserDataJSON         types.String             `tfsdk:"modify_user_data_json"`
	OrgBase                    types.String             `tfsdk:"org_base"`
	OrganizationAttribute      types.String             `tfsdk:"organization_attribute"`
	CreateOrgJSON              types.String             `tfsdk:"create_org_json"`
	UpdateOrgJSON              types.String             `tfsdk:"update_org_json"`
	MaxChangeNumber            types.String             `tfsdk:"max_change_number"`
	IncrementalConfig          types.String             `tfsdk:"incremental_config"`
	CheckForUnique             types.String             `tfsdk:"check_for_unique"`
	ConnectionTimeoutConfig    *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	IsTimeoutConfigValidated   types.Bool               `tfsdk:"is_timeout_config_validated"`
}

// NewADConnectionsDataSource creates a new AD connections data source with default factory
func NewADConnectionsDataSource() datasource.DataSource {
	return &AdConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewADConnectionsDataSourceWithFactory creates a new AD connections data source with custom factory
// Used primarily for testing with mock factories
func NewADConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &AdConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *AdConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *AdConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *AdConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *AdConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_ad_connection_datasource"
}

func ADConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"url":                            schema.StringAttribute{Computed: true},
				"connection_type":                schema.StringAttribute{Computed: true},
				"last_import_time":               schema.StringAttribute{Computed: true},
				"create_account_json":            schema.StringAttribute{Computed: true},
				"disable_account_json":           schema.StringAttribute{Computed: true},
				"group_search_base_dn":           schema.StringAttribute{Computed: true},
				"password_no_of_spl_chars":       schema.StringAttribute{Computed: true},
				"password_no_of_digits":          schema.StringAttribute{Computed: true},
				"status_key_json":                schema.StringAttribute{Computed: true},
				"search_filter":                  schema.StringAttribute{Computed: true},
				"config_json":                    schema.StringAttribute{Computed: true},
				"remove_account_action":          schema.StringAttribute{Computed: true},
				"account_attribute":              schema.StringAttribute{Computed: true},
				"account_name_rule":              schema.StringAttribute{Computed: true},
				"adv_search":                     schema.StringAttribute{Computed: true},
				"ldap_or_ad":                     schema.StringAttribute{Computed: true},
				"entitlement_attribute":          schema.StringAttribute{Computed: true},
				"set_random_password":            schema.StringAttribute{Computed: true},
				"password_min_length":            schema.StringAttribute{Computed: true},
				"password_max_length":            schema.StringAttribute{Computed: true},
				"password_no_of_caps_alpha":      schema.StringAttribute{Computed: true},
				"set_default_page_size":          schema.StringAttribute{Computed: true},
				"is_timeout_supported":           schema.BoolAttribute{Computed: true},
				"reuse_inactive_account":         schema.StringAttribute{Computed: true},
				"import_json":                    schema.StringAttribute{Computed: true},
				"create_update_mappings":         schema.StringAttribute{Computed: true},
				"advance_filter_json":            schema.StringAttribute{Computed: true},
				"org_import_json":                schema.StringAttribute{Computed: true},
				"pam_config":                     schema.StringAttribute{Computed: true},
				"page_size":                      schema.StringAttribute{Computed: true},
				"base":                           schema.StringAttribute{Computed: true},
				"dc_locator":                     schema.StringAttribute{Computed: true},
				"status_threshold_config":        schema.StringAttribute{Computed: true},
				"reset_and_change_password_json": schema.StringAttribute{Computed: true},
				"support_empty_string":           schema.StringAttribute{Computed: true},
				"read_operational_attributes":    schema.StringAttribute{Computed: true},
				"enable_account_json":            schema.StringAttribute{Computed: true},
				"user_attribute":                 schema.StringAttribute{Computed: true},
				"default_user_role":              schema.StringAttribute{Computed: true},
				"endpoints_filter":               schema.StringAttribute{Computed: true},
				"update_account_json":            schema.StringAttribute{Computed: true},
				"reuse_account_json":             schema.StringAttribute{Computed: true},
				"enforce_tree_deletion":          schema.StringAttribute{Computed: true},
				"filter":                         schema.StringAttribute{Computed: true},
				"object_filter":                  schema.StringAttribute{Computed: true},
				"update_user_json":               schema.StringAttribute{Computed: true},
				"save_connection":                schema.StringAttribute{Computed: true},
				"system_name":                    schema.StringAttribute{Computed: true},
				"group_import_mapping":           schema.StringAttribute{Computed: true},
				"unlock_account_json":            schema.StringAttribute{Computed: true},
				"enable_group_management":        schema.StringAttribute{Computed: true},
				"modify_user_data_json":          schema.StringAttribute{Computed: true},
				"org_base":                       schema.StringAttribute{Computed: true},
				"organization_attribute":         schema.StringAttribute{Computed: true},
				"create_org_json":                schema.StringAttribute{Computed: true},
				"update_org_json":                schema.StringAttribute{Computed: true},
				"max_change_number":              schema.StringAttribute{Computed: true},
				"incremental_config":             schema.StringAttribute{Computed: true},
				"check_for_unique":               schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
			},
		},
	}
}

// Schema defines the structure and attributes available for the AD connection data source
func (d *AdConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), ADConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *AdConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeAD, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting AD connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "AD connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := adDatasourceErrorCodes.ProviderConfig()
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
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	opCtx.LogOperationEnd(ctx, "AD connection datasource configured successfully")
}

// Read retrieves AD connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *AdConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ADConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeAD, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting AD connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adDatasourceErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request", errorCode),
		)
		return
	}

	// Update operation context with connection name if available
	connectionName := ""
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		connectionName = state.ConnectionName.ValueString()
		opCtx.ConnectionName = connectionName
		ctx = opCtx.AddContextToLogger(ctx)
	}

	// Prepare connection parameters
	var connectionKey *int64
	if !state.ConnectionKey.IsNull() {
		keyValue := state.ConnectionKey.ValueInt64()
		connectionKey = &keyValue
	}

	// Validate that at least one identifier is provided
	if connectionName == "" && connectionKey == nil {
		errorCode := adDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the AD connection.", errorCode),
		)
		return
	}

	// Execute API call to get AD connection details
	apiResp, err := d.ReadADConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadADConnectionDetails method
		errorCode := adDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read AD connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateADConnectionResponse(apiResp); err != nil {
		errorCode := adDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for AD datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromADConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleADAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for AD connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "AD connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadADConnectionDetails retrieves AD connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *AdConnectionsDataSource) ReadADConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeAD, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting AD connection API call")

	tflog.Debug(logCtx, "Executing API request to get AD connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_ad_connection_datasource", func(token string) error {
		connectionOps := d.connectionFactory.CreateConnectionOperations(d.client.APIBaseURL(), token)

		// Build request with both connectionname and connectionkey
		req := openapi.GetConnectionDetailsRequest{}
		if connectionName != "" {
			req.Connectionname = &connectionName
		}
		if connectionKey != nil {
			keyStr := fmt.Sprintf("%d", *connectionKey)
			req.Connectionkey = &keyStr
		}

		resp, httpResp, err := connectionOps.GetConnectionDetailsDataSource(ctx, req)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := adDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read AD connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeAD, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "AD connection API call completed successfully")

	return apiResp, nil
}

// ValidateADConnectionResponse validates that the API response contains valid AD connection data
// Returns standardized error if validation fails
func (d *AdConnectionsDataSource) ValidateADConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.ADConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - AD connection response is nil")
	}
	return nil
}

// UpdateModelFromADConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *AdConnectionsDataSource) UpdateModelFromADConnectionResponse(state *ADConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseADConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapADConnectionAttributes(state, apiResp)
}

// MapBaseADConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *AdConnectionsDataSource) MapBaseADConnectionFields(state *ADConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.ADConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.ADConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.ADConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-ad-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.ADConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.ADConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.ADConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.ADConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.ADConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADConnectionResponse.Emailtemplate)
}

// MapADConnectionAttributes maps detailed AD connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *AdConnectionsDataSource) MapADConnectionAttributes(state *ADConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.ADConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.ADConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &ADConnectionAttributes{
		URL:                        util.SafeStringDatasource(attrs.URL),
		ConnectionType:             util.SafeStringDatasource(attrs.ConnectionType),
		AdvSearch:                  util.SafeStringDatasource(attrs.ADVSEARCH),
		LastImportTime:             util.SafeStringDatasource(attrs.LAST_IMPORT_TIME),
		CreateAccountJSON:          util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		DisableAccountJSON:         util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON),
		GroupSearchBaseDN:          util.SafeStringDatasource(attrs.GroupSearchBaseDN),
		PasswordNoOfSplChars:       util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS),
		PasswordNoOfDigits:         util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS),
		StatusKeyJSON:              util.SafeStringDatasource(attrs.STATUSKEYJSON),
		SearchFilter:               util.SafeStringDatasource(attrs.SEARCHFILTER),
		ConfigJSON:                 util.SafeStringDatasource(attrs.ConfigJSON),
		RemoveAccountAction:        util.SafeStringDatasource(attrs.REMOVEACCOUNTACTION),
		AccountAttribute:           util.SafeStringDatasource(attrs.ACCOUNT_ATTRIBUTE),
		AccountNameRule:            util.SafeStringDatasource(attrs.ACCOUNTNAMERULE),
		LDAPOrAD:                   util.SafeStringDatasource(attrs.LDAP_OR_AD),
		EntitlementAttribute:       util.SafeStringDatasource(attrs.ENTITLEMENT_ATTRIBUTE),
		SetRandomPassword:          util.SafeStringDatasource(attrs.SETRANDOMPASSWORD),
		PasswordMinLength:          util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH),
		PasswordMaxLength:          util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH),
		PasswordNoOfCapsAlpha:      util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA),
		SetDefaultPageSize:         util.SafeStringDatasource(attrs.SETDEFAULTPAGESIZE),
		IsTimeoutSupported:         util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		ReuseInactiveAccount:       util.SafeStringDatasource(attrs.REUSEINACTIVEACCOUNT),
		ImportJSON:                 util.SafeStringDatasource(attrs.IMPORTJSON),
		CreateUpdateMappings:       util.SafeStringDatasource(attrs.CreateUpdateMappings),
		AdvanceFilterJSON:          util.SafeStringDatasource(attrs.ADVANCE_FILTER_JSON),
		OrgImportJSON:              util.SafeStringDatasource(attrs.ORGIMPORTJSON),
		PAMConfig:                  util.SafeStringDatasource(attrs.PAM_CONFIG),
		PageSize:                   util.SafeStringDatasource(attrs.PAGE_SIZE),
		Base:                       util.SafeStringDatasource(attrs.BASE),
		DCLocator:                  util.SafeStringDatasource(attrs.DC_LOCATOR),
		StatusThresholdConfig:      util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		ResetAndChangePasswordJSON: util.SafeStringDatasource(attrs.RESETANDCHANGEPASSWRDJSON),
		SupportEmptyString:         util.SafeStringDatasource(attrs.SUPPORTEMPTYSTRING),
		ReadOperationalAttributes:  util.SafeStringDatasource(attrs.READ_OPERATIONAL_ATTRIBUTES),
		EnableAccountJSON:          util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON),
		UserAttribute:              util.SafeStringDatasource(attrs.USER_ATTRIBUTE),
		DefaultUserRole:            util.SafeStringDatasource(attrs.DEFAULT_USER_ROLE),
		EndpointsFilter:            util.SafeStringDatasource(attrs.ENDPOINTS_FILTER),
		UpdateAccountJSON:          util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON),
		ReuseAccountJSON:           util.SafeStringDatasource(attrs.REUSEACCOUNTJSON),
		EnforceTreeDeletion:        util.SafeStringDatasource(attrs.ENFORCE_TREE_DELETION),
		Filter:                     util.SafeStringDatasource(attrs.FILTER),
		ObjectFilter:               util.SafeStringDatasource(attrs.OBJECTFILTER),
		UpdateUserJSON:             util.SafeStringDatasource(attrs.UPDATEUSERJSON),
		SaveConnection:             util.SafeStringDatasource(attrs.Saveconnection),
		SystemName:                 util.SafeStringDatasource(attrs.Systemname),
		GroupImportMapping:         util.SafeStringDatasource(attrs.GroupImportMapping),
		UnlockAccountJSON:          util.SafeStringDatasource(attrs.UNLOCKACCOUNTJSON),
		EnableGroupManagement:      util.SafeStringDatasource(attrs.ENABLEGROUPMANAGEMENT),
		ModifyUserDataJSON:         util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		OrgBase:                    util.SafeStringDatasource(attrs.ORG_BASE),
		OrganizationAttribute:      util.SafeStringDatasource(attrs.ORGANIZATION_ATTRIBUTE),
		CreateOrgJSON:              util.SafeStringDatasource(attrs.CREATEORGJSON),
		UpdateOrgJSON:              util.SafeStringDatasource(attrs.UPDATEORGJSON),
		MaxChangeNumber:            util.SafeStringDatasource(attrs.MAX_CHANGENUMBER),
		IncrementalConfig:          util.SafeStringDatasource(attrs.INCREMENTAL_CONFIG),
		CheckForUnique:             util.SafeStringDatasource(attrs.CHECKFORUNIQUE),
		IsTimeoutConfigValidated:   util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
	}

	// Map connection timeout config if present
	d.MapADTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapADTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *AdConnectionsDataSource) MapADTimeoutConfig(attrs *openapi.ADConnectionAttributes, connectionAttrs *ADConnectionAttributes) {
	if attrs.ConnectionTimeoutConfig != nil {
		connectionAttrs.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
			RetryWait:               util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryWait),
			TokenRefreshMaxTryCount: util.SafeInt64(attrs.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
			RetryWaitMaxValue:       util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryWaitMaxValue),
			RetryCount:              util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryCount),
			ReadTimeout:             util.SafeInt64(attrs.ConnectionTimeoutConfig.ReadTimeout),
			ConnectionTimeout:       util.SafeInt64(attrs.ConnectionTimeoutConfig.ConnectionTimeout),
			RetryFailureStatusCode:  util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryFailureStatusCode),
		}
	}
}

// HandleADAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *AdConnectionsDataSource) HandleADAuthenticationLogic(state *ADConnectionDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all connection attributes")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connection_attributes will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing connection attributes from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connection_attributes will be removed from state.",
			)
			state.ConnectionAttributes = nil
		}
	}
}
