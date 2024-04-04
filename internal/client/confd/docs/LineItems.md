# LineItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]LineView**](LineView.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewLineItems

`func NewLineItems(total int32, ) *LineItems`

NewLineItems instantiates a new LineItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemsWithDefaults

`func NewLineItemsWithDefaults() *LineItems`

NewLineItemsWithDefaults instantiates a new LineItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *LineItems) GetItems() []LineView`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LineItems) GetItemsOk() (*[]LineView, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LineItems) SetItems(v []LineView)`

SetItems sets Items field to given value.

### HasItems

`func (o *LineItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *LineItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LineItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LineItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
