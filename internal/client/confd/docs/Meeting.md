# Meeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly]
**Exten** | Pointer to **string** | the external extension to dial to reach this meeting | [optional] [readonly]
**GuestSipAuthorization** | Pointer to **string** | Format: base64(username:password), same as HTTP Basic Auth. | [optional] [readonly]
**IngressHttpUri** | Pointer to **string** | URI to reach this stack (configured by the Ingress HTTP resource) | [optional]
**Name** | Pointer to **string** |  | [optional]
**OwnerUuids** | Pointer to **[]string** |  | [optional]
**Persistent** | Pointer to **bool** | Persistent meetings will not get deleted automatically | [optional]
**RequireAuthorization** | Pointer to **bool** | when &#x60;true&#x60;, the &#x60;guest_sip_authorization&#x60; is always &#x60;null&#x60;. Instead, clients must request an authorization to access the meeting. | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewMeeting

`func NewMeeting() *Meeting`

NewMeeting instantiates a new Meeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingWithDefaults

`func NewMeetingWithDefaults() *Meeting`

NewMeetingWithDefaults instantiates a new Meeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Meeting) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Meeting) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Meeting) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Meeting) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetExten

`func (o *Meeting) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *Meeting) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *Meeting) SetExten(v string)`

SetExten sets Exten field to given value.

### HasExten

`func (o *Meeting) HasExten() bool`

HasExten returns a boolean if a field has been set.

### GetGuestSipAuthorization

`func (o *Meeting) GetGuestSipAuthorization() string`

GetGuestSipAuthorization returns the GuestSipAuthorization field if non-nil, zero value otherwise.

### GetGuestSipAuthorizationOk

`func (o *Meeting) GetGuestSipAuthorizationOk() (*string, bool)`

GetGuestSipAuthorizationOk returns a tuple with the GuestSipAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSipAuthorization

`func (o *Meeting) SetGuestSipAuthorization(v string)`

SetGuestSipAuthorization sets GuestSipAuthorization field to given value.

### HasGuestSipAuthorization

`func (o *Meeting) HasGuestSipAuthorization() bool`

HasGuestSipAuthorization returns a boolean if a field has been set.

### GetIngressHttpUri

`func (o *Meeting) GetIngressHttpUri() string`

GetIngressHttpUri returns the IngressHttpUri field if non-nil, zero value otherwise.

### GetIngressHttpUriOk

`func (o *Meeting) GetIngressHttpUriOk() (*string, bool)`

GetIngressHttpUriOk returns a tuple with the IngressHttpUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressHttpUri

`func (o *Meeting) SetIngressHttpUri(v string)`

SetIngressHttpUri sets IngressHttpUri field to given value.

### HasIngressHttpUri

`func (o *Meeting) HasIngressHttpUri() bool`

HasIngressHttpUri returns a boolean if a field has been set.

### GetName

`func (o *Meeting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Meeting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Meeting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Meeting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerUuids

`func (o *Meeting) GetOwnerUuids() []string`

GetOwnerUuids returns the OwnerUuids field if non-nil, zero value otherwise.

### GetOwnerUuidsOk

`func (o *Meeting) GetOwnerUuidsOk() (*[]string, bool)`

GetOwnerUuidsOk returns a tuple with the OwnerUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUuids

`func (o *Meeting) SetOwnerUuids(v []string)`

SetOwnerUuids sets OwnerUuids field to given value.

### HasOwnerUuids

`func (o *Meeting) HasOwnerUuids() bool`

HasOwnerUuids returns a boolean if a field has been set.

### GetPersistent

`func (o *Meeting) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *Meeting) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *Meeting) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *Meeting) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetRequireAuthorization

`func (o *Meeting) GetRequireAuthorization() bool`

GetRequireAuthorization returns the RequireAuthorization field if non-nil, zero value otherwise.

### GetRequireAuthorizationOk

`func (o *Meeting) GetRequireAuthorizationOk() (*bool, bool)`

GetRequireAuthorizationOk returns a tuple with the RequireAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthorization

`func (o *Meeting) SetRequireAuthorization(v bool)`

SetRequireAuthorization sets RequireAuthorization field to given value.

### HasRequireAuthorization

`func (o *Meeting) HasRequireAuthorization() bool`

HasRequireAuthorization returns a boolean if a field has been set.

### GetUuid

`func (o *Meeting) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Meeting) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Meeting) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Meeting) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
