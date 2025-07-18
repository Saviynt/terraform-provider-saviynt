/*
Saviynt API

API for managing roles in Saviynt.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package roles

import (
	"encoding/json"
)

// checks if the UpdateEnterpriseRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEnterpriseRoleResponse{}

// UpdateEnterpriseRoleResponse struct for UpdateEnterpriseRoleResponse
type UpdateEnterpriseRoleResponse struct {
	Requestid *string `json:"requestid,omitempty"`
	Requestkey *string `json:"requestkey,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewUpdateEnterpriseRoleResponse instantiates a new UpdateEnterpriseRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnterpriseRoleResponse() *UpdateEnterpriseRoleResponse {
	this := UpdateEnterpriseRoleResponse{}
	return &this
}

// NewUpdateEnterpriseRoleResponseWithDefaults instantiates a new UpdateEnterpriseRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnterpriseRoleResponseWithDefaults() *UpdateEnterpriseRoleResponse {
	this := UpdateEnterpriseRoleResponse{}
	return &this
}

// GetRequestid returns the Requestid field value if set, zero value otherwise.
func (o *UpdateEnterpriseRoleResponse) GetRequestid() string {
	if o == nil || IsNil(o.Requestid) {
		var ret string
		return ret
	}
	return *o.Requestid
}

// GetRequestidOk returns a tuple with the Requestid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnterpriseRoleResponse) GetRequestidOk() (*string, bool) {
	if o == nil || IsNil(o.Requestid) {
		return nil, false
	}
	return o.Requestid, true
}

// HasRequestid returns a boolean if a field has been set.
func (o *UpdateEnterpriseRoleResponse) HasRequestid() bool {
	if o != nil && !IsNil(o.Requestid) {
		return true
	}

	return false
}

// SetRequestid gets a reference to the given string and assigns it to the Requestid field.
func (o *UpdateEnterpriseRoleResponse) SetRequestid(v string) {
	o.Requestid = &v
}

// GetRequestkey returns the Requestkey field value if set, zero value otherwise.
func (o *UpdateEnterpriseRoleResponse) GetRequestkey() string {
	if o == nil || IsNil(o.Requestkey) {
		var ret string
		return ret
	}
	return *o.Requestkey
}

// GetRequestkeyOk returns a tuple with the Requestkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnterpriseRoleResponse) GetRequestkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Requestkey) {
		return nil, false
	}
	return o.Requestkey, true
}

// HasRequestkey returns a boolean if a field has been set.
func (o *UpdateEnterpriseRoleResponse) HasRequestkey() bool {
	if o != nil && !IsNil(o.Requestkey) {
		return true
	}

	return false
}

// SetRequestkey gets a reference to the given string and assigns it to the Requestkey field.
func (o *UpdateEnterpriseRoleResponse) SetRequestkey(v string) {
	o.Requestkey = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *UpdateEnterpriseRoleResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnterpriseRoleResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *UpdateEnterpriseRoleResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *UpdateEnterpriseRoleResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateEnterpriseRoleResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnterpriseRoleResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateEnterpriseRoleResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateEnterpriseRoleResponse) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateEnterpriseRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEnterpriseRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requestid) {
		toSerialize["requestid"] = o.Requestid
	}
	if !IsNil(o.Requestkey) {
		toSerialize["requestkey"] = o.Requestkey
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUpdateEnterpriseRoleResponse struct {
	value *UpdateEnterpriseRoleResponse
	isSet bool
}

func (v NullableUpdateEnterpriseRoleResponse) Get() *UpdateEnterpriseRoleResponse {
	return v.value
}

func (v *NullableUpdateEnterpriseRoleResponse) Set(val *UpdateEnterpriseRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnterpriseRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnterpriseRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnterpriseRoleResponse(val *UpdateEnterpriseRoleResponse) *NullableUpdateEnterpriseRoleResponse {
	return &NullableUpdateEnterpriseRoleResponse{value: val, isSet: true}
}

func (v NullableUpdateEnterpriseRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnterpriseRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


