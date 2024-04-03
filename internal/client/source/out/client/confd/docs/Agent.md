# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** |  | [optional]
**Id** | Pointer to **int32** | The id of the agent | [optional] [readonly]
**Lastname** | Pointer to **string** |  | [optional]
**Number** | Pointer to **string** | Agent number. Cannot be modified after creation | [optional]
**Queues** | Pointer to [**[]AgentRelationQueue**](AgentRelationQueue.md) |  | [optional] [readonly]
**Users** | Pointer to [**[]AgentRelationUser**](AgentRelationUser.md) |  | [optional] [readonly]
**Skills** | Pointer to [**[]AgentRelationSkill**](AgentRelationSkill.md) |  | [optional] [readonly]
**Description** | Pointer to **string** | Additional information about the agent | [optional]
**Language** | Pointer to **string** | Language used for the agent menu prompt | [optional]
**Password** | Pointer to **string** | Numeric password used to log agent. | [optional]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *Agent) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Agent) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Agent) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Agent) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetId

`func (o *Agent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastname

`func (o *Agent) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Agent) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Agent) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Agent) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetNumber

`func (o *Agent) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Agent) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Agent) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Agent) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetQueues

`func (o *Agent) GetQueues() []AgentRelationQueue`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *Agent) GetQueuesOk() (*[]AgentRelationQueue, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *Agent) SetQueues(v []AgentRelationQueue)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *Agent) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetUsers

`func (o *Agent) GetUsers() []AgentRelationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Agent) GetUsersOk() (*[]AgentRelationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Agent) SetUsers(v []AgentRelationUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Agent) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetSkills

`func (o *Agent) GetSkills() []AgentRelationSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *Agent) GetSkillsOk() (*[]AgentRelationSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *Agent) SetSkills(v []AgentRelationSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *Agent) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguage

`func (o *Agent) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Agent) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Agent) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Agent) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPassword

`func (o *Agent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Agent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Agent) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Agent) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Agent) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Agent) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Agent) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Agent) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Agent) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Agent) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Agent) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Agent) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
