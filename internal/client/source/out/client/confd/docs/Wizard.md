# Wizard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPassword** | **string** | Accent administrator &#x60;&#x60;root&#x60;&#x60; password |
**Language** | Pointer to **string** | The language in which the Accent will play sounds | [optional] [default to "en_US"]
**License** | **bool** | Accept/decline the GPLv3: <http://www.gnu.org/licenses/gpl-3.0.en.html> |
**Network** | [**WizardNetwork**](WizardNetwork.md) |  |
**Steps** | Pointer to [**WizardSteps**](WizardSteps.md) |  | [optional]
**Timezone** | **string** | System timezone. Example: America/Montreal. For the complete list of supported timezones, see &#x60;&#x60;/usr/share/zoneinfo/&#x60;&#x60; |

## Methods

### NewWizard

`func NewWizard(adminPassword string, license bool, network WizardNetwork, timezone string, ) *Wizard`

NewWizard instantiates a new Wizard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardWithDefaults

`func NewWizardWithDefaults() *Wizard`

NewWizardWithDefaults instantiates a new Wizard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPassword

`func (o *Wizard) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *Wizard) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *Wizard) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### GetLanguage

`func (o *Wizard) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Wizard) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Wizard) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Wizard) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLicense

`func (o *Wizard) GetLicense() bool`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Wizard) GetLicenseOk() (*bool, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Wizard) SetLicense(v bool)`

SetLicense sets License field to given value.

### GetNetwork

`func (o *Wizard) GetNetwork() WizardNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Wizard) GetNetworkOk() (*WizardNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Wizard) SetNetwork(v WizardNetwork)`

SetNetwork sets Network field to given value.

### GetSteps

`func (o *Wizard) GetSteps() WizardSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Wizard) GetStepsOk() (*WizardSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Wizard) SetSteps(v WizardSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Wizard) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTimezone

`func (o *Wizard) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Wizard) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Wizard) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
