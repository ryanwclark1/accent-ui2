# QueueFallbacks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusyDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**CongestionDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**FailDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**NoanswerDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]

## Methods

### NewQueueFallbacks

`func NewQueueFallbacks() *QueueFallbacks`

NewQueueFallbacks instantiates a new QueueFallbacks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueFallbacksWithDefaults

`func NewQueueFallbacksWithDefaults() *QueueFallbacks`

NewQueueFallbacksWithDefaults instantiates a new QueueFallbacks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusyDestination

`func (o *QueueFallbacks) GetBusyDestination() DestinationType`

GetBusyDestination returns the BusyDestination field if non-nil, zero value otherwise.

### GetBusyDestinationOk

`func (o *QueueFallbacks) GetBusyDestinationOk() (*DestinationType, bool)`

GetBusyDestinationOk returns a tuple with the BusyDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyDestination

`func (o *QueueFallbacks) SetBusyDestination(v DestinationType)`

SetBusyDestination sets BusyDestination field to given value.

### HasBusyDestination

`func (o *QueueFallbacks) HasBusyDestination() bool`

HasBusyDestination returns a boolean if a field has been set.

### GetCongestionDestination

`func (o *QueueFallbacks) GetCongestionDestination() DestinationType`

GetCongestionDestination returns the CongestionDestination field if non-nil, zero value otherwise.

### GetCongestionDestinationOk

`func (o *QueueFallbacks) GetCongestionDestinationOk() (*DestinationType, bool)`

GetCongestionDestinationOk returns a tuple with the CongestionDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionDestination

`func (o *QueueFallbacks) SetCongestionDestination(v DestinationType)`

SetCongestionDestination sets CongestionDestination field to given value.

### HasCongestionDestination

`func (o *QueueFallbacks) HasCongestionDestination() bool`

HasCongestionDestination returns a boolean if a field has been set.

### GetFailDestination

`func (o *QueueFallbacks) GetFailDestination() DestinationType`

GetFailDestination returns the FailDestination field if non-nil, zero value otherwise.

### GetFailDestinationOk

`func (o *QueueFallbacks) GetFailDestinationOk() (*DestinationType, bool)`

GetFailDestinationOk returns a tuple with the FailDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDestination

`func (o *QueueFallbacks) SetFailDestination(v DestinationType)`

SetFailDestination sets FailDestination field to given value.

### HasFailDestination

`func (o *QueueFallbacks) HasFailDestination() bool`

HasFailDestination returns a boolean if a field has been set.

### GetNoanswerDestination

`func (o *QueueFallbacks) GetNoanswerDestination() DestinationType`

GetNoanswerDestination returns the NoanswerDestination field if non-nil, zero value otherwise.

### GetNoanswerDestinationOk

`func (o *QueueFallbacks) GetNoanswerDestinationOk() (*DestinationType, bool)`

GetNoanswerDestinationOk returns a tuple with the NoanswerDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoanswerDestination

`func (o *QueueFallbacks) SetNoanswerDestination(v DestinationType)`

SetNoanswerDestination sets NoanswerDestination field to given value.

### HasNoanswerDestination

`func (o *QueueFallbacks) HasNoanswerDestination() bool`

HasNoanswerDestination returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
