# UserSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**HTTPServiceConfig**](HTTPServiceConfig.md) |  |
**Events** | **[]string** |  |
**Name** | **string** |  |
**Service** | **string** | Known services: http. The service may be arbitrary, but it must be bound to an installed plugin in order to be effective.  |
**Tags** | Pointer to **map[string]interface{}** | Arbitrary key-value storage for this subscription. May be used to tag subscriptions. PUT replaces all metadata. | [optional]

## Methods

### NewUserSubscriptionRequest

`func NewUserSubscriptionRequest(config HTTPServiceConfig, events []string, name string, service string, ) *UserSubscriptionRequest`

NewUserSubscriptionRequest instantiates a new UserSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscriptionRequestWithDefaults

`func NewUserSubscriptionRequestWithDefaults() *UserSubscriptionRequest`

NewUserSubscriptionRequestWithDefaults instantiates a new UserSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UserSubscriptionRequest) GetConfig() HTTPServiceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UserSubscriptionRequest) GetConfigOk() (*HTTPServiceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UserSubscriptionRequest) SetConfig(v HTTPServiceConfig)`

SetConfig sets Config field to given value.

### GetEvents

`func (o *UserSubscriptionRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UserSubscriptionRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UserSubscriptionRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### GetName

`func (o *UserSubscriptionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSubscriptionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSubscriptionRequest) SetName(v string)`

SetName sets Name field to given value.

### GetService

`func (o *UserSubscriptionRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UserSubscriptionRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UserSubscriptionRequest) SetService(v string)`

SetService sets Service field to given value.

### GetTags

`func (o *UserSubscriptionRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserSubscriptionRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserSubscriptionRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserSubscriptionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
