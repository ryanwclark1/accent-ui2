# DestinationSwitchboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchboardUuid** | **string** | The UUID of the switchboard. |
**Type** | **string** | MUST be &#39;switchboard&#39; |

## Methods

### NewDestinationSwitchboard

`func NewDestinationSwitchboard(switchboardUuid string, type_ string, ) *DestinationSwitchboard`

NewDestinationSwitchboard instantiates a new DestinationSwitchboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSwitchboardWithDefaults

`func NewDestinationSwitchboardWithDefaults() *DestinationSwitchboard`

NewDestinationSwitchboardWithDefaults instantiates a new DestinationSwitchboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchboardUuid

`func (o *DestinationSwitchboard) GetSwitchboardUuid() string`

GetSwitchboardUuid returns the SwitchboardUuid field if non-nil, zero value otherwise.

### GetSwitchboardUuidOk

`func (o *DestinationSwitchboard) GetSwitchboardUuidOk() (*string, bool)`

GetSwitchboardUuidOk returns a tuple with the SwitchboardUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboardUuid

`func (o *DestinationSwitchboard) SetSwitchboardUuid(v string)`

SetSwitchboardUuid sets SwitchboardUuid field to given value.

### GetType

`func (o *DestinationSwitchboard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationSwitchboard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationSwitchboard) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
