# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAccentVersion** | Pointer to **string** | The maximum Accent version with which this plugin works | [optional]
**MinAccentVersion** | Pointer to **string** | The minimum Accent version with which this plugin works | [optional]
**Upgradable** | Pointer to **bool** | An indication wether installing this version would be an upgrade on not. Unstalled plugins are marked as upgradable. | [optional]
**Version** | Pointer to **string** | The plugin version | [optional]

## Methods

### NewVersionInfo

`func NewVersionInfo() *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAccentVersion

`func (o *VersionInfo) GetMaxAccentVersion() string`

GetMaxAccentVersion returns the MaxAccentVersion field if non-nil, zero value otherwise.

### GetMaxAccentVersionOk

`func (o *VersionInfo) GetMaxAccentVersionOk() (*string, bool)`

GetMaxAccentVersionOk returns a tuple with the MaxAccentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccentVersion

`func (o *VersionInfo) SetMaxAccentVersion(v string)`

SetMaxAccentVersion sets MaxAccentVersion field to given value.

### HasMaxAccentVersion

`func (o *VersionInfo) HasMaxAccentVersion() bool`

HasMaxAccentVersion returns a boolean if a field has been set.

### GetMinAccentVersion

`func (o *VersionInfo) GetMinAccentVersion() string`

GetMinAccentVersion returns the MinAccentVersion field if non-nil, zero value otherwise.

### GetMinAccentVersionOk

`func (o *VersionInfo) GetMinAccentVersionOk() (*string, bool)`

GetMinAccentVersionOk returns a tuple with the MinAccentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAccentVersion

`func (o *VersionInfo) SetMinAccentVersion(v string)`

SetMinAccentVersion sets MinAccentVersion field to given value.

### HasMinAccentVersion

`func (o *VersionInfo) HasMinAccentVersion() bool`

HasMinAccentVersion returns a boolean if a field has been set.

### GetUpgradable

`func (o *VersionInfo) GetUpgradable() bool`

GetUpgradable returns the Upgradable field if non-nil, zero value otherwise.

### GetUpgradableOk

`func (o *VersionInfo) GetUpgradableOk() (*bool, bool)`

GetUpgradableOk returns a tuple with the Upgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradable

`func (o *VersionInfo) SetUpgradable(v bool)`

SetUpgradable sets Upgradable field to given value.

### HasUpgradable

`func (o *VersionInfo) HasUpgradable() bool`

HasUpgradable returns a boolean if a field has been set.

### GetVersion

`func (o *VersionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
