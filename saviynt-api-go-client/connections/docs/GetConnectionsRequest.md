# GetConnectionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectionname** | Pointer to **string** | Specify the connection name for which you want to view the details. | [optional] 
**Connectiontype** | Pointer to **string** | Specify a connection type to view all the connection in EIC for the connection type. | [optional] 
**Max** | Pointer to **string** | Specify the maximum number of results that can be returned in the response. | [optional] 
**Offset** | Pointer to **string** | Specify the number of rows of the result to skip before any rows are retrieved. | [optional] 

## Methods

### NewGetConnectionsRequest

`func NewGetConnectionsRequest() *GetConnectionsRequest`

NewGetConnectionsRequest instantiates a new GetConnectionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionsRequestWithDefaults

`func NewGetConnectionsRequestWithDefaults() *GetConnectionsRequest`

NewGetConnectionsRequestWithDefaults instantiates a new GetConnectionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionname

`func (o *GetConnectionsRequest) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *GetConnectionsRequest) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *GetConnectionsRequest) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *GetConnectionsRequest) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetConnectiontype

`func (o *GetConnectionsRequest) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *GetConnectionsRequest) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *GetConnectionsRequest) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.

### HasConnectiontype

`func (o *GetConnectionsRequest) HasConnectiontype() bool`

HasConnectiontype returns a boolean if a field has been set.

### GetMax

`func (o *GetConnectionsRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetConnectionsRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetConnectionsRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetConnectionsRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetConnectionsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetConnectionsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetConnectionsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetConnectionsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


