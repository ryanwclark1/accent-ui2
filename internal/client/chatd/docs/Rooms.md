# Rooms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** |  | [optional]
**Items** | Pointer to [**[]Room**](Room.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewRooms

`func NewRooms() *Rooms`

NewRooms instantiates a new Rooms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomsWithDefaults

`func NewRoomsWithDefaults() *Rooms`

NewRoomsWithDefaults instantiates a new Rooms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *Rooms) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *Rooms) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *Rooms) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *Rooms) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *Rooms) GetItems() []Room`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Rooms) GetItemsOk() (*[]Room, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Rooms) SetItems(v []Room)`

SetItems sets Items field to given value.

### HasItems

`func (o *Rooms) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *Rooms) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Rooms) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Rooms) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Rooms) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
