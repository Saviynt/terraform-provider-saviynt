# GetKeyStoreCertificateDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**CertificateDetails** | Pointer to [**[]CertificateDetail**](CertificateDetail.md) |  | [optional] 

## Methods

### NewGetKeyStoreCertificateDetailsResponse

`func NewGetKeyStoreCertificateDetailsResponse() *GetKeyStoreCertificateDetailsResponse`

NewGetKeyStoreCertificateDetailsResponse instantiates a new GetKeyStoreCertificateDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyStoreCertificateDetailsResponseWithDefaults

`func NewGetKeyStoreCertificateDetailsResponseWithDefaults() *GetKeyStoreCertificateDetailsResponse`

NewGetKeyStoreCertificateDetailsResponseWithDefaults instantiates a new GetKeyStoreCertificateDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *GetKeyStoreCertificateDetailsResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetKeyStoreCertificateDetailsResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetKeyStoreCertificateDetailsResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetKeyStoreCertificateDetailsResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetKeyStoreCertificateDetailsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetKeyStoreCertificateDetailsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetKeyStoreCertificateDetailsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetKeyStoreCertificateDetailsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCertificateDetails

`func (o *GetKeyStoreCertificateDetailsResponse) GetCertificateDetails() []CertificateDetail`

GetCertificateDetails returns the CertificateDetails field if non-nil, zero value otherwise.

### GetCertificateDetailsOk

`func (o *GetKeyStoreCertificateDetailsResponse) GetCertificateDetailsOk() (*[]CertificateDetail, bool)`

GetCertificateDetailsOk returns a tuple with the CertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateDetails

`func (o *GetKeyStoreCertificateDetailsResponse) SetCertificateDetails(v []CertificateDetail)`

SetCertificateDetails sets CertificateDetails field to given value.

### HasCertificateDetails

`func (o *GetKeyStoreCertificateDetailsResponse) HasCertificateDetails() bool`

HasCertificateDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


