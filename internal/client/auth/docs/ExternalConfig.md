# ExternalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID for the given authentication type. Required only for &#x60;google&#x60; and &#x60;microsoft&#x60; authentication types.  | [optional]
**ClientSecret** | Pointer to **string** | Client secret for the given authentication type. Required only for &#x60;google&#x60; and &#x60;microsoft&#x60; authentication types.  | [optional]
**FcmApiKey** | Pointer to **string** | The API key to use for Firebase Cloud Messaging | [optional]
**IosApnCertificate** | Pointer to **string** | Public certificate to use for Apple Push Notification Service | [optional]
**IosApnPrivate** | Pointer to **bool** | Private key to use for Apple Push Notification Service | [optional]
**UseSandbox** | Pointer to **bool** | Whether to use sandbox for Apple Push Notification Service | [optional]

## Methods

### NewExternalConfig

`func NewExternalConfig() *ExternalConfig`

NewExternalConfig instantiates a new ExternalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConfigWithDefaults

`func NewExternalConfigWithDefaults() *ExternalConfig`

NewExternalConfigWithDefaults instantiates a new ExternalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ExternalConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ExternalConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ExternalConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ExternalConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ExternalConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ExternalConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ExternalConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ExternalConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetFcmApiKey

`func (o *ExternalConfig) GetFcmApiKey() string`

GetFcmApiKey returns the FcmApiKey field if non-nil, zero value otherwise.

### GetFcmApiKeyOk

`func (o *ExternalConfig) GetFcmApiKeyOk() (*string, bool)`

GetFcmApiKeyOk returns a tuple with the FcmApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmApiKey

`func (o *ExternalConfig) SetFcmApiKey(v string)`

SetFcmApiKey sets FcmApiKey field to given value.

### HasFcmApiKey

`func (o *ExternalConfig) HasFcmApiKey() bool`

HasFcmApiKey returns a boolean if a field has been set.

### GetIosApnCertificate

`func (o *ExternalConfig) GetIosApnCertificate() string`

GetIosApnCertificate returns the IosApnCertificate field if non-nil, zero value otherwise.

### GetIosApnCertificateOk

`func (o *ExternalConfig) GetIosApnCertificateOk() (*string, bool)`

GetIosApnCertificateOk returns a tuple with the IosApnCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosApnCertificate

`func (o *ExternalConfig) SetIosApnCertificate(v string)`

SetIosApnCertificate sets IosApnCertificate field to given value.

### HasIosApnCertificate

`func (o *ExternalConfig) HasIosApnCertificate() bool`

HasIosApnCertificate returns a boolean if a field has been set.

### GetIosApnPrivate

`func (o *ExternalConfig) GetIosApnPrivate() bool`

GetIosApnPrivate returns the IosApnPrivate field if non-nil, zero value otherwise.

### GetIosApnPrivateOk

`func (o *ExternalConfig) GetIosApnPrivateOk() (*bool, bool)`

GetIosApnPrivateOk returns a tuple with the IosApnPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosApnPrivate

`func (o *ExternalConfig) SetIosApnPrivate(v bool)`

SetIosApnPrivate sets IosApnPrivate field to given value.

### HasIosApnPrivate

`func (o *ExternalConfig) HasIosApnPrivate() bool`

HasIosApnPrivate returns a boolean if a field has been set.

### GetUseSandbox

`func (o *ExternalConfig) GetUseSandbox() bool`

GetUseSandbox returns the UseSandbox field if non-nil, zero value otherwise.

### GetUseSandboxOk

`func (o *ExternalConfig) GetUseSandboxOk() (*bool, bool)`

GetUseSandboxOk returns a tuple with the UseSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSandbox

`func (o *ExternalConfig) SetUseSandbox(v bool)`

SetUseSandbox sets UseSandbox field to given value.

### HasUseSandbox

`func (o *ExternalConfig) HasUseSandbox() bool`

HasUseSandbox returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
