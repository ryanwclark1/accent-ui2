# SkillItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Skill**](Skill.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewSkillItems

`func NewSkillItems(total int32, ) *SkillItems`

NewSkillItems instantiates a new SkillItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillItemsWithDefaults

`func NewSkillItemsWithDefaults() *SkillItems`

NewSkillItemsWithDefaults instantiates a new SkillItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SkillItems) GetItems() []Skill`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SkillItems) GetItemsOk() (*[]Skill, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SkillItems) SetItems(v []Skill)`

SetItems sets Items field to given value.

### HasItems

`func (o *SkillItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *SkillItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SkillItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SkillItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
