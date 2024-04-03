# SubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] [readonly]
**Total** | Pointer to **int32** |  | [optional] [readonly]

## Methods

### NewSubscriptionList

`func NewSubscriptionList() *SubscriptionList`

NewSubscriptionList instantiates a new SubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListWithDefaults

`func NewSubscriptionListWithDefaults() *SubscriptionList`

NewSubscriptionListWithDefaults instantiates a new SubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SubscriptionList) GetItems() []Subscription`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubscriptionList) GetItemsOk() (*[]Subscription, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubscriptionList) SetItems(v []Subscription)`

SetItems sets Items field to given value.

### HasItems

`func (o *SubscriptionList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *SubscriptionList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SubscriptionList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SubscriptionList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SubscriptionList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
