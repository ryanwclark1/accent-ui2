# TimezoneItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Timezone**](Timezone.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewTimezoneItems

`func NewTimezoneItems() *TimezoneItems`

NewTimezoneItems instantiates a new TimezoneItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneItemsWithDefaults

`func NewTimezoneItemsWithDefaults() *TimezoneItems`

NewTimezoneItemsWithDefaults instantiates a new TimezoneItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TimezoneItems) GetItems() []Timezone`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TimezoneItems) GetItemsOk() (*[]Timezone, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TimezoneItems) SetItems(v []Timezone)`

SetItems sets Items field to given value.

### HasItems

`func (o *TimezoneItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *TimezoneItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TimezoneItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TimezoneItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TimezoneItems) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
