# UserRelationLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Line ID | [optional] [readonly]
**Name** | Pointer to **string** | The name of the line | [optional] [readonly]
**EndpointCustom** | Pointer to [**EndpointCustomRelationBase**](EndpointCustomRelationBase.md) |  | [optional]
**EndpointSccp** | Pointer to [**EndpointSccpRelationBase**](EndpointSccpRelationBase.md) |  | [optional]
**EndpointSip** | Pointer to [**EndpointSipRelationBase**](EndpointSipRelationBase.md) |  | [optional]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]

## Methods

### NewUserRelationLine

`func NewUserRelationLine() *UserRelationLine`

NewUserRelationLine instantiates a new UserRelationLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelationLineWithDefaults

`func NewUserRelationLineWithDefaults() *UserRelationLine`

NewUserRelationLineWithDefaults instantiates a new UserRelationLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRelationLine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRelationLine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRelationLine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserRelationLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserRelationLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRelationLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRelationLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRelationLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *UserRelationLine) GetEndpointCustom() EndpointCustomRelationBase`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *UserRelationLine) GetEndpointCustomOk() (*EndpointCustomRelationBase, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *UserRelationLine) SetEndpointCustom(v EndpointCustomRelationBase)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *UserRelationLine) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

### GetEndpointSccp

`func (o *UserRelationLine) GetEndpointSccp() EndpointSccpRelationBase`

GetEndpointSccp returns the EndpointSccp field if non-nil, zero value otherwise.

### GetEndpointSccpOk

`func (o *UserRelationLine) GetEndpointSccpOk() (*EndpointSccpRelationBase, bool)`

GetEndpointSccpOk returns a tuple with the EndpointSccp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSccp

`func (o *UserRelationLine) SetEndpointSccp(v EndpointSccpRelationBase)`

SetEndpointSccp sets EndpointSccp field to given value.

### HasEndpointSccp

`func (o *UserRelationLine) HasEndpointSccp() bool`

HasEndpointSccp returns a boolean if a field has been set.

### GetEndpointSip

`func (o *UserRelationLine) GetEndpointSip() EndpointSipRelationBase`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *UserRelationLine) GetEndpointSipOk() (*EndpointSipRelationBase, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *UserRelationLine) SetEndpointSip(v EndpointSipRelationBase)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *UserRelationLine) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

### GetExtensions

`func (o *UserRelationLine) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UserRelationLine) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UserRelationLine) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UserRelationLine) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
