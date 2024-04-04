# ContextRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the context | [optional] [readonly]
**Name** | Pointer to **string** | The name used by Asterisk | [optional] [readonly]
**Uuid** | Pointer to **string** | The UUID of the context | [optional] [readonly]

## Methods

### NewContextRelationBase

`func NewContextRelationBase() *ContextRelationBase`

NewContextRelationBase instantiates a new ContextRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRelationBaseWithDefaults

`func NewContextRelationBaseWithDefaults() *ContextRelationBase`

NewContextRelationBaseWithDefaults instantiates a new ContextRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContextRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContextRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContextRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextRelationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContextRelationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *ContextRelationBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ContextRelationBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ContextRelationBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ContextRelationBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
