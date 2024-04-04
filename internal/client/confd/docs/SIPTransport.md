# SIPTransport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**Options** | **[][]string** |  |
**Uuid** | Pointer to **string** |  | [optional] [readonly]

## Methods

### NewSIPTransport

`func NewSIPTransport(name string, options [][]string, ) *SIPTransport`

NewSIPTransport instantiates a new SIPTransport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIPTransportWithDefaults

`func NewSIPTransportWithDefaults() *SIPTransport`

NewSIPTransportWithDefaults instantiates a new SIPTransport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SIPTransport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SIPTransport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SIPTransport) SetName(v string)`

SetName sets Name field to given value.

### GetOptions

`func (o *SIPTransport) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SIPTransport) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SIPTransport) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### GetUuid

`func (o *SIPTransport) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SIPTransport) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SIPTransport) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SIPTransport) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
