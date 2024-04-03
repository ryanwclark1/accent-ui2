# MeetingUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**Persistent** | Pointer to **bool** | Persistent meetings will not get deleted automatically | [optional]

## Methods

### NewMeetingUserRequest

`func NewMeetingUserRequest() *MeetingUserRequest`

NewMeetingUserRequest instantiates a new MeetingUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingUserRequestWithDefaults

`func NewMeetingUserRequestWithDefaults() *MeetingUserRequest`

NewMeetingUserRequestWithDefaults instantiates a new MeetingUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MeetingUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeetingUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersistent

`func (o *MeetingUserRequest) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *MeetingUserRequest) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *MeetingUserRequest) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *MeetingUserRequest) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
