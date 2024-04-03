# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the group | [optional] [readonly]
**Name** | Pointer to **string** | The name of the group | [optional] [readonly]
**Uuid** | Pointer to **string** | Group UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]
**Fallbacks** | Pointer to [**GroupFallbacks**](GroupFallbacks.md) |  | [optional]
**Incalls** | Pointer to [**[]GroupRelationIncall**](GroupRelationIncall.md) |  | [optional] [readonly]
**Members** | Pointer to [**GroupRelationMemberUsers**](GroupRelationMemberUsers.md) |  | [optional]
**Schedules** | Pointer to [**[]ScheduleRelationBase**](ScheduleRelationBase.md) |  | [optional] [readonly]
**CallPermissions** | Pointer to [**[]CallPermissionRelationBase**](CallPermissionRelationBase.md) |  | [optional] [readonly]
**CallerIdMode** | Pointer to **string** | How the caller_id_name will be treated | [optional]
**CallerIdName** | Pointer to **string** | Name to display | [optional]
**Enabled** | Pointer to **bool** | Enable/Disable the group | [optional] [default to true]
**MarkAnsweredElsewhere** | Pointer to **bool** | Mark all calls as \&quot;answered elsewhere\&quot; when cancelled | [optional] [default to false]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**PreprocessSubroutine** | Pointer to **string** |  | [optional]
**RetryDelay** | Pointer to **int32** | Number of seconds before the member of group will ring again | [optional]
**RingInUse** | Pointer to **bool** | Notify the member even if it already in communication | [optional] [default to true]
**RingStrategy** | Pointer to **string** |  | [optional] [default to "all"]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timeout** | Pointer to **int32** | Number of seconds the group will ring before falling back | [optional]
**UserTimeout** | Pointer to **int32** | Number of seconds the member of group will ring | [optional]

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *Group) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Group) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Group) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Group) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExtensions

`func (o *Group) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Group) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Group) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Group) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetFallbacks

`func (o *Group) GetFallbacks() GroupFallbacks`

GetFallbacks returns the Fallbacks field if non-nil, zero value otherwise.

### GetFallbacksOk

`func (o *Group) GetFallbacksOk() (*GroupFallbacks, bool)`

GetFallbacksOk returns a tuple with the Fallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbacks

`func (o *Group) SetFallbacks(v GroupFallbacks)`

SetFallbacks sets Fallbacks field to given value.

### HasFallbacks

`func (o *Group) HasFallbacks() bool`

HasFallbacks returns a boolean if a field has been set.

### GetIncalls

`func (o *Group) GetIncalls() []GroupRelationIncall`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *Group) GetIncallsOk() (*[]GroupRelationIncall, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *Group) SetIncalls(v []GroupRelationIncall)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *Group) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetMembers

`func (o *Group) GetMembers() GroupRelationMemberUsers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() (*GroupRelationMemberUsers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Group) SetMembers(v GroupRelationMemberUsers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSchedules

`func (o *Group) GetSchedules() []ScheduleRelationBase`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *Group) GetSchedulesOk() (*[]ScheduleRelationBase, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *Group) SetSchedules(v []ScheduleRelationBase)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *Group) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetCallPermissions

`func (o *Group) GetCallPermissions() []CallPermissionRelationBase`

GetCallPermissions returns the CallPermissions field if non-nil, zero value otherwise.

### GetCallPermissionsOk

`func (o *Group) GetCallPermissionsOk() (*[]CallPermissionRelationBase, bool)`

GetCallPermissionsOk returns a tuple with the CallPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissions

`func (o *Group) SetCallPermissions(v []CallPermissionRelationBase)`

SetCallPermissions sets CallPermissions field to given value.

### HasCallPermissions

`func (o *Group) HasCallPermissions() bool`

HasCallPermissions returns a boolean if a field has been set.

### GetCallerIdMode

`func (o *Group) GetCallerIdMode() string`

GetCallerIdMode returns the CallerIdMode field if non-nil, zero value otherwise.

### GetCallerIdModeOk

`func (o *Group) GetCallerIdModeOk() (*string, bool)`

GetCallerIdModeOk returns a tuple with the CallerIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdMode

`func (o *Group) SetCallerIdMode(v string)`

SetCallerIdMode sets CallerIdMode field to given value.

### HasCallerIdMode

`func (o *Group) HasCallerIdMode() bool`

HasCallerIdMode returns a boolean if a field has been set.

### GetCallerIdName

`func (o *Group) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Group) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Group) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Group) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetEnabled

`func (o *Group) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Group) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Group) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Group) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMarkAnsweredElsewhere

`func (o *Group) GetMarkAnsweredElsewhere() bool`

GetMarkAnsweredElsewhere returns the MarkAnsweredElsewhere field if non-nil, zero value otherwise.

### GetMarkAnsweredElsewhereOk

`func (o *Group) GetMarkAnsweredElsewhereOk() (*bool, bool)`

GetMarkAnsweredElsewhereOk returns a tuple with the MarkAnsweredElsewhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAnsweredElsewhere

`func (o *Group) SetMarkAnsweredElsewhere(v bool)`

SetMarkAnsweredElsewhere sets MarkAnsweredElsewhere field to given value.

### HasMarkAnsweredElsewhere

`func (o *Group) HasMarkAnsweredElsewhere() bool`

HasMarkAnsweredElsewhere returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *Group) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *Group) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *Group) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *Group) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Group) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Group) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Group) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Group) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRetryDelay

`func (o *Group) GetRetryDelay() int32`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *Group) GetRetryDelayOk() (*int32, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *Group) SetRetryDelay(v int32)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *Group) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetRingInUse

`func (o *Group) GetRingInUse() bool`

GetRingInUse returns the RingInUse field if non-nil, zero value otherwise.

### GetRingInUseOk

`func (o *Group) GetRingInUseOk() (*bool, bool)`

GetRingInUseOk returns a tuple with the RingInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingInUse

`func (o *Group) SetRingInUse(v bool)`

SetRingInUse sets RingInUse field to given value.

### HasRingInUse

`func (o *Group) HasRingInUse() bool`

HasRingInUse returns a boolean if a field has been set.

### GetRingStrategy

`func (o *Group) GetRingStrategy() string`

GetRingStrategy returns the RingStrategy field if non-nil, zero value otherwise.

### GetRingStrategyOk

`func (o *Group) GetRingStrategyOk() (*string, bool)`

GetRingStrategyOk returns a tuple with the RingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingStrategy

`func (o *Group) SetRingStrategy(v string)`

SetRingStrategy sets RingStrategy field to given value.

### HasRingStrategy

`func (o *Group) HasRingStrategy() bool`

HasRingStrategy returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Group) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Group) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Group) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Group) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *Group) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Group) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Group) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Group) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUserTimeout

`func (o *Group) GetUserTimeout() int32`

GetUserTimeout returns the UserTimeout field if non-nil, zero value otherwise.

### GetUserTimeoutOk

`func (o *Group) GetUserTimeoutOk() (*int32, bool)`

GetUserTimeoutOk returns a tuple with the UserTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTimeout

`func (o *Group) SetUserTimeout(v int32)`

SetUserTimeout sets UserTimeout field to given value.

### HasUserTimeout

`func (o *Group) HasUserTimeout() bool`

HasUserTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
