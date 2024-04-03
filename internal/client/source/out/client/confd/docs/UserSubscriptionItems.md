# UserSubscriptionItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserSubscription**](UserSubscription.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewUserSubscriptionItems

`func NewUserSubscriptionItems(total int32, ) *UserSubscriptionItems`

NewUserSubscriptionItems instantiates a new UserSubscriptionItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscriptionItemsWithDefaults

`func NewUserSubscriptionItemsWithDefaults() *UserSubscriptionItems`

NewUserSubscriptionItemsWithDefaults instantiates a new UserSubscriptionItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserSubscriptionItems) GetItems() []UserSubscription`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSubscriptionItems) GetItemsOk() (*[]UserSubscription, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSubscriptionItems) SetItems(v []UserSubscription)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserSubscriptionItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *UserSubscriptionItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserSubscriptionItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserSubscriptionItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
