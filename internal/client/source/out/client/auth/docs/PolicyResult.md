# PolicyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **[]string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Name** | **string** |  |
**Shared** | Pointer to **bool** | Should be shared to sub-tenants or not. Cannot be changed after creation When shared is &#x60;true&#x60;, then all tenants below this policy&#39;s tenant will see it as their own policy with the attribute &#x60;read_only: true&#x60;.  Using &#x60;shared&#x60; attribute will add uniqueness constraints for the slug among all policies&#39; sub-tenants.  | [optional]
**Slug** | Pointer to **string** | A unique, human readable identifier for this policy | [optional]
**ReadOnly** | Pointer to **bool** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewPolicyResult

`func NewPolicyResult(name string, ) *PolicyResult`

NewPolicyResult instantiates a new PolicyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResultWithDefaults

`func NewPolicyResultWithDefaults() *PolicyResult`

NewPolicyResultWithDefaults instantiates a new PolicyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PolicyResult) GetAcl() []string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PolicyResult) GetAclOk() (*[]string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PolicyResult) SetAcl(v []string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PolicyResult) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyResult) SetName(v string)`

SetName sets Name field to given value.

### GetShared

`func (o *PolicyResult) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *PolicyResult) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *PolicyResult) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *PolicyResult) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSlug

`func (o *PolicyResult) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PolicyResult) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PolicyResult) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PolicyResult) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetReadOnly

`func (o *PolicyResult) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PolicyResult) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PolicyResult) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PolicyResult) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetUuid

`func (o *PolicyResult) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PolicyResult) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PolicyResult) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PolicyResult) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
