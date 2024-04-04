# PluginMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the plugin | [optional]
**Namespace** | Pointer to **string** | The namespace of the plugin | [optional]
**Version** | Pointer to **string** | The version of the installed version | [optional]

## Methods

### NewPluginMetadata

`func NewPluginMetadata() *PluginMetadata`

NewPluginMetadata instantiates a new PluginMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginMetadataWithDefaults

`func NewPluginMetadataWithDefaults() *PluginMetadata`

NewPluginMetadataWithDefaults instantiates a new PluginMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PluginMetadata) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PluginMetadata) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PluginMetadata) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PluginMetadata) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVersion

`func (o *PluginMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
