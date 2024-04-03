# Voicemail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to [**[]VoicemailFolder**](VoicemailFolder.md) | The voicemail&#39;s folders | [optional]
**Id** | Pointer to **int32** | The voicemail&#39;s ID | [optional]
**Name** | Pointer to **string** | The voicemail&#39;s name | [optional]
**Number** | Pointer to **string** | The voicemail&#39;s number | [optional]

## Methods

### NewVoicemail

`func NewVoicemail() *Voicemail`

NewVoicemail instantiates a new Voicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailWithDefaults

`func NewVoicemailWithDefaults() *Voicemail`

NewVoicemailWithDefaults instantiates a new Voicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *Voicemail) GetFolders() []VoicemailFolder`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *Voicemail) GetFoldersOk() (*[]VoicemailFolder, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *Voicemail) SetFolders(v []VoicemailFolder)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *Voicemail) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetId

`func (o *Voicemail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Voicemail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Voicemail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Voicemail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Voicemail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Voicemail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Voicemail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Voicemail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *Voicemail) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Voicemail) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Voicemail) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Voicemail) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
