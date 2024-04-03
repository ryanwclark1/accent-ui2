# Voicemail

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
**Id** | Pointer to **int32** | Voicemail ID | [optional] [readonly]
**Name** | Pointer to **string** | Voicemail name | [optional]
**Users** | Pointer to [**[]UserRelationBase**](UserRelationBase.md) |  | [optional] [readonly]

## Methods

### NewVoicemail

`func NewVoicemail(context string, number string, ) *Voicemail`

NewVoicemail instantiates a new Voicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailWithDefaults

`func NewVoicemailWithDefaults() *Voicemail`

NewVoicemailWithDefaults instantiates a new Voicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPassword

`func (o *Voicemail) GetAskPassword() bool`

GetAskPassword returns the AskPassword field if non-nil, zero value otherwise.

### GetAskPasswordOk

`func (o *Voicemail) GetAskPasswordOk() (*bool, bool)`

GetAskPasswordOk returns a tuple with the AskPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPassword

`func (o *Voicemail) SetAskPassword(v bool)`

SetAskPassword sets AskPassword field to given value.

### HasAskPassword

`func (o *Voicemail) HasAskPassword() bool`

HasAskPassword returns a boolean if a field has been set.

### GetAttachAudio

`func (o *Voicemail) GetAttachAudio() bool`

GetAttachAudio returns the AttachAudio field if non-nil, zero value otherwise.

### GetAttachAudioOk

`func (o *Voicemail) GetAttachAudioOk() (*bool, bool)`

GetAttachAudioOk returns a tuple with the AttachAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAudio

`func (o *Voicemail) SetAttachAudio(v bool)`

SetAttachAudio sets AttachAudio field to given value.

### HasAttachAudio

`func (o *Voicemail) HasAttachAudio() bool`

HasAttachAudio returns a boolean if a field has been set.

### GetContext

`func (o *Voicemail) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Voicemail) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Voicemail) SetContext(v string)`

SetContext sets Context field to given value.

### GetDeleteMessages

`func (o *Voicemail) GetDeleteMessages() bool`

GetDeleteMessages returns the DeleteMessages field if non-nil, zero value otherwise.

### GetDeleteMessagesOk

`func (o *Voicemail) GetDeleteMessagesOk() (*bool, bool)`

GetDeleteMessagesOk returns a tuple with the DeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessages

`func (o *Voicemail) SetDeleteMessages(v bool)`

SetDeleteMessages sets DeleteMessages field to given value.

### HasDeleteMessages

`func (o *Voicemail) HasDeleteMessages() bool`

HasDeleteMessages returns a boolean if a field has been set.

### GetEmail

`func (o *Voicemail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Voicemail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Voicemail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Voicemail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *Voicemail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Voicemail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Voicemail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Voicemail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLanguage

`func (o *Voicemail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Voicemail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Voicemail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Voicemail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMaxMessages

`func (o *Voicemail) GetMaxMessages() int32`

GetMaxMessages returns the MaxMessages field if non-nil, zero value otherwise.

### GetMaxMessagesOk

`func (o *Voicemail) GetMaxMessagesOk() (*int32, bool)`

GetMaxMessagesOk returns a tuple with the MaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessages

`func (o *Voicemail) SetMaxMessages(v int32)`

SetMaxMessages sets MaxMessages field to given value.

### HasMaxMessages

`func (o *Voicemail) HasMaxMessages() bool`

HasMaxMessages returns a boolean if a field has been set.

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

### GetOptions

`func (o *Voicemail) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Voicemail) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Voicemail) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Voicemail) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPager

`func (o *Voicemail) GetPager() string`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *Voicemail) GetPagerOk() (*string, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *Voicemail) SetPager(v string)`

SetPager sets Pager field to given value.

### HasPager

`func (o *Voicemail) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPassword

`func (o *Voicemail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Voicemail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Voicemail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Voicemail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Voicemail) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Voicemail) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Voicemail) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Voicemail) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *Voicemail) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Voicemail) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Voicemail) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Voicemail) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

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

### GetUsers

`func (o *Voicemail) GetUsers() []UserRelationBase`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Voicemail) GetUsersOk() (*[]UserRelationBase, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Voicemail) SetUsers(v []UserRelationBase)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Voicemail) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
