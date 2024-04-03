# SoundFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formats** | Pointer to [**[]SoundFormat**](SoundFormat.md) | The audio file formats | [optional] [readonly]
**Name** | Pointer to **string** | The name of the file | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant of the file | [optional]

## Methods

### NewSoundFile

`func NewSoundFile() *SoundFile`

NewSoundFile instantiates a new SoundFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundFileWithDefaults

`func NewSoundFileWithDefaults() *SoundFile`

NewSoundFileWithDefaults instantiates a new SoundFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormats

`func (o *SoundFile) GetFormats() []SoundFormat`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *SoundFile) GetFormatsOk() (*[]SoundFormat, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *SoundFile) SetFormats(v []SoundFormat)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *SoundFile) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### GetName

`func (o *SoundFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoundFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoundFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoundFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *SoundFile) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *SoundFile) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *SoundFile) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *SoundFile) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
