# ProvisioningNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisionHost** | Pointer to **string** | The hostname or IP address used for HTTP and TFTP provisioning requests for DHCP integration. | [optional]
**ProvisionHttpBaseUrl** | Pointer to **string** | The base URL on which the provisioning server will be accessible from outside the network. | [optional]
**ProvisionHttpPort** | Pointer to **int32** | The port used by the HTTP provisioning server. | [optional]

## Methods

### NewProvisioningNetworking

`func NewProvisioningNetworking() *ProvisioningNetworking`

NewProvisioningNetworking instantiates a new ProvisioningNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningNetworkingWithDefaults

`func NewProvisioningNetworkingWithDefaults() *ProvisioningNetworking`

NewProvisioningNetworkingWithDefaults instantiates a new ProvisioningNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisionHost

`func (o *ProvisioningNetworking) GetProvisionHost() string`

GetProvisionHost returns the ProvisionHost field if non-nil, zero value otherwise.

### GetProvisionHostOk

`func (o *ProvisioningNetworking) GetProvisionHostOk() (*string, bool)`

GetProvisionHostOk returns a tuple with the ProvisionHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionHost

`func (o *ProvisioningNetworking) SetProvisionHost(v string)`

SetProvisionHost sets ProvisionHost field to given value.

### HasProvisionHost

`func (o *ProvisioningNetworking) HasProvisionHost() bool`

HasProvisionHost returns a boolean if a field has been set.

### GetProvisionHttpBaseUrl

`func (o *ProvisioningNetworking) GetProvisionHttpBaseUrl() string`

GetProvisionHttpBaseUrl returns the ProvisionHttpBaseUrl field if non-nil, zero value otherwise.

### GetProvisionHttpBaseUrlOk

`func (o *ProvisioningNetworking) GetProvisionHttpBaseUrlOk() (*string, bool)`

GetProvisionHttpBaseUrlOk returns a tuple with the ProvisionHttpBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionHttpBaseUrl

`func (o *ProvisioningNetworking) SetProvisionHttpBaseUrl(v string)`

SetProvisionHttpBaseUrl sets ProvisionHttpBaseUrl field to given value.

### HasProvisionHttpBaseUrl

`func (o *ProvisioningNetworking) HasProvisionHttpBaseUrl() bool`

HasProvisionHttpBaseUrl returns a boolean if a field has been set.

### GetProvisionHttpPort

`func (o *ProvisioningNetworking) GetProvisionHttpPort() int32`

GetProvisionHttpPort returns the ProvisionHttpPort field if non-nil, zero value otherwise.

### GetProvisionHttpPortOk

`func (o *ProvisioningNetworking) GetProvisionHttpPortOk() (*int32, bool)`

GetProvisionHttpPortOk returns a tuple with the ProvisionHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionHttpPort

`func (o *ProvisioningNetworking) SetProvisionHttpPort(v int32)`

SetProvisionHttpPort sets ProvisionHttpPort field to given value.

### HasProvisionHttpPort

`func (o *ProvisioningNetworking) HasProvisionHttpPort() bool`

HasProvisionHttpPort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
