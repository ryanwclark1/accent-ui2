# CreateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The &#x60;access_type&#x60; indicates whether your application can refresh the tokens when the user is not present at the browser. Valid parameter values are *online*, which is the default value, and *offline* Only one refresh token will be created for a given user with a given &#x60;client_id&#x60;. The old refresh for &#x60;client_id&#x60; will be revoken when creating a new one. The *client_id* field is required when using the &#x60;access_type&#x60; *offline*  | [optional] [default to "online"]
**Backend** | Pointer to **string** |  | [optional] [default to "accent_user"]
**ClientId** | Pointer to **string** | The &#x60;client_id&#x60; is used in conjunction with the &#x60;access_type&#x60; *offline* to known for which application a refresh token has been emitted. *Required when using &#x60;access_type: offline&#x60;*  | [optional]
**DomainName** | Pointer to **string** | The &#x60;domain_name&#x60; must match a tenant&#39;s domain_name entry to find the appropriate ldap configuration.  | [optional]
**Expiration** | Pointer to **int32** | Expiration time in seconds. | [optional]
**RefreshToken** | Pointer to **string** | The &#x60;refresh_token&#x60; can be used to get a new access token without using the username/password. This is useful for client application that should not store the username and password once the user has logged in a first time.  | [optional]

## Methods

### NewCreateTokenRequest

`func NewCreateTokenRequest() *CreateTokenRequest`

NewCreateTokenRequest instantiates a new CreateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenRequestWithDefaults

`func NewCreateTokenRequestWithDefaults() *CreateTokenRequest`

NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CreateTokenRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateTokenRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateTokenRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *CreateTokenRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetBackend

`func (o *CreateTokenRequest) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *CreateTokenRequest) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *CreateTokenRequest) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *CreateTokenRequest) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetClientId

`func (o *CreateTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDomainName

`func (o *CreateTokenRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateTokenRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateTokenRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateTokenRequest) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetExpiration

`func (o *CreateTokenRequest) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreateTokenRequest) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreateTokenRequest) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CreateTokenRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRefreshToken

`func (o *CreateTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *CreateTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *CreateTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *CreateTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
