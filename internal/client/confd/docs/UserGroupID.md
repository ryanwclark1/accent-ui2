# UserGroupID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewUserGroupID

`func NewUserGroupID() *UserGroupID`

NewUserGroupID instantiates a new UserGroupID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupIDWithDefaults

`func NewUserGroupIDWithDefaults() *UserGroupID`

NewUserGroupIDWithDefaults instantiates a new UserGroupID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupID) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupID) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupID) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupID) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UserGroupID) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserGroupID) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserGroupID) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserGroupID) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
