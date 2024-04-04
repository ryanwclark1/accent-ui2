# ParkingLotItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ParkingLot**](ParkingLot.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewParkingLotItems

`func NewParkingLotItems(total int32, ) *ParkingLotItems`

NewParkingLotItems instantiates a new ParkingLotItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParkingLotItemsWithDefaults

`func NewParkingLotItemsWithDefaults() *ParkingLotItems`

NewParkingLotItemsWithDefaults instantiates a new ParkingLotItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ParkingLotItems) GetItems() []ParkingLot`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ParkingLotItems) GetItemsOk() (*[]ParkingLot, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ParkingLotItems) SetItems(v []ParkingLot)`

SetItems sets Items field to given value.

### HasItems

`func (o *ParkingLotItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *ParkingLotItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ParkingLotItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ParkingLotItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
