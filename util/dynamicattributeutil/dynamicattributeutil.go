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

package dynamicattributeutil

var AttributeTypeMap = map[string]string{
	"NUMBER":          "NUMBER",
	"STRING":          "STRING",
	"ENUM":            "ENUM",
	"BOOLEAN":         "BOOLEAN",
	"MULTIPLE":        "MULTIPLE SELECT FROM LIST",
	"SQL MULTISELECT": "MULTIPLE SELECT FROM SQL QUERY",
	"SQL ENUM":        "SINGLE SELECT FROM SQL QUERY",
	"PASSWORD":        "PASSWORD",
	"LARGE TEXT":      "LARGE TEXT",
	"CHECKBOX":        "CHECK BOX",
	"DATE":            "DATE",
}

func TranslateValue(input string, valueMap map[string]string) string {
	if input == "" {
		return ""
	}
	if val, ok := valueMap[input]; ok {
		return val
	}
	return input
}
