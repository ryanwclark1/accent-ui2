# AdhocConferenceCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostCallId** | Pointer to **string** | The call_id of the host call | [optional]
**ParticipantCallIds** | Pointer to **[]string** | The call_id of the participating calls in this conference, excluding the host. | [optional]

## Methods

### NewAdhocConferenceCreation

`func NewAdhocConferenceCreation() *AdhocConferenceCreation`

NewAdhocConferenceCreation instantiates a new AdhocConferenceCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdhocConferenceCreationWithDefaults

`func NewAdhocConferenceCreationWithDefaults() *AdhocConferenceCreation`

NewAdhocConferenceCreationWithDefaults instantiates a new AdhocConferenceCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostCallId

`func (o *AdhocConferenceCreation) GetHostCallId() string`

GetHostCallId returns the HostCallId field if non-nil, zero value otherwise.

### GetHostCallIdOk

`func (o *AdhocConferenceCreation) GetHostCallIdOk() (*string, bool)`

GetHostCallIdOk returns a tuple with the HostCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCallId

`func (o *AdhocConferenceCreation) SetHostCallId(v string)`

SetHostCallId sets HostCallId field to given value.

### HasHostCallId

`func (o *AdhocConferenceCreation) HasHostCallId() bool`

HasHostCallId returns a boolean if a field has been set.

### GetParticipantCallIds

`func (o *AdhocConferenceCreation) GetParticipantCallIds() []string`

GetParticipantCallIds returns the ParticipantCallIds field if non-nil, zero value otherwise.

### GetParticipantCallIdsOk

`func (o *AdhocConferenceCreation) GetParticipantCallIdsOk() (*[]string, bool)`

GetParticipantCallIdsOk returns a tuple with the ParticipantCallIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCallIds

`func (o *AdhocConferenceCreation) SetParticipantCallIds(v []string)`

SetParticipantCallIds sets ParticipantCallIds field to given value.

### HasParticipantCallIds

`func (o *AdhocConferenceCreation) HasParticipantCallIds() bool`

HasParticipantCallIds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
