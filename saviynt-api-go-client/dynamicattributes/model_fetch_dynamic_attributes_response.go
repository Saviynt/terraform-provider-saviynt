/*
Dynamic Attribute Management API

Use this API to manage dynamic attributes in Saviynt Identity Cloud.  The Authorization header must have \"Bearer {token}\".

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynamicattributes

import (
	"encoding/json"
)

// checks if the FetchDynamicAttributesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchDynamicAttributesResponse{}

// FetchDynamicAttributesResponse struct for FetchDynamicAttributesResponse
type FetchDynamicAttributesResponse struct {
	// A message indicating the outcome of the operation.
	Msg *string `json:"msg,omitempty"`
	// An error code where '0' signifies success and '1' signifies an unsuccessful operation.
	Errorcode *string `json:"errorcode,omitempty"`
	// Total number of records displayed.
	Displaycount *int32 `json:"displaycount,omitempty"`
	// Total number of records available.
	Totalcount *int32 `json:"totalcount,omitempty"`
	Dynamicattributes *FetchDynamicAttributesResponseDynamicattributes `json:"dynamicattributes,omitempty"`
}

// NewFetchDynamicAttributesResponse instantiates a new FetchDynamicAttributesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchDynamicAttributesResponse() *FetchDynamicAttributesResponse {
	this := FetchDynamicAttributesResponse{}
	return &this
}

// NewFetchDynamicAttributesResponseWithDefaults instantiates a new FetchDynamicAttributesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchDynamicAttributesResponseWithDefaults() *FetchDynamicAttributesResponse {
	this := FetchDynamicAttributesResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *FetchDynamicAttributesResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDynamicAttributesResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *FetchDynamicAttributesResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *FetchDynamicAttributesResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetErrorcode returns the Errorcode field value if set, zero value otherwise.
func (o *FetchDynamicAttributesResponse) GetErrorcode() string {
	if o == nil || IsNil(o.Errorcode) {
		var ret string
		return ret
	}
	return *o.Errorcode
}

// GetErrorcodeOk returns a tuple with the Errorcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDynamicAttributesResponse) GetErrorcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Errorcode) {
		return nil, false
	}
	return o.Errorcode, true
}

// HasErrorcode returns a boolean if a field has been set.
func (o *FetchDynamicAttributesResponse) HasErrorcode() bool {
	if o != nil && !IsNil(o.Errorcode) {
		return true
	}

	return false
}

// SetErrorcode gets a reference to the given string and assigns it to the Errorcode field.
func (o *FetchDynamicAttributesResponse) SetErrorcode(v string) {
	o.Errorcode = &v
}

// GetDisplaycount returns the Displaycount field value if set, zero value otherwise.
func (o *FetchDynamicAttributesResponse) GetDisplaycount() int32 {
	if o == nil || IsNil(o.Displaycount) {
		var ret int32
		return ret
	}
	return *o.Displaycount
}

// GetDisplaycountOk returns a tuple with the Displaycount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDynamicAttributesResponse) GetDisplaycountOk() (*int32, bool) {
	if o == nil || IsNil(o.Displaycount) {
		return nil, false
	}
	return o.Displaycount, true
}

// HasDisplaycount returns a boolean if a field has been set.
func (o *FetchDynamicAttributesResponse) HasDisplaycount() bool {
	if o != nil && !IsNil(o.Displaycount) {
		return true
	}

	return false
}

// SetDisplaycount gets a reference to the given int32 and assigns it to the Displaycount field.
func (o *FetchDynamicAttributesResponse) SetDisplaycount(v int32) {
	o.Displaycount = &v
}

// GetTotalcount returns the Totalcount field value if set, zero value otherwise.
func (o *FetchDynamicAttributesResponse) GetTotalcount() int32 {
	if o == nil || IsNil(o.Totalcount) {
		var ret int32
		return ret
	}
	return *o.Totalcount
}

// GetTotalcountOk returns a tuple with the Totalcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDynamicAttributesResponse) GetTotalcountOk() (*int32, bool) {
	if o == nil || IsNil(o.Totalcount) {
		return nil, false
	}
	return o.Totalcount, true
}

// HasTotalcount returns a boolean if a field has been set.
func (o *FetchDynamicAttributesResponse) HasTotalcount() bool {
	if o != nil && !IsNil(o.Totalcount) {
		return true
	}

	return false
}

// SetTotalcount gets a reference to the given int32 and assigns it to the Totalcount field.
func (o *FetchDynamicAttributesResponse) SetTotalcount(v int32) {
	o.Totalcount = &v
}

// GetDynamicattributes returns the Dynamicattributes field value if set, zero value otherwise.
func (o *FetchDynamicAttributesResponse) GetDynamicattributes() FetchDynamicAttributesResponseDynamicattributes {
	if o == nil || IsNil(o.Dynamicattributes) {
		var ret FetchDynamicAttributesResponseDynamicattributes
		return ret
	}
	return *o.Dynamicattributes
}

// GetDynamicattributesOk returns a tuple with the Dynamicattributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDynamicAttributesResponse) GetDynamicattributesOk() (*FetchDynamicAttributesResponseDynamicattributes, bool) {
	if o == nil || IsNil(o.Dynamicattributes) {
		return nil, false
	}
	return o.Dynamicattributes, true
}

// HasDynamicattributes returns a boolean if a field has been set.
func (o *FetchDynamicAttributesResponse) HasDynamicattributes() bool {
	if o != nil && !IsNil(o.Dynamicattributes) {
		return true
	}

	return false
}

// SetDynamicattributes gets a reference to the given FetchDynamicAttributesResponseDynamicattributes and assigns it to the Dynamicattributes field.
func (o *FetchDynamicAttributesResponse) SetDynamicattributes(v FetchDynamicAttributesResponseDynamicattributes) {
	o.Dynamicattributes = &v
}

func (o FetchDynamicAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchDynamicAttributesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Errorcode) {
		toSerialize["errorcode"] = o.Errorcode
	}
	if !IsNil(o.Displaycount) {
		toSerialize["displaycount"] = o.Displaycount
	}
	if !IsNil(o.Totalcount) {
		toSerialize["totalcount"] = o.Totalcount
	}
	if !IsNil(o.Dynamicattributes) {
		toSerialize["dynamicattributes"] = o.Dynamicattributes
	}
	return toSerialize, nil
}

type NullableFetchDynamicAttributesResponse struct {
	value *FetchDynamicAttributesResponse
	isSet bool
}

func (v NullableFetchDynamicAttributesResponse) Get() *FetchDynamicAttributesResponse {
	return v.value
}

func (v *NullableFetchDynamicAttributesResponse) Set(val *FetchDynamicAttributesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchDynamicAttributesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchDynamicAttributesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchDynamicAttributesResponse(val *FetchDynamicAttributesResponse) *NullableFetchDynamicAttributesResponse {
	return &NullableFetchDynamicAttributesResponse{value: val, isSet: true}
}

func (v NullableFetchDynamicAttributesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchDynamicAttributesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


