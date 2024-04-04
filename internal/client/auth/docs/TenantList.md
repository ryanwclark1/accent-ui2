# TenantList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** | The number of tenants matching the searched term | [optional]
**Items** | Pointer to [**[]TenantResult**](TenantResult.md) | A paginated list of tenants | [optional]
**Total** | Pointer to **int32** | The number of tenants | [optional]

## Methods

### NewTenantList

`func NewTenantList() *TenantList`

NewTenantList instantiates a new TenantList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantListWithDefaults

`func NewTenantListWithDefaults() *TenantList`

NewTenantListWithDefaults instantiates a new TenantList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *TenantList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *TenantList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *TenantList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *TenantList) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *TenantList) GetItems() []TenantResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TenantList) GetItemsOk() (*[]TenantResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TenantList) SetItems(v []TenantResult)`

SetItems sets Items field to given value.

### HasItems

`func (o *TenantList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *TenantList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TenantList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TenantList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TenantList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
