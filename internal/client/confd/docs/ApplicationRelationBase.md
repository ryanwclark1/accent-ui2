# ApplicationRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Application name | [optional]
**Uuid** | Pointer to **string** | Application UUID | [optional] [readonly]

## Methods

### NewApplicationRelationBase

`func NewApplicationRelationBase() *ApplicationRelationBase`

NewApplicationRelationBase instantiates a new ApplicationRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRelationBaseWithDefaults

`func NewApplicationRelationBaseWithDefaults() *ApplicationRelationBase`

NewApplicationRelationBaseWithDefaults instantiates a new ApplicationRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRelationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationRelationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *ApplicationRelationBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApplicationRelationBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApplicationRelationBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApplicationRelationBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
