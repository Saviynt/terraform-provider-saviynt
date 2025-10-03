// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_roles_datasource retrieves role details from the Savi ynt Security Manager.
// The data source supports a single Read operation to look up existing roles with various filters.
package provider

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/endpointsutil"
	"terraform-provider-Saviynt/util/rolesutil"

	openapi "github.com/saviynt/saviynt-api-go-client/roles"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CustomPropertiesModel defines the 60 custom properties that can be reused across structs
type CustomPropertiesModel struct {
	CustomProperty1  types.String `tfsdk:"custom_property1"`
	CustomProperty2  types.String `tfsdk:"custom_property2"`
	CustomProperty3  types.String `tfsdk:"custom_property3"`
	CustomProperty4  types.String `tfsdk:"custom_property4"`
	CustomProperty5  types.String `tfsdk:"custom_property5"`
	CustomProperty6  types.String `tfsdk:"custom_property6"`
	CustomProperty7  types.String `tfsdk:"custom_property7"`
	CustomProperty8  types.String `tfsdk:"custom_property8"`
	CustomProperty9  types.String `tfsdk:"custom_property9"`
	CustomProperty10 types.String `tfsdk:"custom_property10"`
	CustomProperty11 types.String `tfsdk:"custom_property11"`
	CustomProperty12 types.String `tfsdk:"custom_property12"`
	CustomProperty13 types.String `tfsdk:"custom_property13"`
	CustomProperty14 types.String `tfsdk:"custom_property14"`
	CustomProperty15 types.String `tfsdk:"custom_property15"`
	CustomProperty16 types.String `tfsdk:"custom_property16"`
	CustomProperty17 types.String `tfsdk:"custom_property17"`
	CustomProperty18 types.String `tfsdk:"custom_property18"`
	CustomProperty19 types.String `tfsdk:"custom_property19"`
	CustomProperty20 types.String `tfsdk:"custom_property20"`
	CustomProperty21 types.String `tfsdk:"custom_property21"`
	CustomProperty22 types.String `tfsdk:"custom_property22"`
	CustomProperty23 types.String `tfsdk:"custom_property23"`
	CustomProperty24 types.String `tfsdk:"custom_property24"`
	CustomProperty25 types.String `tfsdk:"custom_property25"`
	CustomProperty26 types.String `tfsdk:"custom_property26"`
	CustomProperty27 types.String `tfsdk:"custom_property27"`
	CustomProperty28 types.String `tfsdk:"custom_property28"`
	CustomProperty29 types.String `tfsdk:"custom_property29"`
	CustomProperty30 types.String `tfsdk:"custom_property30"`
	CustomProperty31 types.String `tfsdk:"custom_property31"`
	CustomProperty32 types.String `tfsdk:"custom_property32"`
	CustomProperty33 types.String `tfsdk:"custom_property33"`
	CustomProperty34 types.String `tfsdk:"custom_property34"`
	CustomProperty35 types.String `tfsdk:"custom_property35"`
	CustomProperty36 types.String `tfsdk:"custom_property36"`
	CustomProperty37 types.String `tfsdk:"custom_property37"`
	CustomProperty38 types.String `tfsdk:"custom_property38"`
	CustomProperty39 types.String `tfsdk:"custom_property39"`
	CustomProperty40 types.String `tfsdk:"custom_property40"`
	CustomProperty41 types.String `tfsdk:"custom_property41"`
	CustomProperty42 types.String `tfsdk:"custom_property42"`
	CustomProperty43 types.String `tfsdk:"custom_property43"`
	CustomProperty44 types.String `tfsdk:"custom_property44"`
	CustomProperty45 types.String `tfsdk:"custom_property45"`
	CustomProperty46 types.String `tfsdk:"custom_property46"`
	CustomProperty47 types.String `tfsdk:"custom_property47"`
	CustomProperty48 types.String `tfsdk:"custom_property48"`
	CustomProperty49 types.String `tfsdk:"custom_property49"`
	CustomProperty50 types.String `tfsdk:"custom_property50"`
	CustomProperty51 types.String `tfsdk:"custom_property51"`
	CustomProperty52 types.String `tfsdk:"custom_property52"`
	CustomProperty53 types.String `tfsdk:"custom_property53"`
	CustomProperty54 types.String `tfsdk:"custom_property54"`
	CustomProperty55 types.String `tfsdk:"custom_property55"`
	CustomProperty56 types.String `tfsdk:"custom_property56"`
	CustomProperty57 types.String `tfsdk:"custom_property57"`
	CustomProperty58 types.String `tfsdk:"custom_property58"`
	CustomProperty59 types.String `tfsdk:"custom_property59"`
	CustomProperty60 types.String `tfsdk:"custom_property60"`
}

type RolesDataSource struct {
	client      client.SaviyntClientInterface
	token       string
	provider    client.SaviyntProviderInterface
	roleFactory client.RoleFactoryInterface
}

var _ datasource.DataSource = &RolesDataSource{}
var _ datasource.DataSourceWithConfigure = &RolesDataSource{}

func NewRolesDataSource() datasource.DataSource {
	return &RolesDataSource{
		roleFactory: &client.DefaultRoleFactory{},
	}
}

// NewRolesDataSourceWithFactory creates a new roles data source with custom factory
// Used primarily for testing with mock factories
func NewRolesDataSourceWithFactory(factory client.RoleFactoryInterface) datasource.DataSource {
	return &RolesDataSource{
		roleFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *RolesDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *RolesDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (r *RolesDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// setCustomPropertiesOnRequest sets custom properties on the API request using reflection
// This eliminates the need for 60 individual if-statements
func (d *RolesDataSource) setCustomPropertiesOnRequest(state *RolesDataSourceModel, areq *openapi.GetRolesRequest) {
	// Use reflection to iterate through custom properties
	stateValue := reflect.ValueOf(state).Elem()
	reqValue := reflect.ValueOf(areq) // Don't dereference - methods are on pointer type

	for i := 1; i <= 60; i++ {
		// Get the field from state
		stateFieldName := fmt.Sprintf("CustomProperty%d", i)
		stateField := stateValue.FieldByName(stateFieldName)

		if !stateField.IsValid() {
			continue
		}

		// Get the Terraform types.String value
		tfString := stateField.Interface().(types.String)

		// Check if the field has a value
		if !tfString.IsNull() && tfString.ValueString() != "" {
			// Get the corresponding setter method on the request
			setterMethodName := fmt.Sprintf("SetCustomproperty%d", i)
			setterMethod := reqValue.MethodByName(setterMethodName)

			if setterMethod.IsValid() {
				// Call the setter method with the string value
				args := []reflect.Value{reflect.ValueOf(tfString.ValueString())}
				setterMethod.Call(args)
			}
		}
	}
}

// mapCustomPropertiesFromResponse maps custom properties from API response to Role struct using reflection
// This eliminates the need for 60 individual assignments
func (d *RolesDataSource) mapCustomPropertiesFromResponse(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	// Use reflection to iterate through custom properties
	roleValue := reflect.ValueOf(roleState).Elem()
	itemValue := reflect.ValueOf(item).Elem()

	for i := 1; i <= 60; i++ {
		// Get the field from API response
		itemFieldName := fmt.Sprintf("CustomProperty%d", i)
		itemField := itemValue.FieldByName(itemFieldName)

		if !itemField.IsValid() {
			continue
		}

		// Get the field from role state
		roleFieldName := fmt.Sprintf("CustomProperty%d", i)
		roleField := roleValue.FieldByName(roleFieldName)

		if !roleField.IsValid() || !roleField.CanSet() {
			continue
		}

		// Get the string pointer from the API response with safe type assertion
		stringPtr, ok := itemField.Interface().(*string)
		if !ok {
			continue // Skip if not a *string
		}

		// Use util.SafeString to convert to types.String and set on role state
		safeString := util.SafeString(stringPtr)
		roleField.Set(reflect.ValueOf(safeString))
	}
}

// ParameterMapping defines the mapping between state field and API setter method
type ParameterMapping struct {
	StateFieldName string // Field name in the state struct
	SetterMethod   string // Method name on the API request object
	DefaultValue   string // Optional default value
	Required       bool   // Whether this parameter is required
}

// setRequestParametersFromState sets API request parameters from state using reflection
// This eliminates the need for repetitive if-statements for each parameter
func (d *RolesDataSource) setRequestParametersFromState(state *RolesDataSourceModel, areq *openapi.GetRolesRequest) error {
	// Define the parameter mappings
	parameterMappings := []ParameterMapping{
		{StateFieldName: "RequestedObject", SetterMethod: "SetRequestedObject", DefaultValue: "entitlement", Required: false},
		{StateFieldName: "Username", SetterMethod: "SetUsername", Required: false},
		{StateFieldName: "RoleType", SetterMethod: "SetRoleType", Required: false},
		{StateFieldName: "Requestable", SetterMethod: "SetRequestable", Required: false},
		{StateFieldName: "Status", SetterMethod: "SetStatus", Required: false},
		{StateFieldName: "RoleName", SetterMethod: "SetRoleName", Required: false},
		{StateFieldName: "Description", SetterMethod: "SetDescription", Required: false},
		{StateFieldName: "DisplayName", SetterMethod: "SetDisplayname", Required: false},
		{StateFieldName: "Glossary", SetterMethod: "SetGlossary", Required: false},
		{StateFieldName: "MiningInstance", SetterMethod: "SetMininginstance", Required: false},
		{StateFieldName: "Risk", SetterMethod: "SetRisk", Required: false},
		{StateFieldName: "UpdateUser", SetterMethod: "SetUpdateuser", Required: false},
		{StateFieldName: "SystemId", SetterMethod: "SetSystemid", Required: false},
		{StateFieldName: "SoxCritical", SetterMethod: "SetSoxcritical", Required: false},
		{StateFieldName: "SysCritical", SetterMethod: "SetSyscritical", Required: false},
		{StateFieldName: "Level", SetterMethod: "SetLevel", Required: false},
		{StateFieldName: "Privileged", SetterMethod: "SetPriviliged", Required: false},
		{StateFieldName: "Confidentiality", SetterMethod: "SetConfidentiality", Required: false},
		{StateFieldName: "Max", SetterMethod: "SetMax", Required: false},
		{StateFieldName: "Offset", SetterMethod: "SetOffset", Required: false},
		{StateFieldName: "RoleQuery", SetterMethod: "SetRoleQuery", Required: false},
		{StateFieldName: "HideBlankValues", SetterMethod: "SetHideblankvalues", Required: false},
	}

	// Use reflection to set parameters
	stateValue := reflect.ValueOf(state).Elem()
	reqValue := reflect.ValueOf(areq) // Don't dereference - methods are on pointer type

	for _, mapping := range parameterMappings {
		// Get the field from state
		stateField := stateValue.FieldByName(mapping.StateFieldName)
		if !stateField.IsValid() {
			continue
		}

		// Get the Terraform types.String value
		tfString := stateField.Interface().(types.String)

		// Determine the value to set
		var valueToSet string
		hasValue := false

		if !tfString.IsNull() && tfString.ValueString() != "" {
			valueToSet = tfString.ValueString()
			hasValue = true
		} else if mapping.DefaultValue != "" {
			valueToSet = mapping.DefaultValue
			hasValue = true
		}

		// Set the value if we have one
		if hasValue {
			// Convert human-readable values to numeric values for specific fields
			if mapping.StateFieldName == "RoleType" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.RoleTypeMap)
			} else if mapping.StateFieldName == "Status" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.StatusMap)
			} else if mapping.StateFieldName == "SoxCritical" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.SoxCriticalityMap)
			} else if mapping.StateFieldName == "SysCritical" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.SysCriticalMap)
			} else if mapping.StateFieldName == "Privileged" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.PrivilegedMap)
			} else if mapping.StateFieldName == "Confidentiality" {
				valueToSet = rolesutil.ReverseTranslateValue(valueToSet, rolesutil.ConfidentialityMap)
			}

			setterMethod := reqValue.MethodByName(mapping.SetterMethod)
			if !setterMethod.IsValid() {
				return fmt.Errorf("setter method %s not found on request object", mapping.SetterMethod)
			}

			// Call the setter method with the string value
			args := []reflect.Value{reflect.ValueOf(valueToSet)}
			setterMethod.Call(args)
		}
	}

	return nil
}

type RolesDataSourceModel struct {
	Roledetails     []Role       `tfsdk:"roledetails"`
	Authenticate    types.Bool   `tfsdk:"authenticate"`
	DisplayCount    types.Int64  `tfsdk:"display_count"`
	ErrorCode       types.String `tfsdk:"error_code"`
	TotalCount      types.Int64  `tfsdk:"total_count"`
	Message         types.String `tfsdk:"message"`
	RequestedObject types.String `tfsdk:"requested_object"`
	Username        types.String `tfsdk:"username"`
	RoleType        types.String `tfsdk:"role_type"`
	Requestable     types.String `tfsdk:"requestable"`
	Status          types.String `tfsdk:"status"`
	RoleName        types.String `tfsdk:"role_name"`
	Description     types.String `tfsdk:"description"`
	DisplayName     types.String `tfsdk:"display_name"`
	Glossary        types.String `tfsdk:"glossary"`
	MiningInstance  types.String `tfsdk:"mining_instance"`
	Risk            types.String `tfsdk:"risk"`
	UpdateUser      types.String `tfsdk:"update_user"`
	SystemId        types.String `tfsdk:"system_id"`
	SoxCritical     types.String `tfsdk:"sox_critical"`
	SysCritical     types.String `tfsdk:"sys_critical"`
	Level           types.String `tfsdk:"level"`
	Privileged      types.String `tfsdk:"privileged"`
	Confidentiality types.String `tfsdk:"confidentiality"`
	Max             types.String `tfsdk:"max"`
	Offset          types.String `tfsdk:"offset"`
	RoleQuery       types.String `tfsdk:"role_query"`
	HideBlankValues types.String `tfsdk:"hide_blank_values"`
	// Embed the reusable custom properties model
	CustomPropertiesModel
}

type RoleOwner struct {
	OwnerName types.String `tfsdk:"owner_name"`
	Rank      types.String `tfsdk:"rank"`
}

type EntitlementDetail struct {
	Endpoint            types.String `tfsdk:"endpoint"`
	EntitlementTypeName types.String `tfsdk:"entitlement_type_name"`
	EntitlementValue    types.String `tfsdk:"entitlement_value"`
}

type UserDetail struct {
	UserName types.String `tfsdk:"user_name"`
}
type Role struct {
	RoleKey                  types.Int64         `tfsdk:"role_key"`
	UpdateDate               types.String        `tfsdk:"update_date"`
	RoleType                 types.String        `tfsdk:"role_type"`
	Version                  types.Int64         `tfsdk:"version"`
	RoleName                 types.String        `tfsdk:"role_name"`
	Description              types.String        `tfsdk:"description"`
	Glossary                 types.String        `tfsdk:"glossary"`
	Privileged               types.String        `tfsdk:"privileged"`
	Status                   types.String        `tfsdk:"status"`
	ShowDynamicAttrs         types.String        `tfsdk:"show_dynamic_attrs"`
	DefaultTimeFrameHrs      types.String        `tfsdk:"default_time_frame_hrs"`
	MaxTimeFrameHrs          types.String        `tfsdk:"max_time_frame_hrs"`
	Confidentiality          types.String        `tfsdk:"confidentiality"`
	SoxCritical              types.String        `tfsdk:"sox_critical"`
	SysCritical              types.String        `tfsdk:"sys_critical"`
	Requestable              types.String        `tfsdk:"requestable"`
	DisplayName              types.String        `tfsdk:"display_name"`
	UpdateUser               types.String        `tfsdk:"update_user"`
	EntitlementValueKey      types.String        `tfsdk:"entitlement_value_key"`
	RoleState                types.String        `tfsdk:"role_state"`
	EndpointKey              types.String        `tfsdk:"endpoint_key"`
	LastReviewedCampaignName types.String        `tfsdk:"last_reviewed_campaign_name"`
	LastReviewedBy           types.String        `tfsdk:"last_reviewed_by"`
	Risk                     types.String        `tfsdk:"risk"`
	Owners                   []RoleOwner         `tfsdk:"owners"`
	UserDetails              []UserDetail        `tfsdk:"user_details"`
	EntitlementDetails       []EntitlementDetail `tfsdk:"entitlement_details"`
	// Embed the reusable custom properties model
	CustomPropertiesModel
}

func RoleResultSchema() map[string]schema.Attribute {
	attrs := map[string]schema.Attribute{
		"role_key": schema.Int64Attribute{
			Computed:    true,
			Description: "Unique numeric identifier (key) of the role in Saviynt. This is automatically generated by the system and used for internal references.",
		},
		"update_date": schema.StringAttribute{
			Computed:    true,
			Description: "Date and time when the role was last updated. This shows the most recent modification timestamp.",
		},
		"role_type": schema.StringAttribute{
			Computed:    true,
			Description: "Type of the role. Values: 'ENABLER', 'TRANSACTIONAL', 'FIREFIGHTER', 'ENTERPRISE', 'APPLICATION', 'ENTITLEMENT'.",
		},
		"version": schema.Int64Attribute{
			Computed:    true,
			Description: "Version number of the role. This tracks the number of times the role has been modified. Can be returned as either integer or string from the API and is converted to int64.",
		},
		"role_name": schema.StringAttribute{
			Computed:    true,
			Description: "Name of the role. This is the unique identifier used to reference the role in Saviynt.",
		},
		"description": schema.StringAttribute{
			Computed:    true,
			Description: "Description of the role. Provides detailed information about the role's purpose, responsibilities, and access it grants.",
		},
		"glossary": schema.StringAttribute{
			Computed:    true,
			Description: "Glossary information for the role. Provides additional context and definitions related to the role.",
		},
		"privileged": schema.StringAttribute{
			Computed:    true,
			Description: "Privileged criticality of the role. Describes privileges assigned to the role and amount of risk to provide access. Values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
		},
		"status": schema.StringAttribute{
			Computed:    true,
			Description: "Status of the role. Indicates the current state of the role (e.g., 'Active', 'Inactive'). Values: '1' (Active), '0' (Inactive).",
		},
		"show_dynamic_attrs": schema.StringAttribute{
			Computed:    true,
			Description: "Displays the dynamic attributes associated with the role. For example, there is a Dynamic Attribute A, which is Boolean set as true and false.",
		},
		"default_time_frame_hrs": schema.StringAttribute{
			Computed:    true,
			Description: "Default time frame (in hours) to request access for a role. This defines how long users will have access when assigned this role.",
		},
		"max_time_frame_hrs": schema.StringAttribute{
			Computed:    true,
			Description: "Maximum time frame (in hours) allowed for role access requests. This sets the upper limit for how long users can request access to this role.",
		},
		"confidentiality": schema.StringAttribute{
			Computed:    true,
			Description: "Confidentiality level of the role. Values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
		},
		"sox_critical": schema.StringAttribute{
			Computed:    true,
			Description: "SOX criticality of the role. Values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
		},
		"sys_critical": schema.StringAttribute{
			Computed:    true,
			Description: "SYS criticality of the role. Values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
		},
		"requestable": schema.StringAttribute{
			Computed:    true,
			Description: "Indicates if the role is requestable by users. Values: 'true' (role can be requested), 'false' (role cannot be requested).",
		},
		"display_name": schema.StringAttribute{
			Computed:    true,
			Description: "Display name of the role. This is a user-friendly name that can be different from the role_name.",
		},
		"update_user": schema.StringAttribute{
			Computed:    true,
			Description: "User ID of the user who last updated the role.",
		},
		"entitlement_value_key": schema.StringAttribute{
			Computed:    true,
			Description: "Entitlement value key associated with the role.",
		},
		"role_state": schema.StringAttribute{
			Computed:    true,
			Description: "State of the role. Indicates the current operational state or lifecycle stage of the role.",
		},
		"endpoint_key": schema.StringAttribute{
			Computed:    true,
			Description: "Endpoint key associated with the role.",
		},
		"last_reviewed_campaign_name": schema.StringAttribute{
			Computed:    true,
			Description: "Name of the last campaign that reviewed this role.",
		},
		"last_reviewed_by": schema.StringAttribute{
			Computed:    true,
			Description: "User who last reviewed the role.",
		},
		"risk": schema.StringAttribute{
			Computed:    true,
			Description: "Risk level of the role required during separation of duties. Values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
		},
		"entitlement_details": schema.ListNestedAttribute{
			Computed:    true,
			Description: "List of entitlement details associated with the role.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Description: "Name of the endpoint.",
					},
					"entitlement_type_name": schema.StringAttribute{
						Computed:    true,
						Description: "Type name of the entitlement.",
					},
					"entitlement_value": schema.StringAttribute{
						Computed:    true,
						Description: "Specific value of the entitlement.",
					},
				},
			},
		},
		"user_details": schema.ListNestedAttribute{
			Computed:    true,
			Description: "List of users assigned to the role. Shows all users who currently have this role and inherit its entitlements and permissions.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"user_name": schema.StringAttribute{
						Computed:    true,
						Description: "Username of the user assigned to the role. This is the unique identifier for the user in Saviynt.",
					},
				},
			},
		},
		"owners": schema.ListNestedAttribute{
			Computed:    true,
			Description: "List of role owners with their respective ranks. Owners are responsible for role governance and approvals.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"owner_name": schema.StringAttribute{
						Computed:    true,
						Description: "Username of the role owner. This is a valid Saviynt user responsible for managing the role.",
					},
					"rank": schema.StringAttribute{
						Computed:    true,
						Description: "Rank of the owner (1-27, where 1 is highest priority). This determines the owner's priority in approval workflows. Special ranks: 26 (Primary Certifier), 27 (Secondary Certifier).",
					},
				},
			},
		},
	}
	// Add custom properties 1-60
	for i := 1; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		attrs[key] = schema.StringAttribute{
			Computed:    true,
			Description: fmt.Sprintf("Custom property %d value for the role.", i),
		}
	}

	return attrs
}

func (d *RolesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_roles_datasource"
}

func (d *RolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.RoleDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"requested_object": schema.StringAttribute{
				Optional:    true,
				Description: "Request body parameter that contains users and/or entitlement_values. Use comma-separated values like 'entitlement' to specify what objects to include in the response.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by username. When specified, returns only roles associated with this specific user.",
			},
			"role_type": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by type. Valid values: 'ENABLER', 'TRANSACTIONAL', 'FIREFIGHTER', 'ENTERPRISE', 'APPLICATION', 'ENTITLEMENT'. Use the human-readable names to filter results.",
			},
			"requestable": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by requestable status. Valid values: 'true' (show only requestable roles), 'false' (show only non-requestable roles).",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by status. Valid values: 'Active', 'Inactive'. Use string values to filter results.",
			},
			"role_name": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by role name. When specified, returns only roles that match this exact role name.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by description. When specified, returns only roles that match this description text.",
			},
			"display_name": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by display name. When specified, returns only roles that match this display name.",
			},
			"glossary": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by glossary information. When specified, returns only roles that match this glossary text.",
			},
			"mining_instance": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by mining instance name. When specified, returns only roles associated with this mining instance.",
			},
			"risk": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by risk level. Valid values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'. When specified, returns only roles with this risk level.",
			},
			"update_user": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by the user ID who last updated the role. When specified, returns only roles last modified by this user ID.",
			},
			"system_id": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by system ID. When specified, returns only roles associated with this system identifier.",
			},
			"sox_critical": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by SOX criticality. Valid values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'. When specified, returns only roles with this SOX criticality level.",
			},
			"sys_critical": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by SYS criticality. Valid values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'. When specified, returns only roles with this SYS criticality level.",
			},
			"level": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by hierarchy level. When specified, returns only roles with this hierarchy level in the organizational structure.",
			},
			"privileged": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by privileged criticality. Valid values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'. When specified, returns only roles with this privileged criticality level.",
			},
			"confidentiality": schema.StringAttribute{
				Optional:    true,
				Description: "Filter roles by confidentiality level. Valid values: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'. When specified, returns only roles with this confidentiality level.",
			},
			"role_query": schema.StringAttribute{
				Optional:    true,
				Description: "SQL-like query to filter roles (e.g., 'r.role_name = 'Admin'')",
			},
			"hide_blank_values": schema.StringAttribute{
				Optional:    true,
				Description: "Hide blank values (e.g., true or false)",
			},
			"max": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum number of records to return",
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Description: "Offset for pagination",
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "API response message",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of records returned in the response",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Total count of available records",
			},
			"roledetails": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of roles retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: RoleResultSchema(),
				},
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "Controls visibility of sensitive data in Terraform state. When 'false', role details are omitted from state to prevent sensitive data exposure. When 'true', all role details are returned in state.",
			},
		},
	}

	// Add custom property filter attributes
	for i := 1; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Description: fmt.Sprintf("Custom property %d for additional metadata", i),
		}
	}
}
func (d *RolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting roles datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed - expected *SaviyntProvider, got different type")
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *SaviyntProvider, got different type",
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}
	tflog.Debug(ctx, "Roles datasource configured successfully")
}

func (d *RolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state RolesDataSourceModel

	tflog.Debug(ctx, "Starting roles datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to extract configuration from request")
		return
	}

	// Execute API call to get roles details
	rolesResponse, err := d.ReadRolesDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read roles details", map[string]interface{}{"error": err.Error()})
		resp.Diagnostics.AddError(
			"Roles Read Failed",
			fmt.Sprintf("Failed to read roles: %s", err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateRolesResponse(rolesResponse); err != nil {
		tflog.Error(ctx, "Invalid roles response", map[string]interface{}{"error": err.Error()})
		resp.Diagnostics.AddError(
			"Invalid Roles Response",
			fmt.Sprintf("API response validation failed: %s", err.Error()),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromRolesResponse(&state, rolesResponse)

	// Handle conditional request attributes - only keep attributes that were set by user
	d.HandleConditionalRequestAttributes(ctx, req, &state)

	// Handle authentication logic
	d.HandleRolesAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set state for roles datasource")
		return
	}

	tflog.Info(ctx, "Roles datasource read completed successfully", map[string]interface{}{
		"total_roles": len(state.Roledetails),
		"total_count": state.TotalCount.ValueInt64(),
	})
}

// ReadRolesDetails retrieves roles details from Saviynt API
// Handles role filtering and parameter configuration using factory pattern
// Returns standardized errors with proper correlation tracking
func (d *RolesDataSource) ReadRolesDetails(ctx context.Context, state *RolesDataSourceModel) (*openapi.GetRolesResponse, error) {
	tflog.Debug(ctx, "Starting roles API business logic")

	areq := openapi.GetRolesRequest{}

	// Set all request parameters using reflection-based helper
	if err := d.setRequestParametersFromState(state, &areq); err != nil {
		return nil, fmt.Errorf("failed to set request parameters: %v", err)
	}

	// Add custom properties 1-60 using reflection-based helper
	d.setCustomPropertiesOnRequest(state, &areq)

	tflog.Debug(ctx, "Executing roles API request")
	var rolesResponse *openapi.GetRolesResponse
	var finalHttpResp *http.Response
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "get_roles_datasource", func(token string) error {
		roleOps := d.roleFactory.CreateRoleOperations(d.client.APIBaseURL(), token)
		resp, httpResp, err := roleOps.GetRoles(ctx, areq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		rolesResponse = resp
		finalHttpResp = httpResp  // Capture final HTTP response
		return err
	})

	// Handle errors using helper functions
	if rolesutil.RoleHandleHTTPError(ctx, finalHttpResp, err, "reading roles", &diag.Diagnostics{}) {
		return nil, fmt.Errorf("HTTP error during roles read")
	}

	if rolesResponse != nil && rolesutil.RoleHandleAPIError(ctx, rolesResponse.ErrorCode, rolesResponse.Msg, "reading roles", &diag.Diagnostics{}) {
		return nil, fmt.Errorf("API error during roles read: %s", util.SafeDeref(rolesResponse.Msg))
	}

	roleCount := 0
	if rolesResponse != nil && rolesResponse.Roledetails != nil {
		roleCount = len(rolesResponse.Roledetails)
	}

	logData := map[string]interface{}{
		"role_count": roleCount,
	}
	if finalHttpResp != nil {
		logData["status_code"] = finalHttpResp.StatusCode
	}

	tflog.Debug(ctx, "Roles API request completed successfully", logData)

	return rolesResponse, nil
}

// ValidateRolesResponse validates that the API response contains valid roles data
// Returns standardized error if validation fails
func (d *RolesDataSource) ValidateRolesResponse(rolesResponse *openapi.GetRolesResponse) error {
	if rolesResponse == nil {
		return fmt.Errorf("roles response is nil")
	}

	if rolesResponse.ErrorCode != nil && *rolesResponse.ErrorCode != "0" {
		return fmt.Errorf("API returned error code: %s", *rolesResponse.ErrorCode)
	}

	return nil
}

// UpdateModelFromRolesResponse maps API response data to the Terraform state model
// It handles both basic response fields and detailed role processing
func (d *RolesDataSource) UpdateModelFromRolesResponse(state *RolesDataSourceModel, rolesResponse *openapi.GetRolesResponse) {
	// Map basic response fields
	d.MapBasicRolesFields(state, rolesResponse)

	// Map role details
	d.MapRolesDetails(state, rolesResponse)
}

// MapBasicRolesFields maps basic response fields from API response to state model
// These are common fields available for all roles responses
func (d *RolesDataSource) MapBasicRolesFields(state *RolesDataSourceModel, rolesResponse *openapi.GetRolesResponse) {
	state.Message = util.SafeStringDatasource(rolesResponse.Msg)
	state.DisplayCount = util.SafeInt64(rolesResponse.DisplayCount)
	state.ErrorCode = util.SafeStringDatasource(rolesResponse.ErrorCode)
	state.TotalCount = util.SafeInt64(rolesResponse.TotalCount)
}

// MapRolesDetails maps detailed role information from API response to state model
// Handles complex role processing including custom properties, entitlements, owners, etc.
func (d *RolesDataSource) MapRolesDetails(state *RolesDataSourceModel, rolesResponse *openapi.GetRolesResponse) {
	if rolesResponse.Roledetails == nil {
		tflog.Info(context.Background(), "No role details found in API response - setting empty role list")
		state.Roledetails = []Role{}
		return
	}

	var roles []Role
	for _, item := range rolesResponse.Roledetails {
		roleState := d.ProcessSingleRole(item)
		roles = append(roles, roleState)
	}
	state.Roledetails = roles
}

// ProcessSingleRole processes a single role item from the API response
// Handles all the complex mapping logic for individual roles
func (d *RolesDataSource) ProcessSingleRole(item openapi.GetRoleDetailsResponse) Role {
	roleState := Role{
		RoleKey:                  util.SafeInt64(item.RoleKey),
		UpdateDate:               util.SafeStringDatasource(item.Updatedate),
		RoleType:                 util.SafeStringDatasource(item.Roletype),
		RoleName:                 util.SafeStringDatasource(item.RoleName),
		Description:              util.SafeStringDatasource(item.Description),
		Glossary:                 util.SafeStringDatasource(item.Glossary),
		Status:                   util.SafeStringDatasource(item.Status),
		ShowDynamicAttrs:         util.SafeStringDatasource(item.ShowDynamicAttrs),
		DefaultTimeFrameHrs:      util.SafeStringDatasource(item.DefaultTimeFrameHrs),
		MaxTimeFrameHrs:          util.SafeStringDatasource(item.MaxTimeFrameHrs),
		Requestable:              util.SafeStringDatasource(item.Requestable),
		DisplayName:              util.SafeStringDatasource(item.Displayname),
		UpdateUser:               util.SafeStringDatasource(item.Updateuser),
		EntitlementValueKey:      util.SafeStringDatasource(item.EntitlementValueKey),
		RoleState:                util.SafeStringDatasource(item.RoleState),
		EndpointKey:              util.SafeStringDatasource(item.Endpointkey),
		LastReviewedCampaignName: util.SafeStringDatasource(item.LastReviewedCampaignName),
		LastReviewedBy:           util.SafeStringDatasource(item.LastReviewedBy),
	}

	// Map custom properties using reflection-based helper
	d.mapCustomPropertiesFromResponse(&roleState, &item)

	// Process version field with special handling
	d.ProcessRoleVersion(&roleState, &item)

	// Process role attributes (sox critical, sys critical, etc.)
	d.ProcessRoleAttributes(&roleState, &item)

	// Process complex nested objects
	d.ProcessRoleEntitlements(&roleState, &item)
	d.ProcessRoleUsers(&roleState, &item)
	d.ProcessRoleOwners(&roleState, &item)

	return roleState
}

// ProcessRoleVersion handles the complex version field processing
func (d *RolesDataSource) ProcessRoleVersion(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	if item.Version != nil {
		if item.Version.Int32 != nil {
			roleState.Version = types.Int64Value(int64(*item.Version.Int32))
		} else if item.Version.String != nil {
			val, err := strconv.ParseInt(*item.Version.String, 10, 64)
			if err == nil {
				roleState.Version = types.Int64Value(val)
			} else {
				roleState.Version = types.Int64Null()
			}
		} else {
			roleState.Version = types.Int64Null()
		}
	} else {
		roleState.Version = types.Int64Null()
	}
}

// ProcessRoleAttributes handles role attributes mapping (sox critical, sys critical, etc.)
func (d *RolesDataSource) ProcessRoleAttributes(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	if item.Soxcritical != nil {
		roleState.SoxCritical = types.StringValue(endpointsutil.TranslateValue(*item.Soxcritical, rolesutil.SoxCriticalityMap))
	}
	if item.Syscritical != nil {
		roleState.SysCritical = types.StringValue(endpointsutil.TranslateValue(*item.Syscritical, rolesutil.SysCriticalMap))
	}
	if item.Priviliged != nil {
		roleState.Privileged = types.StringValue(endpointsutil.TranslateValue(*item.Priviliged, rolesutil.PrivilegedMap))
	}
	if item.Confidentiality != nil {
		roleState.Confidentiality = types.StringValue(endpointsutil.TranslateValue(*item.Confidentiality, rolesutil.ConfidentialityMap))
	}
	if item.Risk != nil {
		roleState.Risk = types.StringValue(endpointsutil.TranslateValue(*item.Risk, rolesutil.RiskMap))
	}
}

// ProcessRoleEntitlements handles entitlement details processing
func (d *RolesDataSource) ProcessRoleEntitlements(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	var entitlementDetails []EntitlementDetail
	for _, ent := range item.EntitlementDetails {
		entitlementDetails = append(entitlementDetails, EntitlementDetail{
			Endpoint:            util.SafeStringDatasource(ent.Endpoint),
			EntitlementTypeName: util.SafeStringDatasource(ent.EntitlementTypeName),
			EntitlementValue:    util.SafeStringDatasource(ent.EntitlementValue),
		})
	}
	roleState.EntitlementDetails = entitlementDetails
}

// ProcessRoleUsers handles user details processing
func (d *RolesDataSource) ProcessRoleUsers(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	var users []UserDetail
	for _, user := range item.UserDetails {
		if user.GetUserDetailsResponse != nil && user.GetUserDetailsResponse.Username != nil {
			users = append(users, UserDetail{
				UserName: util.SafeStringDatasource(user.GetUserDetailsResponse.Username),
			})
		}
	}
	roleState.UserDetails = users
}

// ProcessRoleOwners handles owner details processing
func (d *RolesDataSource) ProcessRoleOwners(roleState *Role, item *openapi.GetRoleDetailsResponse) {
	var owners []RoleOwner
	if item.Owner.ArrayOfGetRoleOwnersResponse == nil {
		tflog.Debug(context.Background(), "No owners found for role", map[string]interface{}{"role_name": roleState.RoleName})
	} else {
		for _, owner := range *item.Owner.ArrayOfGetRoleOwnersResponse {
			owners = append(owners, RoleOwner{
				OwnerName: util.SafeStringDatasource(owner.Ownername),
				Rank:      util.SafeStringDatasource(owner.Rank),
			})
		}
	}
	roleState.Owners = owners
}

// HandleConditionalRequestAttributes sets only the request attributes that were provided by the user
// This prevents unused attributes from appearing in the Terraform state
func (d *RolesDataSource) HandleConditionalRequestAttributes(ctx context.Context, req datasource.ReadRequest, state *RolesDataSourceModel) {
	// Get the original configuration from the request
	var config RolesDataSourceModel
	configDiagnostics := req.Config.Get(ctx, &config)
	if configDiagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting conditional state management for request attributes")

	// Define attribute mappings for conditional state management
	attributeMappings := []struct {
		name        string
		configField *types.String
		stateField  *types.String
	}{
		{"requested_object", &config.RequestedObject, &state.RequestedObject},
		{"username", &config.Username, &state.Username},
		{"role_type", &config.RoleType, &state.RoleType},
		{"requestable", &config.Requestable, &state.Requestable},
		{"status", &config.Status, &state.Status},
		{"role_name", &config.RoleName, &state.RoleName},
		{"description", &config.Description, &state.Description},
		{"display_name", &config.DisplayName, &state.DisplayName},
		{"glossary", &config.Glossary, &state.Glossary},
		{"mining_instance", &config.MiningInstance, &state.MiningInstance},
		{"risk", &config.Risk, &state.Risk},
		{"update_user", &config.UpdateUser, &state.UpdateUser},
		{"system_id", &config.SystemId, &state.SystemId},
		{"sox_critical", &config.SoxCritical, &state.SoxCritical},
		{"sys_critical", &config.SysCritical, &state.SysCritical},
		{"level", &config.Level, &state.Level},
		{"privileged", &config.Privileged, &state.Privileged},
		{"confidentiality", &config.Confidentiality, &state.Confidentiality},
		{"max", &config.Max, &state.Max},
		{"offset", &config.Offset, &state.Offset},
		{"role_query", &config.RoleQuery, &state.RoleQuery},
		{"hide_blank_values", &config.HideBlankValues, &state.HideBlankValues},
	}

	// Apply conditional logic: if config field is null, set state field to null
	nulledCount := 0
	for _, mapping := range attributeMappings {
		if mapping.configField.IsNull() {
			*mapping.stateField = types.StringNull()
			nulledCount++
		}
	}

	tflog.Debug(ctx, "Conditional state management completed", map[string]interface{}{
		"attributes_nulled": nulledCount,
		"total_attributes":  len(attributeMappings),
	})

	// Handle custom properties using reflection
	d.HandleConditionalCustomProperties(&config, state)
}

// HandleConditionalCustomProperties handles custom properties conditionally
func (d *RolesDataSource) HandleConditionalCustomProperties(config *RolesDataSourceModel, state *RolesDataSourceModel) {
	// Use reflection to handle custom properties dynamically
	configValue := reflect.ValueOf(config).Elem()
	stateValue := reflect.ValueOf(state).Elem()

	nulledCustomProps := 0
	for i := 1; i <= 60; i++ {
		fieldName := fmt.Sprintf("CustomProperty%d", i)

		configField := configValue.FieldByName(fieldName)
		stateField := stateValue.FieldByName(fieldName)

		if configField.IsValid() && stateField.IsValid() && stateField.CanSet() {
			configTfString := configField.Interface().(types.String)
			if configTfString.IsNull() {
				stateField.Set(reflect.ValueOf(types.StringNull()))
				nulledCustomProps++
			}
		}
	}

	if nulledCustomProps > 0 {
		tflog.Debug(context.Background(), "Custom properties conditional state management completed", map[string]interface{}{
			"custom_properties_nulled": nulledCustomProps,
		})
	}
}

// HandleRolesAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, roledetails are removed from state to prevent sensitive data exposure
// When authenticate=true, all roledetails are returned in state
func (d *RolesDataSource) HandleRolesAuthenticationLogic(state *RolesDataSourceModel, resp *datasource.ReadResponse) {
	// Check if role details are empty regardless of authenticate value
	if len(state.Roledetails) == 0 {
		tflog.Warn(context.Background(), "No role details found from API")

		// Use API message if available, otherwise use default message
		message := "No role details were returned from the API. Please check the filter attributes set in the datasource configuration."
		if !state.Message.IsNull() && state.Message.ValueString() != "" {
			message = state.Message.ValueString()
		}

		resp.Diagnostics.AddError(
			"No Data Found",
			message,
		)
	}

	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all role details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all role details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing role details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; role details will be removed from state.",
			)
			state.Roledetails = nil
		}
	}
}
