# UserGroupsID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]UserGroupID**](UserGroupID.md) |  | [optional]

## Methods

### NewUserGroupsID

`func NewUserGroupsID() *UserGroupsID`

NewUserGroupsID instantiates a new UserGroupsID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupsIDWithDefaults

`func NewUserGroupsIDWithDefaults() *UserGroupsID`

NewUserGroupsIDWithDefaults instantiates a new UserGroupsID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *UserGroupsID) GetGroups() []UserGroupID`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserGroupsID) GetGroupsOk() (*[]UserGroupID, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserGroupsID) SetGroups(v []UserGroupID)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserGroupsID) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
