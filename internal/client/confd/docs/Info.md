# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID of the Accent server | [optional]
**AccentVersion** | Pointer to **string** | Version of the Accent server | [optional]

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Info) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Info) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Info) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Info) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccentVersion

`func (o *Info) GetAccentVersion() string`

GetAccentVersion returns the AccentVersion field if non-nil, zero value otherwise.

### GetAccentVersionOk

`func (o *Info) GetAccentVersionOk() (*string, bool)`

GetAccentVersionOk returns a tuple with the AccentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentVersion

`func (o *Info) SetAccentVersion(v string)`

SetAccentVersion sets AccentVersion field to given value.

### HasAccentVersion

`func (o *Info) HasAccentVersion() bool`

HasAccentVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
