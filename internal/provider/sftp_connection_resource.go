// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_sftp_connection_resource manages SFTP connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new SFTP connector using the supplied configuration.
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
var _ resource.Resource = &SftpConnectionResource{}
var _ resource.ResourceWithImportState = &SftpConnectionResource{}

// Initialize error codes for SFTP Connection operations
var sftpErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSFTP)

type SFTPConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                    types.String `tfsdk:"id"`
	HostName              types.String `tfsdk:"host_name"`
	PortNumber            types.String `tfsdk:"port_number"`
	Username              types.String `tfsdk:"username"`
	AuthCredentialType    types.String `tfsdk:"auth_credential_type"`
	AuthCredentialValue   types.String `tfsdk:"auth_credential_value"`
	AuthCredentialValueWo types.String `tfsdk:"auth_credential_value_wo"`
	Passphrase            types.String `tfsdk:"passphrase"`
	PassphraseWo          types.String `tfsdk:"passphrase_wo"`
	FilesToGet            types.String `tfsdk:"files_to_get"`
	FilesToPut            types.String `tfsdk:"files_to_put"`
	PamConfig             types.String `tfsdk:"pam_config"`
}

type SftpConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewSFTPConnectionResource() resource.Resource {
	return &SftpConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewSFTPConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &SftpConnectionResource{
		connectionFactory: factory,
	}
}

func (r *SftpConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_sftp_connection_resource"
}

func SFTPConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"host_name": schema.StringAttribute{
			Required:    true,
			Description: "SFTP server hostname or IP address. Example: \"sftp.example.com\"",
		},
		"port_number": schema.StringAttribute{
			Required:    true,
			Description: "SFTP server port number. Default is 22. Example: \"22\"",
		},
		"username": schema.StringAttribute{
			Required:    true,
			Description: "Username for SFTP authentication. Example: \"sftpuser\"",
		},
		"auth_credential_type": schema.StringAttribute{
			Required:    true,
			Description: "Type of authentication (password, key, etc.). Example: \"password\"",
		},
		"auth_credential_value": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Authentication credential (password or private key path).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("auth_credential_value_wo")),
			},
		},
		"auth_credential_value_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Authentication credential (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("auth_credential_value")),
			},
		},
		"passphrase": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Passphrase for encrypted private key.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("passphrase_wo")),
			},
		},
		"passphrase_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Passphrase for encrypted private key (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("passphrase")),
			},
		},
		"files_to_get": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Files to download from SFTP server. Example: \"*.csv\"",
		},
		"files_to_put": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Files to upload to SFTP server. Example: \"upload/*.txt\"",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "PAM configuration for SFTP connection.",
		},
	}
}

func (r *SftpConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SFTPConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), SFTPConnectorResourceSchema()),
	}
}

func (r *SftpConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "SFTP connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := sftpErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client, token, and provider reference from the provider state
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "SFTP connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *SftpConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *SftpConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *SftpConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *SftpConnectionResource) BuildSFTPConnector(plan *SFTPConnectorResourceModel, config *SFTPConnectorResourceModel) openapi.SFTPConnector {
	var authCredentialValue string
	if !config.AuthCredentialValue.IsNull() && !config.AuthCredentialValue.IsUnknown() {
		authCredentialValue = config.AuthCredentialValue.ValueString()
	} else if !config.AuthCredentialValueWo.IsNull() && !config.AuthCredentialValueWo.IsUnknown() {
		authCredentialValue = config.AuthCredentialValueWo.ValueString()
	}

	var passphrase string
	if !config.Passphrase.IsNull() && !config.Passphrase.IsUnknown() {
		passphrase = config.Passphrase.ValueString()
	} else if !config.PassphraseWo.IsNull() && !config.PassphraseWo.IsUnknown() {
		passphrase = config.PassphraseWo.ValueString()
	}

	sftpConn := openapi.SFTPConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "SFTP",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			ConnectionDescription: plan.Description.ValueStringPointer(),
			Defaultsavroles:       util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:         util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required fields
		HOST_NAME:             plan.HostName.ValueString(),
		PORT_NUMBER:           plan.PortNumber.ValueString(),
		USERNAME:              plan.Username.ValueString(),
		AUTH_CREDENTIAL_TYPE:  plan.AuthCredentialType.ValueString(),
		AUTH_CREDENTIAL_VALUE: authCredentialValue,
		//optional fields
		PASSPHRASE:   util.StringPointerOrEmpty(types.StringValue(passphrase)),
		FILES_TO_GET: util.StringPointerOrEmpty(plan.FilesToGet),
		FILES_TO_PUT: util.StringPointerOrEmpty(plan.FilesToPut),
		PAM_CONFIG:   util.StringPointerOrEmpty(plan.PamConfig),
	}

	if plan.VaultConnection.ValueString() != "" {
		sftpConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		sftpConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		sftpConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return sftpConn
}

func (r *SftpConnectionResource) UpdateModelFromCreateResponse(plan *SFTPConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))

	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())

	plan.FilesToGet = util.SafeStringDatasource(plan.FilesToGet.ValueStringPointer())
	plan.FilesToPut = util.SafeStringDatasource(plan.FilesToPut.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
}

func (r *SftpConnectionResource) CreateSFTPConnection(ctx context.Context, plan *SFTPConnectorResourceModel, config *SFTPConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SFTP connection creation")

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
		errorCode := sftpErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err, nil)
		return nil, fmt.Errorf("[%s] Failed to check existing connection: %w", errorCode, err)
	}

	if existingResource != nil &&
		existingResource.SFTPConnectionResponse != nil &&
		existingResource.SFTPConnectionResponse.Errorcode != nil &&
		*existingResource.SFTPConnectionResponse.Errorcode == 0 {

		errorCode := sftpErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists.Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "create", connectionName, nil)
	}

	// Build SFTP connection create request
	tflog.Debug(ctx, "Building SFTP connection create request")

	// if (config.Password.IsNull() || config.Password.IsUnknown()) && (config.PasswordWo.IsNull() || config.PasswordWo.IsUnknown()) {
	// 	return nil, fmt.Errorf("either password or password_wo must be set")
	// }

	sftpConn := r.BuildSFTPConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		SFTPConnector: &sftpConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_sftp_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sftpErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create SFTP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := sftpErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "SFTP connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SFTP connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SftpConnectionResource) ReadSFTPConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SFTP connection read operation")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute read operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_sftp_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sftpErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read SFTP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateSFTPConnectionResponse(apiResp); err != nil {
		errorCode := sftpErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for SFTP datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.SFTPConnectionResponse != nil && apiResp.SFTPConnectionResponse.Errorcode != nil && *apiResp.SFTPConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.SFTPConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.SFTPConnectionResponse.Msg))
		errorCode := sftpErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "SFTP connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.SFTPConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.SFTPConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SFTP connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.SFTPConnectionResponse != nil && apiResp.SFTPConnectionResponse.Connectionkey != nil {
				return *apiResp.SFTPConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SftpConnectionResource) UpdateSFTPConnection(ctx context.Context, plan *SFTPConnectorResourceModel, config *SFTPConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SFTP connection update")

	// Build SFTP connection update request
	tflog.Debug(logCtx, "Building SFTP connection update request")

	sftpConn := r.BuildSFTPConnector(plan, config)

	updateReq := openapi.CreateOrUpdateRequest{
		SFTPConnector: &sftpConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_sftp_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := sftpErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update SFTP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := sftpErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "SFTP connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSFTP, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SFTP connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *SftpConnectionResource) UpdateModelFromReadResponse(state *SFTPConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.SFTPConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.SFTPConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionname)

	state.Description = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Emailtemplate)

	state.HostName = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.HOST_NAME)
	state.PortNumber = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.PORT_NUMBER)
	state.Username = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.USERNAME)
	state.AuthCredentialType = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.AUTH_CREDENTIAL_TYPE)
	state.FilesToGet = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.FILES_TO_GET)
	state.FilesToPut = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.FILES_TO_PUT)
	state.PamConfig = util.SafeStringDatasource(apiResp.SFTPConnectionResponse.Connectionattributes.PAM_CONFIG)
	// Note: AUTH_CREDENTIAL_VALUE and PASSPHRASE are not read from API for security reasons
}

func (r *SftpConnectionResource) ValidateSFTPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SFTPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - SFTP connection response is nil")
	}
	return nil
}

func (r *SftpConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config SFTPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpErrorCodes.PlanExtraction()
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
		errorCode := sftpErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateSFTPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "SFTP connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"SFTP Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "SFTP connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *SftpConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SFTPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadSFTPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "SFTP connection read failed", "", err)
		resp.Diagnostics.AddError(
			"SFTP Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)
	apiMessage := util.SafeDeref(apiResp.SFTPConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.SFTPConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SFTP connection resource read completed successfully")
}

func (r *SftpConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config SFTPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpErrorCodes.StateExtraction()
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
		errorCode := sftpErrorCodes.PlanExtraction()
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
		errorCode := sftpErrorCodes.ConfigExtraction()
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
		errorCode := sftpErrorCodes.NameImmutable()
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
	updateResp, err := r.UpdateSFTPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "SFTP connection update failed", "", err)
		resp.Diagnostics.AddError(
			"SFTP Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadSFTPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated SFTP connection", "", err)
		resp.Diagnostics.AddError(
			"SFTP Connection Post-Update Read Failed",
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
		errorCode := sftpErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SFTP connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *SftpConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *SftpConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an SFTP connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "SFTP connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
