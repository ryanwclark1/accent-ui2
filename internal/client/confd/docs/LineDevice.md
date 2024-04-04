# LineDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | Device&#39;s ID | [optional] [readonly]
**LineId** | Pointer to **int32** | Line&#39;s ID | [optional] [readonly]

## Methods

### NewLineDevice

`func NewLineDevice() *LineDevice`

NewLineDevice instantiates a new LineDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineDeviceWithDefaults

`func NewLineDeviceWithDefaults() *LineDevice`

NewLineDeviceWithDefaults instantiates a new LineDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *LineDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LineDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LineDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *LineDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLineId

`func (o *LineDevice) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *LineDevice) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *LineDevice) SetLineId(v int32)`

SetLineId sets LineId field to given value.

### HasLineId

`func (o *LineDevice) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
