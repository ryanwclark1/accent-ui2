# MarketPluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the plugin | [optional]
**Namespace** | Pointer to **string** | The namespace of the plugin | [optional]
**Versions** | Pointer to [**[]VersionInfo**](VersionInfo.md) | Version specific information | [optional]

## Methods

### NewMarketPluginList

`func NewMarketPluginList() *MarketPluginList`

NewMarketPluginList instantiates a new MarketPluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPluginListWithDefaults

`func NewMarketPluginListWithDefaults() *MarketPluginList`

NewMarketPluginListWithDefaults instantiates a new MarketPluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketPluginList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketPluginList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketPluginList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketPluginList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *MarketPluginList) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *MarketPluginList) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *MarketPluginList) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *MarketPluginList) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVersions

`func (o *MarketPluginList) GetVersions() []VersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MarketPluginList) GetVersionsOk() (*[]VersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MarketPluginList) SetVersions(v []VersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MarketPluginList) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
