# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional]
**Exten** | Pointer to **string** |  | [optional]
**Id** | Pointer to **int32** | Extension ID | [optional] [readonly]
**Conference** | Pointer to [**ConferenceRelationBase**](ConferenceRelationBase.md) |  | [optional]
**Group** | Pointer to [**GroupRelationBase**](GroupRelationBase.md) |  | [optional]
**Incall** | Pointer to [**IncallRelationBase**](IncallRelationBase.md) |  | [optional]
**Lines** | Pointer to [**[]LineRelationBase**](LineRelationBase.md) |  | [optional] [readonly]
**Outcall** | Pointer to [**OutcallRelationBase**](OutcallRelationBase.md) |  | [optional]
**ParkingLot** | Pointer to [**ParkingLotRelationBase**](ParkingLotRelationBase.md) |  | [optional]
**Queue** | Pointer to [**QueueRelationBase**](QueueRelationBase.md) |  | [optional]
**Commented** | Pointer to **bool** | If True the extension is disabled. Deprecated, use enabled instead | [optional]
**Enabled** | Pointer to **bool** | If False the extension is disabled. | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewExtension

`func NewExtension() *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Extension) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Extension) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Extension) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Extension) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExten

`func (o *Extension) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *Extension) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *Extension) SetExten(v string)`

SetExten sets Exten field to given value.

### HasExten

`func (o *Extension) HasExten() bool`

HasExten returns a boolean if a field has been set.

### GetId

`func (o *Extension) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Extension) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Extension) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Extension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConference

`func (o *Extension) GetConference() ConferenceRelationBase`

GetConference returns the Conference field if non-nil, zero value otherwise.

### GetConferenceOk

`func (o *Extension) GetConferenceOk() (*ConferenceRelationBase, bool)`

GetConferenceOk returns a tuple with the Conference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConference

`func (o *Extension) SetConference(v ConferenceRelationBase)`

SetConference sets Conference field to given value.

### HasConference

`func (o *Extension) HasConference() bool`

HasConference returns a boolean if a field has been set.

### GetGroup

`func (o *Extension) GetGroup() GroupRelationBase`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Extension) GetGroupOk() (*GroupRelationBase, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Extension) SetGroup(v GroupRelationBase)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Extension) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIncall

`func (o *Extension) GetIncall() IncallRelationBase`

GetIncall returns the Incall field if non-nil, zero value otherwise.

### GetIncallOk

`func (o *Extension) GetIncallOk() (*IncallRelationBase, bool)`

GetIncallOk returns a tuple with the Incall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncall

`func (o *Extension) SetIncall(v IncallRelationBase)`

SetIncall sets Incall field to given value.

### HasIncall

`func (o *Extension) HasIncall() bool`

HasIncall returns a boolean if a field has been set.

### GetLines

`func (o *Extension) GetLines() []LineRelationBase`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Extension) GetLinesOk() (*[]LineRelationBase, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Extension) SetLines(v []LineRelationBase)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Extension) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetOutcall

`func (o *Extension) GetOutcall() OutcallRelationBase`

GetOutcall returns the Outcall field if non-nil, zero value otherwise.

### GetOutcallOk

`func (o *Extension) GetOutcallOk() (*OutcallRelationBase, bool)`

GetOutcallOk returns a tuple with the Outcall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcall

`func (o *Extension) SetOutcall(v OutcallRelationBase)`

SetOutcall sets Outcall field to given value.

### HasOutcall

`func (o *Extension) HasOutcall() bool`

HasOutcall returns a boolean if a field has been set.

### GetParkingLot

`func (o *Extension) GetParkingLot() ParkingLotRelationBase`

GetParkingLot returns the ParkingLot field if non-nil, zero value otherwise.

### GetParkingLotOk

`func (o *Extension) GetParkingLotOk() (*ParkingLotRelationBase, bool)`

GetParkingLotOk returns a tuple with the ParkingLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingLot

`func (o *Extension) SetParkingLot(v ParkingLotRelationBase)`

SetParkingLot sets ParkingLot field to given value.

### HasParkingLot

`func (o *Extension) HasParkingLot() bool`

HasParkingLot returns a boolean if a field has been set.

### GetQueue

`func (o *Extension) GetQueue() QueueRelationBase`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *Extension) GetQueueOk() (*QueueRelationBase, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *Extension) SetQueue(v QueueRelationBase)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *Extension) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetCommented

`func (o *Extension) GetCommented() bool`

GetCommented returns the Commented field if non-nil, zero value otherwise.

### GetCommentedOk

`func (o *Extension) GetCommentedOk() (*bool, bool)`

GetCommentedOk returns a tuple with the Commented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommented

`func (o *Extension) SetCommented(v bool)`

SetCommented sets Commented field to given value.

### HasCommented

`func (o *Extension) HasCommented() bool`

HasCommented returns a boolean if a field has been set.

### GetEnabled

`func (o *Extension) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Extension) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Extension) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Extension) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Extension) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Extension) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Extension) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Extension) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
