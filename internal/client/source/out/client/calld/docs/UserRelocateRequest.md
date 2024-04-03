# UserRelocateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAnswer** | Pointer to **bool** | Inform the destination phone that it should answer automatically. Limitation: this does not work on SCCP phones. | [optional]
**Completions** | Pointer to [**[]RelocateCompletion**](RelocateCompletion.md) |  | [optional]
**Destination** | **string** | What kind of destination the relocated call should be connected |
**InitiatorCall** | **string** | Call ID of the relocate initiator. This call must be owned by the authenticated user. |
**Location** | Pointer to [**UserRelocateLocation**](UserRelocateLocation.md) |  | [optional]
**Timeout** | Pointer to **int32** | Number of seconds to wait for the recipient to answer | [optional]

## Methods

### NewUserRelocateRequest

`func NewUserRelocateRequest(destination string, initiatorCall string, ) *UserRelocateRequest`

NewUserRelocateRequest instantiates a new UserRelocateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelocateRequestWithDefaults

`func NewUserRelocateRequestWithDefaults() *UserRelocateRequest`

NewUserRelocateRequestWithDefaults instantiates a new UserRelocateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAnswer

`func (o *UserRelocateRequest) GetAutoAnswer() bool`

GetAutoAnswer returns the AutoAnswer field if non-nil, zero value otherwise.

### GetAutoAnswerOk

`func (o *UserRelocateRequest) GetAutoAnswerOk() (*bool, bool)`

GetAutoAnswerOk returns a tuple with the AutoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswer

`func (o *UserRelocateRequest) SetAutoAnswer(v bool)`

SetAutoAnswer sets AutoAnswer field to given value.

### HasAutoAnswer

`func (o *UserRelocateRequest) HasAutoAnswer() bool`

HasAutoAnswer returns a boolean if a field has been set.

### GetCompletions

`func (o *UserRelocateRequest) GetCompletions() []RelocateCompletion`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *UserRelocateRequest) GetCompletionsOk() (*[]RelocateCompletion, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *UserRelocateRequest) SetCompletions(v []RelocateCompletion)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *UserRelocateRequest) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetDestination

`func (o *UserRelocateRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UserRelocateRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UserRelocateRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### GetInitiatorCall

`func (o *UserRelocateRequest) GetInitiatorCall() string`

GetInitiatorCall returns the InitiatorCall field if non-nil, zero value otherwise.

### GetInitiatorCallOk

`func (o *UserRelocateRequest) GetInitiatorCallOk() (*string, bool)`

GetInitiatorCallOk returns a tuple with the InitiatorCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCall

`func (o *UserRelocateRequest) SetInitiatorCall(v string)`

SetInitiatorCall sets InitiatorCall field to given value.

### GetLocation

`func (o *UserRelocateRequest) GetLocation() UserRelocateLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserRelocateRequest) GetLocationOk() (*UserRelocateLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserRelocateRequest) SetLocation(v UserRelocateLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UserRelocateRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTimeout

`func (o *UserRelocateRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UserRelocateRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UserRelocateRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UserRelocateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
