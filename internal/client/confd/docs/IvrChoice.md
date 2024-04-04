# IvrChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**DestinationType**](DestinationType.md) |  |
**Exten** | **string** | The extension the caller has to dial to select this choice. Can be an extension pattern |

## Methods

### NewIvrChoice

`func NewIvrChoice(destination DestinationType, exten string, ) *IvrChoice`

NewIvrChoice instantiates a new IvrChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIvrChoiceWithDefaults

`func NewIvrChoiceWithDefaults() *IvrChoice`

NewIvrChoiceWithDefaults instantiates a new IvrChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *IvrChoice) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *IvrChoice) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *IvrChoice) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.

### GetExten

`func (o *IvrChoice) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *IvrChoice) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *IvrChoice) SetExten(v string)`

SetExten sets Exten field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
