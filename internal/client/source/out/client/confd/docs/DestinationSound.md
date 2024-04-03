# DestinationSound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The filename of the sound. The file MUST be imported by to webi in the playback directory. The extension of file SHOULD be not present. |
**NoAnswer** | Pointer to **bool** | Play this sound without answering the call | [optional]
**Skip** | Pointer to **bool** | Do not play this sound if the call is not answered | [optional]
**Type** | **string** | MUST be &#39;sound&#39; |

## Methods

### NewDestinationSound

`func NewDestinationSound(filename string, type_ string, ) *DestinationSound`

NewDestinationSound instantiates a new DestinationSound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSoundWithDefaults

`func NewDestinationSoundWithDefaults() *DestinationSound`

NewDestinationSoundWithDefaults instantiates a new DestinationSound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *DestinationSound) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DestinationSound) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DestinationSound) SetFilename(v string)`

SetFilename sets Filename field to given value.

### GetNoAnswer

`func (o *DestinationSound) GetNoAnswer() bool`

GetNoAnswer returns the NoAnswer field if non-nil, zero value otherwise.

### GetNoAnswerOk

`func (o *DestinationSound) GetNoAnswerOk() (*bool, bool)`

GetNoAnswerOk returns a tuple with the NoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnswer

`func (o *DestinationSound) SetNoAnswer(v bool)`

SetNoAnswer sets NoAnswer field to given value.

### HasNoAnswer

`func (o *DestinationSound) HasNoAnswer() bool`

HasNoAnswer returns a boolean if a field has been set.

### GetSkip

`func (o *DestinationSound) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *DestinationSound) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *DestinationSound) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *DestinationSound) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetType

`func (o *DestinationSound) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationSound) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationSound) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
