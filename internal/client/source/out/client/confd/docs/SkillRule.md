# SkillRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | **string** | The name to identify the skill rule |
**Rules** | Pointer to [**[]SkillRuleRule**](SkillRuleRule.md) |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewSkillRule

`func NewSkillRule(name string, ) *SkillRule`

NewSkillRule instantiates a new SkillRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillRuleWithDefaults

`func NewSkillRuleWithDefaults() *SkillRule`

NewSkillRuleWithDefaults instantiates a new SkillRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SkillRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SkillRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillRule) SetName(v string)`

SetName sets Name field to given value.

### GetRules

`func (o *SkillRule) GetRules() []SkillRuleRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SkillRule) GetRulesOk() (*[]SkillRuleRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SkillRule) SetRules(v []SkillRuleRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SkillRule) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTenantUuid

`func (o *SkillRule) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *SkillRule) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *SkillRule) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *SkillRule) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
