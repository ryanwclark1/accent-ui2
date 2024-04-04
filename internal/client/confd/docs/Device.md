# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional]
**Id** | Pointer to **string** | Device ID | [optional] [readonly]
**Ip** | Pointer to **string** | IP address (10.0.0.0) | [optional]
**IsNew** | Pointer to **bool** | Indicates if the device is a new device, ie in the master tenant | [optional] [readonly]
**Mac** | Pointer to **string** | MAC address (aa:bb:cc:dd:ee:ff) | [optional]
**Model** | Pointer to **string** | Device model | [optional]
**Options** | Pointer to [**DeviceOptions**](DeviceOptions.md) |  | [optional]
**Plugin** | Pointer to **string** | Provisioning plugin to be used by device | [optional]
**Sn** | Pointer to **string** | Serial number | [optional]
**Status** | Pointer to **string** | Device status. autoprov: Device can be provisionned using a provisioning code, configured: Device is configured and ready to be used, not_configured : Device has not been completely configured | [optional] [readonly] [default to "not_configured"]
**TemplateId** | Pointer to **string** | ID of the device template. All device using a device template will have a certain number of common parameters preconfigured for the device | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
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

### GetOptions

`func (o *Device) GetOptions() DeviceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Device) GetOptionsOk() (*DeviceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Device) SetOptions(v DeviceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Device) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

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

### GetSn

`func (o *Device) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *Device) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *Device) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *Device) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *Device) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Device) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Device) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Device) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

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
