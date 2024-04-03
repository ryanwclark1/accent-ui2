# ScheduleRelationQueues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incalls** | Pointer to [**[]QueueRelationBase**](QueueRelationBase.md) |  | [optional] [readonly]

## Methods

### NewScheduleRelationQueues

`func NewScheduleRelationQueues() *ScheduleRelationQueues`

NewScheduleRelationQueues instantiates a new ScheduleRelationQueues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRelationQueuesWithDefaults

`func NewScheduleRelationQueuesWithDefaults() *ScheduleRelationQueues`

NewScheduleRelationQueuesWithDefaults instantiates a new ScheduleRelationQueues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncalls

`func (o *ScheduleRelationQueues) GetIncalls() []QueueRelationBase`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *ScheduleRelationQueues) GetIncallsOk() (*[]QueueRelationBase, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *ScheduleRelationQueues) SetIncalls(v []QueueRelationBase)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *ScheduleRelationQueues) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
