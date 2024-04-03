# ContextItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Context**](Context.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewContextItems

`func NewContextItems(total int32, ) *ContextItems`

NewContextItems instantiates a new ContextItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextItemsWithDefaults

`func NewContextItemsWithDefaults() *ContextItems`

NewContextItemsWithDefaults instantiates a new ContextItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContextItems) GetItems() []Context`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContextItems) GetItemsOk() (*[]Context, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContextItems) SetItems(v []Context)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContextItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *ContextItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContextItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContextItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
