# DestinationApplicationVoicemail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | **string** | MUST be &#39;voicemail&#39; |
**Context** | **int32** | The context of the application |

## Methods

### NewDestinationApplicationVoicemail

`func NewDestinationApplicationVoicemail(type_ string, application string, context int32, ) *DestinationApplicationVoicemail`

NewDestinationApplicationVoicemail instantiates a new DestinationApplicationVoicemail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationVoicemailWithDefaults

`func NewDestinationApplicationVoicemailWithDefaults() *DestinationApplicationVoicemail`

NewDestinationApplicationVoicemailWithDefaults instantiates a new DestinationApplicationVoicemail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationVoicemail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationVoicemail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationVoicemail) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationVoicemail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationVoicemail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationVoicemail) SetApplication(v string)`

SetApplication sets Application field to given value.

### GetContext

`func (o *DestinationApplicationVoicemail) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationApplicationVoicemail) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationApplicationVoicemail) SetContext(v int32)`

SetContext sets Context field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
