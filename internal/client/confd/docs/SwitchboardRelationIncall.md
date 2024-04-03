# SwitchboardRelationIncall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the incoming call | [optional] [readonly]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]

## Methods

### NewSwitchboardRelationIncall

`func NewSwitchboardRelationIncall() *SwitchboardRelationIncall`

NewSwitchboardRelationIncall instantiates a new SwitchboardRelationIncall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchboardRelationIncallWithDefaults

`func NewSwitchboardRelationIncallWithDefaults() *SwitchboardRelationIncall`

NewSwitchboardRelationIncallWithDefaults instantiates a new SwitchboardRelationIncall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SwitchboardRelationIncall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwitchboardRelationIncall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwitchboardRelationIncall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SwitchboardRelationIncall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtensions

`func (o *SwitchboardRelationIncall) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SwitchboardRelationIncall) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SwitchboardRelationIncall) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SwitchboardRelationIncall) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
