# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | Pointer to **string** | The name to identify the skill | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Agents** | Pointer to [**[]SkillRelationAgent**](SkillRelationAgent.md) |  | [optional] [readonly]
**Description** | Pointer to **string** |  | [optional]

## Methods

### NewSkill

`func NewSkill() *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Skill) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Skill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Skill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Skill) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Skill) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Skill) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Skill) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetAgents

`func (o *Skill) GetAgents() []SkillRelationAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *Skill) GetAgentsOk() (*[]SkillRelationAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *Skill) SetAgents(v []SkillRelationAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *Skill) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetDescription

`func (o *Skill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Skill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Skill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Skill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
