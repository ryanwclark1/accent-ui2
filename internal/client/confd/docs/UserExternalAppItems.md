# UserExternalAppItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserExternalApp**](UserExternalApp.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewUserExternalAppItems

`func NewUserExternalAppItems(total int32, ) *UserExternalAppItems`

NewUserExternalAppItems instantiates a new UserExternalAppItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExternalAppItemsWithDefaults

`func NewUserExternalAppItemsWithDefaults() *UserExternalAppItems`

NewUserExternalAppItemsWithDefaults instantiates a new UserExternalAppItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserExternalAppItems) GetItems() []UserExternalApp`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserExternalAppItems) GetItemsOk() (*[]UserExternalApp, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserExternalAppItems) SetItems(v []UserExternalApp)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserExternalAppItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *UserExternalAppItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserExternalAppItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserExternalAppItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
