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

// checks if the BaseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseUser{}

// BaseUser struct for BaseUser
type BaseUser struct {
	// User firstname
	Firstname *string `json:"firstname,omitempty"`
	// User lastname
	Lastname *string `json:"lastname,omitempty"`
	// User UUID. This ID is globally unique across multiple Accent instances
	Uuid *string `json:"uuid,omitempty"`
	// Overwrite all passwords set in call permissions associated to the user
	CallPermissionPassword *string `json:"call_permission_password,omitempty"`
	// Record all calls made by this user (deprecated). If set, all `call_record_*_enabled` will be affected
	CallRecordEnabled *bool `json:"call_record_enabled,omitempty"`
	// Record all external calls received by this user.
	CallRecordIncomingExternalEnabled *bool `json:"call_record_incoming_external_enabled,omitempty"`
	// Record all internal calls received by this user.
	CallRecordIncomingInternalEnabled *bool `json:"call_record_incoming_internal_enabled,omitempty"`
	// Record all external calls made by this user
	CallRecordOutgoingExternalEnabled *bool `json:"call_record_outgoing_external_enabled,omitempty"`
	// Record all internal calls received by this user
	CallRecordOutgoingInternalEnabled *bool `json:"call_record_outgoing_internal_enabled,omitempty"`
	// Authorize call transfers
	CallTransferEnabled *bool `json:"call_transfer_enabled,omitempty"`
	// Name that appears on the phone when calling. Formatted as \"Firstname Lastname\" < number >
	CallerId *string `json:"caller_id,omitempty"`
	// Additional information about the user
	Description *string `json:"description,omitempty"`
	// Authorize to hangup with DTMF
	DtmfHangupEnabled *bool `json:"dtmf_hangup_enabled,omitempty"`
	// User's email. Used for directories (i.e. by accent-dird)
	Email *string `json:"email,omitempty"`
	// Enable/Disable the user
	Enabled *bool `json:"enabled,omitempty"`
	// User ID
	Id *int32 `json:"id,omitempty"`
	// User language (e.g. \"en_US\")
	Language *string `json:"language,omitempty"`
	// Phone number for the user’s mobile device
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty"`
	// Name of the MOH category to use for music on hold
	MusicOnHold *string `json:"music_on_hold,omitempty"`
	// Allow user to record a call by pressing *3
	OnlineCallRecordEnabled *bool `json:"online_call_record_enabled,omitempty"`
	// Name that appears on the phone when calling
	OutgoingCallerId *string `json:"outgoing_caller_id,omitempty"`
	// Password for connecting to the CTI (deprecated)
	Password *string `json:"password,omitempty"`
	// Name of the subroutine to execute in asterisk before receiving a call
	PreprocessSubroutine *string `json:"preprocess_subroutine,omitempty"`
	// Number of seconds a user's phone will ring before falling back
	RingSeconds *int32 `json:"ring_seconds,omitempty"`
	// Number of allowed simultaneous calls on a user's phone
	SimultaneousCalls *int32 `json:"simultaneous_calls,omitempty"`
	// Activate presence sharing in the accent client
	SupervisionEnabled *bool `json:"supervision_enabled,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// User timezone
	Timezone *string `json:"timezone,omitempty"`
	// A custom field which purpose is left to the client. If the user has no userfield, then this field is an empty string.
	Userfield *string `json:"userfield,omitempty"`
	// Username for connecting to the CTI (deprecated)
	Username *string `json:"username,omitempty"`
}

// NewBaseUser instantiates a new BaseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseUser() *BaseUser {
	this := BaseUser{}
	var callRecordEnabled bool = false
	this.CallRecordEnabled = &callRecordEnabled
	var callRecordIncomingExternalEnabled bool = false
	this.CallRecordIncomingExternalEnabled = &callRecordIncomingExternalEnabled
	var callRecordIncomingInternalEnabled bool = false
	this.CallRecordIncomingInternalEnabled = &callRecordIncomingInternalEnabled
	var callRecordOutgoingExternalEnabled bool = false
	this.CallRecordOutgoingExternalEnabled = &callRecordOutgoingExternalEnabled
	var callRecordOutgoingInternalEnabled bool = false
	this.CallRecordOutgoingInternalEnabled = &callRecordOutgoingInternalEnabled
	var callTransferEnabled bool = false
	this.CallTransferEnabled = &callTransferEnabled
	var dtmfHangupEnabled bool = false
	this.DtmfHangupEnabled = &dtmfHangupEnabled
	var enabled bool = true
	this.Enabled = &enabled
	var onlineCallRecordEnabled bool = false
	this.OnlineCallRecordEnabled = &onlineCallRecordEnabled
	var supervisionEnabled bool = true
	this.SupervisionEnabled = &supervisionEnabled
	return &this
}

// NewBaseUserWithDefaults instantiates a new BaseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseUserWithDefaults() *BaseUser {
	this := BaseUser{}
	var callRecordEnabled bool = false
	this.CallRecordEnabled = &callRecordEnabled
	var callRecordIncomingExternalEnabled bool = false
	this.CallRecordIncomingExternalEnabled = &callRecordIncomingExternalEnabled
	var callRecordIncomingInternalEnabled bool = false
	this.CallRecordIncomingInternalEnabled = &callRecordIncomingInternalEnabled
	var callRecordOutgoingExternalEnabled bool = false
	this.CallRecordOutgoingExternalEnabled = &callRecordOutgoingExternalEnabled
	var callRecordOutgoingInternalEnabled bool = false
	this.CallRecordOutgoingInternalEnabled = &callRecordOutgoingInternalEnabled
	var callTransferEnabled bool = false
	this.CallTransferEnabled = &callTransferEnabled
	var dtmfHangupEnabled bool = false
	this.DtmfHangupEnabled = &dtmfHangupEnabled
	var enabled bool = true
	this.Enabled = &enabled
	var onlineCallRecordEnabled bool = false
	this.OnlineCallRecordEnabled = &onlineCallRecordEnabled
	var supervisionEnabled bool = true
	this.SupervisionEnabled = &supervisionEnabled
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *BaseUser) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *BaseUser) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *BaseUser) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *BaseUser) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *BaseUser) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *BaseUser) SetLastname(v string) {
	o.Lastname = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *BaseUser) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *BaseUser) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *BaseUser) SetUuid(v string) {
	o.Uuid = &v
}

// GetCallPermissionPassword returns the CallPermissionPassword field value if set, zero value otherwise.
func (o *BaseUser) GetCallPermissionPassword() string {
	if o == nil || IsNil(o.CallPermissionPassword) {
		var ret string
		return ret
	}
	return *o.CallPermissionPassword
}

// GetCallPermissionPasswordOk returns a tuple with the CallPermissionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallPermissionPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CallPermissionPassword) {
		return nil, false
	}
	return o.CallPermissionPassword, true
}

// HasCallPermissionPassword returns a boolean if a field has been set.
func (o *BaseUser) HasCallPermissionPassword() bool {
	if o != nil && !IsNil(o.CallPermissionPassword) {
		return true
	}

	return false
}

// SetCallPermissionPassword gets a reference to the given string and assigns it to the CallPermissionPassword field.
func (o *BaseUser) SetCallPermissionPassword(v string) {
	o.CallPermissionPassword = &v
}

// GetCallRecordEnabled returns the CallRecordEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallRecordEnabled() bool {
	if o == nil || IsNil(o.CallRecordEnabled) {
		var ret bool
		return ret
	}
	return *o.CallRecordEnabled
}

// GetCallRecordEnabledOk returns a tuple with the CallRecordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallRecordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecordEnabled) {
		return nil, false
	}
	return o.CallRecordEnabled, true
}

// HasCallRecordEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallRecordEnabled() bool {
	if o != nil && !IsNil(o.CallRecordEnabled) {
		return true
	}

	return false
}

// SetCallRecordEnabled gets a reference to the given bool and assigns it to the CallRecordEnabled field.
func (o *BaseUser) SetCallRecordEnabled(v bool) {
	o.CallRecordEnabled = &v
}

// GetCallRecordIncomingExternalEnabled returns the CallRecordIncomingExternalEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallRecordIncomingExternalEnabled() bool {
	if o == nil || IsNil(o.CallRecordIncomingExternalEnabled) {
		var ret bool
		return ret
	}
	return *o.CallRecordIncomingExternalEnabled
}

// GetCallRecordIncomingExternalEnabledOk returns a tuple with the CallRecordIncomingExternalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallRecordIncomingExternalEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecordIncomingExternalEnabled) {
		return nil, false
	}
	return o.CallRecordIncomingExternalEnabled, true
}

// HasCallRecordIncomingExternalEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallRecordIncomingExternalEnabled() bool {
	if o != nil && !IsNil(o.CallRecordIncomingExternalEnabled) {
		return true
	}

	return false
}

// SetCallRecordIncomingExternalEnabled gets a reference to the given bool and assigns it to the CallRecordIncomingExternalEnabled field.
func (o *BaseUser) SetCallRecordIncomingExternalEnabled(v bool) {
	o.CallRecordIncomingExternalEnabled = &v
}

// GetCallRecordIncomingInternalEnabled returns the CallRecordIncomingInternalEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallRecordIncomingInternalEnabled() bool {
	if o == nil || IsNil(o.CallRecordIncomingInternalEnabled) {
		var ret bool
		return ret
	}
	return *o.CallRecordIncomingInternalEnabled
}

// GetCallRecordIncomingInternalEnabledOk returns a tuple with the CallRecordIncomingInternalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallRecordIncomingInternalEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecordIncomingInternalEnabled) {
		return nil, false
	}
	return o.CallRecordIncomingInternalEnabled, true
}

// HasCallRecordIncomingInternalEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallRecordIncomingInternalEnabled() bool {
	if o != nil && !IsNil(o.CallRecordIncomingInternalEnabled) {
		return true
	}

	return false
}

// SetCallRecordIncomingInternalEnabled gets a reference to the given bool and assigns it to the CallRecordIncomingInternalEnabled field.
func (o *BaseUser) SetCallRecordIncomingInternalEnabled(v bool) {
	o.CallRecordIncomingInternalEnabled = &v
}

// GetCallRecordOutgoingExternalEnabled returns the CallRecordOutgoingExternalEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallRecordOutgoingExternalEnabled() bool {
	if o == nil || IsNil(o.CallRecordOutgoingExternalEnabled) {
		var ret bool
		return ret
	}
	return *o.CallRecordOutgoingExternalEnabled
}

// GetCallRecordOutgoingExternalEnabledOk returns a tuple with the CallRecordOutgoingExternalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallRecordOutgoingExternalEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecordOutgoingExternalEnabled) {
		return nil, false
	}
	return o.CallRecordOutgoingExternalEnabled, true
}

// HasCallRecordOutgoingExternalEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallRecordOutgoingExternalEnabled() bool {
	if o != nil && !IsNil(o.CallRecordOutgoingExternalEnabled) {
		return true
	}

	return false
}

// SetCallRecordOutgoingExternalEnabled gets a reference to the given bool and assigns it to the CallRecordOutgoingExternalEnabled field.
func (o *BaseUser) SetCallRecordOutgoingExternalEnabled(v bool) {
	o.CallRecordOutgoingExternalEnabled = &v
}

// GetCallRecordOutgoingInternalEnabled returns the CallRecordOutgoingInternalEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallRecordOutgoingInternalEnabled() bool {
	if o == nil || IsNil(o.CallRecordOutgoingInternalEnabled) {
		var ret bool
		return ret
	}
	return *o.CallRecordOutgoingInternalEnabled
}

// GetCallRecordOutgoingInternalEnabledOk returns a tuple with the CallRecordOutgoingInternalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallRecordOutgoingInternalEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecordOutgoingInternalEnabled) {
		return nil, false
	}
	return o.CallRecordOutgoingInternalEnabled, true
}

// HasCallRecordOutgoingInternalEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallRecordOutgoingInternalEnabled() bool {
	if o != nil && !IsNil(o.CallRecordOutgoingInternalEnabled) {
		return true
	}

	return false
}

// SetCallRecordOutgoingInternalEnabled gets a reference to the given bool and assigns it to the CallRecordOutgoingInternalEnabled field.
func (o *BaseUser) SetCallRecordOutgoingInternalEnabled(v bool) {
	o.CallRecordOutgoingInternalEnabled = &v
}

// GetCallTransferEnabled returns the CallTransferEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetCallTransferEnabled() bool {
	if o == nil || IsNil(o.CallTransferEnabled) {
		var ret bool
		return ret
	}
	return *o.CallTransferEnabled
}

// GetCallTransferEnabledOk returns a tuple with the CallTransferEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallTransferEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallTransferEnabled) {
		return nil, false
	}
	return o.CallTransferEnabled, true
}

// HasCallTransferEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasCallTransferEnabled() bool {
	if o != nil && !IsNil(o.CallTransferEnabled) {
		return true
	}

	return false
}

// SetCallTransferEnabled gets a reference to the given bool and assigns it to the CallTransferEnabled field.
func (o *BaseUser) SetCallTransferEnabled(v bool) {
	o.CallTransferEnabled = &v
}

// GetCallerId returns the CallerId field value if set, zero value otherwise.
func (o *BaseUser) GetCallerId() string {
	if o == nil || IsNil(o.CallerId) {
		var ret string
		return ret
	}
	return *o.CallerId
}

// GetCallerIdOk returns a tuple with the CallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetCallerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallerId) {
		return nil, false
	}
	return o.CallerId, true
}

// HasCallerId returns a boolean if a field has been set.
func (o *BaseUser) HasCallerId() bool {
	if o != nil && !IsNil(o.CallerId) {
		return true
	}

	return false
}

// SetCallerId gets a reference to the given string and assigns it to the CallerId field.
func (o *BaseUser) SetCallerId(v string) {
	o.CallerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseUser) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseUser) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseUser) SetDescription(v string) {
	o.Description = &v
}

// GetDtmfHangupEnabled returns the DtmfHangupEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetDtmfHangupEnabled() bool {
	if o == nil || IsNil(o.DtmfHangupEnabled) {
		var ret bool
		return ret
	}
	return *o.DtmfHangupEnabled
}

// GetDtmfHangupEnabledOk returns a tuple with the DtmfHangupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetDtmfHangupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DtmfHangupEnabled) {
		return nil, false
	}
	return o.DtmfHangupEnabled, true
}

// HasDtmfHangupEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasDtmfHangupEnabled() bool {
	if o != nil && !IsNil(o.DtmfHangupEnabled) {
		return true
	}

	return false
}

// SetDtmfHangupEnabled gets a reference to the given bool and assigns it to the DtmfHangupEnabled field.
func (o *BaseUser) SetDtmfHangupEnabled(v bool) {
	o.DtmfHangupEnabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BaseUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BaseUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BaseUser) SetEmail(v string) {
	o.Email = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BaseUser) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BaseUser) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseUser) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BaseUser) SetId(v int32) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *BaseUser) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *BaseUser) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *BaseUser) SetLanguage(v string) {
	o.Language = &v
}

// GetMobilePhoneNumber returns the MobilePhoneNumber field value if set, zero value otherwise.
func (o *BaseUser) GetMobilePhoneNumber() string {
	if o == nil || IsNil(o.MobilePhoneNumber) {
		var ret string
		return ret
	}
	return *o.MobilePhoneNumber
}

// GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetMobilePhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.MobilePhoneNumber) {
		return nil, false
	}
	return o.MobilePhoneNumber, true
}

// HasMobilePhoneNumber returns a boolean if a field has been set.
func (o *BaseUser) HasMobilePhoneNumber() bool {
	if o != nil && !IsNil(o.MobilePhoneNumber) {
		return true
	}

	return false
}

// SetMobilePhoneNumber gets a reference to the given string and assigns it to the MobilePhoneNumber field.
func (o *BaseUser) SetMobilePhoneNumber(v string) {
	o.MobilePhoneNumber = &v
}

// GetMusicOnHold returns the MusicOnHold field value if set, zero value otherwise.
func (o *BaseUser) GetMusicOnHold() string {
	if o == nil || IsNil(o.MusicOnHold) {
		var ret string
		return ret
	}
	return *o.MusicOnHold
}

// GetMusicOnHoldOk returns a tuple with the MusicOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetMusicOnHoldOk() (*string, bool) {
	if o == nil || IsNil(o.MusicOnHold) {
		return nil, false
	}
	return o.MusicOnHold, true
}

// HasMusicOnHold returns a boolean if a field has been set.
func (o *BaseUser) HasMusicOnHold() bool {
	if o != nil && !IsNil(o.MusicOnHold) {
		return true
	}

	return false
}

// SetMusicOnHold gets a reference to the given string and assigns it to the MusicOnHold field.
func (o *BaseUser) SetMusicOnHold(v string) {
	o.MusicOnHold = &v
}

// GetOnlineCallRecordEnabled returns the OnlineCallRecordEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetOnlineCallRecordEnabled() bool {
	if o == nil || IsNil(o.OnlineCallRecordEnabled) {
		var ret bool
		return ret
	}
	return *o.OnlineCallRecordEnabled
}

// GetOnlineCallRecordEnabledOk returns a tuple with the OnlineCallRecordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetOnlineCallRecordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlineCallRecordEnabled) {
		return nil, false
	}
	return o.OnlineCallRecordEnabled, true
}

// HasOnlineCallRecordEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasOnlineCallRecordEnabled() bool {
	if o != nil && !IsNil(o.OnlineCallRecordEnabled) {
		return true
	}

	return false
}

// SetOnlineCallRecordEnabled gets a reference to the given bool and assigns it to the OnlineCallRecordEnabled field.
func (o *BaseUser) SetOnlineCallRecordEnabled(v bool) {
	o.OnlineCallRecordEnabled = &v
}

// GetOutgoingCallerId returns the OutgoingCallerId field value if set, zero value otherwise.
func (o *BaseUser) GetOutgoingCallerId() string {
	if o == nil || IsNil(o.OutgoingCallerId) {
		var ret string
		return ret
	}
	return *o.OutgoingCallerId
}

// GetOutgoingCallerIdOk returns a tuple with the OutgoingCallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetOutgoingCallerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutgoingCallerId) {
		return nil, false
	}
	return o.OutgoingCallerId, true
}

// HasOutgoingCallerId returns a boolean if a field has been set.
func (o *BaseUser) HasOutgoingCallerId() bool {
	if o != nil && !IsNil(o.OutgoingCallerId) {
		return true
	}

	return false
}

// SetOutgoingCallerId gets a reference to the given string and assigns it to the OutgoingCallerId field.
func (o *BaseUser) SetOutgoingCallerId(v string) {
	o.OutgoingCallerId = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BaseUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BaseUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BaseUser) SetPassword(v string) {
	o.Password = &v
}

// GetPreprocessSubroutine returns the PreprocessSubroutine field value if set, zero value otherwise.
func (o *BaseUser) GetPreprocessSubroutine() string {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		var ret string
		return ret
	}
	return *o.PreprocessSubroutine
}

// GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetPreprocessSubroutineOk() (*string, bool) {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		return nil, false
	}
	return o.PreprocessSubroutine, true
}

// HasPreprocessSubroutine returns a boolean if a field has been set.
func (o *BaseUser) HasPreprocessSubroutine() bool {
	if o != nil && !IsNil(o.PreprocessSubroutine) {
		return true
	}

	return false
}

// SetPreprocessSubroutine gets a reference to the given string and assigns it to the PreprocessSubroutine field.
func (o *BaseUser) SetPreprocessSubroutine(v string) {
	o.PreprocessSubroutine = &v
}

// GetRingSeconds returns the RingSeconds field value if set, zero value otherwise.
func (o *BaseUser) GetRingSeconds() int32 {
	if o == nil || IsNil(o.RingSeconds) {
		var ret int32
		return ret
	}
	return *o.RingSeconds
}

// GetRingSecondsOk returns a tuple with the RingSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetRingSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.RingSeconds) {
		return nil, false
	}
	return o.RingSeconds, true
}

// HasRingSeconds returns a boolean if a field has been set.
func (o *BaseUser) HasRingSeconds() bool {
	if o != nil && !IsNil(o.RingSeconds) {
		return true
	}

	return false
}

// SetRingSeconds gets a reference to the given int32 and assigns it to the RingSeconds field.
func (o *BaseUser) SetRingSeconds(v int32) {
	o.RingSeconds = &v
}

// GetSimultaneousCalls returns the SimultaneousCalls field value if set, zero value otherwise.
func (o *BaseUser) GetSimultaneousCalls() int32 {
	if o == nil || IsNil(o.SimultaneousCalls) {
		var ret int32
		return ret
	}
	return *o.SimultaneousCalls
}

// GetSimultaneousCallsOk returns a tuple with the SimultaneousCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetSimultaneousCallsOk() (*int32, bool) {
	if o == nil || IsNil(o.SimultaneousCalls) {
		return nil, false
	}
	return o.SimultaneousCalls, true
}

// HasSimultaneousCalls returns a boolean if a field has been set.
func (o *BaseUser) HasSimultaneousCalls() bool {
	if o != nil && !IsNil(o.SimultaneousCalls) {
		return true
	}

	return false
}

// SetSimultaneousCalls gets a reference to the given int32 and assigns it to the SimultaneousCalls field.
func (o *BaseUser) SetSimultaneousCalls(v int32) {
	o.SimultaneousCalls = &v
}

// GetSupervisionEnabled returns the SupervisionEnabled field value if set, zero value otherwise.
func (o *BaseUser) GetSupervisionEnabled() bool {
	if o == nil || IsNil(o.SupervisionEnabled) {
		var ret bool
		return ret
	}
	return *o.SupervisionEnabled
}

// GetSupervisionEnabledOk returns a tuple with the SupervisionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetSupervisionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SupervisionEnabled) {
		return nil, false
	}
	return o.SupervisionEnabled, true
}

// HasSupervisionEnabled returns a boolean if a field has been set.
func (o *BaseUser) HasSupervisionEnabled() bool {
	if o != nil && !IsNil(o.SupervisionEnabled) {
		return true
	}

	return false
}

// SetSupervisionEnabled gets a reference to the given bool and assigns it to the SupervisionEnabled field.
func (o *BaseUser) SetSupervisionEnabled(v bool) {
	o.SupervisionEnabled = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *BaseUser) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *BaseUser) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *BaseUser) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *BaseUser) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *BaseUser) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *BaseUser) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUserfield returns the Userfield field value if set, zero value otherwise.
func (o *BaseUser) GetUserfield() string {
	if o == nil || IsNil(o.Userfield) {
		var ret string
		return ret
	}
	return *o.Userfield
}

// GetUserfieldOk returns a tuple with the Userfield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetUserfieldOk() (*string, bool) {
	if o == nil || IsNil(o.Userfield) {
		return nil, false
	}
	return o.Userfield, true
}

// HasUserfield returns a boolean if a field has been set.
func (o *BaseUser) HasUserfield() bool {
	if o != nil && !IsNil(o.Userfield) {
		return true
	}

	return false
}

// SetUserfield gets a reference to the given string and assigns it to the Userfield field.
func (o *BaseUser) SetUserfield(v string) {
	o.Userfield = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *BaseUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *BaseUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *BaseUser) SetUsername(v string) {
	o.Username = &v
}

func (o BaseUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.CallPermissionPassword) {
		toSerialize["call_permission_password"] = o.CallPermissionPassword
	}
	if !IsNil(o.CallRecordEnabled) {
		toSerialize["call_record_enabled"] = o.CallRecordEnabled
	}
	if !IsNil(o.CallRecordIncomingExternalEnabled) {
		toSerialize["call_record_incoming_external_enabled"] = o.CallRecordIncomingExternalEnabled
	}
	if !IsNil(o.CallRecordIncomingInternalEnabled) {
		toSerialize["call_record_incoming_internal_enabled"] = o.CallRecordIncomingInternalEnabled
	}
	if !IsNil(o.CallRecordOutgoingExternalEnabled) {
		toSerialize["call_record_outgoing_external_enabled"] = o.CallRecordOutgoingExternalEnabled
	}
	if !IsNil(o.CallRecordOutgoingInternalEnabled) {
		toSerialize["call_record_outgoing_internal_enabled"] = o.CallRecordOutgoingInternalEnabled
	}
	if !IsNil(o.CallTransferEnabled) {
		toSerialize["call_transfer_enabled"] = o.CallTransferEnabled
	}
	if !IsNil(o.CallerId) {
		toSerialize["caller_id"] = o.CallerId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DtmfHangupEnabled) {
		toSerialize["dtmf_hangup_enabled"] = o.DtmfHangupEnabled
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.MobilePhoneNumber) {
		toSerialize["mobile_phone_number"] = o.MobilePhoneNumber
	}
	if !IsNil(o.MusicOnHold) {
		toSerialize["music_on_hold"] = o.MusicOnHold
	}
	if !IsNil(o.OnlineCallRecordEnabled) {
		toSerialize["online_call_record_enabled"] = o.OnlineCallRecordEnabled
	}
	if !IsNil(o.OutgoingCallerId) {
		toSerialize["outgoing_caller_id"] = o.OutgoingCallerId
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PreprocessSubroutine) {
		toSerialize["preprocess_subroutine"] = o.PreprocessSubroutine
	}
	if !IsNil(o.RingSeconds) {
		toSerialize["ring_seconds"] = o.RingSeconds
	}
	if !IsNil(o.SimultaneousCalls) {
		toSerialize["simultaneous_calls"] = o.SimultaneousCalls
	}
	if !IsNil(o.SupervisionEnabled) {
		toSerialize["supervision_enabled"] = o.SupervisionEnabled
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Userfield) {
		toSerialize["userfield"] = o.Userfield
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableBaseUser struct {
	value *BaseUser
	isSet bool
}

func (v NullableBaseUser) Get() *BaseUser {
	return v.value
}

func (v *NullableBaseUser) Set(val *BaseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseUser(val *BaseUser) *NullableBaseUser {
	return &NullableBaseUser{value: val, isSet: true}
}

func (v NullableBaseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
