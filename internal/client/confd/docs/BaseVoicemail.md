# BaseVoicemail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPassword** | Pointer to **bool** | Ask for password when accessing the voicemail menu | [optional] [default to true]
**AttachAudio** | Pointer to **bool** | Attach an audio file of the recorded message when sending an email | [optional]
**Context** | **string** | Voicemail context |
**DeleteMessages** | Pointer to **bool** | Delete messages after reception.  This can only be set along with &#39;email&#39; and &#39;attach_audio&#39; configured.  | [optional] [default to false]
**Email** | Pointer to **string** | Email address. Notifications and audio attachments will be sent to this address | [optional]
**Enabled** | Pointer to **bool** | Voicemail can be used when it is enabled | [optional] [default to true]
**Language** | Pointer to **string** | Language used for the voicemail menu prompt | [optional]
**MaxMessages** | Pointer to **int32** | Maximum number of messages to store | [optional]
**Number** | **string** | Mailbox number, for example &#x60;0001&#x60; |
**Options** | Pointer to **[][]string** | Advanced configuration options. Options are appended at the end of a  voicemail line in the file &#39;voicemail.conf&#39; used by asterisk. Please consult the asterisk documentation for further details on available parameters. Options must have the following the form:  &#x60;&#x60;&#x60; {   \&quot;options\&quot;: [     [\&quot;name1\&quot;, \&quot;value1\&quot;],     [\&quot;name2\&quot;, \&quot;value2\&quot;]   ] } &#x60;&#x60;&#x60;  The resulting configuration in voicemail.conf will have the following form:  &#x60;&#x60;&#x60; 1000 &#x3D;&gt; ,Firstname Lastname,,,name1&#x3D;value1|name2&#x3D;value2 &#x60;&#x60;&#x60;  | [optional]
**Pager** | Pointer to **string** | Email address. Summarized notifications will be sent to this address | [optional]
**Password** | Pointer to **string** | Numeric password used to access the voicemail | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timezone** | Pointer to **string** | Timezone used for announcing at what time a message was recorded | [optional] [default to "The system default timezone"]

## Methods

### NewBaseVoicemail

`func NewBaseVoicemail(context string, number string, ) *BaseVoicemail`

NewBaseVoicemail instantiates a new BaseVoicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseVoicemailWithDefaults

`func NewBaseVoicemailWithDefaults() *BaseVoicemail`

NewBaseVoicemailWithDefaults instantiates a new BaseVoicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPassword

`func (o *BaseVoicemail) GetAskPassword() bool`

GetAskPassword returns the AskPassword field if non-nil, zero value otherwise.

### GetAskPasswordOk

`func (o *BaseVoicemail) GetAskPasswordOk() (*bool, bool)`

GetAskPasswordOk returns a tuple with the AskPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPassword

`func (o *BaseVoicemail) SetAskPassword(v bool)`

SetAskPassword sets AskPassword field to given value.

### HasAskPassword

`func (o *BaseVoicemail) HasAskPassword() bool`

HasAskPassword returns a boolean if a field has been set.

### GetAttachAudio

`func (o *BaseVoicemail) GetAttachAudio() bool`

GetAttachAudio returns the AttachAudio field if non-nil, zero value otherwise.

### GetAttachAudioOk

`func (o *BaseVoicemail) GetAttachAudioOk() (*bool, bool)`

GetAttachAudioOk returns a tuple with the AttachAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAudio

`func (o *BaseVoicemail) SetAttachAudio(v bool)`

SetAttachAudio sets AttachAudio field to given value.

### HasAttachAudio

`func (o *BaseVoicemail) HasAttachAudio() bool`

HasAttachAudio returns a boolean if a field has been set.

### GetContext

`func (o *BaseVoicemail) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BaseVoicemail) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BaseVoicemail) SetContext(v string)`

SetContext sets Context field to given value.

### GetDeleteMessages

`func (o *BaseVoicemail) GetDeleteMessages() bool`

GetDeleteMessages returns the DeleteMessages field if non-nil, zero value otherwise.

### GetDeleteMessagesOk

`func (o *BaseVoicemail) GetDeleteMessagesOk() (*bool, bool)`

GetDeleteMessagesOk returns a tuple with the DeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessages

`func (o *BaseVoicemail) SetDeleteMessages(v bool)`

SetDeleteMessages sets DeleteMessages field to given value.

### HasDeleteMessages

`func (o *BaseVoicemail) HasDeleteMessages() bool`

HasDeleteMessages returns a boolean if a field has been set.

### GetEmail

`func (o *BaseVoicemail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BaseVoicemail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BaseVoicemail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BaseVoicemail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *BaseVoicemail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseVoicemail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseVoicemail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseVoicemail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLanguage

`func (o *BaseVoicemail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BaseVoicemail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BaseVoicemail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BaseVoicemail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMaxMessages

`func (o *BaseVoicemail) GetMaxMessages() int32`

GetMaxMessages returns the MaxMessages field if non-nil, zero value otherwise.

### GetMaxMessagesOk

`func (o *BaseVoicemail) GetMaxMessagesOk() (*int32, bool)`

GetMaxMessagesOk returns a tuple with the MaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessages

`func (o *BaseVoicemail) SetMaxMessages(v int32)`

SetMaxMessages sets MaxMessages field to given value.

### HasMaxMessages

`func (o *BaseVoicemail) HasMaxMessages() bool`

HasMaxMessages returns a boolean if a field has been set.

### GetNumber

`func (o *BaseVoicemail) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BaseVoicemail) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BaseVoicemail) SetNumber(v string)`

SetNumber sets Number field to given value.

### GetOptions

`func (o *BaseVoicemail) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BaseVoicemail) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BaseVoicemail) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BaseVoicemail) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPager

`func (o *BaseVoicemail) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *BaseVoicemail) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *BaseVoicemail) SetPager(v string)`

SetPager sets Pager field to given value.

### HasPager

`func (o *BaseVoicemail) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPassword

`func (o *BaseVoicemail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BaseVoicemail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BaseVoicemail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BaseVoicemail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTenantUuid

`func (o *BaseVoicemail) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *BaseVoicemail) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *BaseVoicemail) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *BaseVoicemail) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *BaseVoicemail) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BaseVoicemail) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BaseVoicemail) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BaseVoicemail) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
