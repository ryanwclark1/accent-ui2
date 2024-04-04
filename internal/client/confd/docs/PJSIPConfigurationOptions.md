# PJSIPConfigurationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aor** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Auth** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Contact** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**DomainAlias** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Endpoint** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Global** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Registration** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**System** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]
**Transport** | Pointer to [**[]PJSIPConfigurationOption**](PJSIPConfigurationOption.md) | A list of all configuration options for this section | [optional]

## Methods

### NewPJSIPConfigurationOptions

`func NewPJSIPConfigurationOptions() *PJSIPConfigurationOptions`

NewPJSIPConfigurationOptions instantiates a new PJSIPConfigurationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPJSIPConfigurationOptionsWithDefaults

`func NewPJSIPConfigurationOptionsWithDefaults() *PJSIPConfigurationOptions`

NewPJSIPConfigurationOptionsWithDefaults instantiates a new PJSIPConfigurationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAor

`func (o *PJSIPConfigurationOptions) GetAor() []PJSIPConfigurationOption`

GetAor returns the Aor field if non-nil, zero value otherwise.

### GetAorOk

`func (o *PJSIPConfigurationOptions) GetAorOk() (*[]PJSIPConfigurationOption, bool)`

GetAorOk returns a tuple with the Aor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAor

`func (o *PJSIPConfigurationOptions) SetAor(v []PJSIPConfigurationOption)`

SetAor sets Aor field to given value.

### HasAor

`func (o *PJSIPConfigurationOptions) HasAor() bool`

HasAor returns a boolean if a field has been set.

### GetAuth

`func (o *PJSIPConfigurationOptions) GetAuth() []PJSIPConfigurationOption`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *PJSIPConfigurationOptions) GetAuthOk() (*[]PJSIPConfigurationOption, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *PJSIPConfigurationOptions) SetAuth(v []PJSIPConfigurationOption)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *PJSIPConfigurationOptions) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetContact

`func (o *PJSIPConfigurationOptions) GetContact() []PJSIPConfigurationOption`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PJSIPConfigurationOptions) GetContactOk() (*[]PJSIPConfigurationOption, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PJSIPConfigurationOptions) SetContact(v []PJSIPConfigurationOption)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PJSIPConfigurationOptions) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDomainAlias

`func (o *PJSIPConfigurationOptions) GetDomainAlias() []PJSIPConfigurationOption`

GetDomainAlias returns the DomainAlias field if non-nil, zero value otherwise.

### GetDomainAliasOk

`func (o *PJSIPConfigurationOptions) GetDomainAliasOk() (*[]PJSIPConfigurationOption, bool)`

GetDomainAliasOk returns a tuple with the DomainAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAlias

`func (o *PJSIPConfigurationOptions) SetDomainAlias(v []PJSIPConfigurationOption)`

SetDomainAlias sets DomainAlias field to given value.

### HasDomainAlias

`func (o *PJSIPConfigurationOptions) HasDomainAlias() bool`

HasDomainAlias returns a boolean if a field has been set.

### GetEndpoint

`func (o *PJSIPConfigurationOptions) GetEndpoint() []PJSIPConfigurationOption`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PJSIPConfigurationOptions) GetEndpointOk() (*[]PJSIPConfigurationOption, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PJSIPConfigurationOptions) SetEndpoint(v []PJSIPConfigurationOption)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PJSIPConfigurationOptions) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetGlobal

`func (o *PJSIPConfigurationOptions) GetGlobal() []PJSIPConfigurationOption`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *PJSIPConfigurationOptions) GetGlobalOk() (*[]PJSIPConfigurationOption, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *PJSIPConfigurationOptions) SetGlobal(v []PJSIPConfigurationOption)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *PJSIPConfigurationOptions) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetRegistration

`func (o *PJSIPConfigurationOptions) GetRegistration() []PJSIPConfigurationOption`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *PJSIPConfigurationOptions) GetRegistrationOk() (*[]PJSIPConfigurationOption, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *PJSIPConfigurationOptions) SetRegistration(v []PJSIPConfigurationOption)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *PJSIPConfigurationOptions) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSystem

`func (o *PJSIPConfigurationOptions) GetSystem() []PJSIPConfigurationOption`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *PJSIPConfigurationOptions) GetSystemOk() (*[]PJSIPConfigurationOption, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *PJSIPConfigurationOptions) SetSystem(v []PJSIPConfigurationOption)`

SetSystem sets System field to given value.

### HasSystem

`func (o *PJSIPConfigurationOptions) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTransport

`func (o *PJSIPConfigurationOptions) GetTransport() []PJSIPConfigurationOption`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *PJSIPConfigurationOptions) GetTransportOk() (*[]PJSIPConfigurationOption, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *PJSIPConfigurationOptions) SetTransport(v []PJSIPConfigurationOption)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *PJSIPConfigurationOptions) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
