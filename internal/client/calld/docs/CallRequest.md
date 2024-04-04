# CallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**CallRequestDestination**](CallRequestDestination.md) |  |
**Source** | [**CallRequestSource**](CallRequestSource.md) |  |
**Variables** | Pointer to **map[string]interface{}** | Channel variables to set | [optional]

## Methods

### NewCallRequest

`func NewCallRequest(destination CallRequestDestination, source CallRequestSource, ) *CallRequest`

NewCallRequest instantiates a new CallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestWithDefaults

`func NewCallRequestWithDefaults() *CallRequest`

NewCallRequestWithDefaults instantiates a new CallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CallRequest) GetDestination() CallRequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CallRequest) GetDestinationOk() (*CallRequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CallRequest) SetDestination(v CallRequestDestination)`

SetDestination sets Destination field to given value.

### GetSource

`func (o *CallRequest) GetSource() CallRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CallRequest) GetSourceOk() (*CallRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CallRequest) SetSource(v CallRequestSource)`

SetSource sets Source field to given value.

### GetVariables

`func (o *CallRequest) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CallRequest) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CallRequest) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CallRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
