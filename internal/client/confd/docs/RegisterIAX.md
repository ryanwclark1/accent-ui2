# RegisterIAX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the register IAX | [optional] [readonly]
**AuthPassword** | Pointer to **string** | The password to authenticate to the remote_host | [optional]
**AuthUsername** | Pointer to **string** | The username used by the remote_host for the authentication | [optional]
**CallbackContext** | Pointer to **string** | The callback context to use for the register | [optional]
**CallbackExtension** | Pointer to **string** | The callback extension to use for the register | [optional]
**RemoteHost** | **string** | The register domain |
**RemotePort** | Pointer to **int32** | The port of the remote_host | [optional]

## Methods

### NewRegisterIAX

`func NewRegisterIAX(remoteHost string, ) *RegisterIAX`

NewRegisterIAX instantiates a new RegisterIAX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIAXWithDefaults

`func NewRegisterIAXWithDefaults() *RegisterIAX`

NewRegisterIAXWithDefaults instantiates a new RegisterIAX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterIAX) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterIAX) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterIAX) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterIAX) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthPassword

`func (o *RegisterIAX) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *RegisterIAX) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *RegisterIAX) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *RegisterIAX) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthUsername

`func (o *RegisterIAX) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *RegisterIAX) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *RegisterIAX) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *RegisterIAX) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetCallbackContext

`func (o *RegisterIAX) GetCallbackContext() string`

GetCallbackContext returns the CallbackContext field if non-nil, zero value otherwise.

### GetCallbackContextOk

`func (o *RegisterIAX) GetCallbackContextOk() (*string, bool)`

GetCallbackContextOk returns a tuple with the CallbackContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackContext

`func (o *RegisterIAX) SetCallbackContext(v string)`

SetCallbackContext sets CallbackContext field to given value.

### HasCallbackContext

`func (o *RegisterIAX) HasCallbackContext() bool`

HasCallbackContext returns a boolean if a field has been set.

### GetCallbackExtension

`func (o *RegisterIAX) GetCallbackExtension() string`

GetCallbackExtension returns the CallbackExtension field if non-nil, zero value otherwise.

### GetCallbackExtensionOk

`func (o *RegisterIAX) GetCallbackExtensionOk() (*string, bool)`

GetCallbackExtensionOk returns a tuple with the CallbackExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackExtension

`func (o *RegisterIAX) SetCallbackExtension(v string)`

SetCallbackExtension sets CallbackExtension field to given value.

### HasCallbackExtension

`func (o *RegisterIAX) HasCallbackExtension() bool`

HasCallbackExtension returns a boolean if a field has been set.

### GetRemoteHost

`func (o *RegisterIAX) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *RegisterIAX) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *RegisterIAX) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.

### GetRemotePort

`func (o *RegisterIAX) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *RegisterIAX) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *RegisterIAX) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *RegisterIAX) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
