# DeviceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Switchboard** | Pointer to **bool** | Indicate if this device is a switchboard | [optional]

## Methods

### NewDeviceOptions

`func NewDeviceOptions() *DeviceOptions`

NewDeviceOptions instantiates a new DeviceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOptionsWithDefaults

`func NewDeviceOptionsWithDefaults() *DeviceOptions`

NewDeviceOptionsWithDefaults instantiates a new DeviceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchboard

`func (o *DeviceOptions) GetSwitchboard() bool`

GetSwitchboard returns the Switchboard field if non-nil, zero value otherwise.

### GetSwitchboardOk

`func (o *DeviceOptions) GetSwitchboardOk() (*bool, bool)`

GetSwitchboardOk returns a tuple with the Switchboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboard

`func (o *DeviceOptions) SetSwitchboard(v bool)`

SetSwitchboard sets Switchboard field to given value.

### HasSwitchboard

`func (o *DeviceOptions) HasSwitchboard() bool`

HasSwitchboard returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
