# QueueRelationMemberAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]QueueRelationMemberAgent**](QueueRelationMemberAgent.md) |  | [optional] [readonly]
**Users** | Pointer to [**[]QueueRelationMemberUser**](QueueRelationMemberUser.md) |  | [optional] [readonly]

## Methods

### NewQueueRelationMemberAgents

`func NewQueueRelationMemberAgents() *QueueRelationMemberAgents`

NewQueueRelationMemberAgents instantiates a new QueueRelationMemberAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueRelationMemberAgentsWithDefaults

`func NewQueueRelationMemberAgentsWithDefaults() *QueueRelationMemberAgents`

NewQueueRelationMemberAgentsWithDefaults instantiates a new QueueRelationMemberAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *QueueRelationMemberAgents) GetAgents() []QueueRelationMemberAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *QueueRelationMemberAgents) GetAgentsOk() (*[]QueueRelationMemberAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *QueueRelationMemberAgents) SetAgents(v []QueueRelationMemberAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *QueueRelationMemberAgents) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetUsers

`func (o *QueueRelationMemberAgents) GetUsers() []QueueRelationMemberUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *QueueRelationMemberAgents) GetUsersOk() (*[]QueueRelationMemberUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *QueueRelationMemberAgents) SetUsers(v []QueueRelationMemberUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *QueueRelationMemberAgents) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
