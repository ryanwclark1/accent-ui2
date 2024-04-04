# UserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** | The number of users matching the searched term | [optional]
**Items** | Pointer to [**[]UserResult**](UserResult.md) | A paginated list of users | [optional]
**Total** | Pointer to **int32** | The number of users | [optional]

## Methods

### NewUserList

`func NewUserList() *UserList`

NewUserList instantiates a new UserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListWithDefaults

`func NewUserListWithDefaults() *UserList`

NewUserListWithDefaults instantiates a new UserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *UserList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *UserList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *UserList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *UserList) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *UserList) GetItems() []UserResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserList) GetItemsOk() (*[]UserResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserList) SetItems(v []UserResult)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *UserList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UserList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
