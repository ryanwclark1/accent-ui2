# Paging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callers** | Pointer to [**PagingRelationMemberUsers**](PagingRelationMemberUsers.md) |  | [optional]
**Members** | Pointer to [**PagingRelationMemberUsers**](PagingRelationMemberUsers.md) |  | [optional]
**AnnounceCaller** | Pointer to **bool** | Play the announce sound to the caller | [optional] [default to true]
**AnnounceSound** | Pointer to **string** | The sound played to everyone | [optional]
**CallerNotification** | Pointer to **bool** | Play a notification to caller | [optional] [default to true]
**Duplex** | Pointer to **bool** | Duplex audio | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Id** | Pointer to **int32** |  | [optional] [readonly]
**IgnoreForward** | Pointer to **bool** | Ignore attemps to forward the call | [optional] [default to false]
**Name** | Pointer to **string** | The name to identify the paging | [optional]
**Number** | **string** | The number of the paging |
**Record** | Pointer to **bool** | Record the paging | [optional] [default to false]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewPaging

`func NewPaging(number string, ) *Paging`

NewPaging instantiates a new Paging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingWithDefaults

`func NewPagingWithDefaults() *Paging`

NewPagingWithDefaults instantiates a new Paging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallers

`func (o *Paging) GetCallers() PagingRelationMemberUsers`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *Paging) GetCallersOk() (*PagingRelationMemberUsers, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *Paging) SetCallers(v PagingRelationMemberUsers)`

SetCallers sets Callers field to given value.

### HasCallers

`func (o *Paging) HasCallers() bool`

HasCallers returns a boolean if a field has been set.

### GetMembers

`func (o *Paging) GetMembers() PagingRelationMemberUsers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Paging) GetMembersOk() (*PagingRelationMemberUsers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Paging) SetMembers(v PagingRelationMemberUsers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Paging) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAnnounceCaller

`func (o *Paging) GetAnnounceCaller() bool`

GetAnnounceCaller returns the AnnounceCaller field if non-nil, zero value otherwise.

### GetAnnounceCallerOk

`func (o *Paging) GetAnnounceCallerOk() (*bool, bool)`

GetAnnounceCallerOk returns a tuple with the AnnounceCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceCaller

`func (o *Paging) SetAnnounceCaller(v bool)`

SetAnnounceCaller sets AnnounceCaller field to given value.

### HasAnnounceCaller

`func (o *Paging) HasAnnounceCaller() bool`

HasAnnounceCaller returns a boolean if a field has been set.

### GetAnnounceSound

`func (o *Paging) GetAnnounceSound() string`

GetAnnounceSound returns the AnnounceSound field if non-nil, zero value otherwise.

### GetAnnounceSoundOk

`func (o *Paging) GetAnnounceSoundOk() (*string, bool)`

GetAnnounceSoundOk returns a tuple with the AnnounceSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceSound

`func (o *Paging) SetAnnounceSound(v string)`

SetAnnounceSound sets AnnounceSound field to given value.

### HasAnnounceSound

`func (o *Paging) HasAnnounceSound() bool`

HasAnnounceSound returns a boolean if a field has been set.

### GetCallerNotification

`func (o *Paging) GetCallerNotification() bool`

GetCallerNotification returns the CallerNotification field if non-nil, zero value otherwise.

### GetCallerNotificationOk

`func (o *Paging) GetCallerNotificationOk() (*bool, bool)`

GetCallerNotificationOk returns a tuple with the CallerNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerNotification

`func (o *Paging) SetCallerNotification(v bool)`

SetCallerNotification sets CallerNotification field to given value.

### HasCallerNotification

`func (o *Paging) HasCallerNotification() bool`

HasCallerNotification returns a boolean if a field has been set.

### GetDuplex

`func (o *Paging) GetDuplex() bool`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *Paging) GetDuplexOk() (*bool, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *Paging) SetDuplex(v bool)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *Paging) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEnabled

`func (o *Paging) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Paging) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Paging) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Paging) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *Paging) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Paging) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Paging) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Paging) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreForward

`func (o *Paging) GetIgnoreForward() bool`

GetIgnoreForward returns the IgnoreForward field if non-nil, zero value otherwise.

### GetIgnoreForwardOk

`func (o *Paging) GetIgnoreForwardOk() (*bool, bool)`

GetIgnoreForwardOk returns a tuple with the IgnoreForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreForward

`func (o *Paging) SetIgnoreForward(v bool)`

SetIgnoreForward sets IgnoreForward field to given value.

### HasIgnoreForward

`func (o *Paging) HasIgnoreForward() bool`

HasIgnoreForward returns a boolean if a field has been set.

### GetName

`func (o *Paging) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Paging) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Paging) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Paging) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *Paging) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Paging) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Paging) SetNumber(v string)`

SetNumber sets Number field to given value.

### GetRecord

`func (o *Paging) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Paging) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Paging) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *Paging) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Paging) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Paging) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Paging) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Paging) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
