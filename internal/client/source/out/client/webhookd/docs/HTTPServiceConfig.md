# HTTPServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Jinja2 template, where variables come from the event triggering the webhook. For more details, see <https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/templates>. **Default:** the complete event data, JSON-encoded. | [optional]
**ContentType** | Pointer to **string** | Content-Type of the body | [optional]
**Method** | **string** |  |
**Url** | **string** | Jinja2 template, where variables come from the event triggering the webhook. For more details, see <https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/templates> |
**VerifyCertificate** | Pointer to **string** | May be &#x60;true&#x60;, &#x60;false&#x60; or a path to the certificate bundle | [optional] [default to "true"]

## Methods

### NewHTTPServiceConfig

`func NewHTTPServiceConfig(method string, url string, ) *HTTPServiceConfig`

NewHTTPServiceConfig instantiates a new HTTPServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPServiceConfigWithDefaults

`func NewHTTPServiceConfigWithDefaults() *HTTPServiceConfig`

NewHTTPServiceConfigWithDefaults instantiates a new HTTPServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *HTTPServiceConfig) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HTTPServiceConfig) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HTTPServiceConfig) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HTTPServiceConfig) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *HTTPServiceConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HTTPServiceConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HTTPServiceConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HTTPServiceConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMethod

`func (o *HTTPServiceConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HTTPServiceConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HTTPServiceConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### GetUrl

`func (o *HTTPServiceConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HTTPServiceConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HTTPServiceConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### GetVerifyCertificate

`func (o *HTTPServiceConfig) GetVerifyCertificate() string`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *HTTPServiceConfig) GetVerifyCertificateOk() (*string, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *HTTPServiceConfig) SetVerifyCertificate(v string)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *HTTPServiceConfig) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
