# DestinationConference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceId** | **int32** | The id of the conference |
**Type** | **string** | MUST be &#39;conference&#39; |

## Methods

### NewDestinationConference

`func NewDestinationConference(conferenceId int32, type_ string, ) *DestinationConference`

NewDestinationConference instantiates a new DestinationConference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationConferenceWithDefaults

`func NewDestinationConferenceWithDefaults() *DestinationConference`

NewDestinationConferenceWithDefaults instantiates a new DestinationConference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConferenceId

`func (o *DestinationConference) GetConferenceId() int32`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *DestinationConference) GetConferenceIdOk() (*int32, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *DestinationConference) SetConferenceId(v int32)`

SetConferenceId sets ConferenceId field to given value.

### GetType

`func (o *DestinationConference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationConference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationConference) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
