/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

package util

import (
	"fmt"
)

var ProviderDescription = fmt.Sprint("The Saviynt Terraform provider empowers you to leverage Terraform's declarative Infrastructure-as-Code (IaC) capabilities to provision, configure, and manage resources within the Saviynt Identity Cloud.<br/><br/>The provider needs to be configured with the correct credentials in the provider block to be used. For the resources and datasources supported, refer to the navigation menu on the left.")

var SecuritySystemDescription = "Create and manage security systems in Saviynt"
var EndpointDescription = "Create and manage endpoints in Saviynt"
var SecuritySystemDataSourceDescription = "Datasource to retrieve all security systems"
var EndpointDataSourceDescription = "Datasource to retrieve all endpoints"
var ConnDataSourceDescription = "Datasource to retrieve all connections"

var ADConnDescription = "Create and manage AD connector in Saviynt"
var ADSIConnDescription = "Create and manage ADSI connector in Saviynt"
var DBConnDescription = "Create and manage DB connector in Saviynt"
var EntraIDConnDescription = "Create and manage EntraID (AzureAD) connector in Saviynt"
var GithubRestConnDescription = "Create and manage Github Rest connector in Saviynt"
var RestConnDescription = "Create and manage REST connector in Saviynt"
var SAPConnDescription = "Create and manage SAP connector in Saviynt"
var SalesforceConnDescription = "Create and manage Salesforce connector in Saviynt"
var UnixConnDescription = "Create and manage Unix connector in Saviynt"
var WorkdayConnDescription = "Create and manage Workday connector in Saviynt"
var DynamicAttrDescription="Create and manage dynamic attributes in Saviynt"

var ADConnDataSourceDescription = "Retrieve the details for a given AD connector by its name or key"
var ADSIConnDataSourceDescription = "Retrieve the details for a given ADSI connector by its name or key"
var DBConnDataSourceDescription = "Retrieve the details for a given DB connector by its name or key"
var EntraIDConnDataSourceDescription = "Retrieve the details for a given EntraID (AzureAD) connector by its name or key"
var GithubRestConnDataSourceDescription = "Retrieve the details for a given Github REST connector by its name or key"
var RestConnDataSourceDescription = "Retrieve the details for a given REST connector by its name or key"
var SAPConnDataSourceDescription = "Retrieve the details for a given SAP connector by its name or key"
var SalesforceConnDataSourceDescription = "Retrieve the details for a given Salesforce connector by its name or key"
var UnixConnDataSourceDescription = "Retrieve the details for a given Unix connector by its name or key"
var WorkdayConnDataSourceDescription = "Retrieve the details for a given Workday connector by its name or key"

var FileEphemeralResourceDescription = "Provides ephemeral credentials by reading them from a local json file for use by Connector resources."
var EnvEphemeralResourceDescription = "Provides ephemeral credentials by reading them from a environment for use by Connector resources."