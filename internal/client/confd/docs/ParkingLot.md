# ParkingLot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | Pointer to **string** | name to identify the parking lot | [optional]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional] [default to "default"]
**SlotsEnd** | Pointer to **string** | Ending extension to park calls | [optional]
**SlotsStart** | Pointer to **string** | Starting extension to park calls | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timeout** | Pointer to **int32** | Maximum time on parking | [optional]

## Methods

### NewParkingLot

`func NewParkingLot() *ParkingLot`

NewParkingLot instantiates a new ParkingLot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParkingLotWithDefaults

`func NewParkingLotWithDefaults() *ParkingLot`

NewParkingLotWithDefaults instantiates a new ParkingLot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParkingLot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParkingLot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParkingLot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParkingLot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ParkingLot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParkingLot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParkingLot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParkingLot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExtensions

`func (o *ParkingLot) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ParkingLot) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ParkingLot) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ParkingLot) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ParkingLot) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ParkingLot) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ParkingLot) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ParkingLot) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetSlotsEnd

`func (o *ParkingLot) GetSlotsEnd() string`

GetSlotsEnd returns the SlotsEnd field if non-nil, zero value otherwise.

### GetSlotsEndOk

`func (o *ParkingLot) GetSlotsEndOk() (*string, bool)`

GetSlotsEndOk returns a tuple with the SlotsEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotsEnd

`func (o *ParkingLot) SetSlotsEnd(v string)`

SetSlotsEnd sets SlotsEnd field to given value.

### HasSlotsEnd

`func (o *ParkingLot) HasSlotsEnd() bool`

HasSlotsEnd returns a boolean if a field has been set.

### GetSlotsStart

`func (o *ParkingLot) GetSlotsStart() string`

GetSlotsStart returns the SlotsStart field if non-nil, zero value otherwise.

### GetSlotsStartOk

`func (o *ParkingLot) GetSlotsStartOk() (*string, bool)`

GetSlotsStartOk returns a tuple with the SlotsStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotsStart

`func (o *ParkingLot) SetSlotsStart(v string)`

SetSlotsStart sets SlotsStart field to given value.

### HasSlotsStart

`func (o *ParkingLot) HasSlotsStart() bool`

HasSlotsStart returns a boolean if a field has been set.

### GetTenantUuid

`func (o *ParkingLot) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *ParkingLot) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *ParkingLot) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *ParkingLot) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *ParkingLot) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ParkingLot) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ParkingLot) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ParkingLot) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
