# Sound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]SoundFile**](SoundFile.md) | The audio files | [optional] [readonly]
**Name** | **string** | The name of the category (can only by set on create and must be unique) |
**TenantUuid** | Pointer to **string** | The UUID of the tenant of the category | [optional]

## Methods

### NewSound

`func NewSound(name string, ) *Sound`

NewSound instantiates a new Sound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundWithDefaults

`func NewSoundWithDefaults() *Sound`

NewSoundWithDefaults instantiates a new Sound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *Sound) GetFiles() []SoundFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Sound) GetFilesOk() (*[]SoundFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Sound) SetFiles(v []SoundFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Sound) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetName

`func (o *Sound) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sound) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sound) SetName(v string)`

SetName sets Name field to given value.

### GetTenantUuid

`func (o *Sound) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Sound) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Sound) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Sound) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
