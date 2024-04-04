# PluginInstallParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The method used to fetch this plugin |
**Options** | Pointer to **map[string]interface{}** | Method dependant installation options | [optional]

## Methods

### NewPluginInstallParameters

`func NewPluginInstallParameters(method string, ) *PluginInstallParameters`

NewPluginInstallParameters instantiates a new PluginInstallParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInstallParametersWithDefaults

`func NewPluginInstallParametersWithDefaults() *PluginInstallParameters`

NewPluginInstallParametersWithDefaults instantiates a new PluginInstallParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PluginInstallParameters) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PluginInstallParameters) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PluginInstallParameters) SetMethod(v string)`

SetMethod sets Method field to given value.

### GetOptions

`func (o *PluginInstallParameters) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PluginInstallParameters) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PluginInstallParameters) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PluginInstallParameters) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
