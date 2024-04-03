# GroupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**ReadOnly** | Pointer to **bool** |  | [optional]
**Slug** | Pointer to **string** |  | [optional]
**SystemManaged** | Pointer to **bool** | *Deprecated* Please use &#x60;read_only&#x60; | [optional]
**TenantUuid** | Pointer to **string** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewGroupResult

`func NewGroupResult() *GroupResult`

NewGroupResult instantiates a new GroupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResultWithDefaults

`func NewGroupResultWithDefaults() *GroupResult`

NewGroupResultWithDefaults instantiates a new GroupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadOnly

`func (o *GroupResult) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GroupResult) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GroupResult) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GroupResult) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSlug

`func (o *GroupResult) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GroupResult) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GroupResult) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GroupResult) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSystemManaged

`func (o *GroupResult) GetSystemManaged() bool`

GetSystemManaged returns the SystemManaged field if non-nil, zero value otherwise.

### GetSystemManagedOk

`func (o *GroupResult) GetSystemManagedOk() (*bool, bool)`

GetSystemManagedOk returns a tuple with the SystemManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManaged

`func (o *GroupResult) SetSystemManaged(v bool)`

SetSystemManaged sets SystemManaged field to given value.

### HasSystemManaged

`func (o *GroupResult) HasSystemManaged() bool`

HasSystemManaged returns a boolean if a field has been set.

### GetTenantUuid

`func (o *GroupResult) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *GroupResult) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *GroupResult) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *GroupResult) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUuid

`func (o *GroupResult) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupResult) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupResult) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GroupResult) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
