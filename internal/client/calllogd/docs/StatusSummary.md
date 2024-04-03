# StatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusConsumer** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**ServiceToken** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**TaskQueue** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]

## Methods

### NewStatusSummary

`func NewStatusSummary() *StatusSummary`

NewStatusSummary instantiates a new StatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSummaryWithDefaults

`func NewStatusSummaryWithDefaults() *StatusSummary`

NewStatusSummaryWithDefaults instantiates a new StatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusConsumer

`func (o *StatusSummary) GetBusConsumer() ComponentWithStatus`

GetBusConsumer returns the BusConsumer field if non-nil, zero value otherwise.

### GetBusConsumerOk

`func (o *StatusSummary) GetBusConsumerOk() (*ComponentWithStatus, bool)`

GetBusConsumerOk returns a tuple with the BusConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusConsumer

`func (o *StatusSummary) SetBusConsumer(v ComponentWithStatus)`

SetBusConsumer sets BusConsumer field to given value.

### HasBusConsumer

`func (o *StatusSummary) HasBusConsumer() bool`

HasBusConsumer returns a boolean if a field has been set.

### GetServiceToken

`func (o *StatusSummary) GetServiceToken() ComponentWithStatus`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *StatusSummary) GetServiceTokenOk() (*ComponentWithStatus, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *StatusSummary) SetServiceToken(v ComponentWithStatus)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *StatusSummary) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetTaskQueue

`func (o *StatusSummary) GetTaskQueue() ComponentWithStatus`

GetTaskQueue returns the TaskQueue field if non-nil, zero value otherwise.

### GetTaskQueueOk

`func (o *StatusSummary) GetTaskQueueOk() (*ComponentWithStatus, bool)`

GetTaskQueueOk returns a tuple with the TaskQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueue

`func (o *StatusSummary) SetTaskQueue(v ComponentWithStatus)`

SetTaskQueue sets TaskQueue field to given value.

### HasTaskQueue

`func (o *StatusSummary) HasTaskQueue() bool`

HasTaskQueue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
