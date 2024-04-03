# PagingItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Paging**](Paging.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewPagingItems

`func NewPagingItems(total int32, ) *PagingItems`

NewPagingItems instantiates a new PagingItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingItemsWithDefaults

`func NewPagingItemsWithDefaults() *PagingItems`

NewPagingItemsWithDefaults instantiates a new PagingItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PagingItems) GetItems() []Paging`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagingItems) GetItemsOk() (*[]Paging, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagingItems) SetItems(v []Paging)`

SetItems sets Items field to given value.

### HasItems

`func (o *PagingItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *PagingItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PagingItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PagingItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
