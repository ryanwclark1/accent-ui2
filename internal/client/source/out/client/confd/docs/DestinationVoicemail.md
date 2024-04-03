# DestinationVoicemail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Greeting** | Pointer to **string** | Play the specified greeting | [optional]
**SkipInstructions** | Pointer to **bool** | Skip the playback of instructions for leaving a message | [optional]
**Type** | **string** | MUST be &#39;voicemail&#39; |
**VoicemailId** | **int32** | The id of the voicemail |

## Methods

### NewDestinationVoicemail

`func NewDestinationVoicemail(type_ string, voicemailId int32, ) *DestinationVoicemail`

NewDestinationVoicemail instantiates a new DestinationVoicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationVoicemailWithDefaults

`func NewDestinationVoicemailWithDefaults() *DestinationVoicemail`

NewDestinationVoicemailWithDefaults instantiates a new DestinationVoicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreeting

`func (o *DestinationVoicemail) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *DestinationVoicemail) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *DestinationVoicemail) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *DestinationVoicemail) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetSkipInstructions

`func (o *DestinationVoicemail) GetSkipInstructions() bool`

GetSkipInstructions returns the SkipInstructions field if non-nil, zero value otherwise.

### GetSkipInstructionsOk

`func (o *DestinationVoicemail) GetSkipInstructionsOk() (*bool, bool)`

GetSkipInstructionsOk returns a tuple with the SkipInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipInstructions

`func (o *DestinationVoicemail) SetSkipInstructions(v bool)`

SetSkipInstructions sets SkipInstructions field to given value.

### HasSkipInstructions

`func (o *DestinationVoicemail) HasSkipInstructions() bool`

HasSkipInstructions returns a boolean if a field has been set.

### GetType

`func (o *DestinationVoicemail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationVoicemail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationVoicemail) SetType(v string)`

SetType sets Type field to given value.

### GetVoicemailId

`func (o *DestinationVoicemail) GetVoicemailId() int32`

GetVoicemailId returns the VoicemailId field if non-nil, zero value otherwise.

### GetVoicemailIdOk

`func (o *DestinationVoicemail) GetVoicemailIdOk() (*int32, bool)`

GetVoicemailIdOk returns a tuple with the VoicemailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemailId

`func (o *DestinationVoicemail) SetVoicemailId(v int32)`

SetVoicemailId sets VoicemailId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
