# ExternalAuthUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | **int32** | The number of external auth matching the searched term. |
**Items** | [**[]ExternalAuthUser**](ExternalAuthUser.md) | A paginated list of connected external auth users |
**Total** | **int32** | The number of connected external auth users. |

## Methods

### NewExternalAuthUserList

`func NewExternalAuthUserList(filtered int32, items []ExternalAuthUser, total int32, ) *ExternalAuthUserList`

NewExternalAuthUserList instantiates a new ExternalAuthUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAuthUserListWithDefaults

`func NewExternalAuthUserListWithDefaults() *ExternalAuthUserList`

NewExternalAuthUserListWithDefaults instantiates a new ExternalAuthUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *ExternalAuthUserList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *ExternalAuthUserList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *ExternalAuthUserList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### GetItems

`func (o *ExternalAuthUserList) GetItems() []ExternalAuthUser`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExternalAuthUserList) GetItemsOk() (*[]ExternalAuthUser, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExternalAuthUserList) SetItems(v []ExternalAuthUser)`

SetItems sets Items field to given value.

### GetTotal

`func (o *ExternalAuthUserList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExternalAuthUserList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExternalAuthUserList) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
