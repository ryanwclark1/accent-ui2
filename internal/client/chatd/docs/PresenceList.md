# PresenceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** | The number of result matching the searched terms | [optional]
**Items** | Pointer to [**[]Presence**](Presence.md) |  | [optional]
**Total** | Pointer to **int32** | The number of results without filter | [optional]

## Methods

### NewPresenceList

`func NewPresenceList() *PresenceList`

NewPresenceList instantiates a new PresenceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenceListWithDefaults

`func NewPresenceListWithDefaults() *PresenceList`

NewPresenceListWithDefaults instantiates a new PresenceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *PresenceList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *PresenceList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *PresenceList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *PresenceList) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *PresenceList) GetItems() []Presence`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PresenceList) GetItemsOk() (*[]Presence, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PresenceList) SetItems(v []Presence)`

SetItems sets Items field to given value.

### HasItems

`func (o *PresenceList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *PresenceList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PresenceList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PresenceList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PresenceList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
