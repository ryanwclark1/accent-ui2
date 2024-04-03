# GroupRelationMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**GroupRelationMemberUsers**](GroupRelationMemberUsers.md) |  | [optional]

## Methods

### NewGroupRelationMembers

`func NewGroupRelationMembers() *GroupRelationMembers`

NewGroupRelationMembers instantiates a new GroupRelationMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRelationMembersWithDefaults

`func NewGroupRelationMembersWithDefaults() *GroupRelationMembers`

NewGroupRelationMembersWithDefaults instantiates a new GroupRelationMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *GroupRelationMembers) GetMembers() GroupRelationMemberUsers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupRelationMembers) GetMembersOk() (*GroupRelationMemberUsers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupRelationMembers) SetMembers(v GroupRelationMemberUsers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupRelationMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
