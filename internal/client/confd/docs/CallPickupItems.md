# CallPickupItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CallPickup**](CallPickup.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewCallPickupItems

`func NewCallPickupItems(total int32, ) *CallPickupItems`

NewCallPickupItems instantiates a new CallPickupItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPickupItemsWithDefaults

`func NewCallPickupItemsWithDefaults() *CallPickupItems`

NewCallPickupItemsWithDefaults instantiates a new CallPickupItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CallPickupItems) GetItems() []CallPickup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CallPickupItems) GetItemsOk() (*[]CallPickup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CallPickupItems) SetItems(v []CallPickup)`

SetItems sets Items field to given value.

### HasItems

`func (o *CallPickupItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *CallPickupItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CallPickupItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CallPickupItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
