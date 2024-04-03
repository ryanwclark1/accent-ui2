# IncallItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Incall**](Incall.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewIncallItems

`func NewIncallItems(total int32, ) *IncallItems`

NewIncallItems instantiates a new IncallItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncallItemsWithDefaults

`func NewIncallItemsWithDefaults() *IncallItems`

NewIncallItemsWithDefaults instantiates a new IncallItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IncallItems) GetItems() []Incall`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IncallItems) GetItemsOk() (*[]Incall, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IncallItems) SetItems(v []Incall)`

SetItems sets Items field to given value.

### HasItems

`func (o *IncallItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *IncallItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncallItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncallItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
