# Ivr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incalls** | Pointer to [**[]IvrRelationIncall**](IvrRelationIncall.md) |  | [optional] [readonly]
**AbortDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**AbortSound** | Pointer to **string** | The sound played when the caller reach the maximum number of tries. Not used if an abort destination is set | [optional]
**Choices** | Pointer to [**[]IvrChoice**](IvrChoice.md) | The menu&#39;s choices | [optional]
**Description** | Pointer to **string** | Additional information about the IVR | [optional]
**GreetingSound** | Pointer to **string** | The sound played to greet the caller | [optional]
**Id** | Pointer to **int32** | The id of the IVR | [optional] [readonly]
**InvalidDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**InvalidSound** | Pointer to **string** | The sound played when the caller choose an invalid option. Not used if an invalid destination is set | [optional]
**MaxTries** | Pointer to **int32** | The maximum number of tries before aborting the call. Both a timeout and an invalid choice counts toward the number of tries | [optional]
**MenuSound** | Pointer to **string** | The sound played to prompt the caller for input | [optional]
**Name** | Pointer to **string** | The name of the IVR | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timeout** | Pointer to **int32** | Number of seconds to wait after the menu sound is played before either replaying the menu, redirecting the call to the timeout destination (if set) or aborting the call (if the maximum number of tries has been reached) | [optional]
**TimeoutDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]

## Methods

### NewIvr

`func NewIvr() *Ivr`

NewIvr instantiates a new Ivr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIvrWithDefaults

`func NewIvrWithDefaults() *Ivr`

NewIvrWithDefaults instantiates a new Ivr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncalls

`func (o *Ivr) GetIncalls() []IvrRelationIncall`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *Ivr) GetIncallsOk() (*[]IvrRelationIncall, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *Ivr) SetIncalls(v []IvrRelationIncall)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *Ivr) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetAbortDestination

`func (o *Ivr) GetAbortDestination() DestinationType`

GetAbortDestination returns the AbortDestination field if non-nil, zero value otherwise.

### GetAbortDestinationOk

`func (o *Ivr) GetAbortDestinationOk() (*DestinationType, bool)`

GetAbortDestinationOk returns a tuple with the AbortDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortDestination

`func (o *Ivr) SetAbortDestination(v DestinationType)`

SetAbortDestination sets AbortDestination field to given value.

### HasAbortDestination

`func (o *Ivr) HasAbortDestination() bool`

HasAbortDestination returns a boolean if a field has been set.

### GetAbortSound

`func (o *Ivr) GetAbortSound() string`

GetAbortSound returns the AbortSound field if non-nil, zero value otherwise.

### GetAbortSoundOk

`func (o *Ivr) GetAbortSoundOk() (*string, bool)`

GetAbortSoundOk returns a tuple with the AbortSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortSound

`func (o *Ivr) SetAbortSound(v string)`

SetAbortSound sets AbortSound field to given value.

### HasAbortSound

`func (o *Ivr) HasAbortSound() bool`

HasAbortSound returns a boolean if a field has been set.

### GetChoices

`func (o *Ivr) GetChoices() []IvrChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *Ivr) GetChoicesOk() (*[]IvrChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *Ivr) SetChoices(v []IvrChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *Ivr) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDescription

`func (o *Ivr) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ivr) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Ivr) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Ivr) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGreetingSound

`func (o *Ivr) GetGreetingSound() string`

GetGreetingSound returns the GreetingSound field if non-nil, zero value otherwise.

### GetGreetingSoundOk

`func (o *Ivr) GetGreetingSoundOk() (*string, bool)`

GetGreetingSoundOk returns a tuple with the GreetingSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingSound

`func (o *Ivr) SetGreetingSound(v string)`

SetGreetingSound sets GreetingSound field to given value.

### HasGreetingSound

`func (o *Ivr) HasGreetingSound() bool`

HasGreetingSound returns a boolean if a field has been set.

### GetId

`func (o *Ivr) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ivr) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ivr) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Ivr) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvalidDestination

`func (o *Ivr) GetInvalidDestination() DestinationType`

GetInvalidDestination returns the InvalidDestination field if non-nil, zero value otherwise.

### GetInvalidDestinationOk

`func (o *Ivr) GetInvalidDestinationOk() (*DestinationType, bool)`

GetInvalidDestinationOk returns a tuple with the InvalidDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidDestination

`func (o *Ivr) SetInvalidDestination(v DestinationType)`

SetInvalidDestination sets InvalidDestination field to given value.

### HasInvalidDestination

`func (o *Ivr) HasInvalidDestination() bool`

HasInvalidDestination returns a boolean if a field has been set.

### GetInvalidSound

`func (o *Ivr) GetInvalidSound() string`

GetInvalidSound returns the InvalidSound field if non-nil, zero value otherwise.

### GetInvalidSoundOk

`func (o *Ivr) GetInvalidSoundOk() (*string, bool)`

GetInvalidSoundOk returns a tuple with the InvalidSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSound

`func (o *Ivr) SetInvalidSound(v string)`

SetInvalidSound sets InvalidSound field to given value.

### HasInvalidSound

`func (o *Ivr) HasInvalidSound() bool`

HasInvalidSound returns a boolean if a field has been set.

### GetMaxTries

`func (o *Ivr) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *Ivr) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *Ivr) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *Ivr) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetMenuSound

`func (o *Ivr) GetMenuSound() string`

GetMenuSound returns the MenuSound field if non-nil, zero value otherwise.

### GetMenuSoundOk

`func (o *Ivr) GetMenuSoundOk() (*string, bool)`

GetMenuSoundOk returns a tuple with the MenuSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuSound

`func (o *Ivr) SetMenuSound(v string)`

SetMenuSound sets MenuSound field to given value.

### HasMenuSound

`func (o *Ivr) HasMenuSound() bool`

HasMenuSound returns a boolean if a field has been set.

### GetName

`func (o *Ivr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ivr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ivr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ivr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Ivr) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Ivr) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Ivr) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Ivr) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *Ivr) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Ivr) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Ivr) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Ivr) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTimeoutDestination

`func (o *Ivr) GetTimeoutDestination() DestinationType`

GetTimeoutDestination returns the TimeoutDestination field if non-nil, zero value otherwise.

### GetTimeoutDestinationOk

`func (o *Ivr) GetTimeoutDestinationOk() (*DestinationType, bool)`

GetTimeoutDestinationOk returns a tuple with the TimeoutDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutDestination

`func (o *Ivr) SetTimeoutDestination(v DestinationType)`

SetTimeoutDestination sets TimeoutDestination field to given value.

### HasTimeoutDestination

`func (o *Ivr) HasTimeoutDestination() bool`

HasTimeoutDestination returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
