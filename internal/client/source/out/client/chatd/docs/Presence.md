# Presence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | If the user has active connections. | [optional] [readonly]
**DoNotDisturb** | Pointer to **bool** | The \&quot;do not disturb\&quot; status of the user. | [optional] [readonly]
**LastActivity** | Pointer to **string** | The date time in UTC of the user&#39;s last activity. The value is updated when the user changes his state explicitly. The value is NULL for new user. | [optional] [readonly]
**LineState** | Pointer to **string** | The current state of the most prioritize state line. The prioritization of each state is the following: ringing &gt; progressing &gt; holding &gt; talking &gt; available &gt; unavailable | [optional] [readonly]
**Lines** | Pointer to [**[]Line**](Line.md) |  | [optional] [readonly]
**Mobile** | Pointer to **bool** | If the user uses a mobile application and can be considered reachable.  mobile will be true in the following situations  *The user has a mobile refresh token* The user has a mobile session  | [optional] [readonly]
**State** | **string** | The presence state of the user. |
**Status** | Pointer to **string** | An extended description of the user presence. | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant of the user | [optional] [readonly]
**Uuid** | Pointer to **string** | The UUID of the user | [optional] [readonly]

## Methods

### NewPresence

`func NewPresence(state string, ) *Presence`

NewPresence instantiates a new Presence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenceWithDefaults

`func NewPresenceWithDefaults() *Presence`

NewPresenceWithDefaults instantiates a new Presence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *Presence) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Presence) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *Presence) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *Presence) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *Presence) GetDoNotDisturb() bool`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *Presence) GetDoNotDisturbOk() (*bool, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *Presence) SetDoNotDisturb(v bool)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *Presence) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetLastActivity

`func (o *Presence) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *Presence) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *Presence) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *Presence) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetLineState

`func (o *Presence) GetLineState() string`

GetLineState returns the LineState field if non-nil, zero value otherwise.

### GetLineStateOk

`func (o *Presence) GetLineStateOk() (*string, bool)`

GetLineStateOk returns a tuple with the LineState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineState

`func (o *Presence) SetLineState(v string)`

SetLineState sets LineState field to given value.

### HasLineState

`func (o *Presence) HasLineState() bool`

HasLineState returns a boolean if a field has been set.

### GetLines

`func (o *Presence) GetLines() []Line`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Presence) GetLinesOk() (*[]Line, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Presence) SetLines(v []Line)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Presence) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMobile

`func (o *Presence) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *Presence) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *Presence) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *Presence) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetState

`func (o *Presence) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Presence) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Presence) SetState(v string)`

SetState sets State field to given value.

### GetStatus

`func (o *Presence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Presence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Presence) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Presence) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Presence) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Presence) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Presence) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Presence) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUuid

`func (o *Presence) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Presence) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Presence) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Presence) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
