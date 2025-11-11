# ImportTransportPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updateuser** | Pointer to **string** | Username of the user importing the package | [optional] 
**Packagetoimport** | **string** | Complete path of the package that needs to be imported | 
**Businessjustification** | Pointer to **string** | Business justification for the import | [optional] 

## Methods

### NewImportTransportPackageRequest

`func NewImportTransportPackageRequest(packagetoimport string, ) *ImportTransportPackageRequest`

NewImportTransportPackageRequest instantiates a new ImportTransportPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTransportPackageRequestWithDefaults

`func NewImportTransportPackageRequestWithDefaults() *ImportTransportPackageRequest`

NewImportTransportPackageRequestWithDefaults instantiates a new ImportTransportPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateuser

`func (o *ImportTransportPackageRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *ImportTransportPackageRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *ImportTransportPackageRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *ImportTransportPackageRequest) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetPackagetoimport

`func (o *ImportTransportPackageRequest) GetPackagetoimport() string`

GetPackagetoimport returns the Packagetoimport field if non-nil, zero value otherwise.

### GetPackagetoimportOk

`func (o *ImportTransportPackageRequest) GetPackagetoimportOk() (*string, bool)`

GetPackagetoimportOk returns a tuple with the Packagetoimport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagetoimport

`func (o *ImportTransportPackageRequest) SetPackagetoimport(v string)`

SetPackagetoimport sets Packagetoimport field to given value.


### GetBusinessjustification

`func (o *ImportTransportPackageRequest) GetBusinessjustification() string`

GetBusinessjustification returns the Businessjustification field if non-nil, zero value otherwise.

### GetBusinessjustificationOk

`func (o *ImportTransportPackageRequest) GetBusinessjustificationOk() (*string, bool)`

GetBusinessjustificationOk returns a tuple with the Businessjustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessjustification

`func (o *ImportTransportPackageRequest) SetBusinessjustification(v string)`

SetBusinessjustification sets Businessjustification field to given value.

### HasBusinessjustification

`func (o *ImportTransportPackageRequest) HasBusinessjustification() bool`

HasBusinessjustification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


