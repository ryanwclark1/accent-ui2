# LDAPBackendConfigEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDn** | Pointer to **string** | The DN to use to bind the &#x60;accent-auth&#x60; service to the LDAP server. If unspecified, &#x60;accent-auth&#x60; will not bind with a service user but only with the final user account. For this to work though, your users will need to have the right to read their own information, particularly their email address.  | [optional]
**Host** | **string** | The host or IP address of the LDAP server.  |
**Port** | **int32** | The port on which to connect to the LDAP server. |
**ProtocolSecurity** | Pointer to **string** | The layer of security to use for the connection. | [optional] [default to ""]
**ProtocolVersion** | Pointer to **int32** | LDAP protocol version to use | [optional]
**SearchFilters** | Pointer to **string** | Filters for finding a user DN given a service bind is used. Available variables are &#x60;username&#x60;, &#x60;user_login_attribute&#x60; and &#x60;user_email_attribute&#x60;. These variables come from the fields of the same name from the API.  | [optional]
**TenantUuid** | Pointer to **string** |  | [optional] [readonly]
**UserBaseDn** | **string** | The base DN in which users are located |
**UserEmailAttribute** | **string** | The attribute of the email address in the LDAP schema. |
**UserLoginAttribute** | **string** | The attribute that identifies users. It will be prepended to the &#x60;user_base_dn&#x60;.  |
**BindPassword** | Pointer to **string** | The password to use to bind the &#x60;accent-auth&#x60; service to the LDAP server.  | [optional]

## Methods

### NewLDAPBackendConfigEdit

`func NewLDAPBackendConfigEdit(host string, port int32, userBaseDn string, userEmailAttribute string, userLoginAttribute string, ) *LDAPBackendConfigEdit`

NewLDAPBackendConfigEdit instantiates a new LDAPBackendConfigEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPBackendConfigEditWithDefaults

`func NewLDAPBackendConfigEditWithDefaults() *LDAPBackendConfigEdit`

NewLDAPBackendConfigEditWithDefaults instantiates a new LDAPBackendConfigEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDn

`func (o *LDAPBackendConfigEdit) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *LDAPBackendConfigEdit) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *LDAPBackendConfigEdit) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *LDAPBackendConfigEdit) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetHost

`func (o *LDAPBackendConfigEdit) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LDAPBackendConfigEdit) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LDAPBackendConfigEdit) SetHost(v string)`

SetHost sets Host field to given value.

### GetPort

`func (o *LDAPBackendConfigEdit) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPBackendConfigEdit) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPBackendConfigEdit) SetPort(v int32)`

SetPort sets Port field to given value.

### GetProtocolSecurity

`func (o *LDAPBackendConfigEdit) GetProtocolSecurity() string`

GetProtocolSecurity returns the ProtocolSecurity field if non-nil, zero value otherwise.

### GetProtocolSecurityOk

`func (o *LDAPBackendConfigEdit) GetProtocolSecurityOk() (*string, bool)`

GetProtocolSecurityOk returns a tuple with the ProtocolSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolSecurity

`func (o *LDAPBackendConfigEdit) SetProtocolSecurity(v string)`

SetProtocolSecurity sets ProtocolSecurity field to given value.

### HasProtocolSecurity

`func (o *LDAPBackendConfigEdit) HasProtocolSecurity() bool`

HasProtocolSecurity returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *LDAPBackendConfigEdit) GetProtocolVersion() int32`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *LDAPBackendConfigEdit) GetProtocolVersionOk() (*int32, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *LDAPBackendConfigEdit) SetProtocolVersion(v int32)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *LDAPBackendConfigEdit) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetSearchFilters

`func (o *LDAPBackendConfigEdit) GetSearchFilters() string`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *LDAPBackendConfigEdit) GetSearchFiltersOk() (*string, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *LDAPBackendConfigEdit) SetSearchFilters(v string)`

SetSearchFilters sets SearchFilters field to given value.

### HasSearchFilters

`func (o *LDAPBackendConfigEdit) HasSearchFilters() bool`

HasSearchFilters returns a boolean if a field has been set.

### GetTenantUuid

`func (o *LDAPBackendConfigEdit) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *LDAPBackendConfigEdit) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *LDAPBackendConfigEdit) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *LDAPBackendConfigEdit) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUserBaseDn

`func (o *LDAPBackendConfigEdit) GetUserBaseDn() string`

GetUserBaseDn returns the UserBaseDn field if non-nil, zero value otherwise.

### GetUserBaseDnOk

`func (o *LDAPBackendConfigEdit) GetUserBaseDnOk() (*string, bool)`

GetUserBaseDnOk returns a tuple with the UserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDn

`func (o *LDAPBackendConfigEdit) SetUserBaseDn(v string)`

SetUserBaseDn sets UserBaseDn field to given value.

### GetUserEmailAttribute

`func (o *LDAPBackendConfigEdit) GetUserEmailAttribute() string`

GetUserEmailAttribute returns the UserEmailAttribute field if non-nil, zero value otherwise.

### GetUserEmailAttributeOk

`func (o *LDAPBackendConfigEdit) GetUserEmailAttributeOk() (*string, bool)`

GetUserEmailAttributeOk returns a tuple with the UserEmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAttribute

`func (o *LDAPBackendConfigEdit) SetUserEmailAttribute(v string)`

SetUserEmailAttribute sets UserEmailAttribute field to given value.

### GetUserLoginAttribute

`func (o *LDAPBackendConfigEdit) GetUserLoginAttribute() string`

GetUserLoginAttribute returns the UserLoginAttribute field if non-nil, zero value otherwise.

### GetUserLoginAttributeOk

`func (o *LDAPBackendConfigEdit) GetUserLoginAttributeOk() (*string, bool)`

GetUserLoginAttributeOk returns a tuple with the UserLoginAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoginAttribute

`func (o *LDAPBackendConfigEdit) SetUserLoginAttribute(v string)`

SetUserLoginAttribute sets UserLoginAttribute field to given value.

### GetBindPassword

`func (o *LDAPBackendConfigEdit) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *LDAPBackendConfigEdit) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *LDAPBackendConfigEdit) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *LDAPBackendConfigEdit) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
