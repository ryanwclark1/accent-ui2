# IAXCallNumberLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | IPv4 address of the destination |
**Limit** | **string** | Maximum call for the ip_address destination |
**Netmask** | **string** | Netmask of the IP address (e.g. 255.255.255.255) |

## Methods

### NewIAXCallNumberLimits

`func NewIAXCallNumberLimits(ipAddress string, limit string, netmask string, ) *IAXCallNumberLimits`

NewIAXCallNumberLimits instantiates a new IAXCallNumberLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAXCallNumberLimitsWithDefaults

`func NewIAXCallNumberLimitsWithDefaults() *IAXCallNumberLimits`

NewIAXCallNumberLimitsWithDefaults instantiates a new IAXCallNumberLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *IAXCallNumberLimits) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IAXCallNumberLimits) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IAXCallNumberLimits) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### GetLimit

`func (o *IAXCallNumberLimits) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IAXCallNumberLimits) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IAXCallNumberLimits) SetLimit(v string)`

SetLimit sets Limit field to given value.

### GetNetmask

`func (o *IAXCallNumberLimits) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IAXCallNumberLimits) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IAXCallNumberLimits) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
