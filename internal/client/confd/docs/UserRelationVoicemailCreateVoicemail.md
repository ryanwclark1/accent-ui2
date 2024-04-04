# UserRelationVoicemailCreateVoicemail

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
**Id** | Pointer to **int32** | The voicemail is associated if the ID is specified | [optional]

## Methods

### NewUserRelationVoicemailCreateVoicemail

`func NewUserRelationVoicemailCreateVoicemail(context string, number string, ) *UserRelationVoicemailCreateVoicemail`

NewUserRelationVoicemailCreateVoicemail instantiates a new UserRelationVoicemailCreateVoicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelationVoicemailCreateVoicemailWithDefaults

`func NewUserRelationVoicemailCreateVoicemailWithDefaults() *UserRelationVoicemailCreateVoicemail`

NewUserRelationVoicemailCreateVoicemailWithDefaults instantiates a new UserRelationVoicemailCreateVoicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPassword

`func (o *UserRelationVoicemailCreateVoicemail) GetAskPassword() bool`

GetAskPassword returns the AskPassword field if non-nil, zero value otherwise.

### GetAskPasswordOk

`func (o *UserRelationVoicemailCreateVoicemail) GetAskPasswordOk() (*bool, bool)`

GetAskPasswordOk returns a tuple with the AskPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPassword

`func (o *UserRelationVoicemailCreateVoicemail) SetAskPassword(v bool)`

SetAskPassword sets AskPassword field to given value.

### HasAskPassword

`func (o *UserRelationVoicemailCreateVoicemail) HasAskPassword() bool`

HasAskPassword returns a boolean if a field has been set.

### GetAttachAudio

`func (o *UserRelationVoicemailCreateVoicemail) GetAttachAudio() bool`

GetAttachAudio returns the AttachAudio field if non-nil, zero value otherwise.

### GetAttachAudioOk

`func (o *UserRelationVoicemailCreateVoicemail) GetAttachAudioOk() (*bool, bool)`

GetAttachAudioOk returns a tuple with the AttachAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAudio

`func (o *UserRelationVoicemailCreateVoicemail) SetAttachAudio(v bool)`

SetAttachAudio sets AttachAudio field to given value.

### HasAttachAudio

`func (o *UserRelationVoicemailCreateVoicemail) HasAttachAudio() bool`

HasAttachAudio returns a boolean if a field has been set.

### GetContext

`func (o *UserRelationVoicemailCreateVoicemail) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserRelationVoicemailCreateVoicemail) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserRelationVoicemailCreateVoicemail) SetContext(v string)`

SetContext sets Context field to given value.

### GetDeleteMessages

`func (o *UserRelationVoicemailCreateVoicemail) GetDeleteMessages() bool`

GetDeleteMessages returns the DeleteMessages field if non-nil, zero value otherwise.

### GetDeleteMessagesOk

`func (o *UserRelationVoicemailCreateVoicemail) GetDeleteMessagesOk() (*bool, bool)`

GetDeleteMessagesOk returns a tuple with the DeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessages

`func (o *UserRelationVoicemailCreateVoicemail) SetDeleteMessages(v bool)`

SetDeleteMessages sets DeleteMessages field to given value.

### HasDeleteMessages

`func (o *UserRelationVoicemailCreateVoicemail) HasDeleteMessages() bool`

HasDeleteMessages returns a boolean if a field has been set.

### GetEmail

`func (o *UserRelationVoicemailCreateVoicemail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRelationVoicemailCreateVoicemail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRelationVoicemailCreateVoicemail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRelationVoicemailCreateVoicemail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *UserRelationVoicemailCreateVoicemail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserRelationVoicemailCreateVoicemail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserRelationVoicemailCreateVoicemail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserRelationVoicemailCreateVoicemail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLanguage

`func (o *UserRelationVoicemailCreateVoicemail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserRelationVoicemailCreateVoicemail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserRelationVoicemailCreateVoicemail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserRelationVoicemailCreateVoicemail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMaxMessages

`func (o *UserRelationVoicemailCreateVoicemail) GetMaxMessages() int32`

GetMaxMessages returns the MaxMessages field if non-nil, zero value otherwise.

### GetMaxMessagesOk

`func (o *UserRelationVoicemailCreateVoicemail) GetMaxMessagesOk() (*int32, bool)`

GetMaxMessagesOk returns a tuple with the MaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessages

`func (o *UserRelationVoicemailCreateVoicemail) SetMaxMessages(v int32)`

SetMaxMessages sets MaxMessages field to given value.

### HasMaxMessages

`func (o *UserRelationVoicemailCreateVoicemail) HasMaxMessages() bool`

HasMaxMessages returns a boolean if a field has been set.

### GetNumber

`func (o *UserRelationVoicemailCreateVoicemail) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UserRelationVoicemailCreateVoicemail) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UserRelationVoicemailCreateVoicemail) SetNumber(v string)`

SetNumber sets Number field to given value.

### GetOptions

`func (o *UserRelationVoicemailCreateVoicemail) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserRelationVoicemailCreateVoicemail) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserRelationVoicemailCreateVoicemail) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserRelationVoicemailCreateVoicemail) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPager

`func (o *UserRelationVoicemailCreateVoicemail) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *UserRelationVoicemailCreateVoicemail) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *UserRelationVoicemailCreateVoicemail) SetPager(v string)`

SetPager sets Pager field to given value.

### HasPager

`func (o *UserRelationVoicemailCreateVoicemail) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPassword

`func (o *UserRelationVoicemailCreateVoicemail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRelationVoicemailCreateVoicemail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRelationVoicemailCreateVoicemail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserRelationVoicemailCreateVoicemail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTenantUuid

`func (o *UserRelationVoicemailCreateVoicemail) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *UserRelationVoicemailCreateVoicemail) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *UserRelationVoicemailCreateVoicemail) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *UserRelationVoicemailCreateVoicemail) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *UserRelationVoicemailCreateVoicemail) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserRelationVoicemailCreateVoicemail) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserRelationVoicemailCreateVoicemail) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserRelationVoicemailCreateVoicemail) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetId

`func (o *UserRelationVoicemailCreateVoicemail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRelationVoicemailCreateVoicemail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRelationVoicemailCreateVoicemail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserRelationVoicemailCreateVoicemail) HasId() bool`

HasId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
