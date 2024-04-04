# ExternalApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | Arbitrary settings needed by an external app | [optional]
**Label** | Pointer to **string** | The label of the external app | [optional]
**Name** | Pointer to **string** | The name to identify the external app | [optional] [readonly]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewExternalApp

`func NewExternalApp() *ExternalApp`

NewExternalApp instantiates a new ExternalApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAppWithDefaults

`func NewExternalAppWithDefaults() *ExternalApp`

NewExternalAppWithDefaults instantiates a new ExternalApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ExternalApp) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ExternalApp) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ExternalApp) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ExternalApp) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLabel

`func (o *ExternalApp) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalApp) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalApp) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExternalApp) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ExternalApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *ExternalApp) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *ExternalApp) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *ExternalApp) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *ExternalApp) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
