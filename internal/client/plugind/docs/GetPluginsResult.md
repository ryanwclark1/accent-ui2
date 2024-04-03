# GetPluginsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PluginMetadata**](PluginMetadata.md) | A list of plugins | [optional]
**Total** | Pointer to **int32** | The number of plugins installed on the system | [optional]

## Methods

### NewGetPluginsResult

`func NewGetPluginsResult() *GetPluginsResult`

NewGetPluginsResult instantiates a new GetPluginsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPluginsResultWithDefaults

`func NewGetPluginsResultWithDefaults() *GetPluginsResult`

NewGetPluginsResultWithDefaults instantiates a new GetPluginsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetPluginsResult) GetItems() []PluginMetadata`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetPluginsResult) GetItemsOk() (*[]PluginMetadata, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetPluginsResult) SetItems(v []PluginMetadata)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetPluginsResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *GetPluginsResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPluginsResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPluginsResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPluginsResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
