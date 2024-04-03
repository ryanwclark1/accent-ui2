# ContextRangeItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ContextRange**](ContextRange.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewContextRangeItems

`func NewContextRangeItems(total int32, ) *ContextRangeItems`

NewContextRangeItems instantiates a new ContextRangeItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRangeItemsWithDefaults

`func NewContextRangeItemsWithDefaults() *ContextRangeItems`

NewContextRangeItemsWithDefaults instantiates a new ContextRangeItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContextRangeItems) GetItems() []ContextRange`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContextRangeItems) GetItemsOk() (*[]ContextRange, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContextRangeItems) SetItems(v []ContextRange)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContextRangeItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *ContextRangeItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContextRangeItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContextRangeItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
