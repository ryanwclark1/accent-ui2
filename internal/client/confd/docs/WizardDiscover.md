# WizardDiscover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | System domain name | [optional]
**Gateways** | Pointer to [**[]WizardDiscoverGateway**](WizardDiscoverGateway.md) |  | [optional]
**Hostname** | Pointer to **string** | System hostname | [optional]
**Interfaces** | Pointer to [**[]WizardDiscoverInterface**](WizardDiscoverInterface.md) |  | [optional]
**Nameservers** | Pointer to **[]string** | Nameservers from file &#x60;&#x60;/etc/resolv.conf&#x60;&#x60; | [optional]
**Timezone** | Pointer to **string** | System timezone from file &#x60;&#x60;/etc/timezone&#x60;&#x60; | [optional]

## Methods

### NewWizardDiscover

`func NewWizardDiscover() *WizardDiscover`

NewWizardDiscover instantiates a new WizardDiscover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardDiscoverWithDefaults

`func NewWizardDiscoverWithDefaults() *WizardDiscover`

NewWizardDiscoverWithDefaults instantiates a new WizardDiscover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *WizardDiscover) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WizardDiscover) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WizardDiscover) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WizardDiscover) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGateways

`func (o *WizardDiscover) GetGateways() []WizardDiscoverGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *WizardDiscover) GetGatewaysOk() (*[]WizardDiscoverGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *WizardDiscover) SetGateways(v []WizardDiscoverGateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *WizardDiscover) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetHostname

`func (o *WizardDiscover) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WizardDiscover) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WizardDiscover) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WizardDiscover) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInterfaces

`func (o *WizardDiscover) GetInterfaces() []WizardDiscoverInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *WizardDiscover) GetInterfacesOk() (*[]WizardDiscoverInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *WizardDiscover) SetInterfaces(v []WizardDiscoverInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *WizardDiscover) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetNameservers

`func (o *WizardDiscover) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *WizardDiscover) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *WizardDiscover) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *WizardDiscover) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetTimezone

`func (o *WizardDiscover) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WizardDiscover) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WizardDiscover) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *WizardDiscover) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
