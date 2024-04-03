# DeviceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**Device**](Device.md) |  | [optional]

## Methods

### NewDeviceObject

`func NewDeviceObject() *DeviceObject`

NewDeviceObject instantiates a new DeviceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceObjectWithDefaults

`func NewDeviceObjectWithDefaults() *DeviceObject`

NewDeviceObjectWithDefaults instantiates a new DeviceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DeviceObject) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceObject) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceObject) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DeviceObject) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
