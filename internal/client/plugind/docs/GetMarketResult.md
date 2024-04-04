# GetMarketResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | Pointer to **int32** | The number of plugins matching the given search | [optional]
**Items** | Pointer to [**[]MarketPluginList**](MarketPluginList.md) | A list of plugins | [optional]
**Total** | Pointer to **int32** | The number of plugins available on the market | [optional]

## Methods

### NewGetMarketResult

`func NewGetMarketResult() *GetMarketResult`

NewGetMarketResult instantiates a new GetMarketResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketResultWithDefaults

`func NewGetMarketResultWithDefaults() *GetMarketResult`

NewGetMarketResultWithDefaults instantiates a new GetMarketResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *GetMarketResult) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *GetMarketResult) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *GetMarketResult) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *GetMarketResult) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetItems

`func (o *GetMarketResult) GetItems() []MarketPluginList`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetMarketResult) GetItemsOk() (*[]MarketPluginList, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetMarketResult) SetItems(v []MarketPluginList)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetMarketResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *GetMarketResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMarketResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMarketResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMarketResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
