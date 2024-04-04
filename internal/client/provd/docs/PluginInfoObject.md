# PluginInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **map[string]map[string]string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Version** | Pointer to **string** |  | [optional]

## Methods

### NewPluginInfoObject

`func NewPluginInfoObject() *PluginInfoObject`

NewPluginInfoObject instantiates a new PluginInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInfoObjectWithDefaults

`func NewPluginInfoObjectWithDefaults() *PluginInfoObject`

NewPluginInfoObjectWithDefaults instantiates a new PluginInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *PluginInfoObject) GetCapabilities() map[string]map[string]string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PluginInfoObject) GetCapabilitiesOk() (*map[string]map[string]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PluginInfoObject) SetCapabilities(v map[string]map[string]string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PluginInfoObject) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDescription

`func (o *PluginInfoObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginInfoObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginInfoObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginInfoObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PluginInfoObject) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginInfoObject) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginInfoObject) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginInfoObject) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
