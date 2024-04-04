# AgentRelationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]

## Methods

### NewAgentRelationUser

`func NewAgentRelationUser() *AgentRelationUser`

NewAgentRelationUser instantiates a new AgentRelationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRelationUserWithDefaults

`func NewAgentRelationUserWithDefaults() *AgentRelationUser`

NewAgentRelationUserWithDefaults instantiates a new AgentRelationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *AgentRelationUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *AgentRelationUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *AgentRelationUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *AgentRelationUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *AgentRelationUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *AgentRelationUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *AgentRelationUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *AgentRelationUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *AgentRelationUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentRelationUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentRelationUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentRelationUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
