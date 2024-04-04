# DHCP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Is the DHCP service enabled? | [optional]
**NetworkInterfaces** | Pointer to **[]string** | A comma separated list of network interface that the DHCP server listens on | [optional]
**PoolEnd** | Pointer to **string** | The last IP address that can be allocated by DHCP | [optional]
**PoolStart** | Pointer to **string** | The first IP address that can be allocated by DHCP | [optional]

## Methods

### NewDHCP

`func NewDHCP() *DHCP`

NewDHCP instantiates a new DHCP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPWithDefaults

`func NewDHCPWithDefaults() *DHCP`

NewDHCPWithDefaults instantiates a new DHCP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DHCP) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DHCP) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DHCP) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DHCP) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *DHCP) GetNetworkInterfaces() []string`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *DHCP) GetNetworkInterfacesOk() (*[]string, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *DHCP) SetNetworkInterfaces(v []string)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *DHCP) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPoolEnd

`func (o *DHCP) GetPoolEnd() string`

GetPoolEnd returns the PoolEnd field if non-nil, zero value otherwise.

### GetPoolEndOk

`func (o *DHCP) GetPoolEndOk() (*string, bool)`

GetPoolEndOk returns a tuple with the PoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEnd

`func (o *DHCP) SetPoolEnd(v string)`

SetPoolEnd sets PoolEnd field to given value.

### HasPoolEnd

`func (o *DHCP) HasPoolEnd() bool`

HasPoolEnd returns a boolean if a field has been set.

### GetPoolStart

`func (o *DHCP) GetPoolStart() string`

GetPoolStart returns the PoolStart field if non-nil, zero value otherwise.

### GetPoolStartOk

`func (o *DHCP) GetPoolStartOk() (*string, bool)`

GetPoolStartOk returns a tuple with the PoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStart

`func (o *DHCP) SetPoolStart(v string)`

SetPoolStart sets PoolStart field to given value.

### HasPoolStart

`func (o *DHCP) HasPoolStart() bool`

HasPoolStart returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
