# ApplicationNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to [**[]ApplicationNodeCallRequest**](ApplicationNodeCallRequest.md) |  | [optional]

## Methods

### NewApplicationNodeRequest

`func NewApplicationNodeRequest() *ApplicationNodeRequest`

NewApplicationNodeRequest instantiates a new ApplicationNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationNodeRequestWithDefaults

`func NewApplicationNodeRequestWithDefaults() *ApplicationNodeRequest`

NewApplicationNodeRequestWithDefaults instantiates a new ApplicationNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *ApplicationNodeRequest) GetCalls() []ApplicationNodeCallRequest`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ApplicationNodeRequest) GetCallsOk() (*[]ApplicationNodeCallRequest, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ApplicationNodeRequest) SetCalls(v []ApplicationNodeCallRequest)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ApplicationNodeRequest) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
