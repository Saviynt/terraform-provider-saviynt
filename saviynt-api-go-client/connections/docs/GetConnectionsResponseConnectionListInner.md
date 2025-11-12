# GetConnectionsResponseConnectionListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CONNECTIONNAME** | Pointer to **string** | Name of the Connection | [optional] 
**CONNECTIONTYPE** | Pointer to **string** | Type of the Connection | [optional] 
**CONNECTIONDESCRIPTION** | Pointer to **string** | Connection description | [optional] 
**STATUS** | Pointer to **int32** | Connection status | [optional] 
**CREATEDBY** | Pointer to **string** | User who created the connection | [optional] 
**CREATEDON** | Pointer to **string** | Connection creation time | [optional] 
**UPDATEDBY** | Pointer to **string** | User who updated the connection | [optional] 
**UPDATEDON** | Pointer to **string** | Time of connection updation | [optional] 

## Methods

### NewGetConnectionsResponseConnectionListInner

`func NewGetConnectionsResponseConnectionListInner() *GetConnectionsResponseConnectionListInner`

NewGetConnectionsResponseConnectionListInner instantiates a new GetConnectionsResponseConnectionListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionsResponseConnectionListInnerWithDefaults

`func NewGetConnectionsResponseConnectionListInnerWithDefaults() *GetConnectionsResponseConnectionListInner`

NewGetConnectionsResponseConnectionListInnerWithDefaults instantiates a new GetConnectionsResponseConnectionListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCONNECTIONNAME

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONNAME() string`

GetCONNECTIONNAME returns the CONNECTIONNAME field if non-nil, zero value otherwise.

### GetCONNECTIONNAMEOk

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONNAMEOk() (*string, bool)`

GetCONNECTIONNAMEOk returns a tuple with the CONNECTIONNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONNAME

`func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONNAME(v string)`

SetCONNECTIONNAME sets CONNECTIONNAME field to given value.

### HasCONNECTIONNAME

`func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONNAME() bool`

HasCONNECTIONNAME returns a boolean if a field has been set.

### GetCONNECTIONTYPE

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONTYPE() string`

GetCONNECTIONTYPE returns the CONNECTIONTYPE field if non-nil, zero value otherwise.

### GetCONNECTIONTYPEOk

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONTYPEOk() (*string, bool)`

GetCONNECTIONTYPEOk returns a tuple with the CONNECTIONTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONTYPE

`func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONTYPE(v string)`

SetCONNECTIONTYPE sets CONNECTIONTYPE field to given value.

### HasCONNECTIONTYPE

`func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONTYPE() bool`

HasCONNECTIONTYPE returns a boolean if a field has been set.

### GetCONNECTIONDESCRIPTION

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONDESCRIPTION() string`

GetCONNECTIONDESCRIPTION returns the CONNECTIONDESCRIPTION field if non-nil, zero value otherwise.

### GetCONNECTIONDESCRIPTIONOk

`func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONDESCRIPTIONOk() (*string, bool)`

GetCONNECTIONDESCRIPTIONOk returns a tuple with the CONNECTIONDESCRIPTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONDESCRIPTION

`func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONDESCRIPTION(v string)`

SetCONNECTIONDESCRIPTION sets CONNECTIONDESCRIPTION field to given value.

### HasCONNECTIONDESCRIPTION

`func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONDESCRIPTION() bool`

HasCONNECTIONDESCRIPTION returns a boolean if a field has been set.

### GetSTATUS

`func (o *GetConnectionsResponseConnectionListInner) GetSTATUS() int32`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *GetConnectionsResponseConnectionListInner) GetSTATUSOk() (*int32, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *GetConnectionsResponseConnectionListInner) SetSTATUS(v int32)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *GetConnectionsResponseConnectionListInner) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetCREATEDBY

`func (o *GetConnectionsResponseConnectionListInner) GetCREATEDBY() string`

GetCREATEDBY returns the CREATEDBY field if non-nil, zero value otherwise.

### GetCREATEDBYOk

`func (o *GetConnectionsResponseConnectionListInner) GetCREATEDBYOk() (*string, bool)`

GetCREATEDBYOk returns a tuple with the CREATEDBY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEDBY

`func (o *GetConnectionsResponseConnectionListInner) SetCREATEDBY(v string)`

SetCREATEDBY sets CREATEDBY field to given value.

### HasCREATEDBY

`func (o *GetConnectionsResponseConnectionListInner) HasCREATEDBY() bool`

HasCREATEDBY returns a boolean if a field has been set.

### GetCREATEDON

`func (o *GetConnectionsResponseConnectionListInner) GetCREATEDON() string`

GetCREATEDON returns the CREATEDON field if non-nil, zero value otherwise.

### GetCREATEDONOk

`func (o *GetConnectionsResponseConnectionListInner) GetCREATEDONOk() (*string, bool)`

GetCREATEDONOk returns a tuple with the CREATEDON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEDON

`func (o *GetConnectionsResponseConnectionListInner) SetCREATEDON(v string)`

SetCREATEDON sets CREATEDON field to given value.

### HasCREATEDON

`func (o *GetConnectionsResponseConnectionListInner) HasCREATEDON() bool`

HasCREATEDON returns a boolean if a field has been set.

### GetUPDATEDBY

`func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDBY() string`

GetUPDATEDBY returns the UPDATEDBY field if non-nil, zero value otherwise.

### GetUPDATEDBYOk

`func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDBYOk() (*string, bool)`

GetUPDATEDBYOk returns a tuple with the UPDATEDBY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEDBY

`func (o *GetConnectionsResponseConnectionListInner) SetUPDATEDBY(v string)`

SetUPDATEDBY sets UPDATEDBY field to given value.

### HasUPDATEDBY

`func (o *GetConnectionsResponseConnectionListInner) HasUPDATEDBY() bool`

HasUPDATEDBY returns a boolean if a field has been set.

### GetUPDATEDON

`func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDON() string`

GetUPDATEDON returns the UPDATEDON field if non-nil, zero value otherwise.

### GetUPDATEDONOk

`func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDONOk() (*string, bool)`

GetUPDATEDONOk returns a tuple with the UPDATEDON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEDON

`func (o *GetConnectionsResponseConnectionListInner) SetUPDATEDON(v string)`

SetUPDATEDON sets UPDATEDON field to given value.

### HasUPDATEDON

`func (o *GetConnectionsResponseConnectionListInner) HasUPDATEDON() bool`

HasUPDATEDON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


