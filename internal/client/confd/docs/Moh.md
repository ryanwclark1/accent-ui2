# Moh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | The command to run (only used when mode is \&quot;custom\&quot;) | [optional]
**Files** | Pointer to [**[]MohFile**](MohFile.md) | The audio files | [optional] [readonly]
**Label** | Pointer to **string** | The label of the MOH class | [optional]
**Mode** | Pointer to **string** | The play mode of the MOH class. Notice: &#x60;mp3&#x60; is deprecated and should not be used | [optional] [default to "files"]
**Name** | Pointer to **string** | The name used by Asterisk | [optional] [readonly]
**Sort** | Pointer to **string** | The order in which files are played (only used when mode is \&quot;files\&quot;) | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Uuid** | Pointer to **string** | The UUID of the MOH class | [optional] [readonly]

## Methods

### NewMoh

`func NewMoh() *Moh`

NewMoh instantiates a new Moh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMohWithDefaults

`func NewMohWithDefaults() *Moh`

NewMohWithDefaults instantiates a new Moh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *Moh) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Moh) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Moh) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Moh) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetFiles

`func (o *Moh) GetFiles() []MohFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Moh) GetFilesOk() (*[]MohFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Moh) SetFiles(v []MohFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Moh) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLabel

`func (o *Moh) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Moh) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Moh) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Moh) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMode

`func (o *Moh) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Moh) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Moh) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Moh) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *Moh) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Moh) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Moh) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Moh) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSort

`func (o *Moh) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Moh) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Moh) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Moh) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Moh) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Moh) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Moh) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Moh) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUuid

`func (o *Moh) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Moh) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Moh) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Moh) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
