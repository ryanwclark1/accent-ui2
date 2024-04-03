# VoicemailZoneMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message that each user hears. | [optional]
**Name** | **string** | label on which each voicemail is mapped |
**Timezone** | **string** | Timezone as define in &#39;/usr/share/zoneinfo/&#39; |

## Methods

### NewVoicemailZoneMessage

`func NewVoicemailZoneMessage(name string, timezone string, ) *VoicemailZoneMessage`

NewVoicemailZoneMessage instantiates a new VoicemailZoneMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailZoneMessageWithDefaults

`func NewVoicemailZoneMessageWithDefaults() *VoicemailZoneMessage`

NewVoicemailZoneMessageWithDefaults instantiates a new VoicemailZoneMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VoicemailZoneMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VoicemailZoneMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VoicemailZoneMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VoicemailZoneMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *VoicemailZoneMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoicemailZoneMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoicemailZoneMessage) SetName(v string)`

SetName sets Name field to given value.

### GetTimezone

`func (o *VoicemailZoneMessage) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *VoicemailZoneMessage) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *VoicemailZoneMessage) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
