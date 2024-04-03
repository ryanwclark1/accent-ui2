# PluginsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**Voicemails** | Pointer to [**VoicemailsStatus**](VoicemailsStatus.md) |  | [optional]

## Methods

### NewPluginsStatus

`func NewPluginsStatus() *PluginsStatus`

NewPluginsStatus instantiates a new PluginsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsStatusWithDefaults

`func NewPluginsStatusWithDefaults() *PluginsStatus`

NewPluginsStatusWithDefaults instantiates a new PluginsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *PluginsStatus) GetEndpoints() ComponentWithStatus`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PluginsStatus) GetEndpointsOk() (*ComponentWithStatus, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PluginsStatus) SetEndpoints(v ComponentWithStatus)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PluginsStatus) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetVoicemails

`func (o *PluginsStatus) GetVoicemails() VoicemailsStatus`

GetVoicemails returns the Voicemails field if non-nil, zero value otherwise.

### GetVoicemailsOk

`func (o *PluginsStatus) GetVoicemailsOk() (*VoicemailsStatus, bool)`

GetVoicemailsOk returns a tuple with the Voicemails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemails

`func (o *PluginsStatus) SetVoicemails(v VoicemailsStatus)`

SetVoicemails sets Voicemails field to given value.

### HasVoicemails

`func (o *PluginsStatus) HasVoicemails() bool`

HasVoicemails returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
