# HTTPIngress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | The public URI to contact this stack HTTP API |
**Uuid** | Pointer to **string** |  | [optional] [readonly]

## Methods

### NewHTTPIngress

`func NewHTTPIngress(uri string, ) *HTTPIngress`

NewHTTPIngress instantiates a new HTTPIngress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPIngressWithDefaults

`func NewHTTPIngressWithDefaults() *HTTPIngress`

NewHTTPIngressWithDefaults instantiates a new HTTPIngress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *HTTPIngress) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *HTTPIngress) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *HTTPIngress) SetUri(v string)`

SetUri sets Uri field to given value.

### GetUuid

`func (o *HTTPIngress) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HTTPIngress) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HTTPIngress) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HTTPIngress) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
