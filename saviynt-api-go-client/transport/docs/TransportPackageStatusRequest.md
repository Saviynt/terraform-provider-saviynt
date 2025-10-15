# TransportPackageStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | enter the type of transport request (export, import, or transfer) | 
**Filename** | **string** | enter the transport package zip file name in full (including the .zip extension) | 
**Requestid** | Pointer to **string** | (required only if the operation is import) - enter the request ID generated during the submission of the import request | [optional] 

## Methods

### NewTransportPackageStatusRequest

`func NewTransportPackageStatusRequest(operation string, filename string, ) *TransportPackageStatusRequest`

NewTransportPackageStatusRequest instantiates a new TransportPackageStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPackageStatusRequestWithDefaults

`func NewTransportPackageStatusRequestWithDefaults() *TransportPackageStatusRequest`

NewTransportPackageStatusRequestWithDefaults instantiates a new TransportPackageStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *TransportPackageStatusRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TransportPackageStatusRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TransportPackageStatusRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetFilename

`func (o *TransportPackageStatusRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TransportPackageStatusRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TransportPackageStatusRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetRequestid

`func (o *TransportPackageStatusRequest) GetRequestid() string`

GetRequestid returns the Requestid field if non-nil, zero value otherwise.

### GetRequestidOk

`func (o *TransportPackageStatusRequest) GetRequestidOk() (*string, bool)`

GetRequestidOk returns a tuple with the Requestid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestid

`func (o *TransportPackageStatusRequest) SetRequestid(v string)`

SetRequestid sets Requestid field to given value.

### HasRequestid

`func (o *TransportPackageStatusRequest) HasRequestid() bool`

HasRequestid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


