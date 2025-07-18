/*
Dynamic Attribute Management API

Use this API to manage dynamic attributes in Saviynt Identity Cloud.  The Authorization header must have \"Bearer {token}\".

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynamicattributes

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateDynamicAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDynamicAttributeRequest{}

// CreateDynamicAttributeRequest struct for CreateDynamicAttributeRequest
type CreateDynamicAttributeRequest struct {
	// Name of the security systems
	Securitysystem string `json:"securitysystem"`
	// Name of the endpoint
	Endpoint string `json:"endpoint"`
	// Username
	Updateuser string `json:"updateuser"`
	Dynamicattributes []CreateDynamicAttributesInner `json:"dynamicattributes"`
}

type _CreateDynamicAttributeRequest CreateDynamicAttributeRequest

// NewCreateDynamicAttributeRequest instantiates a new CreateDynamicAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDynamicAttributeRequest(securitysystem string, endpoint string, updateuser string, dynamicattributes []CreateDynamicAttributesInner) *CreateDynamicAttributeRequest {
	this := CreateDynamicAttributeRequest{}
	this.Securitysystem = securitysystem
	this.Endpoint = endpoint
	this.Updateuser = updateuser
	this.Dynamicattributes = dynamicattributes
	return &this
}

// NewCreateDynamicAttributeRequestWithDefaults instantiates a new CreateDynamicAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDynamicAttributeRequestWithDefaults() *CreateDynamicAttributeRequest {
	this := CreateDynamicAttributeRequest{}
	return &this
}

// GetSecuritysystem returns the Securitysystem field value
func (o *CreateDynamicAttributeRequest) GetSecuritysystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Securitysystem
}

// GetSecuritysystemOk returns a tuple with the Securitysystem field value
// and a boolean to check if the value has been set.
func (o *CreateDynamicAttributeRequest) GetSecuritysystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Securitysystem, true
}

// SetSecuritysystem sets field value
func (o *CreateDynamicAttributeRequest) SetSecuritysystem(v string) {
	o.Securitysystem = v
}

// GetEndpoint returns the Endpoint field value
func (o *CreateDynamicAttributeRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CreateDynamicAttributeRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CreateDynamicAttributeRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetUpdateuser returns the Updateuser field value
func (o *CreateDynamicAttributeRequest) GetUpdateuser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updateuser
}

// GetUpdateuserOk returns a tuple with the Updateuser field value
// and a boolean to check if the value has been set.
func (o *CreateDynamicAttributeRequest) GetUpdateuserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updateuser, true
}

// SetUpdateuser sets field value
func (o *CreateDynamicAttributeRequest) SetUpdateuser(v string) {
	o.Updateuser = v
}

// GetDynamicattributes returns the Dynamicattributes field value
func (o *CreateDynamicAttributeRequest) GetDynamicattributes() []CreateDynamicAttributesInner {
	if o == nil {
		var ret []CreateDynamicAttributesInner
		return ret
	}

	return o.Dynamicattributes
}

// GetDynamicattributesOk returns a tuple with the Dynamicattributes field value
// and a boolean to check if the value has been set.
func (o *CreateDynamicAttributeRequest) GetDynamicattributesOk() ([]CreateDynamicAttributesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dynamicattributes, true
}

// SetDynamicattributes sets field value
func (o *CreateDynamicAttributeRequest) SetDynamicattributes(v []CreateDynamicAttributesInner) {
	o.Dynamicattributes = v
}

func (o CreateDynamicAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDynamicAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["securitysystem"] = o.Securitysystem
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["updateuser"] = o.Updateuser
	toSerialize["dynamicattributes"] = o.Dynamicattributes
	return toSerialize, nil
}

func (o *CreateDynamicAttributeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"securitysystem",
		"endpoint",
		"updateuser",
		"dynamicattributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateDynamicAttributeRequest := _CreateDynamicAttributeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDynamicAttributeRequest)

	if err != nil {
		return err
	}

	*o = CreateDynamicAttributeRequest(varCreateDynamicAttributeRequest)

	return err
}

type NullableCreateDynamicAttributeRequest struct {
	value *CreateDynamicAttributeRequest
	isSet bool
}

func (v NullableCreateDynamicAttributeRequest) Get() *CreateDynamicAttributeRequest {
	return v.value
}

func (v *NullableCreateDynamicAttributeRequest) Set(val *CreateDynamicAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDynamicAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDynamicAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDynamicAttributeRequest(val *CreateDynamicAttributeRequest) *NullableCreateDynamicAttributeRequest {
	return &NullableCreateDynamicAttributeRequest{value: val, isSet: true}
}

func (v NullableCreateDynamicAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDynamicAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


