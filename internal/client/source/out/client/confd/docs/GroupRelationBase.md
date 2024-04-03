# GroupRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the group | [optional] [readonly]
**Name** | Pointer to **string** | The name of the group | [optional] [readonly]
**Uuid** | Pointer to **string** | Group UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]

## Methods

### NewGroupRelationBase

`func NewGroupRelationBase() *GroupRelationBase`

NewGroupRelationBase instantiates a new GroupRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRelationBaseWithDefaults

`func NewGroupRelationBaseWithDefaults() *GroupRelationBase`

NewGroupRelationBaseWithDefaults instantiates a new GroupRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GroupRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRelationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupRelationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *GroupRelationBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupRelationBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupRelationBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GroupRelationBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
