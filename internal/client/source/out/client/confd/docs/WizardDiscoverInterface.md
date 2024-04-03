# WizardDiscoverInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Interface name (e.g. eth0) | [optional]
**IpAddress** | Pointer to **string** | IPv4 address of the interface | [optional]
**Netmask** | Pointer to **string** | Netmask of the IP address | [optional]

## Methods

### NewWizardDiscoverInterface

`func NewWizardDiscoverInterface() *WizardDiscoverInterface`

NewWizardDiscoverInterface instantiates a new WizardDiscoverInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardDiscoverInterfaceWithDefaults

`func NewWizardDiscoverInterfaceWithDefaults() *WizardDiscoverInterface`

NewWizardDiscoverInterfaceWithDefaults instantiates a new WizardDiscoverInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *WizardDiscoverInterface) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *WizardDiscoverInterface) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *WizardDiscoverInterface) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *WizardDiscoverInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpAddress

`func (o *WizardDiscoverInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *WizardDiscoverInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *WizardDiscoverInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *WizardDiscoverInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *WizardDiscoverInterface) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *WizardDiscoverInterface) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *WizardDiscoverInterface) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *WizardDiscoverInterface) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
