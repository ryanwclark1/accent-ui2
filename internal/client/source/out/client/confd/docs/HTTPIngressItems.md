# HTTPIngressItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]HTTPIngress**](HTTPIngress.md) |  | [optional]

## Methods

### NewHTTPIngressItems

`func NewHTTPIngressItems() *HTTPIngressItems`

NewHTTPIngressItems instantiates a new HTTPIngressItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPIngressItemsWithDefaults

`func NewHTTPIngressItemsWithDefaults() *HTTPIngressItems`

NewHTTPIngressItemsWithDefaults instantiates a new HTTPIngressItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HTTPIngressItems) GetItems() []HTTPIngress`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HTTPIngressItems) GetItemsOk() (*[]HTTPIngress, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HTTPIngressItems) SetItems(v []HTTPIngress)`

SetItems sets Items field to given value.

### HasItems

`func (o *HTTPIngressItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
