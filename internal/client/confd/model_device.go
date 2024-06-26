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

// checks if the Device type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Device{}

// Device struct for Device
type Device struct {
	Description *string `json:"description,omitempty"`
	// Device ID
	Id *string `json:"id,omitempty"`
	// IP address (10.0.0.0)
	Ip *string `json:"ip,omitempty"`
	// Indicates if the device is a new device, ie in the master tenant
	IsNew *bool `json:"is_new,omitempty"`
	// MAC address (aa:bb:cc:dd:ee:ff)
	Mac *string `json:"mac,omitempty"`
	// Device model
	Model   *string        `json:"model,omitempty"`
	Options *DeviceOptions `json:"options,omitempty"`
	// Provisioning plugin to be used by device
	Plugin *string `json:"plugin,omitempty"`
	// Serial number
	Sn *string `json:"sn,omitempty"`
	// Device status. autoprov: Device can be provisionned using a provisioning code, configured: Device is configured and ready to be used, not_configured : Device has not been completely configured
	Status *string `json:"status,omitempty"`
	// ID of the device template. All device using a device template will have a certain number of common parameters preconfigured for the device
	TemplateId *string `json:"template_id,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// Vendor name
	Vendor *string `json:"vendor,omitempty"`
	// Firmware version
	Version *string `json:"version,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Device) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Device) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Device) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Device) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Device) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Device) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Device) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Device) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Device) SetIp(v string) {
	o.Ip = &v
}

// GetIsNew returns the IsNew field value if set, zero value otherwise.
func (o *Device) GetIsNew() bool {
	if o == nil || IsNil(o.IsNew) {
		var ret bool
		return ret
	}
	return *o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIsNewOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNew) {
		return nil, false
	}
	return o.IsNew, true
}

// HasIsNew returns a boolean if a field has been set.
func (o *Device) HasIsNew() bool {
	if o != nil && !IsNil(o.IsNew) {
		return true
	}

	return false
}

// SetIsNew gets a reference to the given bool and assigns it to the IsNew field.
func (o *Device) SetIsNew(v bool) {
	o.IsNew = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *Device) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *Device) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *Device) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *Device) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *Device) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *Device) SetModel(v string) {
	o.Model = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Device) GetOptions() DeviceOptions {
	if o == nil || IsNil(o.Options) {
		var ret DeviceOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOptionsOk() (*DeviceOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Device) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DeviceOptions and assigns it to the Options field.
func (o *Device) SetOptions(v DeviceOptions) {
	o.Options = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *Device) GetPlugin() string {
	if o == nil || IsNil(o.Plugin) {
		var ret string
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetPluginOk() (*string, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *Device) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given string and assigns it to the Plugin field.
func (o *Device) SetPlugin(v string) {
	o.Plugin = &v
}

// GetSn returns the Sn field value if set, zero value otherwise.
func (o *Device) GetSn() string {
	if o == nil || IsNil(o.Sn) {
		var ret string
		return ret
	}
	return *o.Sn
}

// GetSnOk returns a tuple with the Sn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetSnOk() (*string, bool) {
	if o == nil || IsNil(o.Sn) {
		return nil, false
	}
	return o.Sn, true
}

// HasSn returns a boolean if a field has been set.
func (o *Device) HasSn() bool {
	if o != nil && !IsNil(o.Sn) {
		return true
	}

	return false
}

// SetSn gets a reference to the given string and assigns it to the Sn field.
func (o *Device) SetSn(v string) {
	o.Sn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Device) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Device) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Device) SetStatus(v string) {
	o.Status = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *Device) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *Device) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *Device) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *Device) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *Device) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *Device) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *Device) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *Device) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *Device) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Device) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Device) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Device) SetVersion(v string) {
	o.Version = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Device) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.IsNew) {
		toSerialize["is_new"] = o.IsNew
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Sn) {
		toSerialize["sn"] = o.Sn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
