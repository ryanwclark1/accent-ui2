# GroupMemberUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | priority of user in the group. Only used for linear ring strategy | [optional]
**Uuid** | **string** |  |

## Methods

### NewGroupMemberUser

`func NewGroupMemberUser(uuid string, ) *GroupMemberUser`

NewGroupMemberUser instantiates a new GroupMemberUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberUserWithDefaults

`func NewGroupMemberUserWithDefaults() *GroupMemberUser`

NewGroupMemberUserWithDefaults instantiates a new GroupMemberUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *GroupMemberUser) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GroupMemberUser) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GroupMemberUser) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GroupMemberUser) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUuid

`func (o *GroupMemberUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupMemberUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupMemberUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
