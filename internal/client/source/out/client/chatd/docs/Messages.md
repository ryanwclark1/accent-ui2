# Messages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** |  | [optional]
**Items** | Pointer to [**[]Message**](Message.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewMessages

`func NewMessages() *Messages`

NewMessages instantiates a new Messages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesWithDefaults

`func NewMessagesWithDefaults() *Messages`

NewMessagesWithDefaults instantiates a new Messages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *Messages) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *Messages) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *Messages) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *Messages) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *Messages) GetItems() []Message`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Messages) GetItemsOk() (*[]Message, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Messages) SetItems(v []Message)`

SetItems sets Items field to given value.

### HasItems

`func (o *Messages) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *Messages) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Messages) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Messages) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Messages) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
