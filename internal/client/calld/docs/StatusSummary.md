# StatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ari** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**BusConsumer** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**Plugins** | Pointer to [**PluginsStatus**](PluginsStatus.md) |  | [optional]
**ServiceToken** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]

## Methods

### NewStatusSummary

`func NewStatusSummary() *StatusSummary`

NewStatusSummary instantiates a new StatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSummaryWithDefaults

`func NewStatusSummaryWithDefaults() *StatusSummary`

NewStatusSummaryWithDefaults instantiates a new StatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAri

`func (o *StatusSummary) GetAri() ComponentWithStatus`

GetAri returns the Ari field if non-nil, zero value otherwise.

### GetAriOk

`func (o *StatusSummary) GetAriOk() (*ComponentWithStatus, bool)`

GetAriOk returns a tuple with the Ari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAri

`func (o *StatusSummary) SetAri(v ComponentWithStatus)`

SetAri sets Ari field to given value.

### HasAri

`func (o *StatusSummary) HasAri() bool`

HasAri returns a boolean if a field has been set.

### GetBusConsumer

`func (o *StatusSummary) GetBusConsumer() ComponentWithStatus`

GetBusConsumer returns the BusConsumer field if non-nil, zero value otherwise.

### GetBusConsumerOk

`func (o *StatusSummary) GetBusConsumerOk() (*ComponentWithStatus, bool)`

GetBusConsumerOk returns a tuple with the BusConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusConsumer

`func (o *StatusSummary) SetBusConsumer(v ComponentWithStatus)`

SetBusConsumer sets BusConsumer field to given value.

### HasBusConsumer

`func (o *StatusSummary) HasBusConsumer() bool`

HasBusConsumer returns a boolean if a field has been set.

### GetPlugins

`func (o *StatusSummary) GetPlugins() PluginsStatus`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *StatusSummary) GetPluginsOk() (*PluginsStatus, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *StatusSummary) SetPlugins(v PluginsStatus)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *StatusSummary) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetServiceToken

`func (o *StatusSummary) GetServiceToken() ComponentWithStatus`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *StatusSummary) GetServiceTokenOk() (*ComponentWithStatus, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *StatusSummary) SetServiceToken(v ComponentWithStatus)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *StatusSummary) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
