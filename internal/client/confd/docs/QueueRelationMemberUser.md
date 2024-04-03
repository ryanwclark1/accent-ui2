# QueueRelationMemberUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Priority** | Pointer to **int32** | Priority of user in the queue. Only used for linear ring strategy | [optional]

## Methods

### NewQueueRelationMemberUser

`func NewQueueRelationMemberUser() *QueueRelationMemberUser`

NewQueueRelationMemberUser instantiates a new QueueRelationMemberUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueRelationMemberUserWithDefaults

`func NewQueueRelationMemberUserWithDefaults() *QueueRelationMemberUser`

NewQueueRelationMemberUserWithDefaults instantiates a new QueueRelationMemberUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *QueueRelationMemberUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *QueueRelationMemberUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *QueueRelationMemberUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *QueueRelationMemberUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *QueueRelationMemberUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *QueueRelationMemberUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *QueueRelationMemberUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *QueueRelationMemberUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *QueueRelationMemberUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *QueueRelationMemberUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *QueueRelationMemberUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *QueueRelationMemberUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPriority

`func (o *QueueRelationMemberUser) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueueRelationMemberUser) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueueRelationMemberUser) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueueRelationMemberUser) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
