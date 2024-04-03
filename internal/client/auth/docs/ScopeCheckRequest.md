# ScopeCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **[]string** | Scopes to check against |
**TenantUuid** | Pointer to **string** | If provided, also checks the token against this tenant | [optional]

## Methods

### NewScopeCheckRequest

`func NewScopeCheckRequest(scopes []string, ) *ScopeCheckRequest`

NewScopeCheckRequest instantiates a new ScopeCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeCheckRequestWithDefaults

`func NewScopeCheckRequestWithDefaults() *ScopeCheckRequest`

NewScopeCheckRequestWithDefaults instantiates a new ScopeCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ScopeCheckRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ScopeCheckRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ScopeCheckRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### GetTenantUuid

`func (o *ScopeCheckRequest) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *ScopeCheckRequest) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *ScopeCheckRequest) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *ScopeCheckRequest) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
