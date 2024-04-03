# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the queue | [optional] [readonly]
**Label** | Pointer to **string** | The label of the queue | [optional]
**Name** | Pointer to **string** | The name of the queue. Cannot be modified | [optional]
**Schedules** | Pointer to [**[]ScheduleRelationBase**](ScheduleRelationBase.md) |  | [optional] [readonly]
**Members** | Pointer to [**QueueRelationMemberAgents**](QueueRelationMemberAgents.md) |  | [optional]
**AnnounceHoldTimeOnEntry** | Pointer to **bool** | Announce hold time on entry | [optional] [default to false]
**CallerIdMode** | Pointer to **string** | How the caller_id_name will be treated | [optional]
**CallerIdName** | Pointer to **string** | Name to display | [optional]
**DataQuality** | Pointer to **bool** | Asterisk definition: data-quality (modem) call (minimum delay) | [optional] [default to false]
**DtmfHangupCalleeEnabled** | Pointer to **bool** | Enable DTMF hangup by callee | [optional] [default to false]
**DtmfHangupCallerEnabled** | Pointer to **bool** | Enable DTMF hangup by caller | [optional] [default to false]
**DtmfRecordCalleeEnabled** | Pointer to **bool** | Enable DTMF records by callee | [optional] [default to false]
**DtmfRecordCallerEnabled** | Pointer to **bool** | Enable DTMF records by caller | [optional] [default to false]
**DtmfTransferCalleeEnabled** | Pointer to **bool** | Enable DTMF transfers by callee | [optional] [default to false]
**DtmfTransferCallerEnabled** | Pointer to **bool** | Enable DTMF transfers by caller | [optional] [default to false]
**Enabled** | Pointer to **bool** | Enable/Disable the queue | [optional] [default to true]
**IgnoreForward** | Pointer to **bool** | Ignore call forward requests from members | [optional] [default to false]
**MarkAnsweredElsewhere** | Pointer to **bool** | Mark all calls as \&quot;answered elsewhere\&quot; when cancelled | [optional] [default to true]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**Options** | Pointer to **[][]string** | Advanced configuration options. Options are appended at the end of a queue account in the file &#39;queues.conf&#39; used by asterisk. Please consult the asterisk documentation for further details on available parameters. Because of database limitations, only the following options can be configured:  *announce* context *timeout* monitor-type *monitor-format* queue-youarenext *queue-thereare* queue-callswaiting *queue-holdtime* queue-minutes *queue-seconds* queue-thankyou *queue-reporthold* periodic-announce *announce-frequency* periodic-announce-frequency *announce-round-seconds* announce-holdtime *retry* wrapuptime *maxlen* servicelevel *strategy* joinempty *leavewhenempty* ringinuse *reportholdtime* memberdelay *weight* timeoutrestart *timeoutpriority* autofill *autopause* setinterfacevar *setqueueentryvar* setqueuevar *membermacro* min-announce-frequency *random-periodic-announce* announce-position *announce-position-limit* defaultrule  Options must have the following form: &#x60;&#x60;&#x60; {   \&quot;options\&quot;: [      [\&quot;name1\&quot;, \&quot;value1\&quot;],      [\&quot;name2\&quot;, \&quot;value2\&quot;]   ] } &#x60;&#x60;&#x60;  The resulting configuration in queues.conf will have the following form: &#x60;&#x60;&#x60; [queuename] name1&#x3D;value1 name2&#x3D;value2 &#x60;&#x60;&#x60;  | [optional]
**PreprocessSubroutine** | Pointer to **string** |  | [optional]
**RetryOnTimeout** | Pointer to **bool** | Retry when time has elapsed | [optional] [default to true]
**RingOnHold** | Pointer to **bool** | Ring instead of On-Hold Music | [optional] [default to false]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timeout** | Pointer to **int32** | Number of seconds the queue will ring before falling back | [optional]
**WaitRatioDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**WaitRatioThreshold** | Pointer to **int32** | Wait ratio (waiting calls per logged-in agent) threshold before fallback on &#39;wait_ratio_destination&#39;. Set to &#39;null&#39; to disable it. | [optional]
**WaitTimeDestination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]
**WaitTimeThreshold** | Pointer to **int32** | Wait time threshold before fallback on &#39;wait_time_destination&#39;. Set to &#39;null&#39; to disable it. | [optional]

## Methods

### NewQueue

`func NewQueue() *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Queue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Queue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Queue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Queue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Queue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Queue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Queue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Queue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *Queue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Queue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Queue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Queue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedules

`func (o *Queue) GetSchedules() []ScheduleRelationBase`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *Queue) GetSchedulesOk() (*[]ScheduleRelationBase, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *Queue) SetSchedules(v []ScheduleRelationBase)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *Queue) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetMembers

`func (o *Queue) GetMembers() QueueRelationMemberAgents`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Queue) GetMembersOk() (*QueueRelationMemberAgents, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Queue) SetMembers(v QueueRelationMemberAgents)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Queue) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAnnounceHoldTimeOnEntry

`func (o *Queue) GetAnnounceHoldTimeOnEntry() bool`

GetAnnounceHoldTimeOnEntry returns the AnnounceHoldTimeOnEntry field if non-nil, zero value otherwise.

### GetAnnounceHoldTimeOnEntryOk

`func (o *Queue) GetAnnounceHoldTimeOnEntryOk() (*bool, bool)`

GetAnnounceHoldTimeOnEntryOk returns a tuple with the AnnounceHoldTimeOnEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceHoldTimeOnEntry

`func (o *Queue) SetAnnounceHoldTimeOnEntry(v bool)`

SetAnnounceHoldTimeOnEntry sets AnnounceHoldTimeOnEntry field to given value.

### HasAnnounceHoldTimeOnEntry

`func (o *Queue) HasAnnounceHoldTimeOnEntry() bool`

HasAnnounceHoldTimeOnEntry returns a boolean if a field has been set.

### GetCallerIdMode

`func (o *Queue) GetCallerIdMode() string`

GetCallerIdMode returns the CallerIdMode field if non-nil, zero value otherwise.

### GetCallerIdModeOk

`func (o *Queue) GetCallerIdModeOk() (*string, bool)`

GetCallerIdModeOk returns a tuple with the CallerIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdMode

`func (o *Queue) SetCallerIdMode(v string)`

SetCallerIdMode sets CallerIdMode field to given value.

### HasCallerIdMode

`func (o *Queue) HasCallerIdMode() bool`

HasCallerIdMode returns a boolean if a field has been set.

### GetCallerIdName

`func (o *Queue) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Queue) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Queue) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Queue) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetDataQuality

`func (o *Queue) GetDataQuality() bool`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *Queue) GetDataQualityOk() (*bool, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *Queue) SetDataQuality(v bool)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *Queue) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### GetDtmfHangupCalleeEnabled

`func (o *Queue) GetDtmfHangupCalleeEnabled() bool`

GetDtmfHangupCalleeEnabled returns the DtmfHangupCalleeEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupCalleeEnabledOk

`func (o *Queue) GetDtmfHangupCalleeEnabledOk() (*bool, bool)`

GetDtmfHangupCalleeEnabledOk returns a tuple with the DtmfHangupCalleeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupCalleeEnabled

`func (o *Queue) SetDtmfHangupCalleeEnabled(v bool)`

SetDtmfHangupCalleeEnabled sets DtmfHangupCalleeEnabled field to given value.

### HasDtmfHangupCalleeEnabled

`func (o *Queue) HasDtmfHangupCalleeEnabled() bool`

HasDtmfHangupCalleeEnabled returns a boolean if a field has been set.

### GetDtmfHangupCallerEnabled

`func (o *Queue) GetDtmfHangupCallerEnabled() bool`

GetDtmfHangupCallerEnabled returns the DtmfHangupCallerEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupCallerEnabledOk

`func (o *Queue) GetDtmfHangupCallerEnabledOk() (*bool, bool)`

GetDtmfHangupCallerEnabledOk returns a tuple with the DtmfHangupCallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupCallerEnabled

`func (o *Queue) SetDtmfHangupCallerEnabled(v bool)`

SetDtmfHangupCallerEnabled sets DtmfHangupCallerEnabled field to given value.

### HasDtmfHangupCallerEnabled

`func (o *Queue) HasDtmfHangupCallerEnabled() bool`

HasDtmfHangupCallerEnabled returns a boolean if a field has been set.

### GetDtmfRecordCalleeEnabled

`func (o *Queue) GetDtmfRecordCalleeEnabled() bool`

GetDtmfRecordCalleeEnabled returns the DtmfRecordCalleeEnabled field if non-nil, zero value otherwise.

### GetDtmfRecordCalleeEnabledOk

`func (o *Queue) GetDtmfRecordCalleeEnabledOk() (*bool, bool)`

GetDtmfRecordCalleeEnabledOk returns a tuple with the DtmfRecordCalleeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfRecordCalleeEnabled

`func (o *Queue) SetDtmfRecordCalleeEnabled(v bool)`

SetDtmfRecordCalleeEnabled sets DtmfRecordCalleeEnabled field to given value.

### HasDtmfRecordCalleeEnabled

`func (o *Queue) HasDtmfRecordCalleeEnabled() bool`

HasDtmfRecordCalleeEnabled returns a boolean if a field has been set.

### GetDtmfRecordCallerEnabled

`func (o *Queue) GetDtmfRecordCallerEnabled() bool`

GetDtmfRecordCallerEnabled returns the DtmfRecordCallerEnabled field if non-nil, zero value otherwise.

### GetDtmfRecordCallerEnabledOk

`func (o *Queue) GetDtmfRecordCallerEnabledOk() (*bool, bool)`

GetDtmfRecordCallerEnabledOk returns a tuple with the DtmfRecordCallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfRecordCallerEnabled

`func (o *Queue) SetDtmfRecordCallerEnabled(v bool)`

SetDtmfRecordCallerEnabled sets DtmfRecordCallerEnabled field to given value.

### HasDtmfRecordCallerEnabled

`func (o *Queue) HasDtmfRecordCallerEnabled() bool`

HasDtmfRecordCallerEnabled returns a boolean if a field has been set.

### GetDtmfTransferCalleeEnabled

`func (o *Queue) GetDtmfTransferCalleeEnabled() bool`

GetDtmfTransferCalleeEnabled returns the DtmfTransferCalleeEnabled field if non-nil, zero value otherwise.

### GetDtmfTransferCalleeEnabledOk

`func (o *Queue) GetDtmfTransferCalleeEnabledOk() (*bool, bool)`

GetDtmfTransferCalleeEnabledOk returns a tuple with the DtmfTransferCalleeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfTransferCalleeEnabled

`func (o *Queue) SetDtmfTransferCalleeEnabled(v bool)`

SetDtmfTransferCalleeEnabled sets DtmfTransferCalleeEnabled field to given value.

### HasDtmfTransferCalleeEnabled

`func (o *Queue) HasDtmfTransferCalleeEnabled() bool`

HasDtmfTransferCalleeEnabled returns a boolean if a field has been set.

### GetDtmfTransferCallerEnabled

`func (o *Queue) GetDtmfTransferCallerEnabled() bool`

GetDtmfTransferCallerEnabled returns the DtmfTransferCallerEnabled field if non-nil, zero value otherwise.

### GetDtmfTransferCallerEnabledOk

`func (o *Queue) GetDtmfTransferCallerEnabledOk() (*bool, bool)`

GetDtmfTransferCallerEnabledOk returns a tuple with the DtmfTransferCallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfTransferCallerEnabled

`func (o *Queue) SetDtmfTransferCallerEnabled(v bool)`

SetDtmfTransferCallerEnabled sets DtmfTransferCallerEnabled field to given value.

### HasDtmfTransferCallerEnabled

`func (o *Queue) HasDtmfTransferCallerEnabled() bool`

HasDtmfTransferCallerEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *Queue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Queue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Queue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Queue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIgnoreForward

`func (o *Queue) GetIgnoreForward() bool`

GetIgnoreForward returns the IgnoreForward field if non-nil, zero value otherwise.

### GetIgnoreForwardOk

`func (o *Queue) GetIgnoreForwardOk() (*bool, bool)`

GetIgnoreForwardOk returns a tuple with the IgnoreForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreForward

`func (o *Queue) SetIgnoreForward(v bool)`

SetIgnoreForward sets IgnoreForward field to given value.

### HasIgnoreForward

`func (o *Queue) HasIgnoreForward() bool`

HasIgnoreForward returns a boolean if a field has been set.

### GetMarkAnsweredElsewhere

`func (o *Queue) GetMarkAnsweredElsewhere() bool`

GetMarkAnsweredElsewhere returns the MarkAnsweredElsewhere field if non-nil, zero value otherwise.

### GetMarkAnsweredElsewhereOk

`func (o *Queue) GetMarkAnsweredElsewhereOk() (*bool, bool)`

GetMarkAnsweredElsewhereOk returns a tuple with the MarkAnsweredElsewhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAnsweredElsewhere

`func (o *Queue) SetMarkAnsweredElsewhere(v bool)`

SetMarkAnsweredElsewhere sets MarkAnsweredElsewhere field to given value.

### HasMarkAnsweredElsewhere

`func (o *Queue) HasMarkAnsweredElsewhere() bool`

HasMarkAnsweredElsewhere returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *Queue) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *Queue) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *Queue) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *Queue) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetOptions

`func (o *Queue) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Queue) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Queue) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Queue) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Queue) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Queue) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Queue) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Queue) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRetryOnTimeout

`func (o *Queue) GetRetryOnTimeout() bool`

GetRetryOnTimeout returns the RetryOnTimeout field if non-nil, zero value otherwise.

### GetRetryOnTimeoutOk

`func (o *Queue) GetRetryOnTimeoutOk() (*bool, bool)`

GetRetryOnTimeoutOk returns a tuple with the RetryOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOnTimeout

`func (o *Queue) SetRetryOnTimeout(v bool)`

SetRetryOnTimeout sets RetryOnTimeout field to given value.

### HasRetryOnTimeout

`func (o *Queue) HasRetryOnTimeout() bool`

HasRetryOnTimeout returns a boolean if a field has been set.

### GetRingOnHold

`func (o *Queue) GetRingOnHold() bool`

GetRingOnHold returns the RingOnHold field if non-nil, zero value otherwise.

### GetRingOnHoldOk

`func (o *Queue) GetRingOnHoldOk() (*bool, bool)`

GetRingOnHoldOk returns a tuple with the RingOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingOnHold

`func (o *Queue) SetRingOnHold(v bool)`

SetRingOnHold sets RingOnHold field to given value.

### HasRingOnHold

`func (o *Queue) HasRingOnHold() bool`

HasRingOnHold returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Queue) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Queue) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Queue) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Queue) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *Queue) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Queue) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Queue) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Queue) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWaitRatioDestination

`func (o *Queue) GetWaitRatioDestination() DestinationType`

GetWaitRatioDestination returns the WaitRatioDestination field if non-nil, zero value otherwise.

### GetWaitRatioDestinationOk

`func (o *Queue) GetWaitRatioDestinationOk() (*DestinationType, bool)`

GetWaitRatioDestinationOk returns a tuple with the WaitRatioDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitRatioDestination

`func (o *Queue) SetWaitRatioDestination(v DestinationType)`

SetWaitRatioDestination sets WaitRatioDestination field to given value.

### HasWaitRatioDestination

`func (o *Queue) HasWaitRatioDestination() bool`

HasWaitRatioDestination returns a boolean if a field has been set.

### GetWaitRatioThreshold

`func (o *Queue) GetWaitRatioThreshold() int32`

GetWaitRatioThreshold returns the WaitRatioThreshold field if non-nil, zero value otherwise.

### GetWaitRatioThresholdOk

`func (o *Queue) GetWaitRatioThresholdOk() (*int32, bool)`

GetWaitRatioThresholdOk returns a tuple with the WaitRatioThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitRatioThreshold

`func (o *Queue) SetWaitRatioThreshold(v int32)`

SetWaitRatioThreshold sets WaitRatioThreshold field to given value.

### HasWaitRatioThreshold

`func (o *Queue) HasWaitRatioThreshold() bool`

HasWaitRatioThreshold returns a boolean if a field has been set.

### GetWaitTimeDestination

`func (o *Queue) GetWaitTimeDestination() DestinationType`

GetWaitTimeDestination returns the WaitTimeDestination field if non-nil, zero value otherwise.

### GetWaitTimeDestinationOk

`func (o *Queue) GetWaitTimeDestinationOk() (*DestinationType, bool)`

GetWaitTimeDestinationOk returns a tuple with the WaitTimeDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeDestination

`func (o *Queue) SetWaitTimeDestination(v DestinationType)`

SetWaitTimeDestination sets WaitTimeDestination field to given value.

### HasWaitTimeDestination

`func (o *Queue) HasWaitTimeDestination() bool`

HasWaitTimeDestination returns a boolean if a field has been set.

### GetWaitTimeThreshold

`func (o *Queue) GetWaitTimeThreshold() int32`

GetWaitTimeThreshold returns the WaitTimeThreshold field if non-nil, zero value otherwise.

### GetWaitTimeThresholdOk

`func (o *Queue) GetWaitTimeThresholdOk() (*int32, bool)`

GetWaitTimeThresholdOk returns a tuple with the WaitTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeThreshold

`func (o *Queue) SetWaitTimeThreshold(v int32)`

SetWaitTimeThreshold sets WaitTimeThreshold field to given value.

### HasWaitTimeThreshold

`func (o *Queue) HasWaitTimeThreshold() bool`

HasWaitTimeThreshold returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
