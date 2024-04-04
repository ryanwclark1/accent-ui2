# UserRelationIncall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the incoming call | [optional] [readonly]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]

## Methods

### NewUserRelationIncall

`func NewUserRelationIncall() *UserRelationIncall`

NewUserRelationIncall instantiates a new UserRelationIncall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelationIncallWithDefaults

`func NewUserRelationIncallWithDefaults() *UserRelationIncall`

NewUserRelationIncallWithDefaults instantiates a new UserRelationIncall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRelationIncall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRelationIncall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRelationIncall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserRelationIncall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtensions

`func (o *UserRelationIncall) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UserRelationIncall) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UserRelationIncall) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UserRelationIncall) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
