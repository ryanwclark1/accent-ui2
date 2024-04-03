# DestinationOutcall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exten** | **string** |  |
**OutcallId** | **int32** | The id of the outcall |
**Type** | **string** | MUST be &#39;outcall&#39; |

## Methods

### NewDestinationOutcall

`func NewDestinationOutcall(exten string, outcallId int32, type_ string, ) *DestinationOutcall`

NewDestinationOutcall instantiates a new DestinationOutcall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationOutcallWithDefaults

`func NewDestinationOutcallWithDefaults() *DestinationOutcall`

NewDestinationOutcallWithDefaults instantiates a new DestinationOutcall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExten

`func (o *DestinationOutcall) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *DestinationOutcall) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *DestinationOutcall) SetExten(v string)`

SetExten sets Exten field to given value.

### GetOutcallId

`func (o *DestinationOutcall) GetOutcallId() int32`

GetOutcallId returns the OutcallId field if non-nil, zero value otherwise.

### GetOutcallIdOk

`func (o *DestinationOutcall) GetOutcallIdOk() (*int32, bool)`

GetOutcallIdOk returns a tuple with the OutcallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcallId

`func (o *DestinationOutcall) SetOutcallId(v int32)`

SetOutcallId sets OutcallId field to given value.

### GetType

`func (o *DestinationOutcall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationOutcall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationOutcall) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
