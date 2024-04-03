# SIPTransportItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SIPTransport**](SIPTransport.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewSIPTransportItems

`func NewSIPTransportItems() *SIPTransportItems`

NewSIPTransportItems instantiates a new SIPTransportItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIPTransportItemsWithDefaults

`func NewSIPTransportItemsWithDefaults() *SIPTransportItems`

NewSIPTransportItemsWithDefaults instantiates a new SIPTransportItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SIPTransportItems) GetItems() []SIPTransport`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SIPTransportItems) GetItemsOk() (*[]SIPTransport, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SIPTransportItems) SetItems(v []SIPTransport)`

SetItems sets Items field to given value.

### HasItems

`func (o *SIPTransportItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *SIPTransportItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SIPTransportItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SIPTransportItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SIPTransportItems) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
