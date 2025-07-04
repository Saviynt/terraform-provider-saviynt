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

package rolesutil

var RoleTypeMap = map[string]string{
	"1": "ENABLER",
	"2": "TRANSACTIONAL",
	"3": "FIREFIGHTER",
	"4": "ENTERPRISE",
	"5": "APPLICATION",
	"6": "ENTITLEMENT",
}

var SoxCriticalityMap = map[string]string{
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var SysCriticalMap = map[string]string{
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var PrivilegedMap = map[string]string{
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var ConfidentialityMap = map[string]string{
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}
