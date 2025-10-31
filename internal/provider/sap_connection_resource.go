// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_sap_connection_resource manages Sap connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Sap connector using the supplied configuration.
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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SapConnectionResource{}
var _ resource.ResourceWithImportState = &SapConnectionResource{}

// Initialize error codes for SAP Connection operations
var sapErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSAP)

type SapConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                             types.String `tfsdk:"id"`
	Messageserver                  types.String `tfsdk:"message_server"`
	JcoAshost                      types.String `tfsdk:"jco_ashost"`
	JcoSysnr                       types.String `tfsdk:"jco_sysnr"`
	JcoClient                      types.String `tfsdk:"jco_client"`
	JcoUser                        types.String `tfsdk:"jco_user"`
	Password                       types.String `tfsdk:"password"`
	PasswordWo                     types.String `tfsdk:"password_wo"`
	JcoLang                        types.String `tfsdk:"jco_lang"`
	JcoR3Name                      types.String `tfsdk:"jco_r3name"`
	JcoMshost                      types.String `tfsdk:"jco_mshost"`
	JcoMsserv                      types.String `tfsdk:"jco_msserv"`
	JcoGroup                       types.String `tfsdk:"jco_group"`
	Snc                            types.String `tfsdk:"snc"`
	JcoSncMode                     types.String `tfsdk:"jco_snc_mode"`
	JcoSncPartnername              types.String `tfsdk:"jco_snc_partnername"`
	JcoSncMyname                   types.String `tfsdk:"jco_snc_myname"`
	JcoSncLibrary                  types.String `tfsdk:"jco_snc_library"`
	JcoSncQop                      types.String `tfsdk:"jco_snc_qop"`
	Tables                         types.String `tfsdk:"tables"`
	Systemname                     types.String `tfsdk:"system_name"`
	Terminatedusergroup            types.String `tfsdk:"terminated_user_group"`
	TerminatedUserRoleAction       types.String `tfsdk:"terminated_user_role_action"`
	Createaccountjson              types.String `tfsdk:"create_account_json"`
	ProvJcoAshost                  types.String `tfsdk:"prov_jco_ashost"`
	ProvJcoSysnr                   types.String `tfsdk:"prov_jco_sysnr"`
	ProvJcoClient                  types.String `tfsdk:"prov_jco_client"`
	ProvJcoUser                    types.String `tfsdk:"prov_jco_user"`
	ProvPassword                   types.String `tfsdk:"prov_password"`
	ProvPasswordWo                 types.String `tfsdk:"prov_password_wo"`
	ProvJcoLang                    types.String `tfsdk:"prov_jco_lang"`
	ProvJcoR3Name                  types.String `tfsdk:"prov_jco_r3name"`
	ProvJcoMshost                  types.String `tfsdk:"prov_jco_mshost"`
	ProvJcoMsserv                  types.String `tfsdk:"prov_jco_msserv"`
	ProvJcoGroup                   types.String `tfsdk:"prov_jco_group"`
	ProvCuaEnabled                 types.String `tfsdk:"prov_cua_enabled"`
	ProvCuaSnc                     types.String `tfsdk:"prov_cua_snc"`
	ResetPwdForNewaccount          types.String `tfsdk:"reset_pwd_for_newaccount"`
	Enforcepasswordchange          types.String `tfsdk:"enforce_password_change"`
	PasswordMinLength              types.String `tfsdk:"password_min_length"`
	PasswordMaxLength              types.String `tfsdk:"password_max_length"`
	PasswordNoofcapsalpha          types.String `tfsdk:"password_no_of_caps_alpha"`
	PasswordNoofdigits             types.String `tfsdk:"password_no_of_digits"`
	PasswordNoofsplchars           types.String `tfsdk:"password_no_of_spl_chars"`
	Hanareftablejson               types.String `tfsdk:"hanareftablejson"`
	Enableaccountjson              types.String `tfsdk:"enable_account_json"`
	Updateaccountjson              types.String `tfsdk:"update_account_json"`
	Userimportjson                 types.String `tfsdk:"user_import_json"`
	StatusThresholdConfig          types.String `tfsdk:"status_threshold_config"`
	Setcuasystem                   types.String `tfsdk:"set_cua_system"`
	FirefighteridGrantAccessJson   types.String `tfsdk:"fire_fighter_id_grant_access_json"`
	FirefighteridRevokeAccessJson  types.String `tfsdk:"fire_fighter_id_revoke_access_json"`
	Modifyuserdatajson             types.String `tfsdk:"modify_user_data_json"`
	ExternalSodEvalJson            types.String `tfsdk:"external_sod_eval_json"`
	ExternalSodEvalJsonDetail      types.String `tfsdk:"external_sod_eval_json_detail"`
	LogsTableFilter                types.String `tfsdk:"logs_table_filter"`
	PamConfig                      types.String `tfsdk:"pam_config"`
	SaptableFilterLang             types.String `tfsdk:"saptable_filter_lang"`
	AlternateOutputParameterEtData types.String `tfsdk:"alternate_output_parameter_et_data"`
	AuditLogJson                   types.String `tfsdk:"audit_log_json"`
	EccOrS4Hana                    types.String `tfsdk:"ecc_or_s4hana"`
	DataImportFilter               types.String `tfsdk:"data_import_filter"`
	Configjson                     types.String `tfsdk:"config_json"`

	//25.B.1
	RoleDefaultDate types.String `tfsdk:"role_default_date"`
}

type SapConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	saviyntVersion    string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewSapConnectionResource() resource.Resource {
	return &SapConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewSapConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &SapConnectionResource{
		connectionFactory: factory,
	}
}

func (r *SapConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_sap_connection_resource"
}

func SapConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"message_server": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Messageserver.",
		},
		"jco_ashost": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcoashost.",
		},
		"jco_sysnr": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosysnr.",
		},
		"jco_client": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcoclient.",
		},
		"jco_user": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcouser.",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Password. Either this or password_wo need to be set to configure the password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password_wo")),
			},
		},
		"password_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Password write-only attribute. Either this or password need to be set to configure the password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password")),
			},
		},
		"jco_lang": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcolang.",
		},
		"jco_r3name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcor3name.",
		},
		"jco_mshost": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcomshost.",
		},
		"jco_msserv": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcomsserv.",
		},
		"jco_group": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcogroup.",
		},
		"snc": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Snc.",
		},
		"jco_snc_mode": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosncmode.",
		},
		"jco_snc_partnername": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosncpartnername.",
		},
		"jco_snc_myname": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosncmyname.",
		},
		"jco_snc_library": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosnclibrary.",
		},
		"jco_snc_qop": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Jcosncqop.",
		},
		"tables": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Tables.",
		},
		"system_name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Systemname.",
		},
		"terminated_user_group": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Terminatedusergroup.",
		},
		"terminated_user_role_action": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Terminateduserroleaction.",
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Createaccountjson.",
		},
		"prov_jco_ashost": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcoashost.",
		},
		"prov_jco_sysnr": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcosysnr.",
		},
		"prov_jco_client": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcoclient.",
		},
		"prov_jco_user": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcouser.",
		},
		"prov_password": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Provpassword. Either this field or the prov_password_wo field must be populated to set the prov_password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("prov_password_wo")),
			},
		},
		"prov_password_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Provpassword write-only attribute. Either this field or the prov_password field must be populated to set the prov_password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("prov_password")),
			},
		},
		"prov_jco_lang": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcolang.",
		},
		"prov_jco_r3name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcor3name.",
		},
		"prov_jco_mshost": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcomshost.",
		},
		"prov_jco_msserv": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcomsserv.",
		},
		"prov_jco_group": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provjcogroup.",
		},
		"prov_cua_enabled": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provcuaenabled.",
		},
		"prov_cua_snc": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provcuasnc.",
		},
		"reset_pwd_for_newaccount": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Resetpwdfornewaccount.",
		},
		"enforce_password_change": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Enforcepasswordchange.",
		},
		"password_min_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Passwordminlength.",
		},
		"password_max_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Passwordmaxlength.",
		},
		"password_no_of_caps_alpha": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Passwordnoofcapsalpha.",
		},
		"password_no_of_digits": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Passwordnoofdigits.",
		},
		"password_no_of_spl_chars": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Passwordnoofsplchars.",
		},
		"hanareftablejson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Hanareftablejson.",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Enableaccountjson.",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Updateaccountjson.",
		},
		"user_import_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Userimportjson.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Statusthresholdconfig.",
		},
		"set_cua_system": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Setcuasystem.",
		},
		"fire_fighter_id_grant_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Firefighteridgrantaccessjson.",
		},
		"fire_fighter_id_revoke_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Firefighteridrevokeaccessjson.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Modifyuserdatajson.",
		},
		"external_sod_eval_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Externalsodevaljson.",
		},
		"external_sod_eval_json_detail": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Externalsodevaljsondetail.",
		},
		"logs_table_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Logstablefilter.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Pamconfig.",
		},
		"saptable_filter_lang": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Saptablefilterlang.",
		},
		"alternate_output_parameter_et_data": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Alternateoutputparameteretdata.",
		},
		"audit_log_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Auditlogjson.",
		},
		"ecc_or_s4hana": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Eccors4hana.",
		},
		"data_import_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Data import filter.",
		},
		"config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Config json.",
		},
		"role_default_date": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Role default date.",
		},
	}
}

func (r *SapConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SAPConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), SapConnectorResourceSchema()),
	}
}

func (r *SapConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "SAP connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := sapErrorCodes.ProviderConfig()
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
	r.saviyntVersion = prov.saviyntVersion
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "SAP connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *SapConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *SapConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *SapConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *SapConnectionResource) CreateSAPConnection(ctx context.Context, plan *SapConnectorResourceModel, config *SapConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SAP connection creation")

	// Check if SAP connection already exists (idempotency check) with retry logic
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
		errorCode := sapErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.SAPConnectionResponse != nil &&
		existingResource.SAPConnectionResponse.Errorcode != nil &&
		*existingResource.SAPConnectionResponse.Errorcode == 0 {

		errorCode := sapErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "create", connectionName, nil)
	}

	// Build SAP connection create request
	tflog.Debug(ctx, "Building SAP connection create request")

	sapConn := r.BuildSAPConnector(plan, config)
	sapConnRequest := openapi.CreateOrUpdateRequest{
		SAPConnector: &sapConn,
	}

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_sap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, sapConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sapErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create SAP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := sapErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "SAP connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SAP connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SapConnectionResource) BuildSAPConnector(plan *SapConnectorResourceModel, config *SapConnectorResourceModel) openapi.SAPConnector {
	var password string
	if !config.Password.IsNull() && !config.Password.IsUnknown() {
		password = config.Password.ValueString()
	} else if !config.PasswordWo.IsNull() && !config.PasswordWo.IsUnknown() {
		password = config.PasswordWo.ValueString()
	}

	var provPassword string
	if !config.ProvPassword.IsNull() && !config.ProvPassword.IsUnknown() {
		provPassword = config.ProvPassword.ValueString()
	} else if !config.ProvPasswordWo.IsNull() && !config.ProvPasswordWo.IsUnknown() {
		provPassword = config.ProvPasswordWo.ValueString()
	}

	sapConn := openapi.SAPConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:  "SAP",
			ConnectionName:  plan.ConnectionName.ValueString(),
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		MESSAGESERVER:                      util.StringPointerOrEmpty(plan.Messageserver),
		JCO_ASHOST:                         util.StringPointerOrEmpty(plan.JcoAshost),
		JCO_SYSNR:                          util.StringPointerOrEmpty(plan.JcoSysnr),
		JCO_CLIENT:                         util.StringPointerOrEmpty(plan.JcoClient),
		JCO_USER:                           util.StringPointerOrEmpty(plan.JcoUser),
		PASSWORD:                           util.StringPointerOrEmpty(types.StringValue(password)),
		JCO_LANG:                           util.StringPointerOrEmpty(plan.JcoLang),
		JCOR3NAME:                          util.StringPointerOrEmpty(plan.JcoR3Name),
		JCO_MSHOST:                         util.StringPointerOrEmpty(plan.JcoMshost),
		JCO_MSSERV:                         util.StringPointerOrEmpty(plan.JcoMsserv),
		JCO_GROUP:                          util.StringPointerOrEmpty(plan.JcoGroup),
		SNC:                                util.StringPointerOrEmpty(plan.Snc),
		JCO_SNC_MODE:                       util.StringPointerOrEmpty(plan.JcoSncMode),
		JCO_SNC_PARTNERNAME:                util.StringPointerOrEmpty(plan.JcoSncPartnername),
		JCO_SNC_MYNAME:                     util.StringPointerOrEmpty(plan.JcoSncMyname),
		JCO_SNC_LIBRARY:                    util.StringPointerOrEmpty(plan.JcoSncLibrary),
		JCO_SNC_QOP:                        util.StringPointerOrEmpty(plan.JcoSncQop),
		TABLES:                             util.StringPointerOrEmpty(plan.Tables),
		SYSTEMNAME:                         util.StringPointerOrEmpty(plan.Systemname),
		TERMINATEDUSERGROUP:                util.StringPointerOrEmpty(plan.Terminatedusergroup),
		TERMINATED_USER_ROLE_ACTION:        util.StringPointerOrEmpty(plan.TerminatedUserRoleAction),
		CREATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Createaccountjson),
		PROV_JCO_ASHOST:                    util.StringPointerOrEmpty(plan.ProvJcoAshost),
		PROV_JCO_SYSNR:                     util.StringPointerOrEmpty(plan.ProvJcoSysnr),
		PROV_JCO_CLIENT:                    util.StringPointerOrEmpty(plan.ProvJcoClient),
		PROV_JCO_USER:                      util.StringPointerOrEmpty(plan.ProvJcoUser),
		PROV_PASSWORD:                      util.StringPointerOrEmpty(types.StringValue(provPassword)),
		PROV_JCO_LANG:                      util.StringPointerOrEmpty(plan.ProvJcoLang),
		PROVJCOR3NAME:                      util.StringPointerOrEmpty(plan.ProvJcoR3Name),
		PROV_JCO_MSHOST:                    util.StringPointerOrEmpty(plan.ProvJcoMshost),
		PROV_JCO_MSSERV:                    util.StringPointerOrEmpty(plan.ProvJcoMsserv),
		PROV_JCO_GROUP:                     util.StringPointerOrEmpty(plan.ProvJcoGroup),
		PROV_CUA_ENABLED:                   util.StringPointerOrEmpty(plan.ProvCuaEnabled),
		PROV_CUA_SNC:                       util.StringPointerOrEmpty(plan.ProvCuaSnc),
		RESET_PWD_FOR_NEWACCOUNT:           util.StringPointerOrEmpty(plan.ResetPwdForNewaccount),
		ENFORCEPASSWORDCHANGE:              util.StringPointerOrEmpty(plan.Enforcepasswordchange),
		PASSWORD_MIN_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMinLength),
		PASSWORD_MAX_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMaxLength),
		PASSWORD_NOOFCAPSALPHA:             util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha),
		PASSWORD_NOOFDIGITS:                util.StringPointerOrEmpty(plan.PasswordNoofdigits),
		PASSWORD_NOOFSPLCHARS:              util.StringPointerOrEmpty(plan.PasswordNoofsplchars),
		HANAREFTABLEJSON:                   util.StringPointerOrEmpty(plan.Hanareftablejson),
		ENABLEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Enableaccountjson),
		UPDATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Updateaccountjson),
		USERIMPORTJSON:                     util.StringPointerOrEmpty(plan.Userimportjson),
		STATUS_THRESHOLD_CONFIG:            util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		SETCUASYSTEM:                       util.StringPointerOrEmpty(plan.Setcuasystem),
		FIREFIGHTERID_GRANT_ACCESS_JSON:    util.StringPointerOrEmpty(plan.FirefighteridGrantAccessJson),
		FIREFIGHTERID_REVOKE_ACCESS_JSON:   util.StringPointerOrEmpty(plan.FirefighteridRevokeAccessJson),
		MODIFYUSERDATAJSON:                 util.StringPointerOrEmpty(plan.Modifyuserdatajson),
		EXTERNAL_SOD_EVAL_JSON:             util.StringPointerOrEmpty(plan.ExternalSodEvalJson),
		EXTERNAL_SOD_EVAL_JSON_DETAIL:      util.StringPointerOrEmpty(plan.ExternalSodEvalJsonDetail),
		LOGS_TABLE_FILTER:                  util.StringPointerOrEmpty(plan.LogsTableFilter),
		PAM_CONFIG:                         util.StringPointerOrEmpty(plan.PamConfig),
		SAPTABLE_FILTER_LANG:               util.StringPointerOrEmpty(plan.SaptableFilterLang),
		ALTERNATE_OUTPUT_PARAMETER_ET_DATA: util.StringPointerOrEmpty(plan.AlternateOutputParameterEtData),
		AUDIT_LOG_JSON:                     util.StringPointerOrEmpty(plan.AuditLogJson),
		ECCORS4HANA:                        util.StringPointerOrEmpty(plan.EccOrS4Hana),
		DATA_IMPORT_FILTER:                 util.StringPointerOrEmpty(plan.DataImportFilter),
		ConfigJSON:                         util.StringPointerOrEmpty(plan.Configjson),

		ROLE_DEFAULT_DATE: util.StringPointerOrEmpty(plan.RoleDefaultDate),
	}

	if plan.VaultConnection.ValueString() != "" {
		sapConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		sapConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		sapConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return sapConn
}

func (r *SapConnectionResource) UpdateModelFromCreateResponse(plan *SapConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))

	// Update all optional fields to maintain state
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.Messageserver = util.SafeStringDatasource(plan.Messageserver.ValueStringPointer())
	plan.JcoAshost = util.SafeStringDatasource(plan.JcoAshost.ValueStringPointer())
	plan.JcoSysnr = util.SafeStringDatasource(plan.JcoSysnr.ValueStringPointer())
	plan.JcoClient = util.SafeStringDatasource(plan.JcoClient.ValueStringPointer())
	plan.JcoUser = util.SafeStringDatasource(plan.JcoUser.ValueStringPointer())
	plan.JcoLang = util.SafeStringDatasource(plan.JcoLang.ValueStringPointer())
	plan.JcoR3Name = util.SafeStringDatasource(plan.JcoR3Name.ValueStringPointer())
	plan.JcoMshost = util.SafeStringDatasource(plan.JcoMshost.ValueStringPointer())
	plan.JcoMsserv = util.SafeStringDatasource(plan.JcoMsserv.ValueStringPointer())
	plan.JcoGroup = util.SafeStringDatasource(plan.JcoGroup.ValueStringPointer())
	plan.Snc = util.SafeStringDatasource(plan.Snc.ValueStringPointer())
	plan.JcoSncMode = util.SafeStringDatasource(plan.JcoSncMode.ValueStringPointer())
	plan.JcoSncPartnername = util.SafeStringDatasource(plan.JcoSncPartnername.ValueStringPointer())
	plan.JcoSncMyname = util.SafeStringDatasource(plan.JcoSncMyname.ValueStringPointer())
	plan.JcoSncLibrary = util.SafeStringDatasource(plan.JcoSncLibrary.ValueStringPointer())
	plan.JcoSncQop = util.SafeStringDatasource(plan.JcoSncQop.ValueStringPointer())
	plan.Tables = util.SafeStringDatasource(plan.Tables.ValueStringPointer())
	plan.Systemname = util.SafeStringDatasource(plan.Systemname.ValueStringPointer())
	plan.Terminatedusergroup = util.SafeStringDatasource(plan.Terminatedusergroup.ValueStringPointer())
	plan.TerminatedUserRoleAction = util.SafeStringDatasource(plan.TerminatedUserRoleAction.ValueStringPointer())
	plan.Createaccountjson = util.SafeStringDatasource(plan.Createaccountjson.ValueStringPointer())
	plan.ProvJcoAshost = util.SafeStringDatasource(plan.ProvJcoAshost.ValueStringPointer())
	plan.ProvJcoSysnr = util.SafeStringDatasource(plan.ProvJcoSysnr.ValueStringPointer())
	plan.ProvJcoClient = util.SafeStringDatasource(plan.ProvJcoClient.ValueStringPointer())
	plan.ProvJcoUser = util.SafeStringDatasource(plan.ProvJcoUser.ValueStringPointer())
	plan.ProvJcoLang = util.SafeStringDatasource(plan.ProvJcoLang.ValueStringPointer())
	plan.ProvJcoR3Name = util.SafeStringDatasource(plan.ProvJcoR3Name.ValueStringPointer())
	plan.ProvJcoMshost = util.SafeStringDatasource(plan.ProvJcoMshost.ValueStringPointer())
	plan.ProvJcoMsserv = util.SafeStringDatasource(plan.ProvJcoMsserv.ValueStringPointer())
	plan.ProvJcoGroup = util.SafeStringDatasource(plan.ProvJcoGroup.ValueStringPointer())
	plan.ProvCuaEnabled = util.SafeStringDatasource(plan.ProvCuaEnabled.ValueStringPointer())
	plan.ProvCuaSnc = util.SafeStringDatasource(plan.ProvCuaSnc.ValueStringPointer())
	plan.ResetPwdForNewaccount = util.SafeStringDatasource(plan.ResetPwdForNewaccount.ValueStringPointer())
	plan.Enforcepasswordchange = util.SafeStringDatasource(plan.Enforcepasswordchange.ValueStringPointer())
	plan.PasswordMinLength = util.SafeStringDatasource(plan.PasswordMinLength.ValueStringPointer())
	plan.PasswordMaxLength = util.SafeStringDatasource(plan.PasswordMaxLength.ValueStringPointer())
	plan.PasswordNoofcapsalpha = util.SafeStringDatasource(plan.PasswordNoofcapsalpha.ValueStringPointer())
	plan.PasswordNoofdigits = util.SafeStringDatasource(plan.PasswordNoofdigits.ValueStringPointer())
	plan.PasswordNoofsplchars = util.SafeStringDatasource(plan.PasswordNoofsplchars.ValueStringPointer())
	plan.Hanareftablejson = util.SafeStringDatasource(plan.Hanareftablejson.ValueStringPointer())
	plan.Enableaccountjson = util.SafeStringDatasource(plan.Enableaccountjson.ValueStringPointer())
	plan.Updateaccountjson = util.SafeStringDatasource(plan.Updateaccountjson.ValueStringPointer())
	plan.Userimportjson = util.SafeStringDatasource(plan.Userimportjson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.Setcuasystem = util.SafeStringDatasource(plan.Setcuasystem.ValueStringPointer())
	plan.FirefighteridGrantAccessJson = util.SafeStringDatasource(plan.FirefighteridGrantAccessJson.ValueStringPointer())
	plan.FirefighteridRevokeAccessJson = util.SafeStringDatasource(plan.FirefighteridRevokeAccessJson.ValueStringPointer())
	plan.Modifyuserdatajson = util.SafeStringDatasource(plan.Modifyuserdatajson.ValueStringPointer())
	plan.ExternalSodEvalJson = util.SafeStringDatasource(plan.ExternalSodEvalJson.ValueStringPointer())
	plan.ExternalSodEvalJsonDetail = util.SafeStringDatasource(plan.ExternalSodEvalJsonDetail.ValueStringPointer())
	plan.LogsTableFilter = util.SafeStringDatasource(plan.LogsTableFilter.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.SaptableFilterLang = util.SafeStringDatasource(plan.SaptableFilterLang.ValueStringPointer())
	plan.AlternateOutputParameterEtData = util.SafeStringDatasource(plan.AlternateOutputParameterEtData.ValueStringPointer())
	plan.AuditLogJson = util.SafeStringDatasource(plan.AuditLogJson.ValueStringPointer())
	plan.EccOrS4Hana = util.SafeStringDatasource(plan.EccOrS4Hana.ValueStringPointer())
	plan.DataImportFilter = util.SafeStringDatasource(plan.DataImportFilter.ValueStringPointer())
	plan.Configjson = util.SafeStringDatasource(plan.Configjson.ValueStringPointer())

	plan.RoleDefaultDate = util.SafeStringDatasource(plan.RoleDefaultDate.ValueStringPointer())

	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *SapConnectionResource) ReadSAPConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SAP connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_sap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sapErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read SAP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateSAPConnectionResponse(apiResp); err != nil {
		errorCode := sapErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for SAP datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.SAPConnectionResponse != nil && apiResp.SAPConnectionResponse.Errorcode != nil && *apiResp.SAPConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.SAPConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.SAPConnectionResponse.Msg))
		errorCode := sapErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "SAP connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.SAPConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.SAPConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SAP connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.SAPConnectionResponse != nil && apiResp.SAPConnectionResponse.Connectionkey != nil {
				return *apiResp.SAPConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SapConnectionResource) UpdateModelFromReadResponse(state *SapConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.SAPConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.SAPConnectionResponse.Connectionkey))

	// Update all fields from API response
	state.ConnectionName = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Emailtemplate)
	state.Createaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.AuditLogJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.AUDIT_LOG_JSON)
	state.SaptableFilterLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SAPTABLE_FILTER_LANG)
	state.PasswordNoofsplchars = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	state.Terminatedusergroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATEDUSERGROUP)
	state.LogsTableFilter = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.LOGS_TABLE_FILTER)
	state.EccOrS4Hana = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ECCORS4HANA)
	state.FirefighteridRevokeAccessJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_JSON)
	state.Configjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ConfigJSON)
	state.FirefighteridGrantAccessJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_JSON)
	state.JcoSncLibrary = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_LIBRARY)
	state.JcoR3Name = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCOR3NAME)
	state.ExternalSodEvalJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON)
	state.JcoAshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_ASHOST)
	state.PasswordNoofdigits = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	state.ProvJcoMshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSHOST)
	state.PamConfig = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.JcoSncMyname = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MYNAME)
	state.Enforcepasswordchange = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENFORCEPASSWORDCHANGE)
	state.JcoUser = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_USER)
	state.JcoSncMode = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MODE)
	state.ProvJcoMsserv = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSSERV)
	state.Hanareftablejson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.HANAREFTABLEJSON)
	state.PasswordMinLength = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	state.JcoClient = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_CLIENT)
	state.TerminatedUserRoleAction = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATED_USER_ROLE_ACTION)
	state.ResetPwdForNewaccount = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.RESET_PWD_FOR_NEWACCOUNT)
	state.ProvJcoClient = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_CLIENT)
	state.Snc = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SNC)
	state.JcoMsserv = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSSERV)
	state.ProvCuaSnc = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_SNC)
	state.ProvJcoUser = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_USER)
	state.JcoLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_LANG)
	state.JcoSncPartnername = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_PARTNERNAME)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.ProvJcoSysnr = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_SYSNR)
	state.Setcuasystem = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SETCUASYSTEM)
	state.Messageserver = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MESSAGESERVER)
	state.ProvJcoAshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_ASHOST)
	state.ProvJcoGroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_GROUP)
	state.ProvCuaEnabled = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_ENABLED)
	state.JcoMshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSHOST)
	state.ProvJcoR3Name = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROVJCOR3NAME)
	state.PasswordNoofcapsalpha = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	state.Modifyuserdatajson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.JcoSncQop = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_QOP)
	state.Tables = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TABLES)
	state.ProvJcoLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_LANG)
	state.JcoSysnr = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SYSNR)
	state.ExternalSodEvalJsonDetail = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON_DETAIL)
	state.DataImportFilter = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.DATA_IMPORT_FILTER)
	state.Enableaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.AlternateOutputParameterEtData = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ALTERNATE_OUTPUT_PARAMETER_ET_DATA)
	state.JcoGroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_GROUP)
	state.PasswordMaxLength = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	state.Userimportjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.USERIMPORTJSON)
	state.Systemname = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SYSTEMNAME)
	state.Updateaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)

	state.RoleDefaultDate = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ROLE_DEFAULT_DATE)
}

func (r *SapConnectionResource) ValidateSAPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SAPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - SAP connection response is nil")
	}
	return nil
}

func (r *SapConnectionResource) UpdateSAPConnection(ctx context.Context, plan *SapConnectorResourceModel, config *SapConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SAP connection update")

	// Build SAP connection update request
	tflog.Debug(logCtx, "Building SAP connection update request")

	sapConn := r.BuildSAPConnector(plan, config)

	sapConnRequest := openapi.CreateOrUpdateRequest{
		SAPConnector: &sapConn,
	}

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_sap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, sapConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sapErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update SAP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := sapErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "SAP connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SAP connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SapConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config SapConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := sapErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	util.ValidateAttributeCompatibility(r.saviyntVersion, "SAP", "RoleDefaultDate", plan.RoleDefaultDate.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := sapErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateSAPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "SAP connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"SAP Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "SAP connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *SapConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SapConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := sapErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadSAPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "SAP connection read failed", "", err)
		resp.Diagnostics.AddError(
			"SAP Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	apiMessage := util.SafeDeref(apiResp.SAPConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.SAPConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := sapErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SAP connection resource read completed successfully")
}

func (r *SapConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config SapConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := sapErrorCodes.StateExtraction()
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
		errorCode := sapErrorCodes.PlanExtraction()
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
		errorCode := sapErrorCodes.ConfigExtraction()
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
		errorCode := sapErrorCodes.NameImmutable()
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

	util.ValidateAttributeCompatibility(r.saviyntVersion, "SAP", "RoleDefaultDate", plan.RoleDefaultDate.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	updateResp, err := r.UpdateSAPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "SAP connection update failed", "", err)
		resp.Diagnostics.AddError(
			"SAP Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadSAPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated SAP connection", "", err)
		resp.Diagnostics.AddError(
			"SAP Connection Post-Update Read Failed",
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
		errorCode := sapErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SAP connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *SapConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *SapConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing a SAP connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "SAP connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
