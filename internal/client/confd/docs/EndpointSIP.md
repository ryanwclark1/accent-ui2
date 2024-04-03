# EndpointSIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AorSectionOptions** | Pointer to **[][]string** | A list of PJSIP AOR section options for this endpoint | [optional]
**AsteriskId** | Pointer to **string** | The ID of the Asterisk onto which this configuration applies. | [optional]
**AuthSectionOptions** | Pointer to **[][]string** | A list of PJSIP auth section options for this endpoint | [optional]
**EndpointSectionOptions** | Pointer to **[][]string** | A list of PJSIP endpoint section options for this endpoint | [optional]
**IdentifySectionOptions** | Pointer to **[][]string** | A list of PJSIP identify section options for this endpoint | [optional]
**Label** | Pointer to **string** | The human readable name for this configuration | [optional]
**Name** | Pointer to **string** | The name of the PJSIP entity, auto-generated if missing | [optional]
**OutboundAuthSectionOptions** | Pointer to **[][]string** | A list of PJSIP auth section options for this endpoint | [optional]
**RegistrationOutboundAuthSectionOptions** | Pointer to **[][]string** | A list of PJSIP auth section options for this endpoint | [optional]
**RegistrationSectionOptions** | Pointer to **[][]string** | A list of PJSIP registration section options for this endpoint | [optional]
**Templates** | Pointer to [**[]EndpointSIPTemplatesRelation**](EndpointSIPTemplatesRelation.md) | A list of templates this configuration will inherit from.  The inheritance only applies to option sections. Not to the name, label and other fields. For a given list of templates [A, B, C] A will be applied over B. C will be applied over (A,B) etc.  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Transport** | Pointer to [**SIPTransportRelation**](SIPTransportRelation.md) |  | [optional]
**Uuid** | Pointer to **string** | The UUID of this resource | [optional] [readonly]

## Methods

### NewEndpointSIP

`func NewEndpointSIP() *EndpointSIP`

NewEndpointSIP instantiates a new EndpointSIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointSIPWithDefaults

`func NewEndpointSIPWithDefaults() *EndpointSIP`

NewEndpointSIPWithDefaults instantiates a new EndpointSIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAorSectionOptions

`func (o *EndpointSIP) GetAorSectionOptions() [][]string`

GetAorSectionOptions returns the AorSectionOptions field if non-nil, zero value otherwise.

### GetAorSectionOptionsOk

`func (o *EndpointSIP) GetAorSectionOptionsOk() (*[][]string, bool)`

GetAorSectionOptionsOk returns a tuple with the AorSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAorSectionOptions

`func (o *EndpointSIP) SetAorSectionOptions(v [][]string)`

SetAorSectionOptions sets AorSectionOptions field to given value.

### HasAorSectionOptions

`func (o *EndpointSIP) HasAorSectionOptions() bool`

HasAorSectionOptions returns a boolean if a field has been set.

### GetAsteriskId

`func (o *EndpointSIP) GetAsteriskId() string`

GetAsteriskId returns the AsteriskId field if non-nil, zero value otherwise.

### GetAsteriskIdOk

`func (o *EndpointSIP) GetAsteriskIdOk() (*string, bool)`

GetAsteriskIdOk returns a tuple with the AsteriskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsteriskId

`func (o *EndpointSIP) SetAsteriskId(v string)`

SetAsteriskId sets AsteriskId field to given value.

### HasAsteriskId

`func (o *EndpointSIP) HasAsteriskId() bool`

HasAsteriskId returns a boolean if a field has been set.

### GetAuthSectionOptions

`func (o *EndpointSIP) GetAuthSectionOptions() [][]string`

GetAuthSectionOptions returns the AuthSectionOptions field if non-nil, zero value otherwise.

### GetAuthSectionOptionsOk

`func (o *EndpointSIP) GetAuthSectionOptionsOk() (*[][]string, bool)`

GetAuthSectionOptionsOk returns a tuple with the AuthSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSectionOptions

`func (o *EndpointSIP) SetAuthSectionOptions(v [][]string)`

SetAuthSectionOptions sets AuthSectionOptions field to given value.

### HasAuthSectionOptions

`func (o *EndpointSIP) HasAuthSectionOptions() bool`

HasAuthSectionOptions returns a boolean if a field has been set.

### GetEndpointSectionOptions

`func (o *EndpointSIP) GetEndpointSectionOptions() [][]string`

GetEndpointSectionOptions returns the EndpointSectionOptions field if non-nil, zero value otherwise.

### GetEndpointSectionOptionsOk

`func (o *EndpointSIP) GetEndpointSectionOptionsOk() (*[][]string, bool)`

GetEndpointSectionOptionsOk returns a tuple with the EndpointSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSectionOptions

`func (o *EndpointSIP) SetEndpointSectionOptions(v [][]string)`

SetEndpointSectionOptions sets EndpointSectionOptions field to given value.

### HasEndpointSectionOptions

`func (o *EndpointSIP) HasEndpointSectionOptions() bool`

HasEndpointSectionOptions returns a boolean if a field has been set.

### GetIdentifySectionOptions

`func (o *EndpointSIP) GetIdentifySectionOptions() [][]string`

GetIdentifySectionOptions returns the IdentifySectionOptions field if non-nil, zero value otherwise.

### GetIdentifySectionOptionsOk

`func (o *EndpointSIP) GetIdentifySectionOptionsOk() (*[][]string, bool)`

GetIdentifySectionOptionsOk returns a tuple with the IdentifySectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifySectionOptions

`func (o *EndpointSIP) SetIdentifySectionOptions(v [][]string)`

SetIdentifySectionOptions sets IdentifySectionOptions field to given value.

### HasIdentifySectionOptions

`func (o *EndpointSIP) HasIdentifySectionOptions() bool`

HasIdentifySectionOptions returns a boolean if a field has been set.

### GetLabel

`func (o *EndpointSIP) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointSIP) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointSIP) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointSIP) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *EndpointSIP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointSIP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointSIP) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointSIP) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundAuthSectionOptions

`func (o *EndpointSIP) GetOutboundAuthSectionOptions() [][]string`

GetOutboundAuthSectionOptions returns the OutboundAuthSectionOptions field if non-nil, zero value otherwise.

### GetOutboundAuthSectionOptionsOk

`func (o *EndpointSIP) GetOutboundAuthSectionOptionsOk() (*[][]string, bool)`

GetOutboundAuthSectionOptionsOk returns a tuple with the OutboundAuthSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAuthSectionOptions

`func (o *EndpointSIP) SetOutboundAuthSectionOptions(v [][]string)`

SetOutboundAuthSectionOptions sets OutboundAuthSectionOptions field to given value.

### HasOutboundAuthSectionOptions

`func (o *EndpointSIP) HasOutboundAuthSectionOptions() bool`

HasOutboundAuthSectionOptions returns a boolean if a field has been set.

### GetRegistrationOutboundAuthSectionOptions

`func (o *EndpointSIP) GetRegistrationOutboundAuthSectionOptions() [][]string`

GetRegistrationOutboundAuthSectionOptions returns the RegistrationOutboundAuthSectionOptions field if non-nil, zero value otherwise.

### GetRegistrationOutboundAuthSectionOptionsOk

`func (o *EndpointSIP) GetRegistrationOutboundAuthSectionOptionsOk() (*[][]string, bool)`

GetRegistrationOutboundAuthSectionOptionsOk returns a tuple with the RegistrationOutboundAuthSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOutboundAuthSectionOptions

`func (o *EndpointSIP) SetRegistrationOutboundAuthSectionOptions(v [][]string)`

SetRegistrationOutboundAuthSectionOptions sets RegistrationOutboundAuthSectionOptions field to given value.

### HasRegistrationOutboundAuthSectionOptions

`func (o *EndpointSIP) HasRegistrationOutboundAuthSectionOptions() bool`

HasRegistrationOutboundAuthSectionOptions returns a boolean if a field has been set.

### GetRegistrationSectionOptions

`func (o *EndpointSIP) GetRegistrationSectionOptions() [][]string`

GetRegistrationSectionOptions returns the RegistrationSectionOptions field if non-nil, zero value otherwise.

### GetRegistrationSectionOptionsOk

`func (o *EndpointSIP) GetRegistrationSectionOptionsOk() (*[][]string, bool)`

GetRegistrationSectionOptionsOk returns a tuple with the RegistrationSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSectionOptions

`func (o *EndpointSIP) SetRegistrationSectionOptions(v [][]string)`

SetRegistrationSectionOptions sets RegistrationSectionOptions field to given value.

### HasRegistrationSectionOptions

`func (o *EndpointSIP) HasRegistrationSectionOptions() bool`

HasRegistrationSectionOptions returns a boolean if a field has been set.

### GetTemplates

`func (o *EndpointSIP) GetTemplates() []EndpointSIPTemplatesRelation`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *EndpointSIP) GetTemplatesOk() (*[]EndpointSIPTemplatesRelation, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *EndpointSIP) SetTemplates(v []EndpointSIPTemplatesRelation)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *EndpointSIP) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTenantUuid

`func (o *EndpointSIP) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *EndpointSIP) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *EndpointSIP) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *EndpointSIP) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTransport

`func (o *EndpointSIP) GetTransport() SIPTransportRelation`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *EndpointSIP) GetTransportOk() (*SIPTransportRelation, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *EndpointSIP) SetTransport(v SIPTransportRelation)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *EndpointSIP) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUuid

`func (o *EndpointSIP) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EndpointSIP) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EndpointSIP) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EndpointSIP) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
