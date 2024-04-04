# Incall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the incoming call | [optional] [readonly]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]
**Schedules** | Pointer to [**[]ScheduleRelationBase**](ScheduleRelationBase.md) |  | [optional] [readonly]
**CallerIdMode** | Pointer to **string** | How the caller_id_name will be treated | [optional]
**CallerIdName** | Pointer to **string** | Name to display when calling | [optional]
**Description** | Pointer to **string** | Additional information about the incoming call | [optional]
**Destination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**GreetingSound** | Pointer to **string** | The name of the sound file to be played before redirecting the caller to the destination | [optional]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewIncall

`func NewIncall() *Incall`

NewIncall instantiates a new Incall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncallWithDefaults

`func NewIncallWithDefaults() *Incall`

NewIncallWithDefaults instantiates a new Incall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Incall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Incall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtensions

`func (o *Incall) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Incall) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Incall) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Incall) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetSchedules

`func (o *Incall) GetSchedules() []ScheduleRelationBase`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *Incall) GetSchedulesOk() (*[]ScheduleRelationBase, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *Incall) SetSchedules(v []ScheduleRelationBase)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *Incall) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetCallerIdMode

`func (o *Incall) GetCallerIdMode() string`

GetCallerIdMode returns the CallerIdMode field if non-nil, zero value otherwise.

### GetCallerIdModeOk

`func (o *Incall) GetCallerIdModeOk() (*string, bool)`

GetCallerIdModeOk returns a tuple with the CallerIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdMode

`func (o *Incall) SetCallerIdMode(v string)`

SetCallerIdMode sets CallerIdMode field to given value.

### HasCallerIdMode

`func (o *Incall) HasCallerIdMode() bool`

HasCallerIdMode returns a boolean if a field has been set.

### GetCallerIdName

`func (o *Incall) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Incall) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Incall) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Incall) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetDescription

`func (o *Incall) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Incall) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Incall) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Incall) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestination

`func (o *Incall) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Incall) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Incall) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Incall) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetGreetingSound

`func (o *Incall) GetGreetingSound() string`

GetGreetingSound returns the GreetingSound field if non-nil, zero value otherwise.

### GetGreetingSoundOk

`func (o *Incall) GetGreetingSoundOk() (*string, bool)`

GetGreetingSoundOk returns a tuple with the GreetingSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingSound

`func (o *Incall) SetGreetingSound(v string)`

SetGreetingSound sets GreetingSound field to given value.

### HasGreetingSound

`func (o *Incall) HasGreetingSound() bool`

HasGreetingSound returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Incall) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Incall) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Incall) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Incall) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Incall) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Incall) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Incall) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Incall) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
