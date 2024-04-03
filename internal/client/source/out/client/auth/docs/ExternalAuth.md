# ExternalAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional]
**Enabled** | Pointer to **bool** |  | [optional]
**PluginInfo** | Pointer to **map[string]interface{}** |  | [optional]
**Type** | Pointer to **string** | The external auth type name | [optional]

## Methods

### NewExternalAuth

`func NewExternalAuth() *ExternalAuth`

NewExternalAuth instantiates a new ExternalAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAuthWithDefaults

`func NewExternalAuthWithDefaults() *ExternalAuth`

NewExternalAuthWithDefaults instantiates a new ExternalAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExternalAuth) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalAuth) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalAuth) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ExternalAuth) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEnabled

`func (o *ExternalAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExternalAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExternalAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExternalAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPluginInfo

`func (o *ExternalAuth) GetPluginInfo() map[string]interface{}`

GetPluginInfo returns the PluginInfo field if non-nil, zero value otherwise.

### GetPluginInfoOk

`func (o *ExternalAuth) GetPluginInfoOk() (*map[string]interface{}, bool)`

GetPluginInfoOk returns a tuple with the PluginInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginInfo

`func (o *ExternalAuth) SetPluginInfo(v map[string]interface{})`

SetPluginInfo sets PluginInfo field to given value.

### HasPluginInfo

`func (o *ExternalAuth) HasPluginInfo() bool`

HasPluginInfo returns a boolean if a field has been set.

### GetType

`func (o *ExternalAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalAuth) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
