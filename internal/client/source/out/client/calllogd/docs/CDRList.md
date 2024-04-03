# CDRList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** |  | [optional]
**Items** | Pointer to [**[]CDR**](CDR.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewCDRList

`func NewCDRList() *CDRList`

NewCDRList instantiates a new CDRList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDRListWithDefaults

`func NewCDRListWithDefaults() *CDRList`

NewCDRListWithDefaults instantiates a new CDRList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *CDRList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *CDRList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *CDRList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *CDRList) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *CDRList) GetItems() []CDR`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CDRList) GetItemsOk() (*[]CDR, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CDRList) SetItems(v []CDR)`

SetItems sets Items field to given value.

### HasItems

`func (o *CDRList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *CDRList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CDRList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CDRList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CDRList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
