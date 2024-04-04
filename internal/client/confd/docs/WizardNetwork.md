# WizardNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name |
**Gateway** | **string** | Gateway IPv4 address |
**Hostname** | **string** | System hostname |
**Interface** | **string** | Interface name (e.g. eth0) |
**IpAddress** | **string** | IPv4 address of the VoIP interface (connected to phones) |
**Nameservers** | **[]string** | List of IPv4 addresses. Nameservers are used in resolv.conf. |
**Netmask** | **string** | Netmask of the IP address (e.g. 255.255.0.0) |

## Methods

### NewWizardNetwork

`func NewWizardNetwork(domain string, gateway string, hostname string, interface_ string, ipAddress string, nameservers []string, netmask string, ) *WizardNetwork`

NewWizardNetwork instantiates a new WizardNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardNetworkWithDefaults

`func NewWizardNetworkWithDefaults() *WizardNetwork`

NewWizardNetworkWithDefaults instantiates a new WizardNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *WizardNetwork) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WizardNetwork) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WizardNetwork) SetDomain(v string)`

SetDomain sets Domain field to given value.

### GetGateway

`func (o *WizardNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *WizardNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *WizardNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### GetHostname

`func (o *WizardNetwork) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WizardNetwork) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WizardNetwork) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### GetInterface

`func (o *WizardNetwork) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *WizardNetwork) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *WizardNetwork) SetInterface(v string)`

SetInterface sets Interface field to given value.

### GetIpAddress

`func (o *WizardNetwork) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *WizardNetwork) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *WizardNetwork) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### GetNameservers

`func (o *WizardNetwork) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *WizardNetwork) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *WizardNetwork) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### GetNetmask

`func (o *WizardNetwork) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *WizardNetwork) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *WizardNetwork) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
