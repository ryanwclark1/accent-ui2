/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"encoding/json"
)

// checks if the Queue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Queue{}

// Queue struct for Queue
type Queue struct {
	// The id of the queue
	Id *int32 `json:"id,omitempty"`
	// The label of the queue
	Label *string `json:"label,omitempty"`
	// The name of the queue. Cannot be modified
	Name      *string                    `json:"name,omitempty"`
	Schedules []ScheduleRelationBase     `json:"schedules,omitempty"`
	Members   *QueueRelationMemberAgents `json:"members,omitempty"`
	// Announce hold time on entry
	AnnounceHoldTimeOnEntry *bool `json:"announce_hold_time_on_entry,omitempty"`
	// How the caller_id_name will be treated
	CallerIdMode *string `json:"caller_id_mode,omitempty"`
	// Name to display
	CallerIdName *string `json:"caller_id_name,omitempty"`
	// Asterisk definition: data-quality (modem) call (minimum delay)
	DataQuality *bool `json:"data_quality,omitempty"`
	// Enable DTMF hangup by callee
	DtmfHangupCalleeEnabled *bool `json:"dtmf_hangup_callee_enabled,omitempty"`
	// Enable DTMF hangup by caller
	DtmfHangupCallerEnabled *bool `json:"dtmf_hangup_caller_enabled,omitempty"`
	// Enable DTMF records by callee
	DtmfRecordCalleeEnabled *bool `json:"dtmf_record_callee_enabled,omitempty"`
	// Enable DTMF records by caller
	DtmfRecordCallerEnabled *bool `json:"dtmf_record_caller_enabled,omitempty"`
	// Enable DTMF transfers by callee
	DtmfTransferCalleeEnabled *bool `json:"dtmf_transfer_callee_enabled,omitempty"`
	// Enable DTMF transfers by caller
	DtmfTransferCallerEnabled *bool `json:"dtmf_transfer_caller_enabled,omitempty"`
	// Enable/Disable the queue
	Enabled *bool `json:"enabled,omitempty"`
	// Ignore call forward requests from members
	IgnoreForward *bool `json:"ignore_forward,omitempty"`
	// Mark all calls as \"answered elsewhere\" when cancelled
	MarkAnsweredElsewhere *bool `json:"mark_answered_elsewhere,omitempty"`
	// Name of the MOH category to use for music on hold
	MusicOnHold *string `json:"music_on_hold,omitempty"`
	// Advanced configuration options. Options are appended at the end of a queue account in the file 'queues.conf' used by asterisk. Please consult the asterisk documentation for further details on available parameters. Because of database limitations, only the following options can be configured:  * announce * context * timeout * monitor-type * monitor-format * queue-youarenext * queue-thereare * queue-callswaiting * queue-holdtime * queue-minutes * queue-seconds * queue-thankyou * queue-reporthold * periodic-announce * announce-frequency * periodic-announce-frequency * announce-round-seconds * announce-holdtime * retry * wrapuptime * maxlen * servicelevel * strategy * joinempty * leavewhenempty * ringinuse * reportholdtime * memberdelay * weight * timeoutrestart * timeoutpriority * autofill * autopause * setinterfacevar * setqueueentryvar * setqueuevar * membermacro * min-announce-frequency * random-periodic-announce * announce-position * announce-position-limit * defaultrule  Options must have the following form: ``` {   \"options\": [      [\"name1\", \"value1\"],      [\"name2\", \"value2\"]   ] } ```  The resulting configuration in queues.conf will have the following form: ``` [queuename] name1=value1 name2=value2 ```
	Options              [][]string `json:"options,omitempty"`
	PreprocessSubroutine *string    `json:"preprocess_subroutine,omitempty"`
	// Retry when time has elapsed
	RetryOnTimeout *bool `json:"retry_on_timeout,omitempty"`
	// Ring instead of On-Hold Music
	RingOnHold *bool `json:"ring_on_hold,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// Number of seconds the queue will ring before falling back
	Timeout              *int32           `json:"timeout,omitempty"`
	WaitRatioDestination *DestinationType `json:"wait_ratio_destination,omitempty"`
	// Wait ratio (waiting calls per logged-in agent) threshold before fallback on 'wait_ratio_destination'. Set to 'null' to disable it.
	WaitRatioThreshold  *int32           `json:"wait_ratio_threshold,omitempty"`
	WaitTimeDestination *DestinationType `json:"wait_time_destination,omitempty"`
	// Wait time threshold before fallback on 'wait_time_destination'. Set to 'null' to disable it.
	WaitTimeThreshold *int32 `json:"wait_time_threshold,omitempty"`
}

// NewQueue instantiates a new Queue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueue() *Queue {
	this := Queue{}
	var announceHoldTimeOnEntry bool = false
	this.AnnounceHoldTimeOnEntry = &announceHoldTimeOnEntry
	var dataQuality bool = false
	this.DataQuality = &dataQuality
	var dtmfHangupCalleeEnabled bool = false
	this.DtmfHangupCalleeEnabled = &dtmfHangupCalleeEnabled
	var dtmfHangupCallerEnabled bool = false
	this.DtmfHangupCallerEnabled = &dtmfHangupCallerEnabled
	var dtmfRecordCalleeEnabled bool = false
	this.DtmfRecordCalleeEnabled = &dtmfRecordCalleeEnabled
	var dtmfRecordCallerEnabled bool = false
	this.DtmfRecordCallerEnabled = &dtmfRecordCallerEnabled
	var dtmfTransferCalleeEnabled bool = false
	this.DtmfTransferCalleeEnabled = &dtmfTransferCalleeEnabled
	var dtmfTransferCallerEnabled bool = false
	this.DtmfTransferCallerEnabled = &dtmfTransferCallerEnabled
	var enabled bool = true
	this.Enabled = &enabled
	var ignoreForward bool = false
	this.IgnoreForward = &ignoreForward
	var markAnsweredElsewhere bool = true
	this.MarkAnsweredElsewhere = &markAnsweredElsewhere
	var retryOnTimeout bool = true
	this.RetryOnTimeout = &retryOnTimeout
	var ringOnHold bool = false
	this.RingOnHold = &ringOnHold
	return &this
}

// NewQueueWithDefaults instantiates a new Queue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueWithDefaults() *Queue {
	this := Queue{}
	var announceHoldTimeOnEntry bool = false
	this.AnnounceHoldTimeOnEntry = &announceHoldTimeOnEntry
	var dataQuality bool = false
	this.DataQuality = &dataQuality
	var dtmfHangupCalleeEnabled bool = false
	this.DtmfHangupCalleeEnabled = &dtmfHangupCalleeEnabled
	var dtmfHangupCallerEnabled bool = false
	this.DtmfHangupCallerEnabled = &dtmfHangupCallerEnabled
	var dtmfRecordCalleeEnabled bool = false
	this.DtmfRecordCalleeEnabled = &dtmfRecordCalleeEnabled
	var dtmfRecordCallerEnabled bool = false
	this.DtmfRecordCallerEnabled = &dtmfRecordCallerEnabled
	var dtmfTransferCalleeEnabled bool = false
	this.DtmfTransferCalleeEnabled = &dtmfTransferCalleeEnabled
	var dtmfTransferCallerEnabled bool = false
	this.DtmfTransferCallerEnabled = &dtmfTransferCallerEnabled
	var enabled bool = true
	this.Enabled = &enabled
	var ignoreForward bool = false
	this.IgnoreForward = &ignoreForward
	var markAnsweredElsewhere bool = true
	this.MarkAnsweredElsewhere = &markAnsweredElsewhere
	var retryOnTimeout bool = true
	this.RetryOnTimeout = &retryOnTimeout
	var ringOnHold bool = false
	this.RingOnHold = &ringOnHold
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Queue) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Queue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Queue) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Queue) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Queue) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Queue) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Queue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Queue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Queue) SetName(v string) {
	o.Name = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *Queue) GetSchedules() []ScheduleRelationBase {
	if o == nil || IsNil(o.Schedules) {
		var ret []ScheduleRelationBase
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetSchedulesOk() ([]ScheduleRelationBase, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *Queue) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ScheduleRelationBase and assigns it to the Schedules field.
func (o *Queue) SetSchedules(v []ScheduleRelationBase) {
	o.Schedules = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Queue) GetMembers() QueueRelationMemberAgents {
	if o == nil || IsNil(o.Members) {
		var ret QueueRelationMemberAgents
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetMembersOk() (*QueueRelationMemberAgents, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Queue) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given QueueRelationMemberAgents and assigns it to the Members field.
func (o *Queue) SetMembers(v QueueRelationMemberAgents) {
	o.Members = &v
}

// GetAnnounceHoldTimeOnEntry returns the AnnounceHoldTimeOnEntry field value if set, zero value otherwise.
func (o *Queue) GetAnnounceHoldTimeOnEntry() bool {
	if o == nil || IsNil(o.AnnounceHoldTimeOnEntry) {
		var ret bool
		return ret
	}
	return *o.AnnounceHoldTimeOnEntry
}

// GetAnnounceHoldTimeOnEntryOk returns a tuple with the AnnounceHoldTimeOnEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetAnnounceHoldTimeOnEntryOk() (*bool, bool) {
	if o == nil || IsNil(o.AnnounceHoldTimeOnEntry) {
		return nil, false
	}
	return o.AnnounceHoldTimeOnEntry, true
}

// HasAnnounceHoldTimeOnEntry returns a boolean if a field has been set.
func (o *Queue) HasAnnounceHoldTimeOnEntry() bool {
	if o != nil && !IsNil(o.AnnounceHoldTimeOnEntry) {
		return true
	}

	return false
}

// SetAnnounceHoldTimeOnEntry gets a reference to the given bool and assigns it to the AnnounceHoldTimeOnEntry field.
func (o *Queue) SetAnnounceHoldTimeOnEntry(v bool) {
	o.AnnounceHoldTimeOnEntry = &v
}

// GetCallerIdMode returns the CallerIdMode field value if set, zero value otherwise.
func (o *Queue) GetCallerIdMode() string {
	if o == nil || IsNil(o.CallerIdMode) {
		var ret string
		return ret
	}
	return *o.CallerIdMode
}

// GetCallerIdModeOk returns a tuple with the CallerIdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetCallerIdModeOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdMode) {
		return nil, false
	}
	return o.CallerIdMode, true
}

// HasCallerIdMode returns a boolean if a field has been set.
func (o *Queue) HasCallerIdMode() bool {
	if o != nil && !IsNil(o.CallerIdMode) {
		return true
	}

	return false
}

// SetCallerIdMode gets a reference to the given string and assigns it to the CallerIdMode field.
func (o *Queue) SetCallerIdMode(v string) {
	o.CallerIdMode = &v
}

// GetCallerIdName returns the CallerIdName field value if set, zero value otherwise.
func (o *Queue) GetCallerIdName() string {
	if o == nil || IsNil(o.CallerIdName) {
		var ret string
		return ret
	}
	return *o.CallerIdName
}

// GetCallerIdNameOk returns a tuple with the CallerIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetCallerIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdName) {
		return nil, false
	}
	return o.CallerIdName, true
}

// HasCallerIdName returns a boolean if a field has been set.
func (o *Queue) HasCallerIdName() bool {
	if o != nil && !IsNil(o.CallerIdName) {
		return true
	}

	return false
}

// SetCallerIdName gets a reference to the given string and assigns it to the CallerIdName field.
func (o *Queue) SetCallerIdName(v string) {
	o.CallerIdName = &v
}

// GetDataQuality returns the DataQuality field value if set, zero value otherwise.
func (o *Queue) GetDataQuality() bool {
	if o == nil || IsNil(o.DataQuality) {
		var ret bool
		return ret
	}
	return *o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDataQualityOk() (*bool, bool) {
	if o == nil || IsNil(o.DataQuality) {
		return nil, false
	}
	return o.DataQuality, true
}

// HasDataQuality returns a boolean if a field has been set.
func (o *Queue) HasDataQuality() bool {
	if o != nil && !IsNil(o.DataQuality) {
		return true
	}

	return false
}

// SetDataQuality gets a reference to the given bool and assigns it to the DataQuality field.
func (o *Queue) SetDataQuality(v bool) {
	o.DataQuality = &v
}

// GetDtmfHangupCalleeEnabled returns the DtmfHangupCalleeEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfHangupCalleeEnabled() bool {
	if o == nil || IsNil(o.DtmfHangupCalleeEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfHangupCalleeEnabled
}

// GetDtmfHangupCalleeEnabledOk returns a tuple with the DtmfHangupCalleeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfHangupCalleeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfHangupCalleeEnabled) {
		return nil, false
	}
	return o.DtmfHangupCalleeEnabled, true
}

// HasDtmfHangupCalleeEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfHangupCalleeEnabled() bool {
	if o != nil && !IsNil(o.DtmfHangupCalleeEnabled) {
		return true
	}

	return false
}

// SetDtmfHangupCalleeEnabled gets a reference to the given bool and assigns it to the DtmfHangupCalleeEnabled field.
func (o *Queue) SetDtmfHangupCalleeEnabled(v bool) {
	o.DtmfHangupCalleeEnabled = &v
}

// GetDtmfHangupCallerEnabled returns the DtmfHangupCallerEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfHangupCallerEnabled() bool {
	if o == nil || IsNil(o.DtmfHangupCallerEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfHangupCallerEnabled
}

// GetDtmfHangupCallerEnabledOk returns a tuple with the DtmfHangupCallerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfHangupCallerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfHangupCallerEnabled) {
		return nil, false
	}
	return o.DtmfHangupCallerEnabled, true
}

// HasDtmfHangupCallerEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfHangupCallerEnabled() bool {
	if o != nil && !IsNil(o.DtmfHangupCallerEnabled) {
		return true
	}

	return false
}

// SetDtmfHangupCallerEnabled gets a reference to the given bool and assigns it to the DtmfHangupCallerEnabled field.
func (o *Queue) SetDtmfHangupCallerEnabled(v bool) {
	o.DtmfHangupCallerEnabled = &v
}

// GetDtmfRecordCalleeEnabled returns the DtmfRecordCalleeEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfRecordCalleeEnabled() bool {
	if o == nil || IsNil(o.DtmfRecordCalleeEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfRecordCalleeEnabled
}

// GetDtmfRecordCalleeEnabledOk returns a tuple with the DtmfRecordCalleeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfRecordCalleeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfRecordCalleeEnabled) {
		return nil, false
	}
	return o.DtmfRecordCalleeEnabled, true
}

// HasDtmfRecordCalleeEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfRecordCalleeEnabled() bool {
	if o != nil && !IsNil(o.DtmfRecordCalleeEnabled) {
		return true
	}

	return false
}

// SetDtmfRecordCalleeEnabled gets a reference to the given bool and assigns it to the DtmfRecordCalleeEnabled field.
func (o *Queue) SetDtmfRecordCalleeEnabled(v bool) {
	o.DtmfRecordCalleeEnabled = &v
}

// GetDtmfRecordCallerEnabled returns the DtmfRecordCallerEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfRecordCallerEnabled() bool {
	if o == nil || IsNil(o.DtmfRecordCallerEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfRecordCallerEnabled
}

// GetDtmfRecordCallerEnabledOk returns a tuple with the DtmfRecordCallerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfRecordCallerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfRecordCallerEnabled) {
		return nil, false
	}
	return o.DtmfRecordCallerEnabled, true
}

// HasDtmfRecordCallerEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfRecordCallerEnabled() bool {
	if o != nil && !IsNil(o.DtmfRecordCallerEnabled) {
		return true
	}

	return false
}

// SetDtmfRecordCallerEnabled gets a reference to the given bool and assigns it to the DtmfRecordCallerEnabled field.
func (o *Queue) SetDtmfRecordCallerEnabled(v bool) {
	o.DtmfRecordCallerEnabled = &v
}

// GetDtmfTransferCalleeEnabled returns the DtmfTransferCalleeEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfTransferCalleeEnabled() bool {
	if o == nil || IsNil(o.DtmfTransferCalleeEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfTransferCalleeEnabled
}

// GetDtmfTransferCalleeEnabledOk returns a tuple with the DtmfTransferCalleeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfTransferCalleeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfTransferCalleeEnabled) {
		return nil, false
	}
	return o.DtmfTransferCalleeEnabled, true
}

// HasDtmfTransferCalleeEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfTransferCalleeEnabled() bool {
	if o != nil && !IsNil(o.DtmfTransferCalleeEnabled) {
		return true
	}

	return false
}

// SetDtmfTransferCalleeEnabled gets a reference to the given bool and assigns it to the DtmfTransferCalleeEnabled field.
func (o *Queue) SetDtmfTransferCalleeEnabled(v bool) {
	o.DtmfTransferCalleeEnabled = &v
}

// GetDtmfTransferCallerEnabled returns the DtmfTransferCallerEnabled field value if set, zero value otherwise.
func (o *Queue) GetDtmfTransferCallerEnabled() bool {
	if o == nil || IsNil(o.DtmfTransferCallerEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfTransferCallerEnabled
}

// GetDtmfTransferCallerEnabledOk returns a tuple with the DtmfTransferCallerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDtmfTransferCallerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfTransferCallerEnabled) {
		return nil, false
	}
	return o.DtmfTransferCallerEnabled, true
}

// HasDtmfTransferCallerEnabled returns a boolean if a field has been set.
func (o *Queue) HasDtmfTransferCallerEnabled() bool {
	if o != nil && !IsNil(o.DtmfTransferCallerEnabled) {
		return true
	}

	return false
}

// SetDtmfTransferCallerEnabled gets a reference to the given bool and assigns it to the DtmfTransferCallerEnabled field.
func (o *Queue) SetDtmfTransferCallerEnabled(v bool) {
	o.DtmfTransferCallerEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Queue) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Queue) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Queue) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIgnoreForward returns the IgnoreForward field value if set, zero value otherwise.
func (o *Queue) GetIgnoreForward() bool {
	if o == nil || IsNil(o.IgnoreForward) {
		var ret bool
		return ret
	}
	return *o.IgnoreForward
}

// GetIgnoreForwardOk returns a tuple with the IgnoreForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetIgnoreForwardOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreForward) {
		return nil, false
	}
	return o.IgnoreForward, true
}

// HasIgnoreForward returns a boolean if a field has been set.
func (o *Queue) HasIgnoreForward() bool {
	if o != nil && !IsNil(o.IgnoreForward) {
		return true
	}

	return false
}

// SetIgnoreForward gets a reference to the given bool and assigns it to the IgnoreForward field.
func (o *Queue) SetIgnoreForward(v bool) {
	o.IgnoreForward = &v
}

// GetMarkAnsweredElsewhere returns the MarkAnsweredElsewhere field value if set, zero value otherwise.
func (o *Queue) GetMarkAnsweredElsewhere() bool {
	if o == nil || IsNil(o.MarkAnsweredElsewhere) {
		var ret bool
		return ret
	}
	return *o.MarkAnsweredElsewhere
}

// GetMarkAnsweredElsewhereOk returns a tuple with the MarkAnsweredElsewhere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetMarkAnsweredElsewhereOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkAnsweredElsewhere) {
		return nil, false
	}
	return o.MarkAnsweredElsewhere, true
}

// HasMarkAnsweredElsewhere returns a boolean if a field has been set.
func (o *Queue) HasMarkAnsweredElsewhere() bool {
	if o != nil && !IsNil(o.MarkAnsweredElsewhere) {
		return true
	}

	return false
}

// SetMarkAnsweredElsewhere gets a reference to the given bool and assigns it to the MarkAnsweredElsewhere field.
func (o *Queue) SetMarkAnsweredElsewhere(v bool) {
	o.MarkAnsweredElsewhere = &v
}

// GetMusicOnHold returns the MusicOnHold field value if set, zero value otherwise.
func (o *Queue) GetMusicOnHold() string {
	if o == nil || IsNil(o.MusicOnHold) {
		var ret string
		return ret
	}
	return *o.MusicOnHold
}

// GetMusicOnHoldOk returns a tuple with the MusicOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetMusicOnHoldOk() (*string, bool) {
	if o == nil || IsNil(o.MusicOnHold) {
		return nil, false
	}
	return o.MusicOnHold, true
}

// HasMusicOnHold returns a boolean if a field has been set.
func (o *Queue) HasMusicOnHold() bool {
	if o != nil && !IsNil(o.MusicOnHold) {
		return true
	}

	return false
}

// SetMusicOnHold gets a reference to the given string and assigns it to the MusicOnHold field.
func (o *Queue) SetMusicOnHold(v string) {
	o.MusicOnHold = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Queue) GetOptions() [][]string {
	if o == nil || IsNil(o.Options) {
		var ret [][]string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Queue) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given [][]string and assigns it to the Options field.
func (o *Queue) SetOptions(v [][]string) {
	o.Options = v
}

// GetPreprocessSubroutine returns the PreprocessSubroutine field value if set, zero value otherwise.
func (o *Queue) GetPreprocessSubroutine() string {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		var ret string
		return ret
	}
	return *o.PreprocessSubroutine
}

// GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetPreprocessSubroutineOk() (*string, bool) {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		return nil, false
	}
	return o.PreprocessSubroutine, true
}

// HasPreprocessSubroutine returns a boolean if a field has been set.
func (o *Queue) HasPreprocessSubroutine() bool {
	if o != nil && !IsNil(o.PreprocessSubroutine) {
		return true
	}

	return false
}

// SetPreprocessSubroutine gets a reference to the given string and assigns it to the PreprocessSubroutine field.
func (o *Queue) SetPreprocessSubroutine(v string) {
	o.PreprocessSubroutine = &v
}

// GetRetryOnTimeout returns the RetryOnTimeout field value if set, zero value otherwise.
func (o *Queue) GetRetryOnTimeout() bool {
	if o == nil || IsNil(o.RetryOnTimeout) {
		var ret bool
		return ret
	}
	return *o.RetryOnTimeout
}

// GetRetryOnTimeoutOk returns a tuple with the RetryOnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetRetryOnTimeoutOk() (*bool, bool) {
	if o == nil || IsNil(o.RetryOnTimeout) {
		return nil, false
	}
	return o.RetryOnTimeout, true
}

// HasRetryOnTimeout returns a boolean if a field has been set.
func (o *Queue) HasRetryOnTimeout() bool {
	if o != nil && !IsNil(o.RetryOnTimeout) {
		return true
	}

	return false
}

// SetRetryOnTimeout gets a reference to the given bool and assigns it to the RetryOnTimeout field.
func (o *Queue) SetRetryOnTimeout(v bool) {
	o.RetryOnTimeout = &v
}

// GetRingOnHold returns the RingOnHold field value if set, zero value otherwise.
func (o *Queue) GetRingOnHold() bool {
	if o == nil || IsNil(o.RingOnHold) {
		var ret bool
		return ret
	}
	return *o.RingOnHold
}

// GetRingOnHoldOk returns a tuple with the RingOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetRingOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.RingOnHold) {
		return nil, false
	}
	return o.RingOnHold, true
}

// HasRingOnHold returns a boolean if a field has been set.
func (o *Queue) HasRingOnHold() bool {
	if o != nil && !IsNil(o.RingOnHold) {
		return true
	}

	return false
}

// SetRingOnHold gets a reference to the given bool and assigns it to the RingOnHold field.
func (o *Queue) SetRingOnHold(v bool) {
	o.RingOnHold = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *Queue) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *Queue) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *Queue) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *Queue) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *Queue) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *Queue) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetWaitRatioDestination returns the WaitRatioDestination field value if set, zero value otherwise.
func (o *Queue) GetWaitRatioDestination() DestinationType {
	if o == nil || IsNil(o.WaitRatioDestination) {
		var ret DestinationType
		return ret
	}
	return *o.WaitRatioDestination
}

// GetWaitRatioDestinationOk returns a tuple with the WaitRatioDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetWaitRatioDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.WaitRatioDestination) {
		return nil, false
	}
	return o.WaitRatioDestination, true
}

// HasWaitRatioDestination returns a boolean if a field has been set.
func (o *Queue) HasWaitRatioDestination() bool {
	if o != nil && !IsNil(o.WaitRatioDestination) {
		return true
	}

	return false
}

// SetWaitRatioDestination gets a reference to the given DestinationType and assigns it to the WaitRatioDestination field.
func (o *Queue) SetWaitRatioDestination(v DestinationType) {
	o.WaitRatioDestination = &v
}

// GetWaitRatioThreshold returns the WaitRatioThreshold field value if set, zero value otherwise.
func (o *Queue) GetWaitRatioThreshold() int32 {
	if o == nil || IsNil(o.WaitRatioThreshold) {
		var ret int32
		return ret
	}
	return *o.WaitRatioThreshold
}

// GetWaitRatioThresholdOk returns a tuple with the WaitRatioThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetWaitRatioThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitRatioThreshold) {
		return nil, false
	}
	return o.WaitRatioThreshold, true
}

// HasWaitRatioThreshold returns a boolean if a field has been set.
func (o *Queue) HasWaitRatioThreshold() bool {
	if o != nil && !IsNil(o.WaitRatioThreshold) {
		return true
	}

	return false
}

// SetWaitRatioThreshold gets a reference to the given int32 and assigns it to the WaitRatioThreshold field.
func (o *Queue) SetWaitRatioThreshold(v int32) {
	o.WaitRatioThreshold = &v
}

// GetWaitTimeDestination returns the WaitTimeDestination field value if set, zero value otherwise.
func (o *Queue) GetWaitTimeDestination() DestinationType {
	if o == nil || IsNil(o.WaitTimeDestination) {
		var ret DestinationType
		return ret
	}
	return *o.WaitTimeDestination
}

// GetWaitTimeDestinationOk returns a tuple with the WaitTimeDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetWaitTimeDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.WaitTimeDestination) {
		return nil, false
	}
	return o.WaitTimeDestination, true
}

// HasWaitTimeDestination returns a boolean if a field has been set.
func (o *Queue) HasWaitTimeDestination() bool {
	if o != nil && !IsNil(o.WaitTimeDestination) {
		return true
	}

	return false
}

// SetWaitTimeDestination gets a reference to the given DestinationType and assigns it to the WaitTimeDestination field.
func (o *Queue) SetWaitTimeDestination(v DestinationType) {
	o.WaitTimeDestination = &v
}

// GetWaitTimeThreshold returns the WaitTimeThreshold field value if set, zero value otherwise.
func (o *Queue) GetWaitTimeThreshold() int32 {
	if o == nil || IsNil(o.WaitTimeThreshold) {
		var ret int32
		return ret
	}
	return *o.WaitTimeThreshold
}

// GetWaitTimeThresholdOk returns a tuple with the WaitTimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetWaitTimeThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTimeThreshold) {
		return nil, false
	}
	return o.WaitTimeThreshold, true
}

// HasWaitTimeThreshold returns a boolean if a field has been set.
func (o *Queue) HasWaitTimeThreshold() bool {
	if o != nil && !IsNil(o.WaitTimeThreshold) {
		return true
	}

	return false
}

// SetWaitTimeThreshold gets a reference to the given int32 and assigns it to the WaitTimeThreshold field.
func (o *Queue) SetWaitTimeThreshold(v int32) {
	o.WaitTimeThreshold = &v
}

func (o Queue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Queue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.AnnounceHoldTimeOnEntry) {
		toSerialize["announce_hold_time_on_entry"] = o.AnnounceHoldTimeOnEntry
	}
	if !IsNil(o.CallerIdMode) {
		toSerialize["caller_id_mode"] = o.CallerIdMode
	}
	if !IsNil(o.CallerIdName) {
		toSerialize["caller_id_name"] = o.CallerIdName
	}
	if !IsNil(o.DataQuality) {
		toSerialize["data_quality"] = o.DataQuality
	}
	if !IsNil(o.DtmfHangupCalleeEnabled) {
		toSerialize["dtmf_hangup_callee_enabled"] = o.DtmfHangupCalleeEnabled
	}
	if !IsNil(o.DtmfHangupCallerEnabled) {
		toSerialize["dtmf_hangup_caller_enabled"] = o.DtmfHangupCallerEnabled
	}
	if !IsNil(o.DtmfRecordCalleeEnabled) {
		toSerialize["dtmf_record_callee_enabled"] = o.DtmfRecordCalleeEnabled
	}
	if !IsNil(o.DtmfRecordCallerEnabled) {
		toSerialize["dtmf_record_caller_enabled"] = o.DtmfRecordCallerEnabled
	}
	if !IsNil(o.DtmfTransferCalleeEnabled) {
		toSerialize["dtmf_transfer_callee_enabled"] = o.DtmfTransferCalleeEnabled
	}
	if !IsNil(o.DtmfTransferCallerEnabled) {
		toSerialize["dtmf_transfer_caller_enabled"] = o.DtmfTransferCallerEnabled
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IgnoreForward) {
		toSerialize["ignore_forward"] = o.IgnoreForward
	}
	if !IsNil(o.MarkAnsweredElsewhere) {
		toSerialize["mark_answered_elsewhere"] = o.MarkAnsweredElsewhere
	}
	if !IsNil(o.MusicOnHold) {
		toSerialize["music_on_hold"] = o.MusicOnHold
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.PreprocessSubroutine) {
		toSerialize["preprocess_subroutine"] = o.PreprocessSubroutine
	}
	if !IsNil(o.RetryOnTimeout) {
		toSerialize["retry_on_timeout"] = o.RetryOnTimeout
	}
	if !IsNil(o.RingOnHold) {
		toSerialize["ring_on_hold"] = o.RingOnHold
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.WaitRatioDestination) {
		toSerialize["wait_ratio_destination"] = o.WaitRatioDestination
	}
	if !IsNil(o.WaitRatioThreshold) {
		toSerialize["wait_ratio_threshold"] = o.WaitRatioThreshold
	}
	if !IsNil(o.WaitTimeDestination) {
		toSerialize["wait_time_destination"] = o.WaitTimeDestination
	}
	if !IsNil(o.WaitTimeThreshold) {
		toSerialize["wait_time_threshold"] = o.WaitTimeThreshold
	}
	return toSerialize, nil
}

type NullableQueue struct {
	value *Queue
	isSet bool
}

func (v NullableQueue) Get() *Queue {
	return v.value
}

func (v *NullableQueue) Set(val *Queue) {
	v.value = val
	v.isSet = true
}

func (v NullableQueue) IsSet() bool {
	return v.isSet
}

func (v *NullableQueue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueue(val *Queue) *NullableQueue {
	return &NullableQueue{value: val, isSet: true}
}

func (v NullableQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
