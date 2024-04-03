# QueueMemberAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Penalty** | Pointer to **int32** | Agent&#39;s penalty. A priority used for distributing calls. | [optional]
**Priority** | Pointer to **int32** | Priority of agent in the queue. Only used for linear ring strategy | [optional]

## Methods

### NewQueueMemberAgent

`func NewQueueMemberAgent() *QueueMemberAgent`

NewQueueMemberAgent instantiates a new QueueMemberAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueMemberAgentWithDefaults

`func NewQueueMemberAgentWithDefaults() *QueueMemberAgent`

NewQueueMemberAgentWithDefaults instantiates a new QueueMemberAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPenalty

`func (o *QueueMemberAgent) GetPenalty() int32`

GetPenalty returns the Penalty field if non-nil, zero value otherwise.

### GetPenaltyOk

`func (o *QueueMemberAgent) GetPenaltyOk() (*int32, bool)`

GetPenaltyOk returns a tuple with the Penalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalty

`func (o *QueueMemberAgent) SetPenalty(v int32)`

SetPenalty sets Penalty field to given value.

### HasPenalty

`func (o *QueueMemberAgent) HasPenalty() bool`

HasPenalty returns a boolean if a field has been set.

### GetPriority

`func (o *QueueMemberAgent) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueueMemberAgent) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueueMemberAgent) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueueMemberAgent) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
