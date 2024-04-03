# LoginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** | Context | [optional]
**Extension** | Pointer to **string** | Extension | [optional]

## Methods

### NewLoginInfo

`func NewLoginInfo() *LoginInfo`

NewLoginInfo instantiates a new LoginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginInfoWithDefaults

`func NewLoginInfoWithDefaults() *LoginInfo`

NewLoginInfoWithDefaults instantiates a new LoginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *LoginInfo) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LoginInfo) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LoginInfo) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *LoginInfo) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExtension

`func (o *LoginInfo) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *LoginInfo) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *LoginInfo) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *LoginInfo) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
