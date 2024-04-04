# SoundFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional]
**Language** | Pointer to **string** |  | [optional]
**Path** | Pointer to **string** |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant of the file | [optional]
**Text** | Pointer to **string** |  | [optional]

## Methods

### NewSoundFormat

`func NewSoundFormat() *SoundFormat`

NewSoundFormat instantiates a new SoundFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundFormatWithDefaults

`func NewSoundFormatWithDefaults() *SoundFormat`

NewSoundFormatWithDefaults instantiates a new SoundFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SoundFormat) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SoundFormat) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SoundFormat) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SoundFormat) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLanguage

`func (o *SoundFormat) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SoundFormat) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SoundFormat) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SoundFormat) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPath

`func (o *SoundFormat) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SoundFormat) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SoundFormat) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SoundFormat) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTenantUuid

`func (o *SoundFormat) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *SoundFormat) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *SoundFormat) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *SoundFormat) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetText

`func (o *SoundFormat) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SoundFormat) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SoundFormat) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SoundFormat) HasText() bool`

HasText returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
