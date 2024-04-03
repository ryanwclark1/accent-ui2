# CallPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the call permission | [optional] [readonly]
**Name** | **string** | The name of the call permission |
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Description** | Pointer to **string** | Additional information about the call permission | [optional]
**Enabled** | Pointer to **bool** | Disable or enable the call permission | [optional] [default to true]
**Extensions** | Pointer to **[]string** | Extensions to apply the call permission | [optional]
**Mode** | Pointer to **string** | Determine how the call permission are applied | [optional] [default to "deny"]
**Password** | Pointer to **string** | The password to be able to call extensions | [optional]

## Methods

### NewCallPermission

`func NewCallPermission(name string, ) *CallPermission`

NewCallPermission instantiates a new CallPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPermissionWithDefaults

`func NewCallPermissionWithDefaults() *CallPermission`

NewCallPermissionWithDefaults instantiates a new CallPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallPermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CallPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CallPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallPermission) SetName(v string)`

SetName sets Name field to given value.

### GetTenantUuid

`func (o *CallPermission) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *CallPermission) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *CallPermission) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *CallPermission) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetDescription

`func (o *CallPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CallPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CallPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CallPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CallPermission) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CallPermission) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CallPermission) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CallPermission) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtensions

`func (o *CallPermission) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CallPermission) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CallPermission) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CallPermission) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMode

`func (o *CallPermission) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CallPermission) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CallPermission) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CallPermission) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPassword

`func (o *CallPermission) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CallPermission) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CallPermission) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CallPermission) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
