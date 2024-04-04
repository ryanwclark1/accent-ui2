# RegistrarItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Registrar**](Registrar.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewRegistrarItems

`func NewRegistrarItems() *RegistrarItems`

NewRegistrarItems instantiates a new RegistrarItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrarItemsWithDefaults

`func NewRegistrarItemsWithDefaults() *RegistrarItems`

NewRegistrarItemsWithDefaults instantiates a new RegistrarItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RegistrarItems) GetItems() []Registrar`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RegistrarItems) GetItemsOk() (*[]Registrar, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RegistrarItems) SetItems(v []Registrar)`

SetItems sets Items field to given value.

### HasItems

`func (o *RegistrarItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *RegistrarItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RegistrarItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RegistrarItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RegistrarItems) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
