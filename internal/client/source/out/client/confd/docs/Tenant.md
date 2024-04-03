# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalSipTemplateUuid** | Pointer to **string** |  | [optional] [readonly]
**RegistrationTrunkSipTemplateUuid** | Pointer to **string** |  | [optional] [readonly]
**SipTemplatesGenerated** | Pointer to **bool** | Wether or not the SIP templates have been generated | [optional] [readonly]
**Uuid** | Pointer to **string** | The UUID of the Tenant | [optional] [readonly]
**WebrtcSipTemplateUuid** | Pointer to **string** |  | [optional] [readonly]

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalSipTemplateUuid

`func (o *Tenant) GetGlobalSipTemplateUuid() string`

GetGlobalSipTemplateUuid returns the GlobalSipTemplateUuid field if non-nil, zero value otherwise.

### GetGlobalSipTemplateUuidOk

`func (o *Tenant) GetGlobalSipTemplateUuidOk() (*string, bool)`

GetGlobalSipTemplateUuidOk returns a tuple with the GlobalSipTemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSipTemplateUuid

`func (o *Tenant) SetGlobalSipTemplateUuid(v string)`

SetGlobalSipTemplateUuid sets GlobalSipTemplateUuid field to given value.

### HasGlobalSipTemplateUuid

`func (o *Tenant) HasGlobalSipTemplateUuid() bool`

HasGlobalSipTemplateUuid returns a boolean if a field has been set.

### GetRegistrationTrunkSipTemplateUuid

`func (o *Tenant) GetRegistrationTrunkSipTemplateUuid() string`

GetRegistrationTrunkSipTemplateUuid returns the RegistrationTrunkSipTemplateUuid field if non-nil, zero value otherwise.

### GetRegistrationTrunkSipTemplateUuidOk

`func (o *Tenant) GetRegistrationTrunkSipTemplateUuidOk() (*string, bool)`

GetRegistrationTrunkSipTemplateUuidOk returns a tuple with the RegistrationTrunkSipTemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTrunkSipTemplateUuid

`func (o *Tenant) SetRegistrationTrunkSipTemplateUuid(v string)`

SetRegistrationTrunkSipTemplateUuid sets RegistrationTrunkSipTemplateUuid field to given value.

### HasRegistrationTrunkSipTemplateUuid

`func (o *Tenant) HasRegistrationTrunkSipTemplateUuid() bool`

HasRegistrationTrunkSipTemplateUuid returns a boolean if a field has been set.

### GetSipTemplatesGenerated

`func (o *Tenant) GetSipTemplatesGenerated() bool`

GetSipTemplatesGenerated returns the SipTemplatesGenerated field if non-nil, zero value otherwise.

### GetSipTemplatesGeneratedOk

`func (o *Tenant) GetSipTemplatesGeneratedOk() (*bool, bool)`

GetSipTemplatesGeneratedOk returns a tuple with the SipTemplatesGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTemplatesGenerated

`func (o *Tenant) SetSipTemplatesGenerated(v bool)`

SetSipTemplatesGenerated sets SipTemplatesGenerated field to given value.

### HasSipTemplatesGenerated

`func (o *Tenant) HasSipTemplatesGenerated() bool`

HasSipTemplatesGenerated returns a boolean if a field has been set.

### GetUuid

`func (o *Tenant) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Tenant) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Tenant) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Tenant) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWebrtcSipTemplateUuid

`func (o *Tenant) GetWebrtcSipTemplateUuid() string`

GetWebrtcSipTemplateUuid returns the WebrtcSipTemplateUuid field if non-nil, zero value otherwise.

### GetWebrtcSipTemplateUuidOk

`func (o *Tenant) GetWebrtcSipTemplateUuidOk() (*string, bool)`

GetWebrtcSipTemplateUuidOk returns a tuple with the WebrtcSipTemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebrtcSipTemplateUuid

`func (o *Tenant) SetWebrtcSipTemplateUuid(v string)`

SetWebrtcSipTemplateUuid sets WebrtcSipTemplateUuid field to given value.

### HasWebrtcSipTemplateUuid

`func (o *Tenant) HasWebrtcSipTemplateUuid() bool`

HasWebrtcSipTemplateUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
