# QueueMemberUserRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | Priority of user in the queue. Only used for linear ring strategy | [optional]

## Methods

### NewQueueMemberUserRelationBase

`func NewQueueMemberUserRelationBase() *QueueMemberUserRelationBase`

NewQueueMemberUserRelationBase instantiates a new QueueMemberUserRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueMemberUserRelationBaseWithDefaults

`func NewQueueMemberUserRelationBaseWithDefaults() *QueueMemberUserRelationBase`

NewQueueMemberUserRelationBaseWithDefaults instantiates a new QueueMemberUserRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *QueueMemberUserRelationBase) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueueMemberUserRelationBase) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueueMemberUserRelationBase) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueueMemberUserRelationBase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
