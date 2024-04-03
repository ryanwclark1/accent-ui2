# DHCPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address of the device | [optional]
**Mac** | Pointer to **string** | The MAC address of the device | [optional]
**Op** | Pointer to **string** | The operation to perform | [optional]
**Options** | Pointer to **[]string** |  | [optional]

## Methods

### NewDHCPInfo

`func NewDHCPInfo() *DHCPInfo`

NewDHCPInfo instantiates a new DHCPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPInfoWithDefaults

`func NewDHCPInfoWithDefaults() *DHCPInfo`

NewDHCPInfoWithDefaults instantiates a new DHCPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DHCPInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DHCPInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DHCPInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DHCPInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *DHCPInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DHCPInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DHCPInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DHCPInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOp

`func (o *DHCPInfo) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *DHCPInfo) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *DHCPInfo) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *DHCPInfo) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOptions

`func (o *DHCPInfo) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DHCPInfo) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DHCPInfo) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DHCPInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
