# SubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**HTTPServiceConfig**](HTTPServiceConfig.md) |  |
**Events** | **[]string** |  |
**EventsUserUuid** | Pointer to **string** | Only trigger webhook when an event occurs related to this user. Not compatible with all events. For more details, see: <https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/user_filter>. | [optional]
**EventsAccentUuid** | Pointer to **string** | Only trigger webhook when an event occurs on this Accent. | [optional]
**Name** | **string** |  |
**Service** | **string** | Known services: http. The service may be arbitrary, but it must be bound to an installed plugin in order to be effective.  |
**Tags** | Pointer to **map[string]interface{}** |  | [optional]

## Methods

### NewSubscriptionRequest

`func NewSubscriptionRequest(config HTTPServiceConfig, events []string, name string, service string, ) *SubscriptionRequest`

NewSubscriptionRequest instantiates a new SubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRequestWithDefaults

`func NewSubscriptionRequestWithDefaults() *SubscriptionRequest`

NewSubscriptionRequestWithDefaults instantiates a new SubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SubscriptionRequest) GetConfig() HTTPServiceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SubscriptionRequest) GetConfigOk() (*HTTPServiceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SubscriptionRequest) SetConfig(v HTTPServiceConfig)`

SetConfig sets Config field to given value.

### GetEvents

`func (o *SubscriptionRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubscriptionRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubscriptionRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### GetEventsUserUuid

`func (o *SubscriptionRequest) GetEventsUserUuid() string`

GetEventsUserUuid returns the EventsUserUuid field if non-nil, zero value otherwise.

### GetEventsUserUuidOk

`func (o *SubscriptionRequest) GetEventsUserUuidOk() (*string, bool)`

GetEventsUserUuidOk returns a tuple with the EventsUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUserUuid

`func (o *SubscriptionRequest) SetEventsUserUuid(v string)`

SetEventsUserUuid sets EventsUserUuid field to given value.

### HasEventsUserUuid

`func (o *SubscriptionRequest) HasEventsUserUuid() bool`

HasEventsUserUuid returns a boolean if a field has been set.

### GetEventsAccentUuid

`func (o *SubscriptionRequest) GetEventsAccentUuid() string`

GetEventsAccentUuid returns the EventsAccentUuid field if non-nil, zero value otherwise.

### GetEventsAccentUuidOk

`func (o *SubscriptionRequest) GetEventsAccentUuidOk() (*string, bool)`

GetEventsAccentUuidOk returns a tuple with the EventsAccentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsAccentUuid

`func (o *SubscriptionRequest) SetEventsAccentUuid(v string)`

SetEventsAccentUuid sets EventsAccentUuid field to given value.

### HasEventsAccentUuid

`func (o *SubscriptionRequest) HasEventsAccentUuid() bool`

HasEventsAccentUuid returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionRequest) SetName(v string)`

SetName sets Name field to given value.

### GetService

`func (o *SubscriptionRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SubscriptionRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SubscriptionRequest) SetService(v string)`

SetService sets Service field to given value.

### GetTags

`func (o *SubscriptionRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SubscriptionRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SubscriptionRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SubscriptionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
