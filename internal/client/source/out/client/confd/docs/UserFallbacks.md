# UserFallbacks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusyDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**CongestionDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**FailDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**NoanswerDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]

## Methods

### NewUserFallbacks

`func NewUserFallbacks() *UserFallbacks`

NewUserFallbacks instantiates a new UserFallbacks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFallbacksWithDefaults

`func NewUserFallbacksWithDefaults() *UserFallbacks`

NewUserFallbacksWithDefaults instantiates a new UserFallbacks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusyDestination

`func (o *UserFallbacks) GetBusyDestination() DestinationType`

GetBusyDestination returns the BusyDestination field if non-nil, zero value otherwise.

### GetBusyDestinationOk

`func (o *UserFallbacks) GetBusyDestinationOk() (*DestinationType, bool)`

GetBusyDestinationOk returns a tuple with the BusyDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyDestination

`func (o *UserFallbacks) SetBusyDestination(v DestinationType)`

SetBusyDestination sets BusyDestination field to given value.

### HasBusyDestination

`func (o *UserFallbacks) HasBusyDestination() bool`

HasBusyDestination returns a boolean if a field has been set.

### GetCongestionDestination

`func (o *UserFallbacks) GetCongestionDestination() DestinationType`

GetCongestionDestination returns the CongestionDestination field if non-nil, zero value otherwise.

### GetCongestionDestinationOk

`func (o *UserFallbacks) GetCongestionDestinationOk() (*DestinationType, bool)`

GetCongestionDestinationOk returns a tuple with the CongestionDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionDestination

`func (o *UserFallbacks) SetCongestionDestination(v DestinationType)`

SetCongestionDestination sets CongestionDestination field to given value.

### HasCongestionDestination

`func (o *UserFallbacks) HasCongestionDestination() bool`

HasCongestionDestination returns a boolean if a field has been set.

### GetFailDestination

`func (o *UserFallbacks) GetFailDestination() DestinationType`

GetFailDestination returns the FailDestination field if non-nil, zero value otherwise.

### GetFailDestinationOk

`func (o *UserFallbacks) GetFailDestinationOk() (*DestinationType, bool)`

GetFailDestinationOk returns a tuple with the FailDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDestination

`func (o *UserFallbacks) SetFailDestination(v DestinationType)`

SetFailDestination sets FailDestination field to given value.

### HasFailDestination

`func (o *UserFallbacks) HasFailDestination() bool`

HasFailDestination returns a boolean if a field has been set.

### GetNoanswerDestination

`func (o *UserFallbacks) GetNoanswerDestination() DestinationType`

GetNoanswerDestination returns the NoanswerDestination field if non-nil, zero value otherwise.

### GetNoanswerDestinationOk

`func (o *UserFallbacks) GetNoanswerDestinationOk() (*DestinationType, bool)`

GetNoanswerDestinationOk returns a tuple with the NoanswerDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoanswerDestination

`func (o *UserFallbacks) SetNoanswerDestination(v DestinationType)`

SetNoanswerDestination sets NoanswerDestination field to given value.

### HasNoanswerDestination

`func (o *UserFallbacks) HasNoanswerDestination() bool`

HasNoanswerDestination returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
