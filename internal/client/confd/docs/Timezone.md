# Timezone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | Pointer to **string** |  | [optional]

## Methods

### NewTimezone

`func NewTimezone() *Timezone`

NewTimezone instantiates a new Timezone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneWithDefaults

`func NewTimezoneWithDefaults() *Timezone`

NewTimezoneWithDefaults instantiates a new Timezone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *Timezone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *Timezone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *Timezone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *Timezone) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
