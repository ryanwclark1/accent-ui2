# DestinationIVR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IvrId** | **int32** | The id of the IVR |
**Type** | **string** | MUST be &#39;ivr&#39; |

## Methods

### NewDestinationIVR

`func NewDestinationIVR(ivrId int32, type_ string, ) *DestinationIVR`

NewDestinationIVR instantiates a new DestinationIVR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationIVRWithDefaults

`func NewDestinationIVRWithDefaults() *DestinationIVR`

NewDestinationIVRWithDefaults instantiates a new DestinationIVR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIvrId

`func (o *DestinationIVR) GetIvrId() int32`

GetIvrId returns the IvrId field if non-nil, zero value otherwise.

### GetIvrIdOk

`func (o *DestinationIVR) GetIvrIdOk() (*int32, bool)`

GetIvrIdOk returns a tuple with the IvrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvrId

`func (o *DestinationIVR) SetIvrId(v int32)`

SetIvrId sets IvrId field to given value.

### GetType

`func (o *DestinationIVR) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationIVR) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationIVR) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
