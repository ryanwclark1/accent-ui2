# EmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRewritingRules** | Pointer to [**[]RewritingRule**](RewritingRule.md) | Rules for local address to external address rewriting | [optional]
**DomainName** | Pointer to **string** | The domain name of the current mailing server, i.e. &#x60;example.com&#x60; | [optional]
**FallbackSmtpHost** | Pointer to **string** | Fallback relay server hostname or address. It is possible to specify the port, for example &#x60;domain.com:587&#x60;  | [optional]
**From** | Pointer to **string** | The e-mail address or domain name to use in the &#x60;From&#x60; header for local services  | [optional]
**SmtpHost** | Pointer to **string** | E-mail relay server hostname or address. It is possible to specify the port, for example &#x60;domain.com:587&#x60;  | [optional]

## Methods

### NewEmailConfig

`func NewEmailConfig() *EmailConfig`

NewEmailConfig instantiates a new EmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigWithDefaults

`func NewEmailConfigWithDefaults() *EmailConfig`

NewEmailConfigWithDefaults instantiates a new EmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRewritingRules

`func (o *EmailConfig) GetAddressRewritingRules() []RewritingRule`

GetAddressRewritingRules returns the AddressRewritingRules field if non-nil, zero value otherwise.

### GetAddressRewritingRulesOk

`func (o *EmailConfig) GetAddressRewritingRulesOk() (*[]RewritingRule, bool)`

GetAddressRewritingRulesOk returns a tuple with the AddressRewritingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRewritingRules

`func (o *EmailConfig) SetAddressRewritingRules(v []RewritingRule)`

SetAddressRewritingRules sets AddressRewritingRules field to given value.

### HasAddressRewritingRules

`func (o *EmailConfig) HasAddressRewritingRules() bool`

HasAddressRewritingRules returns a boolean if a field has been set.

### GetDomainName

`func (o *EmailConfig) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *EmailConfig) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *EmailConfig) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *EmailConfig) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetFallbackSmtpHost

`func (o *EmailConfig) GetFallbackSmtpHost() string`

GetFallbackSmtpHost returns the FallbackSmtpHost field if non-nil, zero value otherwise.

### GetFallbackSmtpHostOk

`func (o *EmailConfig) GetFallbackSmtpHostOk() (*string, bool)`

GetFallbackSmtpHostOk returns a tuple with the FallbackSmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackSmtpHost

`func (o *EmailConfig) SetFallbackSmtpHost(v string)`

SetFallbackSmtpHost sets FallbackSmtpHost field to given value.

### HasFallbackSmtpHost

`func (o *EmailConfig) HasFallbackSmtpHost() bool`

HasFallbackSmtpHost returns a boolean if a field has been set.

### GetFrom

`func (o *EmailConfig) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailConfig) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailConfig) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailConfig) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSmtpHost

`func (o *EmailConfig) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *EmailConfig) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *EmailConfig) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *EmailConfig) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
