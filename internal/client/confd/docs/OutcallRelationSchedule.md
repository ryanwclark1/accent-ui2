# OutcallRelationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | Pointer to **string** | The name to identify the schedule | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewOutcallRelationSchedule

`func NewOutcallRelationSchedule() *OutcallRelationSchedule`

NewOutcallRelationSchedule instantiates a new OutcallRelationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcallRelationScheduleWithDefaults

`func NewOutcallRelationScheduleWithDefaults() *OutcallRelationSchedule`

NewOutcallRelationScheduleWithDefaults instantiates a new OutcallRelationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutcallRelationSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutcallRelationSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutcallRelationSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OutcallRelationSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OutcallRelationSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutcallRelationSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutcallRelationSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OutcallRelationSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *OutcallRelationSchedule) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *OutcallRelationSchedule) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *OutcallRelationSchedule) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *OutcallRelationSchedule) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
