# GetEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username to filter entitlements | [optional] 
**Entitlementtype** | Pointer to **string** | Type of entitlement to filter | [optional] 
**EntitlementValue** | Pointer to **string** | Entitlement value | [optional] 
**Endpoint** | Pointer to **string** | Name of the endpoint to filter entitlements | [optional] 
**RequestedObject** | Pointer to **string** | Type of object requested | [optional] 
**Max** | Pointer to **int32** | Maximum number of results to return | [optional] 
**Offset** | Pointer to **int32** | Number of results to skip | [optional] 
**EntitlementResponseFields** | Pointer to **string** | Comma-separated list of entitlement fields to return | [optional] 
**UserResponseFields** | Pointer to **string** | Comma-separated list of user fields to return | [optional] 
**Userfiltercriteria** | Pointer to **string** | Filter criteria for users | [optional] 
**Accountname** | Pointer to **string** | Account name to filter by | [optional] 
**Entownerwithrank** | Pointer to **string** | If \&quot;true\&quot;, returns the list of owners with owner rank for every entitlementValue | [optional] 
**Returnentitlementmap** | Pointer to **string** | If true, entitlementmap details will be returned (default is false) | [optional] 
**Exactmatch** | Pointer to **string** | Default is true. If given as false, it will search based on similar matches | [optional] 
**Entitlementfiltercriteria** | Pointer to **string** | Filter criteria for entitlements | [optional] 
**EntQuery** | Pointer to **string** | Query to support only entitlement_values parameters. Sample - \&quot;ent.description &#x3D; &#39;Desc&#39; or ent.displayname like &#39;%display%&#39;\&quot; | [optional] 

## Methods

### NewGetEntitlementRequest

`func NewGetEntitlementRequest() *GetEntitlementRequest`

NewGetEntitlementRequest instantiates a new GetEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlementRequestWithDefaults

`func NewGetEntitlementRequestWithDefaults() *GetEntitlementRequest`

NewGetEntitlementRequestWithDefaults instantiates a new GetEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GetEntitlementRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetEntitlementRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetEntitlementRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetEntitlementRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEntitlementtype

`func (o *GetEntitlementRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *GetEntitlementRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *GetEntitlementRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.

### HasEntitlementtype

`func (o *GetEntitlementRequest) HasEntitlementtype() bool`

HasEntitlementtype returns a boolean if a field has been set.

### GetEntitlementValue

`func (o *GetEntitlementRequest) GetEntitlementValue() string`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *GetEntitlementRequest) GetEntitlementValueOk() (*string, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *GetEntitlementRequest) SetEntitlementValue(v string)`

SetEntitlementValue sets EntitlementValue field to given value.

### HasEntitlementValue

`func (o *GetEntitlementRequest) HasEntitlementValue() bool`

HasEntitlementValue returns a boolean if a field has been set.

### GetEndpoint

`func (o *GetEntitlementRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetEntitlementRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetEntitlementRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GetEntitlementRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetRequestedObject

`func (o *GetEntitlementRequest) GetRequestedObject() string`

GetRequestedObject returns the RequestedObject field if non-nil, zero value otherwise.

### GetRequestedObjectOk

`func (o *GetEntitlementRequest) GetRequestedObjectOk() (*string, bool)`

GetRequestedObjectOk returns a tuple with the RequestedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedObject

`func (o *GetEntitlementRequest) SetRequestedObject(v string)`

SetRequestedObject sets RequestedObject field to given value.

### HasRequestedObject

`func (o *GetEntitlementRequest) HasRequestedObject() bool`

HasRequestedObject returns a boolean if a field has been set.

### GetMax

`func (o *GetEntitlementRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetEntitlementRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetEntitlementRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetEntitlementRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetEntitlementRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetEntitlementRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetEntitlementRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetEntitlementRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetEntitlementResponseFields

`func (o *GetEntitlementRequest) GetEntitlementResponseFields() string`

GetEntitlementResponseFields returns the EntitlementResponseFields field if non-nil, zero value otherwise.

### GetEntitlementResponseFieldsOk

`func (o *GetEntitlementRequest) GetEntitlementResponseFieldsOk() (*string, bool)`

GetEntitlementResponseFieldsOk returns a tuple with the EntitlementResponseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementResponseFields

`func (o *GetEntitlementRequest) SetEntitlementResponseFields(v string)`

SetEntitlementResponseFields sets EntitlementResponseFields field to given value.

### HasEntitlementResponseFields

`func (o *GetEntitlementRequest) HasEntitlementResponseFields() bool`

HasEntitlementResponseFields returns a boolean if a field has been set.

### GetUserResponseFields

`func (o *GetEntitlementRequest) GetUserResponseFields() string`

GetUserResponseFields returns the UserResponseFields field if non-nil, zero value otherwise.

### GetUserResponseFieldsOk

`func (o *GetEntitlementRequest) GetUserResponseFieldsOk() (*string, bool)`

GetUserResponseFieldsOk returns a tuple with the UserResponseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResponseFields

`func (o *GetEntitlementRequest) SetUserResponseFields(v string)`

SetUserResponseFields sets UserResponseFields field to given value.

### HasUserResponseFields

`func (o *GetEntitlementRequest) HasUserResponseFields() bool`

HasUserResponseFields returns a boolean if a field has been set.

### GetUserfiltercriteria

`func (o *GetEntitlementRequest) GetUserfiltercriteria() string`

GetUserfiltercriteria returns the Userfiltercriteria field if non-nil, zero value otherwise.

### GetUserfiltercriteriaOk

`func (o *GetEntitlementRequest) GetUserfiltercriteriaOk() (*string, bool)`

GetUserfiltercriteriaOk returns a tuple with the Userfiltercriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfiltercriteria

`func (o *GetEntitlementRequest) SetUserfiltercriteria(v string)`

SetUserfiltercriteria sets Userfiltercriteria field to given value.

### HasUserfiltercriteria

`func (o *GetEntitlementRequest) HasUserfiltercriteria() bool`

HasUserfiltercriteria returns a boolean if a field has been set.

### GetAccountname

`func (o *GetEntitlementRequest) GetAccountname() string`

GetAccountname returns the Accountname field if non-nil, zero value otherwise.

### GetAccountnameOk

`func (o *GetEntitlementRequest) GetAccountnameOk() (*string, bool)`

GetAccountnameOk returns a tuple with the Accountname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountname

`func (o *GetEntitlementRequest) SetAccountname(v string)`

SetAccountname sets Accountname field to given value.

### HasAccountname

`func (o *GetEntitlementRequest) HasAccountname() bool`

HasAccountname returns a boolean if a field has been set.

### GetEntownerwithrank

`func (o *GetEntitlementRequest) GetEntownerwithrank() string`

GetEntownerwithrank returns the Entownerwithrank field if non-nil, zero value otherwise.

### GetEntownerwithrankOk

`func (o *GetEntitlementRequest) GetEntownerwithrankOk() (*string, bool)`

GetEntownerwithrankOk returns a tuple with the Entownerwithrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntownerwithrank

`func (o *GetEntitlementRequest) SetEntownerwithrank(v string)`

SetEntownerwithrank sets Entownerwithrank field to given value.

### HasEntownerwithrank

`func (o *GetEntitlementRequest) HasEntownerwithrank() bool`

HasEntownerwithrank returns a boolean if a field has been set.

### GetReturnentitlementmap

`func (o *GetEntitlementRequest) GetReturnentitlementmap() string`

GetReturnentitlementmap returns the Returnentitlementmap field if non-nil, zero value otherwise.

### GetReturnentitlementmapOk

`func (o *GetEntitlementRequest) GetReturnentitlementmapOk() (*string, bool)`

GetReturnentitlementmapOk returns a tuple with the Returnentitlementmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnentitlementmap

`func (o *GetEntitlementRequest) SetReturnentitlementmap(v string)`

SetReturnentitlementmap sets Returnentitlementmap field to given value.

### HasReturnentitlementmap

`func (o *GetEntitlementRequest) HasReturnentitlementmap() bool`

HasReturnentitlementmap returns a boolean if a field has been set.

### GetExactmatch

`func (o *GetEntitlementRequest) GetExactmatch() string`

GetExactmatch returns the Exactmatch field if non-nil, zero value otherwise.

### GetExactmatchOk

`func (o *GetEntitlementRequest) GetExactmatchOk() (*string, bool)`

GetExactmatchOk returns a tuple with the Exactmatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactmatch

`func (o *GetEntitlementRequest) SetExactmatch(v string)`

SetExactmatch sets Exactmatch field to given value.

### HasExactmatch

`func (o *GetEntitlementRequest) HasExactmatch() bool`

HasExactmatch returns a boolean if a field has been set.

### GetEntitlementfiltercriteria

`func (o *GetEntitlementRequest) GetEntitlementfiltercriteria() string`

GetEntitlementfiltercriteria returns the Entitlementfiltercriteria field if non-nil, zero value otherwise.

### GetEntitlementfiltercriteriaOk

`func (o *GetEntitlementRequest) GetEntitlementfiltercriteriaOk() (*string, bool)`

GetEntitlementfiltercriteriaOk returns a tuple with the Entitlementfiltercriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementfiltercriteria

`func (o *GetEntitlementRequest) SetEntitlementfiltercriteria(v string)`

SetEntitlementfiltercriteria sets Entitlementfiltercriteria field to given value.

### HasEntitlementfiltercriteria

`func (o *GetEntitlementRequest) HasEntitlementfiltercriteria() bool`

HasEntitlementfiltercriteria returns a boolean if a field has been set.

### GetEntQuery

`func (o *GetEntitlementRequest) GetEntQuery() string`

GetEntQuery returns the EntQuery field if non-nil, zero value otherwise.

### GetEntQueryOk

`func (o *GetEntitlementRequest) GetEntQueryOk() (*string, bool)`

GetEntQueryOk returns a tuple with the EntQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntQuery

`func (o *GetEntitlementRequest) SetEntQuery(v string)`

SetEntQuery sets EntQuery field to given value.

### HasEntQuery

`func (o *GetEntitlementRequest) HasEntQuery() bool`

HasEntQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


