# WizardConfigured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurable** | Pointer to **bool** | Whether all services which the wizard depends on are started or not | [optional]
**ConfigurableStatus** | Pointer to **map[string]interface{}** | A mapping of all dependencies and there statuses | [optional]
**Configured** | Pointer to **bool** | Whether Accent has already been configured or not | [optional]

## Methods

### NewWizardConfigured

`func NewWizardConfigured() *WizardConfigured`

NewWizardConfigured instantiates a new WizardConfigured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardConfiguredWithDefaults

`func NewWizardConfiguredWithDefaults() *WizardConfigured`

NewWizardConfiguredWithDefaults instantiates a new WizardConfigured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurable

`func (o *WizardConfigured) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *WizardConfigured) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *WizardConfigured) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *WizardConfigured) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetConfigurableStatus

`func (o *WizardConfigured) GetConfigurableStatus() map[string]interface{}`

GetConfigurableStatus returns the ConfigurableStatus field if non-nil, zero value otherwise.

### GetConfigurableStatusOk

`func (o *WizardConfigured) GetConfigurableStatusOk() (*map[string]interface{}, bool)`

GetConfigurableStatusOk returns a tuple with the ConfigurableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableStatus

`func (o *WizardConfigured) SetConfigurableStatus(v map[string]interface{})`

SetConfigurableStatus sets ConfigurableStatus field to given value.

### HasConfigurableStatus

`func (o *WizardConfigured) HasConfigurableStatus() bool`

HasConfigurableStatus returns a boolean if a field has been set.

### GetConfigured

`func (o *WizardConfigured) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *WizardConfigured) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *WizardConfigured) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *WizardConfigured) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
