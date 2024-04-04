# FtpProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Param** | Pointer to **string** | The proxy for FTP requests. Format is &#x60;http://[user:password@]host:port&#x60; | [optional]

## Methods

### NewFtpProxy

`func NewFtpProxy() *FtpProxy`

NewFtpProxy instantiates a new FtpProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtpProxyWithDefaults

`func NewFtpProxyWithDefaults() *FtpProxy`

NewFtpProxyWithDefaults instantiates a new FtpProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParam

`func (o *FtpProxy) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *FtpProxy) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *FtpProxy) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *FtpProxy) HasParam() bool`

HasParam returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
