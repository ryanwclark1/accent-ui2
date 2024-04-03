# GroupItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Group**](Group.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewGroupItems

`func NewGroupItems(total int32, ) *GroupItems`

NewGroupItems instantiates a new GroupItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupItemsWithDefaults

`func NewGroupItemsWithDefaults() *GroupItems`

NewGroupItemsWithDefaults instantiates a new GroupItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GroupItems) GetItems() []Group`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupItems) GetItemsOk() (*[]Group, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupItems) SetItems(v []Group)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *GroupItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GroupItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GroupItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
