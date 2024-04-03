# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletable** | Pointer to **bool** |  | [optional]
**Id** | Pointer to **string** | The unique configuration id | [optional]
**ParentIds** | Pointer to **[]string** |  | [optional]
**RawConfig** | Pointer to [**RawConfigurationObject**](RawConfigurationObject.md) |  | [optional]

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletable

`func (o *Config) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *Config) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *Config) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *Config) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetId

`func (o *Config) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Config) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Config) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Config) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentIds

`func (o *Config) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *Config) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *Config) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.

### HasParentIds

`func (o *Config) HasParentIds() bool`

HasParentIds returns a boolean if a field has been set.

### GetRawConfig

`func (o *Config) GetRawConfig() RawConfigurationObject`

GetRawConfig returns the RawConfig field if non-nil, zero value otherwise.

### GetRawConfigOk

`func (o *Config) GetRawConfigOk() (*RawConfigurationObject, bool)`

GetRawConfigOk returns a tuple with the RawConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawConfig

`func (o *Config) SetRawConfig(v RawConfigurationObject)`

SetRawConfig sets RawConfig field to given value.

### HasRawConfig

`func (o *Config) HasRawConfig() bool`

HasRawConfig returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
