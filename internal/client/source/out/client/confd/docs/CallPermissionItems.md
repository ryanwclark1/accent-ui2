# CallPermissionItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CallPermission**](CallPermission.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewCallPermissionItems

`func NewCallPermissionItems(total int32, ) *CallPermissionItems`

NewCallPermissionItems instantiates a new CallPermissionItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPermissionItemsWithDefaults

`func NewCallPermissionItemsWithDefaults() *CallPermissionItems`

NewCallPermissionItemsWithDefaults instantiates a new CallPermissionItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CallPermissionItems) GetItems() []CallPermission`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CallPermissionItems) GetItemsOk() (*[]CallPermission, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CallPermissionItems) SetItems(v []CallPermission)`

SetItems sets Items field to given value.

### HasItems

`func (o *CallPermissionItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *CallPermissionItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CallPermissionItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CallPermissionItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
