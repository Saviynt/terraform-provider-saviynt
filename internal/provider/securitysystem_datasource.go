// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_security_system_datasource retrieves security system details from the Saviynt Security Manager.
// The data source supports filtering and pagination with comprehensive error handling and authentication control.
package provider

import (
	"context"
	"fmt"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

var _ datasource.DataSource = &securitySystemsDataSource{}

// Initialize error codes for Security System datasource operations
var securitySystemDatasourceErrorCodes = errorsutil.NewSecuritySystemErrorCodeGenerator()

// SecuritySystemsDataSource defines the data source
type securitySystemsDataSource struct {
	client                client.SaviyntClientInterface
	token                 string
	provider              client.SaviyntProviderInterface
	securitySystemFactory client.SecuritySystemFactoryInterface
}

type SecuritySystemsDataSourceModel struct {
	ID             types.String            `tfsdk:"id"`
	Systemname     types.String            `tfsdk:"systemname"`
	Authenticate   types.Bool              `tfsdk:"authenticate"`
	Max            types.Int64             `tfsdk:"max"`
	Offset         types.Int64             `tfsdk:"offset"`
	Connectionname types.String            `tfsdk:"connectionname"`
	ConnectionType types.String            `tfsdk:"connection_type"`
	Msg            types.String            `tfsdk:"msg"`
	DisplayCount   types.Int64             `tfsdk:"display_count"`
	ErrorCode      types.String            `tfsdk:"error_code"`
	TotalCount     types.Int64             `tfsdk:"total_count"`
	Results        []SecuritySystemDetails `tfsdk:"results"`
}

// SecuritySystemDetails represents a single security system details object.
type SecuritySystemDetails struct {
	Systemname                         types.String   `tfsdk:"systemname"`
	ConnectionType                     types.String   `tfsdk:"connection_type"`
	DisplayName                        types.String   `tfsdk:"display_name"`
	Hostname                           types.String   `tfsdk:"hostname"`
	Port                               types.String   `tfsdk:"port"`
	AccessAddWorkflow                  types.String   `tfsdk:"access_add_workflow"`
	AccessRemoveWorkflow               types.String   `tfsdk:"access_remove_workflow"`
	AddServiceAccountWorkflow          types.String   `tfsdk:"add_service_account_workflow"`
	RemoveServiceAccountWorkflow       types.String   `tfsdk:"remove_service_account_workflow"`
	Connectionparameters               types.String   `tfsdk:"connection_parameters"`
	AutomatedProvisioning              types.String   `tfsdk:"automated_provisioning"`
	UseOpenConnector                   types.String   `tfsdk:"use_open_connector"`
	ManageEntity                       types.String   `tfsdk:"manage_entity"`
	PersistentData                     types.String   `tfsdk:"persistent_data"`
	DefaultSystem                      types.String   `tfsdk:"default_system"`
	ReconApplication                   types.String   `tfsdk:"recon_application"`
	InstantProvision                   types.String   `tfsdk:"instant_provision"`
	ProvisioningTries                  types.String   `tfsdk:"provisioning_tries"`
	ProvisioningComments               types.String   `tfsdk:"provisioning_comments"`
	ProposedAccountOwnersWorkflow      types.String   `tfsdk:"proposed_account_owners_workflow"`
	FirefighterIDWorkflow              types.String   `tfsdk:"firefighterid_workflow"`
	FirefighterIDRequestAccessWorkflow types.String   `tfsdk:"firefighterid_request_access_workflow"`
	PolicyRule                         types.String   `tfsdk:"policy_rule"`
	PolicyRuleServiceAccount           types.String   `tfsdk:"policy_rule_service_account"`
	Connectionname                     types.String   `tfsdk:"connectionname1"`
	ProvisioningConnection             types.String   `tfsdk:"provisioning_connection"`
	ServiceDeskConnection              types.String   `tfsdk:"service_desk_connection"`
	ExternalRiskConnectionJson         types.String   `tfsdk:"external_risk_connection_json"`
	Connection                         types.String   `tfsdk:"connection"`
	CreateDate                         types.String   `tfsdk:"create_date"`
	UpdateDate                         types.String   `tfsdk:"update_date"`
	Endpoints                          types.String   `tfsdk:"endpoints"`
	CreatedBy                          types.String   `tfsdk:"created_by"`
	UpdatedBy                          types.String   `tfsdk:"updated_by"`
	Status                             types.String   `tfsdk:"status"`
	CreatedFrom                        types.String   `tfsdk:"created_from"`
	InherentSodReportFields            []types.String `tfsdk:"inherent_sod_report_fields"`
}

var _ datasource.DataSource = &securitySystemsDataSource{}

// NewSecuritySystemsDataSource creates a new security systems data source with default factory
func NewSecuritySystemsDataSource() datasource.DataSource {
	return &securitySystemsDataSource{
		securitySystemFactory: &client.DefaultSecuritySystemFactory{},
	}
}

// NewSecuritySystemsDataSourceWithFactory creates a new security systems data source with custom factory
// Used primarily for testing with mock factories
func NewSecuritySystemsDataSourceWithFactory(factory client.SecuritySystemFactoryInterface) datasource.DataSource {
	return &securitySystemsDataSource{
		securitySystemFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *securitySystemsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *securitySystemsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *securitySystemsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

func (d *securitySystemsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_security_system_datasource"
}

func (d *securitySystemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SecuritySystemDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Resource ID.",
			},
			"systemname": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the security systeme.",
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "If false, do not store security system details in state",
			},
			"max": schema.Int64Attribute{
				Optional:    true,
				Description: "Maximum number of security systems to return in the response.",
			},
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: "Number of security systems to skip before returning results (for pagination).",
			},
			"connectionname": schema.StringAttribute{
				Optional:    true,
				Description: "Filter security systems by connection name.",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Filter security systems by connection type (e.g., AD, REST, DB).",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of items currently displayed (e.g., on the current page or view).",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of items available in the dataset, irrespective of the current display settings.",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of security systems retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_name":                          schema.StringAttribute{Computed: true, Description: "Specify a user-friendly display name that is shown on the the user interface."},
						"hostname":                              schema.StringAttribute{Computed: true, Description: "Security system for which you want to create an endpoint."},
						"connection_type":                       schema.StringAttribute{Computed: true, Description: "Specify a connection type to view all connections in EIC for the connection type."},
						"systemname":                            schema.StringAttribute{Computed: true, Description: "Specify the security system name."},
						"access_add_workflow":                   schema.StringAttribute{Computed: true, Description: "Specify the workflow used for approvals for an access request (account, entitlements, role, etc.)."},
						"access_remove_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow used when revoking access from accounts, entitlements, or performing other de-provisioning tasks."},
						"add_service_account_workflow":          schema.StringAttribute{Computed: true, Description: "Workflow for adding a service account."},
						"remove_service_account_workflow":       schema.StringAttribute{Computed: true, Description: "Workflow for removing a service account."},
						"connection_parameters":                 schema.StringAttribute{Computed: true, Description: "Query or parameters to restrict endpoint access to specific users."},
						"automated_provisioning":                schema.StringAttribute{Computed: true, Description: "Enables automated provisioning if set to true."},
						"use_open_connector":                    schema.StringAttribute{Computed: true, Description: "Enables connectivity using open-source connectors such as REST if set to true."},
						"manage_entity":                         schema.StringAttribute{Computed: true, Description: "Indicates if entity management is enabled for the security system."},
						"persistent_data":                       schema.StringAttribute{Computed: true, Description: "Indicates whether persistent data storage is enabled for the security system."},
						"default_system":                        schema.StringAttribute{Computed: true, Description: "Sets this security system as the default system for account searches when set to true."},
						"recon_application":                     schema.StringAttribute{Computed: true, Description: "Enables importing data from endpoints associated with the security system."},
						"instant_provision":                     schema.StringAttribute{Computed: true, Description: "Prevents users from submitting duplicate provisioning requests if set to true."},
						"provisioning_tries":                    schema.StringAttribute{Computed: true, Description: "Number of attempts allowed for provisioning actions."},
						"provisioning_comments":                 schema.StringAttribute{Computed: true, Description: "Comments relevant to provisioning actions."},
						"proposed_account_owners_workflow":      schema.StringAttribute{Computed: true, Description: "Workflow for assigning proposed account owners."},
						"firefighterid_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow for handling firefighter ID requests."},
						"firefighterid_request_access_workflow": schema.StringAttribute{Computed: true, Description: "Workflow for requesting access to firefighter IDs."},
						"policy_rule":                           schema.StringAttribute{Computed: true, Description: "Password policy assigned for the security system."},
						"policy_rule_service_account":           schema.StringAttribute{Computed: true, Description: "Password policy applied to service accounts."},
						"connectionname1":                       schema.StringAttribute{Computed: true, Description: "Name of connection used for reconciling identity objects from third-party applications."},
						"provisioning_connection":               schema.StringAttribute{Computed: true, Description: "Dedicated connection for provisioning and de-provisioning tasks."},
						"service_desk_connection":               schema.StringAttribute{Computed: true, Description: "Connection to service desk or ticketing system integration."},
						"external_risk_connection_json":         schema.StringAttribute{Computed: true, Description: "JSON configuration for external risk connections (e.g., SAP)."},
						"connection":                            schema.StringAttribute{Computed: true, Description: "Primary connection used by the security system."},
						"create_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating when the security system was created."},
						"update_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating the last update to the security system."},
						"endpoints":                             schema.StringAttribute{Computed: true, Description: "Endpoints associated with the security system."},
						"created_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who created the security system."},
						"updated_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who last updated the security system."},
						"status":                                schema.StringAttribute{Computed: true, Description: "Current status of the security system (e.g., enabled, disabled)."},
						"created_from":                          schema.StringAttribute{Computed: true, Description: "Origin or method through which the security system was created."},
						"port":                                  schema.StringAttribute{Computed: true, Description: "Port information or description for the endpoint."},
						"inherent_sod_report_fields": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "List of fields used in filtering Segregation of Duties (SOD) reports.",
						},
					},
				},
			},
		},
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *securitySystemsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateSecuritySystemOperationContext("configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting security systems datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Security systems datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryConfiguration, 1)
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetSecuritySystemErrorMessage(errorsutil.ErrSecuritySystemProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	opCtx.LogOperationEnd(ctx, "Security systems datasource configured successfully")
}

// Read retrieves security systems from Saviynt and populates the Terraform state
// Supports filtering and pagination with comprehensive error handling
func (d *securitySystemsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SecuritySystemsDataSourceModel

	opCtx := errorsutil.CreateSecuritySystemOperationContext("datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting security systems datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryConfiguration, 2)
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetSecuritySystemErrorMessage(errorsutil.ErrSecuritySystemConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request", errorCode),
		)
		return
	}

	// Execute API call to get security systems
	apiResp, err := d.ReadSecuritySystems(ctx, &state)
	if err != nil {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryAPIOperation, 1)
		opCtx.LogOperationError(ctx, "Failed to read security systems", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetSecuritySystemErrorMessage(errorsutil.ErrSecuritySystemReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateSecuritySystemsResponse(apiResp); err != nil {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryAPIOperation, 2)
		opCtx.LogOperationError(ctx, "Invalid API response", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetSecuritySystemErrorMessage(errorsutil.ErrSecuritySystemAPIError),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromSecuritySystemsResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleSecuritySystemsAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryStateManagement, 1)
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetSecuritySystemErrorMessage(errorsutil.ErrSecuritySystemStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for security systems datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Security systems datasource read completed successfully",
		map[string]interface{}{
			"result_count": len(state.Results),
		})
}

// ReadSecuritySystems retrieves security systems from Saviynt API
// Handles filtering and pagination using factory pattern with proper error handling
func (d *securitySystemsDataSource) ReadSecuritySystems(ctx context.Context, state *SecuritySystemsDataSourceModel) (*openapi.GetSecuritySystems200Response, error) {
	opCtx := errorsutil.CreateSecuritySystemOperationContext("api_read", "")
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting security systems API call")

	tflog.Debug(logCtx, "Executing API request to get security systems")

	var apiResp *openapi.GetSecuritySystems200Response

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_security_systems_datasource", func(token string) error {
		securitySystemOps := d.securitySystemFactory.CreateSecuritySystemOperations(d.client.APIBaseURL(), token)

		// Create API request with flexible parameter handling
		apiReq := securitySystemOps.GetSecuritySystemsRequest(ctx)

		// Apply all user input filters
		if !state.Systemname.IsNull() && state.Systemname.ValueString() != "" {
			apiReq = apiReq.Systemname(state.Systemname.ValueString())
		}
		if !state.Max.IsNull() {
			apiReq = apiReq.Max(int32(state.Max.ValueInt64()))
		}
		if !state.Offset.IsNull() {
			apiReq = apiReq.Offset(int32(state.Offset.ValueInt64()))
		}
		if !state.Connectionname.IsNull() && state.Connectionname.ValueString() != "" {
			apiReq = apiReq.Connectionname(state.Connectionname.ValueString())
		}
		if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
			apiReq = apiReq.ConnectionType(state.ConnectionType.ValueString())
		}

		resp, httpResp, err := apiReq.Execute()
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := securitySystemDatasourceErrorCodes.GenerateErrorCode(errorsutil.SecuritySystemCategoryAPIOperation, 3)
		opCtx.LogOperationError(logCtx, "Failed to read security systems", errorCode, err)
		return nil, errorsutil.CreateSecuritySystemStandardError(errorsutil.ErrSecuritySystemReadFailed, "api_read", "", err)
	}

	opCtx.LogOperationEnd(logCtx, "Security systems API call completed successfully")

	return apiResp, nil
}

// ValidateSecuritySystemsResponse validates that the API response contains valid data
// Returns standardized error if validation fails
func (d *securitySystemsDataSource) ValidateSecuritySystemsResponse(apiResp *openapi.GetSecuritySystems200Response) error {
	if apiResp == nil {
		return fmt.Errorf("API response is nil")
	}
	return nil
}

// UpdateModelFromSecuritySystemsResponse maps API response data to the Terraform state model
// Handles both metadata and security system details mapping
func (d *securitySystemsDataSource) UpdateModelFromSecuritySystemsResponse(state *SecuritySystemsDataSourceModel, apiResp *openapi.GetSecuritySystems200Response) {
	// Set resource ID
	state.ID = types.StringValue("ds-security-systems")

	// Map response metadata
	state.Msg = util.SafeStringDatasource(apiResp.Msg)
	state.ErrorCode = util.SafeStringDatasource(apiResp.ErrorCode)

	if apiResp.DisplayCount != nil {
		state.DisplayCount = types.Int64Value(int64(*apiResp.DisplayCount))
	} else {
		state.DisplayCount = types.Int64Value(0)
	}

	if apiResp.TotalCount != nil {
		state.TotalCount = types.Int64Value(int64(*apiResp.TotalCount))
	} else {
		state.TotalCount = types.Int64Value(0)
	}

	// Map security systems details
	if apiResp.SecuritySystemDetails != nil {
		state.Results = make([]SecuritySystemDetails, 0, len(apiResp.SecuritySystemDetails))
		for _, item := range apiResp.SecuritySystemDetails {
			securitySystemState := SecuritySystemDetails{
				ConnectionType:                     util.SafeString(item.ConnectionType),
				Connectionname:                     util.SafeString(item.Connectionname),
				DefaultSystem:                      util.SafeString(item.DefaultSystem),
				DisplayName:                        util.SafeString(item.DisplayName),
				Hostname:                           util.SafeString(item.Hostname),
				Systemname:                         util.SafeString(item.Systemname),
				AccessAddWorkflow:                  util.SafeString(item.AccessAddWorkflow),
				AccessRemoveWorkflow:               util.SafeString(item.AccessRemoveWorkflow),
				AddServiceAccountWorkflow:          util.SafeString(item.AddServiceAccountWorkflow),
				RemoveServiceAccountWorkflow:       util.SafeString(item.RemoveServiceAccountWorkflow),
				Connectionparameters:               util.SafeString(item.Connectionparameters),
				ProvisioningConnection:             util.SafeString(item.ProvisioningConnection),
				Connection:                         util.SafeString(item.Connection),
				CreateDate:                         util.SafeString(item.CreateDate),
				UpdateDate:                         util.SafeString(item.UpdateDate),
				Endpoints:                          util.SafeString(item.Endpoints),
				UseOpenConnector:                   util.SafeString(item.Useopenconnector),
				ReconApplication:                   util.SafeString(item.ReconApplication),
				AutomatedProvisioning:              util.SafeString(item.AutomatedProvisioning),
				InstantProvision:                   util.SafeString(item.Instantprovision),
				ProvisioningComments:               util.SafeString(item.Provisioningcomments),
				ProvisioningTries:                  util.SafeString(item.ProvisioningTries),
				ProposedAccountOwnersWorkflow:      util.SafeString(item.ProposedAccountOwnersworkflow),
				FirefighterIDWorkflow:              util.SafeString(item.FirefighteridWorkflow),
				FirefighterIDRequestAccessWorkflow: util.SafeString(item.FirefighteridRequestAccessWorkflow),
				PolicyRuleServiceAccount:           util.SafeString(item.PolicyRuleServiceAccount),
				ServiceDeskConnection:              util.SafeString(item.ServiceDeskConnection),
				ExternalRiskConnectionJson:         util.SafeString(item.ExternalRiskConnectionJson),
				CreatedBy:                          util.SafeString(item.CreatedBy),
				UpdatedBy:                          util.SafeString(item.UpdatedBy),
				Status:                             util.SafeString(item.Status),
				CreatedFrom:                        util.SafeString(item.CreatedFrom),
				PolicyRule:                         util.SafeString(item.PolicyRule),
				Port:                               util.SafeString(item.Port),
				InherentSodReportFields:            util.StringsToTypeStrings(item.InherentSODReportFields),
			}
			state.Results = append(state.Results, securitySystemState)
		}
	}
}

// HandleSecuritySystemsAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all results are returned in state
func (d *securitySystemsDataSource) HandleSecuritySystemsAuthenticationLogic(state *SecuritySystemsDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all security system details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all security system details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing security system details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; security system details will be removed from state.",
			)
			state.Results = nil
		}
	}
}
