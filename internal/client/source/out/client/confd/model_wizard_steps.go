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

// checks if the WizardSteps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardSteps{}

// WizardSteps struct for WizardSteps
type WizardSteps struct {
	// Create admin in accent-auth?
	Admin *bool `json:"admin,omitempty"`
	// Generate /etc/accent/common.conf
	Commonconf *bool `json:"commonconf,omitempty"`
	// Initialize Accent database?
	Database *bool `json:"database,omitempty"`
	// Modify /etc/hosts?
	ManageHostsFile *bool `json:"manage_hosts_file,omitempty"`
	// Modify /etc/resolv.conf?
	ManageResolvFile *bool `json:"manage_resolv_file,omitempty"`
	// Enable/start Accent services?
	ManageServices *bool `json:"manage_services,omitempty"`
	// Initialize accent-provd database?
	Provisioning *bool `json:"provisioning,omitempty"`
}

// NewWizardSteps instantiates a new WizardSteps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardSteps() *WizardSteps {
	this := WizardSteps{}
	var admin bool = true
	this.Admin = &admin
	var commonconf bool = true
	this.Commonconf = &commonconf
	var database bool = true
	this.Database = &database
	var manageHostsFile bool = true
	this.ManageHostsFile = &manageHostsFile
	var manageResolvFile bool = true
	this.ManageResolvFile = &manageResolvFile
	var manageServices bool = true
	this.ManageServices = &manageServices
	var provisioning bool = true
	this.Provisioning = &provisioning
	return &this
}

// NewWizardStepsWithDefaults instantiates a new WizardSteps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardStepsWithDefaults() *WizardSteps {
	this := WizardSteps{}
	var admin bool = true
	this.Admin = &admin
	var commonconf bool = true
	this.Commonconf = &commonconf
	var database bool = true
	this.Database = &database
	var manageHostsFile bool = true
	this.ManageHostsFile = &manageHostsFile
	var manageResolvFile bool = true
	this.ManageResolvFile = &manageResolvFile
	var manageServices bool = true
	this.ManageServices = &manageServices
	var provisioning bool = true
	this.Provisioning = &provisioning
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *WizardSteps) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *WizardSteps) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *WizardSteps) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCommonconf returns the Commonconf field value if set, zero value otherwise.
func (o *WizardSteps) GetCommonconf() bool {
	if o == nil || IsNil(o.Commonconf) {
		var ret bool
		return ret
	}
	return *o.Commonconf
}

// GetCommonconfOk returns a tuple with the Commonconf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetCommonconfOk() (*bool, bool) {
	if o == nil || IsNil(o.Commonconf) {
		return nil, false
	}
	return o.Commonconf, true
}

// HasCommonconf returns a boolean if a field has been set.
func (o *WizardSteps) HasCommonconf() bool {
	if o != nil && !IsNil(o.Commonconf) {
		return true
	}

	return false
}

// SetCommonconf gets a reference to the given bool and assigns it to the Commonconf field.
func (o *WizardSteps) SetCommonconf(v bool) {
	o.Commonconf = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *WizardSteps) GetDatabase() bool {
	if o == nil || IsNil(o.Database) {
		var ret bool
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetDatabaseOk() (*bool, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *WizardSteps) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given bool and assigns it to the Database field.
func (o *WizardSteps) SetDatabase(v bool) {
	o.Database = &v
}

// GetManageHostsFile returns the ManageHostsFile field value if set, zero value otherwise.
func (o *WizardSteps) GetManageHostsFile() bool {
	if o == nil || IsNil(o.ManageHostsFile) {
		var ret bool
		return ret
	}
	return *o.ManageHostsFile
}

// GetManageHostsFileOk returns a tuple with the ManageHostsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetManageHostsFileOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageHostsFile) {
		return nil, false
	}
	return o.ManageHostsFile, true
}

// HasManageHostsFile returns a boolean if a field has been set.
func (o *WizardSteps) HasManageHostsFile() bool {
	if o != nil && !IsNil(o.ManageHostsFile) {
		return true
	}

	return false
}

// SetManageHostsFile gets a reference to the given bool and assigns it to the ManageHostsFile field.
func (o *WizardSteps) SetManageHostsFile(v bool) {
	o.ManageHostsFile = &v
}

// GetManageResolvFile returns the ManageResolvFile field value if set, zero value otherwise.
func (o *WizardSteps) GetManageResolvFile() bool {
	if o == nil || IsNil(o.ManageResolvFile) {
		var ret bool
		return ret
	}
	return *o.ManageResolvFile
}

// GetManageResolvFileOk returns a tuple with the ManageResolvFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetManageResolvFileOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageResolvFile) {
		return nil, false
	}
	return o.ManageResolvFile, true
}

// HasManageResolvFile returns a boolean if a field has been set.
func (o *WizardSteps) HasManageResolvFile() bool {
	if o != nil && !IsNil(o.ManageResolvFile) {
		return true
	}

	return false
}

// SetManageResolvFile gets a reference to the given bool and assigns it to the ManageResolvFile field.
func (o *WizardSteps) SetManageResolvFile(v bool) {
	o.ManageResolvFile = &v
}

// GetManageServices returns the ManageServices field value if set, zero value otherwise.
func (o *WizardSteps) GetManageServices() bool {
	if o == nil || IsNil(o.ManageServices) {
		var ret bool
		return ret
	}
	return *o.ManageServices
}

// GetManageServicesOk returns a tuple with the ManageServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetManageServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageServices) {
		return nil, false
	}
	return o.ManageServices, true
}

// HasManageServices returns a boolean if a field has been set.
func (o *WizardSteps) HasManageServices() bool {
	if o != nil && !IsNil(o.ManageServices) {
		return true
	}

	return false
}

// SetManageServices gets a reference to the given bool and assigns it to the ManageServices field.
func (o *WizardSteps) SetManageServices(v bool) {
	o.ManageServices = &v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *WizardSteps) GetProvisioning() bool {
	if o == nil || IsNil(o.Provisioning) {
		var ret bool
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardSteps) GetProvisioningOk() (*bool, bool) {
	if o == nil || IsNil(o.Provisioning) {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *WizardSteps) HasProvisioning() bool {
	if o != nil && !IsNil(o.Provisioning) {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given bool and assigns it to the Provisioning field.
func (o *WizardSteps) SetProvisioning(v bool) {
	o.Provisioning = &v
}

func (o WizardSteps) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardSteps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Commonconf) {
		toSerialize["commonconf"] = o.Commonconf
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.ManageHostsFile) {
		toSerialize["manage_hosts_file"] = o.ManageHostsFile
	}
	if !IsNil(o.ManageResolvFile) {
		toSerialize["manage_resolv_file"] = o.ManageResolvFile
	}
	if !IsNil(o.ManageServices) {
		toSerialize["manage_services"] = o.ManageServices
	}
	if !IsNil(o.Provisioning) {
		toSerialize["provisioning"] = o.Provisioning
	}
	return toSerialize, nil
}

type NullableWizardSteps struct {
	value *WizardSteps
	isSet bool
}

func (v NullableWizardSteps) Get() *WizardSteps {
	return v.value
}

func (v *NullableWizardSteps) Set(val *WizardSteps) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardSteps) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardSteps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardSteps(val *WizardSteps) *NullableWizardSteps {
	return &NullableWizardSteps{value: val, isSet: true}
}

func (v NullableWizardSteps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardSteps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
