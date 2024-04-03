# DestinationCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command to execute |
**Type** | **string** | MUST be &#39;custom&#39; |

## Methods

### NewDestinationCustom

`func NewDestinationCustom(command string, type_ string, ) *DestinationCustom`

NewDestinationCustom instantiates a new DestinationCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationCustomWithDefaults

`func NewDestinationCustomWithDefaults() *DestinationCustom`

NewDestinationCustomWithDefaults instantiates a new DestinationCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *DestinationCustom) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DestinationCustom) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DestinationCustom) SetCommand(v string)`

SetCommand sets Command field to given value.

### GetType

`func (o *DestinationCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationCustom) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
