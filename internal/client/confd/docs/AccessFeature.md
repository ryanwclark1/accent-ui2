# AccessFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional]
**Feature** | **string** | The feature to limit the access to |
**Host** | **string** | The host or subnet string (e.g. 10.0.0.0/24) |
**Id** | Pointer to **int32** | The access_feature ID | [optional]

## Methods

### NewAccessFeature

`func NewAccessFeature(feature string, host string, ) *AccessFeature`

NewAccessFeature instantiates a new AccessFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessFeatureWithDefaults

`func NewAccessFeatureWithDefaults() *AccessFeature`

NewAccessFeatureWithDefaults instantiates a new AccessFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AccessFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccessFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccessFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccessFeature) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeature

`func (o *AccessFeature) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *AccessFeature) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *AccessFeature) SetFeature(v string)`

SetFeature sets Feature field to given value.

### GetHost

`func (o *AccessFeature) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AccessFeature) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AccessFeature) SetHost(v string)`

SetHost sets Host field to given value.

### GetId

`func (o *AccessFeature) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessFeature) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessFeature) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessFeature) HasId() bool`

HasId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
