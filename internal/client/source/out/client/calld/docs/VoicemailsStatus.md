# VoicemailsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StatusValue**](StatusValue.md) |  | [optional]
**CacheItems** | Pointer to **int32** |  | [optional]

## Methods

### NewVoicemailsStatus

`func NewVoicemailsStatus() *VoicemailsStatus`

NewVoicemailsStatus instantiates a new VoicemailsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailsStatusWithDefaults

`func NewVoicemailsStatusWithDefaults() *VoicemailsStatus`

NewVoicemailsStatusWithDefaults instantiates a new VoicemailsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VoicemailsStatus) GetStatus() StatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoicemailsStatus) GetStatusOk() (*StatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoicemailsStatus) SetStatus(v StatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoicemailsStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCacheItems

`func (o *VoicemailsStatus) GetCacheItems() int32`

GetCacheItems returns the CacheItems field if non-nil, zero value otherwise.

### GetCacheItemsOk

`func (o *VoicemailsStatus) GetCacheItemsOk() (*int32, bool)`

GetCacheItemsOk returns a tuple with the CacheItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheItems

`func (o *VoicemailsStatus) SetCacheItems(v int32)`

SetCacheItems sets CacheItems field to given value.

### HasCacheItems

`func (o *VoicemailsStatus) HasCacheItems() bool`

HasCacheItems returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
