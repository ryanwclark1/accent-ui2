# QueueMemberUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | Priority of user in the queue. Only used for linear ring strategy | [optional]

## Methods

### NewQueueMemberUser

`func NewQueueMemberUser() *QueueMemberUser`

NewQueueMemberUser instantiates a new QueueMemberUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueMemberUserWithDefaults

`func NewQueueMemberUserWithDefaults() *QueueMemberUser`

NewQueueMemberUserWithDefaults instantiates a new QueueMemberUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *QueueMemberUser) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueueMemberUser) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueueMemberUser) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueueMemberUser) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
