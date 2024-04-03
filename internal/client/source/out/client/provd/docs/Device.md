# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **string** | Indicated how the device was added | [optional]
**Config** | Pointer to **string** | ID of the device configuration. Generally the same as the device ID, except when in autoprov | [optional]
**Configured** | Pointer to **bool** |  | [optional] [readonly]
**Description** | Pointer to **string** |  | [optional]
**Id** | Pointer to **string** | Device ID | [optional]
**Ip** | Pointer to **string** | IP address (10.0.0.0) | [optional]
**IsNew** | Pointer to **bool** | Indicates if the device is a new device, ie in the master tenant | [optional] [readonly]
**Mac** | Pointer to **string** | MAC address (aa:bb:cc:dd:ee:ff) | [optional]
**Model** | Pointer to **string** | Device model | [optional]
**Plugin** | Pointer to **string** | Provisioning plugin used by the device | [optional]
**RemoteStateSipUsername** | Pointer to **string** |  | [optional]
**TenantUuid** | Pointer to **string** | The tenant UUID, defining the ownership of this device | [optional] [readonly]
**Vendor** | Pointer to **string** | Vendor name | [optional]
**Version** | Pointer to **string** | Firmware version | [optional]

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *Device) GetAdded() string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *Device) GetAddedOk() (*string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *Device) SetAdded(v string)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *Device) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetConfig

`func (o *Device) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Device) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Device) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Device) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigured

`func (o *Device) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *Device) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *Device) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *Device) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDescription

`func (o *Device) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Device) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Device) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Device) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *Device) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Device) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Device) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Device) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsNew

`func (o *Device) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *Device) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *Device) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *Device) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetMac

`func (o *Device) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Device) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Device) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Device) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Device) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Device) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPlugin

`func (o *Device) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *Device) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *Device) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *Device) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetRemoteStateSipUsername

`func (o *Device) GetRemoteStateSipUsername() string`

GetRemoteStateSipUsername returns the RemoteStateSipUsername field if non-nil, zero value otherwise.

### GetRemoteStateSipUsernameOk

`func (o *Device) GetRemoteStateSipUsernameOk() (*string, bool)`

GetRemoteStateSipUsernameOk returns a tuple with the RemoteStateSipUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStateSipUsername

`func (o *Device) SetRemoteStateSipUsername(v string)`

SetRemoteStateSipUsername sets RemoteStateSipUsername field to given value.

### HasRemoteStateSipUsername

`func (o *Device) HasRemoteStateSipUsername() bool`

HasRemoteStateSipUsername returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Device) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Device) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Device) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Device) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetVendor

`func (o *Device) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Device) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Device) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Device) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *Device) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Device) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Device) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Device) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
