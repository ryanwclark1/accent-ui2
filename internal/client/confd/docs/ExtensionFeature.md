# ExtensionFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly]
**Enabled** | Pointer to **bool** | If False the extension feature is disabled | [optional]
**Exten** | **string** |  |
**Feature** | Pointer to **string** | The feature of the extension | [optional] [readonly]
**Uuid** | Pointer to **string** | Extension UUID | [optional] [readonly]

## Methods

### NewExtensionFeature

`func NewExtensionFeature(exten string, ) *ExtensionFeature`

NewExtensionFeature instantiates a new ExtensionFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionFeatureWithDefaults

`func NewExtensionFeatureWithDefaults() *ExtensionFeature`

NewExtensionFeatureWithDefaults instantiates a new ExtensionFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ExtensionFeature) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ExtensionFeature) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ExtensionFeature) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ExtensionFeature) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtensionFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtensionFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtensionFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtensionFeature) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExten

`func (o *ExtensionFeature) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *ExtensionFeature) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *ExtensionFeature) SetExten(v string)`

SetExten sets Exten field to given value.

### GetFeature

`func (o *ExtensionFeature) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ExtensionFeature) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ExtensionFeature) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ExtensionFeature) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetUuid

`func (o *ExtensionFeature) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExtensionFeature) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExtensionFeature) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExtensionFeature) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
