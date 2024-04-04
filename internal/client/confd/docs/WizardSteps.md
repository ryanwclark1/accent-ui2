# WizardSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Create admin in accent-auth? | [optional] [default to true]
**Commonconf** | Pointer to **bool** | Generate /etc/accent/common.conf | [optional] [default to true]
**Database** | Pointer to **bool** | Initialize Accent database? | [optional] [default to true]
**ManageHostsFile** | Pointer to **bool** | Modify /etc/hosts? | [optional] [default to true]
**ManageResolvFile** | Pointer to **bool** | Modify /etc/resolv.conf? | [optional] [default to true]
**ManageServices** | Pointer to **bool** | Enable/start Accent services? | [optional] [default to true]
**Provisioning** | Pointer to **bool** | Initialize accent-provd database? | [optional] [default to true]

## Methods

### NewWizardSteps

`func NewWizardSteps() *WizardSteps`

NewWizardSteps instantiates a new WizardSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardStepsWithDefaults

`func NewWizardStepsWithDefaults() *WizardSteps`

NewWizardStepsWithDefaults instantiates a new WizardSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *WizardSteps) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *WizardSteps) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *WizardSteps) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *WizardSteps) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCommonconf

`func (o *WizardSteps) GetCommonconf() bool`

GetCommonconf returns the Commonconf field if non-nil, zero value otherwise.

### GetCommonconfOk

`func (o *WizardSteps) GetCommonconfOk() (*bool, bool)`

GetCommonconfOk returns a tuple with the Commonconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonconf

`func (o *WizardSteps) SetCommonconf(v bool)`

SetCommonconf sets Commonconf field to given value.

### HasCommonconf

`func (o *WizardSteps) HasCommonconf() bool`

HasCommonconf returns a boolean if a field has been set.

### GetDatabase

`func (o *WizardSteps) GetDatabase() bool`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *WizardSteps) GetDatabaseOk() (*bool, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *WizardSteps) SetDatabase(v bool)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *WizardSteps) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetManageHostsFile

`func (o *WizardSteps) GetManageHostsFile() bool`

GetManageHostsFile returns the ManageHostsFile field if non-nil, zero value otherwise.

### GetManageHostsFileOk

`func (o *WizardSteps) GetManageHostsFileOk() (*bool, bool)`

GetManageHostsFileOk returns a tuple with the ManageHostsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageHostsFile

`func (o *WizardSteps) SetManageHostsFile(v bool)`

SetManageHostsFile sets ManageHostsFile field to given value.

### HasManageHostsFile

`func (o *WizardSteps) HasManageHostsFile() bool`

HasManageHostsFile returns a boolean if a field has been set.

### GetManageResolvFile

`func (o *WizardSteps) GetManageResolvFile() bool`

GetManageResolvFile returns the ManageResolvFile field if non-nil, zero value otherwise.

### GetManageResolvFileOk

`func (o *WizardSteps) GetManageResolvFileOk() (*bool, bool)`

GetManageResolvFileOk returns a tuple with the ManageResolvFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageResolvFile

`func (o *WizardSteps) SetManageResolvFile(v bool)`

SetManageResolvFile sets ManageResolvFile field to given value.

### HasManageResolvFile

`func (o *WizardSteps) HasManageResolvFile() bool`

HasManageResolvFile returns a boolean if a field has been set.

### GetManageServices

`func (o *WizardSteps) GetManageServices() bool`

GetManageServices returns the ManageServices field if non-nil, zero value otherwise.

### GetManageServicesOk

`func (o *WizardSteps) GetManageServicesOk() (*bool, bool)`

GetManageServicesOk returns a tuple with the ManageServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageServices

`func (o *WizardSteps) SetManageServices(v bool)`

SetManageServices sets ManageServices field to given value.

### HasManageServices

`func (o *WizardSteps) HasManageServices() bool`

HasManageServices returns a boolean if a field has been set.

### GetProvisioning

`func (o *WizardSteps) GetProvisioning() bool`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *WizardSteps) GetProvisioningOk() (*bool, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *WizardSteps) SetProvisioning(v bool)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *WizardSteps) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
