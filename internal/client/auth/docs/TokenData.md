# TokenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **[]string** | The list of allowed accesses for this token | [optional]
**AuthId** | Pointer to **string** | The unique identifier retrieved from the backend | [optional]
**ExpiresAt** | Pointer to **string** |  | [optional]
**IssuedAt** | Pointer to **string** |  | [optional]
**Metadata** | Pointer to **map[string]interface{}** | Information owned by accent-auth about this user | [optional]
**SessionUuid** | Pointer to **string** |  | [optional]
**Token** | Pointer to **string** |  | [optional]
**UtcExpiresAt** | Pointer to **string** |  | [optional]
**UtcIssuedAt** | Pointer to **string** |  | [optional]
**AccentUserUuid** | Pointer to **string** | The UUID of the matching accent-confd user if there is one. This field can be null. This field should NOT be used anymore, the \&quot;pbx_user_uuid\&quot; in the metadata field is the prefered method to access this information.  | [optional]
**AccentUuid** | Pointer to **string** |  | [optional]

## Methods

### NewTokenData

`func NewTokenData() *TokenData`

NewTokenData instantiates a new TokenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataWithDefaults

`func NewTokenDataWithDefaults() *TokenData`

NewTokenDataWithDefaults instantiates a new TokenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *TokenData) GetAcl() []string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *TokenData) GetAclOk() (*[]string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *TokenData) SetAcl(v []string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *TokenData) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAuthId

`func (o *TokenData) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *TokenData) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *TokenData) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *TokenData) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenData) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIssuedAt

`func (o *TokenData) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *TokenData) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *TokenData) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *TokenData) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *TokenData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TokenData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSessionUuid

`func (o *TokenData) GetSessionUuid() string`

GetSessionUuid returns the SessionUuid field if non-nil, zero value otherwise.

### GetSessionUuidOk

`func (o *TokenData) GetSessionUuidOk() (*string, bool)`

GetSessionUuidOk returns a tuple with the SessionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUuid

`func (o *TokenData) SetSessionUuid(v string)`

SetSessionUuid sets SessionUuid field to given value.

### HasSessionUuid

`func (o *TokenData) HasSessionUuid() bool`

HasSessionUuid returns a boolean if a field has been set.

### GetToken

`func (o *TokenData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenData) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUtcExpiresAt

`func (o *TokenData) GetUtcExpiresAt() string`

GetUtcExpiresAt returns the UtcExpiresAt field if non-nil, zero value otherwise.

### GetUtcExpiresAtOk

`func (o *TokenData) GetUtcExpiresAtOk() (*string, bool)`

GetUtcExpiresAtOk returns a tuple with the UtcExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcExpiresAt

`func (o *TokenData) SetUtcExpiresAt(v string)`

SetUtcExpiresAt sets UtcExpiresAt field to given value.

### HasUtcExpiresAt

`func (o *TokenData) HasUtcExpiresAt() bool`

HasUtcExpiresAt returns a boolean if a field has been set.

### GetUtcIssuedAt

`func (o *TokenData) GetUtcIssuedAt() string`

GetUtcIssuedAt returns the UtcIssuedAt field if non-nil, zero value otherwise.

### GetUtcIssuedAtOk

`func (o *TokenData) GetUtcIssuedAtOk() (*string, bool)`

GetUtcIssuedAtOk returns a tuple with the UtcIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcIssuedAt

`func (o *TokenData) SetUtcIssuedAt(v string)`

SetUtcIssuedAt sets UtcIssuedAt field to given value.

### HasUtcIssuedAt

`func (o *TokenData) HasUtcIssuedAt() bool`

HasUtcIssuedAt returns a boolean if a field has been set.

### GetAccentUserUuid

`func (o *TokenData) GetAccentUserUuid() string`

GetAccentUserUuid returns the AccentUserUuid field if non-nil, zero value otherwise.

### GetAccentUserUuidOk

`func (o *TokenData) GetAccentUserUuidOk() (*string, bool)`

GetAccentUserUuidOk returns a tuple with the AccentUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentUserUuid

`func (o *TokenData) SetAccentUserUuid(v string)`

SetAccentUserUuid sets AccentUserUuid field to given value.

### HasAccentUserUuid

`func (o *TokenData) HasAccentUserUuid() bool`

HasAccentUserUuid returns a boolean if a field has been set.

### GetAccentUuid

`func (o *TokenData) GetAccentUuid() string`

GetAccentUuid returns the AccentUuid field if non-nil, zero value otherwise.

### GetAccentUuidOk

`func (o *TokenData) GetAccentUuidOk() (*string, bool)`

GetAccentUuidOk returns a tuple with the AccentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentUuid

`func (o *TokenData) SetAccentUuid(v string)`

SetAccentUuid sets AccentUuid field to given value.

### HasAccentUuid

`func (o *TokenData) HasAccentUuid() bool`

HasAccentUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
