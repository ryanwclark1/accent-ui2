# HTTPIngressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | The public URI to contact this stack HTTP API |

## Methods

### NewHTTPIngressRequest

`func NewHTTPIngressRequest(uri string, ) *HTTPIngressRequest`

NewHTTPIngressRequest instantiates a new HTTPIngressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPIngressRequestWithDefaults

`func NewHTTPIngressRequestWithDefaults() *HTTPIngressRequest`

NewHTTPIngressRequestWithDefaults instantiates a new HTTPIngressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *HTTPIngressRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *HTTPIngressRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *HTTPIngressRequest) SetUri(v string)`

SetUri sets Uri field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
