# UserExternalApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | Arbitrary settings needed by an external app | [optional]
**Label** | Pointer to **string** | The label of the external app | [optional]
**Name** | Pointer to **string** | The name to identify the external app | [optional] [readonly]

## Methods

### NewUserExternalApp

`func NewUserExternalApp() *UserExternalApp`

NewUserExternalApp instantiates a new UserExternalApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExternalAppWithDefaults

`func NewUserExternalAppWithDefaults() *UserExternalApp`

NewUserExternalAppWithDefaults instantiates a new UserExternalApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *UserExternalApp) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UserExternalApp) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UserExternalApp) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UserExternalApp) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLabel

`func (o *UserExternalApp) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserExternalApp) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserExternalApp) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UserExternalApp) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UserExternalApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserExternalApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserExternalApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserExternalApp) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
