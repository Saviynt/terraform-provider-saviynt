# CreateOrUpdateTriggersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | [**[]TriggerItem**](TriggerItem.md) |  | 

## Methods

### NewCreateOrUpdateTriggersRequest

`func NewCreateOrUpdateTriggersRequest(triggers []TriggerItem, ) *CreateOrUpdateTriggersRequest`

NewCreateOrUpdateTriggersRequest instantiates a new CreateOrUpdateTriggersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateTriggersRequestWithDefaults

`func NewCreateOrUpdateTriggersRequestWithDefaults() *CreateOrUpdateTriggersRequest`

NewCreateOrUpdateTriggersRequestWithDefaults instantiates a new CreateOrUpdateTriggersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *CreateOrUpdateTriggersRequest) GetTriggers() []TriggerItem`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CreateOrUpdateTriggersRequest) GetTriggersOk() (*[]TriggerItem, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CreateOrUpdateTriggersRequest) SetTriggers(v []TriggerItem)`

SetTriggers sets Triggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


