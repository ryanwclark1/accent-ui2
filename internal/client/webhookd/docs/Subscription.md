# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]string** |  |
**Events** | **[]string** |  |
**EventsUserUuid** | Pointer to **string** | Only trigger webhook when an event occurs related to this user. Not compatible with all events. For more details, see: <https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/user_filter> | [optional]
**Name** | **string** |  |
**OwnerUserUuid** | Pointer to **string** | The user who owns this subscription. Admin-created subscriptions are not owned. | [optional] [readonly]
**Service** | **string** |  |
**Uuid** | Pointer to **string** |  | [optional] [readonly]

## Methods

### NewSubscription

`func NewSubscription(config map[string]string, events []string, name string, service string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Subscription) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Subscription) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Subscription) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### GetEvents

`func (o *Subscription) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Subscription) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Subscription) SetEvents(v []string)`

SetEvents sets Events field to given value.

### GetEventsUserUuid

`func (o *Subscription) GetEventsUserUuid() string`

GetEventsUserUuid returns the EventsUserUuid field if non-nil, zero value otherwise.

### GetEventsUserUuidOk

`func (o *Subscription) GetEventsUserUuidOk() (*string, bool)`

GetEventsUserUuidOk returns a tuple with the EventsUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUserUuid

`func (o *Subscription) SetEventsUserUuid(v string)`

SetEventsUserUuid sets EventsUserUuid field to given value.

### HasEventsUserUuid

`func (o *Subscription) HasEventsUserUuid() bool`

HasEventsUserUuid returns a boolean if a field has been set.

### GetName

`func (o *Subscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscription) SetName(v string)`

SetName sets Name field to given value.

### GetOwnerUserUuid

`func (o *Subscription) GetOwnerUserUuid() string`

GetOwnerUserUuid returns the OwnerUserUuid field if non-nil, zero value otherwise.

### GetOwnerUserUuidOk

`func (o *Subscription) GetOwnerUserUuidOk() (*string, bool)`

GetOwnerUserUuidOk returns a tuple with the OwnerUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserUuid

`func (o *Subscription) SetOwnerUserUuid(v string)`

SetOwnerUserUuid sets OwnerUserUuid field to given value.

### HasOwnerUserUuid

`func (o *Subscription) HasOwnerUserUuid() bool`

HasOwnerUserUuid returns a boolean if a field has been set.

### GetService

`func (o *Subscription) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Subscription) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Subscription) SetService(v string)`

SetService sets Service field to given value.

### GetUuid

`func (o *Subscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Subscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Subscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Subscription) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
