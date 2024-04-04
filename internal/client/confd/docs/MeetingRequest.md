# MeetingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**Persistent** | Pointer to **bool** | Persistent meetings will not get deleted automatically | [optional]
**OwnerUuids** | Pointer to **[]string** |  | [optional]

## Methods

### NewMeetingRequest

`func NewMeetingRequest() *MeetingRequest`

NewMeetingRequest instantiates a new MeetingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingRequestWithDefaults

`func NewMeetingRequestWithDefaults() *MeetingRequest`

NewMeetingRequestWithDefaults instantiates a new MeetingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MeetingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeetingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersistent

`func (o *MeetingRequest) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *MeetingRequest) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *MeetingRequest) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *MeetingRequest) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetOwnerUuids

`func (o *MeetingRequest) GetOwnerUuids() []string`

GetOwnerUuids returns the OwnerUuids field if non-nil, zero value otherwise.

### GetOwnerUuidsOk

`func (o *MeetingRequest) GetOwnerUuidsOk() (*[]string, bool)`

GetOwnerUuidsOk returns a tuple with the OwnerUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUuids

`func (o *MeetingRequest) SetOwnerUuids(v []string)`

SetOwnerUuids sets OwnerUuids field to given value.

### HasOwnerUuids

`func (o *MeetingRequest) HasOwnerUuids() bool`

HasOwnerUuids returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
