# ExportTransportPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updateuser** | Pointer to **string** | Username of the user exporting the package | [optional] 
**Transportowner** | Pointer to **string** | Option to transport owners for selected objects | [optional] 
**Transportmembers** | Pointer to **string** | Option to transport members for selected objects such as SAV role | [optional] 
**Exportonline** | **string** | Determines if package needs to be exported online | 
**Exportpath** | **string** | Local path where export package will be generated (required if exportonline is false) | 
**Environmentname** | Pointer to **string** | Name of the environment (required if exportonline is true) | [optional] 
**Objectstoexport** | [**ExportTransportPackageRequestObjectstoexport**](ExportTransportPackageRequestObjectstoexport.md) |  | 
**Businessjustification** | Pointer to **string** | Business justification for the export | [optional] 

## Methods

### NewExportTransportPackageRequest

`func NewExportTransportPackageRequest(exportonline string, exportpath string, objectstoexport ExportTransportPackageRequestObjectstoexport, ) *ExportTransportPackageRequest`

NewExportTransportPackageRequest instantiates a new ExportTransportPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTransportPackageRequestWithDefaults

`func NewExportTransportPackageRequestWithDefaults() *ExportTransportPackageRequest`

NewExportTransportPackageRequestWithDefaults instantiates a new ExportTransportPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateuser

`func (o *ExportTransportPackageRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *ExportTransportPackageRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *ExportTransportPackageRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *ExportTransportPackageRequest) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetTransportowner

`func (o *ExportTransportPackageRequest) GetTransportowner() string`

GetTransportowner returns the Transportowner field if non-nil, zero value otherwise.

### GetTransportownerOk

`func (o *ExportTransportPackageRequest) GetTransportownerOk() (*string, bool)`

GetTransportownerOk returns a tuple with the Transportowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportowner

`func (o *ExportTransportPackageRequest) SetTransportowner(v string)`

SetTransportowner sets Transportowner field to given value.

### HasTransportowner

`func (o *ExportTransportPackageRequest) HasTransportowner() bool`

HasTransportowner returns a boolean if a field has been set.

### GetTransportmembers

`func (o *ExportTransportPackageRequest) GetTransportmembers() string`

GetTransportmembers returns the Transportmembers field if non-nil, zero value otherwise.

### GetTransportmembersOk

`func (o *ExportTransportPackageRequest) GetTransportmembersOk() (*string, bool)`

GetTransportmembersOk returns a tuple with the Transportmembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportmembers

`func (o *ExportTransportPackageRequest) SetTransportmembers(v string)`

SetTransportmembers sets Transportmembers field to given value.

### HasTransportmembers

`func (o *ExportTransportPackageRequest) HasTransportmembers() bool`

HasTransportmembers returns a boolean if a field has been set.

### GetExportonline

`func (o *ExportTransportPackageRequest) GetExportonline() string`

GetExportonline returns the Exportonline field if non-nil, zero value otherwise.

### GetExportonlineOk

`func (o *ExportTransportPackageRequest) GetExportonlineOk() (*string, bool)`

GetExportonlineOk returns a tuple with the Exportonline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportonline

`func (o *ExportTransportPackageRequest) SetExportonline(v string)`

SetExportonline sets Exportonline field to given value.


### GetExportpath

`func (o *ExportTransportPackageRequest) GetExportpath() string`

GetExportpath returns the Exportpath field if non-nil, zero value otherwise.

### GetExportpathOk

`func (o *ExportTransportPackageRequest) GetExportpathOk() (*string, bool)`

GetExportpathOk returns a tuple with the Exportpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportpath

`func (o *ExportTransportPackageRequest) SetExportpath(v string)`

SetExportpath sets Exportpath field to given value.


### GetEnvironmentname

`func (o *ExportTransportPackageRequest) GetEnvironmentname() string`

GetEnvironmentname returns the Environmentname field if non-nil, zero value otherwise.

### GetEnvironmentnameOk

`func (o *ExportTransportPackageRequest) GetEnvironmentnameOk() (*string, bool)`

GetEnvironmentnameOk returns a tuple with the Environmentname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentname

`func (o *ExportTransportPackageRequest) SetEnvironmentname(v string)`

SetEnvironmentname sets Environmentname field to given value.

### HasEnvironmentname

`func (o *ExportTransportPackageRequest) HasEnvironmentname() bool`

HasEnvironmentname returns a boolean if a field has been set.

### GetObjectstoexport

`func (o *ExportTransportPackageRequest) GetObjectstoexport() ExportTransportPackageRequestObjectstoexport`

GetObjectstoexport returns the Objectstoexport field if non-nil, zero value otherwise.

### GetObjectstoexportOk

`func (o *ExportTransportPackageRequest) GetObjectstoexportOk() (*ExportTransportPackageRequestObjectstoexport, bool)`

GetObjectstoexportOk returns a tuple with the Objectstoexport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectstoexport

`func (o *ExportTransportPackageRequest) SetObjectstoexport(v ExportTransportPackageRequestObjectstoexport)`

SetObjectstoexport sets Objectstoexport field to given value.


### GetBusinessjustification

`func (o *ExportTransportPackageRequest) GetBusinessjustification() string`

GetBusinessjustification returns the Businessjustification field if non-nil, zero value otherwise.

### GetBusinessjustificationOk

`func (o *ExportTransportPackageRequest) GetBusinessjustificationOk() (*string, bool)`

GetBusinessjustificationOk returns a tuple with the Businessjustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessjustification

`func (o *ExportTransportPackageRequest) SetBusinessjustification(v string)`

SetBusinessjustification sets Businessjustification field to given value.

### HasBusinessjustification

`func (o *ExportTransportPackageRequest) HasBusinessjustification() bool`

HasBusinessjustification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


