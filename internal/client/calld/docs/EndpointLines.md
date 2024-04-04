# EndpointLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** | The number of line endpoint matching the searched terms | [optional]
**Items** | Pointer to [**[]EndpointLine**](EndpointLine.md) |  | [optional]
**Total** | Pointer to **int32** | The number of line endpoint | [optional]

## Methods

### NewEndpointLines

`func NewEndpointLines() *EndpointLines`

NewEndpointLines instantiates a new EndpointLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointLinesWithDefaults

`func NewEndpointLinesWithDefaults() *EndpointLines`

NewEndpointLinesWithDefaults instantiates a new EndpointLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *EndpointLines) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *EndpointLines) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *EndpointLines) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *EndpointLines) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *EndpointLines) GetItems() []EndpointLine`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EndpointLines) GetItemsOk() (*[]EndpointLine, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EndpointLines) SetItems(v []EndpointLine)`

SetItems sets Items field to given value.

### HasItems

`func (o *EndpointLines) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *EndpointLines) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EndpointLines) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EndpointLines) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EndpointLines) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
