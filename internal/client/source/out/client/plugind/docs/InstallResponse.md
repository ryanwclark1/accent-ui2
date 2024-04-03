# InstallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | A UUID associated to this plugin installation | [optional]

## Methods

### NewInstallResponse

`func NewInstallResponse() *InstallResponse`

NewInstallResponse instantiates a new InstallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallResponseWithDefaults

`func NewInstallResponseWithDefaults() *InstallResponse`

NewInstallResponseWithDefaults instantiates a new InstallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *InstallResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstallResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstallResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstallResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
