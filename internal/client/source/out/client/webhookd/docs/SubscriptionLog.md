# SubscriptionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | The current attempts | [optional]
**Detail** | Pointer to [**HTTPServiceLog**](HTTPServiceLog.md) |  | [optional]
**EndedAt** | Pointer to **time.Time** |  | [optional]
**Event** | Pointer to **string** |  | [optional]
**MaxAttempts** | Pointer to **int32** | Limit of number of attempts | [optional]
**StartedAt** | Pointer to **time.Time** |  | [optional]
**Status** | Pointer to **string** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewSubscriptionLog

`func NewSubscriptionLog() *SubscriptionLog`

NewSubscriptionLog instantiates a new SubscriptionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionLogWithDefaults

`func NewSubscriptionLogWithDefaults() *SubscriptionLog`

NewSubscriptionLogWithDefaults instantiates a new SubscriptionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *SubscriptionLog) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *SubscriptionLog) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *SubscriptionLog) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *SubscriptionLog) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetDetail

`func (o *SubscriptionLog) GetDetail() HTTPServiceLog`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SubscriptionLog) GetDetailOk() (*HTTPServiceLog, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SubscriptionLog) SetDetail(v HTTPServiceLog)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SubscriptionLog) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetEndedAt

`func (o *SubscriptionLog) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SubscriptionLog) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SubscriptionLog) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *SubscriptionLog) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvent

`func (o *SubscriptionLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SubscriptionLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SubscriptionLog) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *SubscriptionLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *SubscriptionLog) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *SubscriptionLog) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *SubscriptionLog) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *SubscriptionLog) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetStartedAt

`func (o *SubscriptionLog) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SubscriptionLog) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SubscriptionLog) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SubscriptionLog) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *SubscriptionLog) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SubscriptionLog) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SubscriptionLog) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SubscriptionLog) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
