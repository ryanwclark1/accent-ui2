# LocationLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | the SIP contact to use when multiple registers are used | [optional]
**LineId** | **int32** | the ID of the line where the relocated call should be connected |

## Methods

### NewLocationLine

`func NewLocationLine(lineId int32, ) *LocationLine`

NewLocationLine instantiates a new LocationLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationLineWithDefaults

`func NewLocationLineWithDefaults() *LocationLine`

NewLocationLineWithDefaults instantiates a new LocationLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *LocationLine) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *LocationLine) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *LocationLine) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *LocationLine) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetLineId

`func (o *LocationLine) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *LocationLine) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *LocationLine) SetLineId(v int32)`

SetLineId sets LineId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
