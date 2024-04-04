# CDR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **time.Time** |  | [optional]
**Answered** | Pointer to **bool** |  | [optional]
**CallDirection** | Pointer to **string** |  | [optional]
**DestinationDetails** | Pointer to **map[string]interface{}** | Contains the &#x60;type&#x60; of the called destination; which can be either &#x60;user&#x60;, &#x60;conference&#x60;, &#x60;meeting&#x60;, or &#x60;unknown&#x60; by default. Also contains useful information about the destination (&#x60;id&#x60; and &#x60;name&#x60;). | [optional]
**DestinationExtension** | Pointer to **string** |  | [optional]
**DestinationInternalContext** | Pointer to **string** |  | [optional]
**DestinationInternalExtension** | Pointer to **string** | the internal extension of the line that answers | [optional]
**DestinationInternalTenantUuid** | Pointer to **string** |  | [optional]
**DestinationLineId** | Pointer to **int32** |  | [optional]
**DestinationName** | Pointer to **string** |  | [optional]
**DestinationTenantUuid** | Pointer to **string** |  | [optional]
**DestinationUserUuid** | Pointer to **string** |  | [optional]
**Duration** | Pointer to **int32** | Duration of the call, in seconds. | [optional]
**End** | Pointer to **time.Time** |  | [optional]
**Id** | Pointer to **int32** |  | [optional]
**Recordings** | Pointer to [**[]Recording**](Recording.md) |  | [optional]
**RequestedContext** | Pointer to **string** |  | [optional]
**RequestedExtension** | Pointer to **string** |  | [optional]
**RequestedInternalContext** | Pointer to **string** |  | [optional]
**RequestedInternalExtension** | Pointer to **string** | the internal extension of the first line to ring | [optional]
**RequestedInternalTenantUuid** | Pointer to **string** |  | [optional]
**RequestedName** | Pointer to **string** |  | [optional]
**RequestedTenantUuid** | Pointer to **string** |  | [optional]
**SourceExtension** | Pointer to **string** |  | [optional]
**SourceInternalContext** | Pointer to **string** |  | [optional]
**SourceInternalExtension** | Pointer to **string** | the internal extension of the line that placed the call | [optional]
**SourceInternalName** | Pointer to **string** |  | [optional]
**SourceInternalTenantUuid** | Pointer to **string** |  | [optional]
**SourceLineId** | Pointer to **int32** |  | [optional]
**SourceName** | Pointer to **string** |  | [optional]
**SourceTenantUuid** | Pointer to **string** |  | [optional]
**SourceUserUuid** | Pointer to **string** |  | [optional]
**Start** | Pointer to **time.Time** |  | [optional]
**Tags** | Pointer to **[]string** |  | [optional]

## Methods

### NewCDR

`func NewCDR() *CDR`

NewCDR instantiates a new CDR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDRWithDefaults

`func NewCDRWithDefaults() *CDR`

NewCDRWithDefaults instantiates a new CDR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *CDR) GetAnswer() time.Time`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *CDR) GetAnswerOk() (*time.Time, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *CDR) SetAnswer(v time.Time)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *CDR) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAnswered

`func (o *CDR) GetAnswered() bool`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *CDR) GetAnsweredOk() (*bool, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *CDR) SetAnswered(v bool)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *CDR) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetCallDirection

`func (o *CDR) GetCallDirection() string`

GetCallDirection returns the CallDirection field if non-nil, zero value otherwise.

### GetCallDirectionOk

`func (o *CDR) GetCallDirectionOk() (*string, bool)`

GetCallDirectionOk returns a tuple with the CallDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDirection

`func (o *CDR) SetCallDirection(v string)`

SetCallDirection sets CallDirection field to given value.

### HasCallDirection

`func (o *CDR) HasCallDirection() bool`

HasCallDirection returns a boolean if a field has been set.

### GetDestinationDetails

`func (o *CDR) GetDestinationDetails() map[string]interface{}`

GetDestinationDetails returns the DestinationDetails field if non-nil, zero value otherwise.

### GetDestinationDetailsOk

`func (o *CDR) GetDestinationDetailsOk() (*map[string]interface{}, bool)`

GetDestinationDetailsOk returns a tuple with the DestinationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDetails

`func (o *CDR) SetDestinationDetails(v map[string]interface{})`

SetDestinationDetails sets DestinationDetails field to given value.

### HasDestinationDetails

`func (o *CDR) HasDestinationDetails() bool`

HasDestinationDetails returns a boolean if a field has been set.

### GetDestinationExtension

`func (o *CDR) GetDestinationExtension() string`

GetDestinationExtension returns the DestinationExtension field if non-nil, zero value otherwise.

### GetDestinationExtensionOk

`func (o *CDR) GetDestinationExtensionOk() (*string, bool)`

GetDestinationExtensionOk returns a tuple with the DestinationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationExtension

`func (o *CDR) SetDestinationExtension(v string)`

SetDestinationExtension sets DestinationExtension field to given value.

### HasDestinationExtension

`func (o *CDR) HasDestinationExtension() bool`

HasDestinationExtension returns a boolean if a field has been set.

### GetDestinationInternalContext

`func (o *CDR) GetDestinationInternalContext() string`

GetDestinationInternalContext returns the DestinationInternalContext field if non-nil, zero value otherwise.

### GetDestinationInternalContextOk

`func (o *CDR) GetDestinationInternalContextOk() (*string, bool)`

GetDestinationInternalContextOk returns a tuple with the DestinationInternalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationInternalContext

`func (o *CDR) SetDestinationInternalContext(v string)`

SetDestinationInternalContext sets DestinationInternalContext field to given value.

### HasDestinationInternalContext

`func (o *CDR) HasDestinationInternalContext() bool`

HasDestinationInternalContext returns a boolean if a field has been set.

### GetDestinationInternalExtension

`func (o *CDR) GetDestinationInternalExtension() string`

GetDestinationInternalExtension returns the DestinationInternalExtension field if non-nil, zero value otherwise.

### GetDestinationInternalExtensionOk

`func (o *CDR) GetDestinationInternalExtensionOk() (*string, bool)`

GetDestinationInternalExtensionOk returns a tuple with the DestinationInternalExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationInternalExtension

`func (o *CDR) SetDestinationInternalExtension(v string)`

SetDestinationInternalExtension sets DestinationInternalExtension field to given value.

### HasDestinationInternalExtension

`func (o *CDR) HasDestinationInternalExtension() bool`

HasDestinationInternalExtension returns a boolean if a field has been set.

### GetDestinationInternalTenantUuid

`func (o *CDR) GetDestinationInternalTenantUuid() string`

GetDestinationInternalTenantUuid returns the DestinationInternalTenantUuid field if non-nil, zero value otherwise.

### GetDestinationInternalTenantUuidOk

`func (o *CDR) GetDestinationInternalTenantUuidOk() (*string, bool)`

GetDestinationInternalTenantUuidOk returns a tuple with the DestinationInternalTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationInternalTenantUuid

`func (o *CDR) SetDestinationInternalTenantUuid(v string)`

SetDestinationInternalTenantUuid sets DestinationInternalTenantUuid field to given value.

### HasDestinationInternalTenantUuid

`func (o *CDR) HasDestinationInternalTenantUuid() bool`

HasDestinationInternalTenantUuid returns a boolean if a field has been set.

### GetDestinationLineId

`func (o *CDR) GetDestinationLineId() int32`

GetDestinationLineId returns the DestinationLineId field if non-nil, zero value otherwise.

### GetDestinationLineIdOk

`func (o *CDR) GetDestinationLineIdOk() (*int32, bool)`

GetDestinationLineIdOk returns a tuple with the DestinationLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLineId

`func (o *CDR) SetDestinationLineId(v int32)`

SetDestinationLineId sets DestinationLineId field to given value.

### HasDestinationLineId

`func (o *CDR) HasDestinationLineId() bool`

HasDestinationLineId returns a boolean if a field has been set.

### GetDestinationName

`func (o *CDR) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *CDR) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *CDR) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *CDR) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDestinationTenantUuid

`func (o *CDR) GetDestinationTenantUuid() string`

GetDestinationTenantUuid returns the DestinationTenantUuid field if non-nil, zero value otherwise.

### GetDestinationTenantUuidOk

`func (o *CDR) GetDestinationTenantUuidOk() (*string, bool)`

GetDestinationTenantUuidOk returns a tuple with the DestinationTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTenantUuid

`func (o *CDR) SetDestinationTenantUuid(v string)`

SetDestinationTenantUuid sets DestinationTenantUuid field to given value.

### HasDestinationTenantUuid

`func (o *CDR) HasDestinationTenantUuid() bool`

HasDestinationTenantUuid returns a boolean if a field has been set.

### GetDestinationUserUuid

`func (o *CDR) GetDestinationUserUuid() string`

GetDestinationUserUuid returns the DestinationUserUuid field if non-nil, zero value otherwise.

### GetDestinationUserUuidOk

`func (o *CDR) GetDestinationUserUuidOk() (*string, bool)`

GetDestinationUserUuidOk returns a tuple with the DestinationUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUserUuid

`func (o *CDR) SetDestinationUserUuid(v string)`

SetDestinationUserUuid sets DestinationUserUuid field to given value.

### HasDestinationUserUuid

`func (o *CDR) HasDestinationUserUuid() bool`

HasDestinationUserUuid returns a boolean if a field has been set.

### GetDuration

`func (o *CDR) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CDR) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CDR) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CDR) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *CDR) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CDR) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CDR) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CDR) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *CDR) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CDR) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CDR) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CDR) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordings

`func (o *CDR) GetRecordings() []Recording`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *CDR) GetRecordingsOk() (*[]Recording, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *CDR) SetRecordings(v []Recording)`

SetRecordings sets Recordings field to given value.

### HasRecordings

`func (o *CDR) HasRecordings() bool`

HasRecordings returns a boolean if a field has been set.

### GetRequestedContext

`func (o *CDR) GetRequestedContext() string`

GetRequestedContext returns the RequestedContext field if non-nil, zero value otherwise.

### GetRequestedContextOk

`func (o *CDR) GetRequestedContextOk() (*string, bool)`

GetRequestedContextOk returns a tuple with the RequestedContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedContext

`func (o *CDR) SetRequestedContext(v string)`

SetRequestedContext sets RequestedContext field to given value.

### HasRequestedContext

`func (o *CDR) HasRequestedContext() bool`

HasRequestedContext returns a boolean if a field has been set.

### GetRequestedExtension

`func (o *CDR) GetRequestedExtension() string`

GetRequestedExtension returns the RequestedExtension field if non-nil, zero value otherwise.

### GetRequestedExtensionOk

`func (o *CDR) GetRequestedExtensionOk() (*string, bool)`

GetRequestedExtensionOk returns a tuple with the RequestedExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedExtension

`func (o *CDR) SetRequestedExtension(v string)`

SetRequestedExtension sets RequestedExtension field to given value.

### HasRequestedExtension

`func (o *CDR) HasRequestedExtension() bool`

HasRequestedExtension returns a boolean if a field has been set.

### GetRequestedInternalContext

`func (o *CDR) GetRequestedInternalContext() string`

GetRequestedInternalContext returns the RequestedInternalContext field if non-nil, zero value otherwise.

### GetRequestedInternalContextOk

`func (o *CDR) GetRequestedInternalContextOk() (*string, bool)`

GetRequestedInternalContextOk returns a tuple with the RequestedInternalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInternalContext

`func (o *CDR) SetRequestedInternalContext(v string)`

SetRequestedInternalContext sets RequestedInternalContext field to given value.

### HasRequestedInternalContext

`func (o *CDR) HasRequestedInternalContext() bool`

HasRequestedInternalContext returns a boolean if a field has been set.

### GetRequestedInternalExtension

`func (o *CDR) GetRequestedInternalExtension() string`

GetRequestedInternalExtension returns the RequestedInternalExtension field if non-nil, zero value otherwise.

### GetRequestedInternalExtensionOk

`func (o *CDR) GetRequestedInternalExtensionOk() (*string, bool)`

GetRequestedInternalExtensionOk returns a tuple with the RequestedInternalExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInternalExtension

`func (o *CDR) SetRequestedInternalExtension(v string)`

SetRequestedInternalExtension sets RequestedInternalExtension field to given value.

### HasRequestedInternalExtension

`func (o *CDR) HasRequestedInternalExtension() bool`

HasRequestedInternalExtension returns a boolean if a field has been set.

### GetRequestedInternalTenantUuid

`func (o *CDR) GetRequestedInternalTenantUuid() string`

GetRequestedInternalTenantUuid returns the RequestedInternalTenantUuid field if non-nil, zero value otherwise.

### GetRequestedInternalTenantUuidOk

`func (o *CDR) GetRequestedInternalTenantUuidOk() (*string, bool)`

GetRequestedInternalTenantUuidOk returns a tuple with the RequestedInternalTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInternalTenantUuid

`func (o *CDR) SetRequestedInternalTenantUuid(v string)`

SetRequestedInternalTenantUuid sets RequestedInternalTenantUuid field to given value.

### HasRequestedInternalTenantUuid

`func (o *CDR) HasRequestedInternalTenantUuid() bool`

HasRequestedInternalTenantUuid returns a boolean if a field has been set.

### GetRequestedName

`func (o *CDR) GetRequestedName() string`

GetRequestedName returns the RequestedName field if non-nil, zero value otherwise.

### GetRequestedNameOk

`func (o *CDR) GetRequestedNameOk() (*string, bool)`

GetRequestedNameOk returns a tuple with the RequestedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedName

`func (o *CDR) SetRequestedName(v string)`

SetRequestedName sets RequestedName field to given value.

### HasRequestedName

`func (o *CDR) HasRequestedName() bool`

HasRequestedName returns a boolean if a field has been set.

### GetRequestedTenantUuid

`func (o *CDR) GetRequestedTenantUuid() string`

GetRequestedTenantUuid returns the RequestedTenantUuid field if non-nil, zero value otherwise.

### GetRequestedTenantUuidOk

`func (o *CDR) GetRequestedTenantUuidOk() (*string, bool)`

GetRequestedTenantUuidOk returns a tuple with the RequestedTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTenantUuid

`func (o *CDR) SetRequestedTenantUuid(v string)`

SetRequestedTenantUuid sets RequestedTenantUuid field to given value.

### HasRequestedTenantUuid

`func (o *CDR) HasRequestedTenantUuid() bool`

HasRequestedTenantUuid returns a boolean if a field has been set.

### GetSourceExtension

`func (o *CDR) GetSourceExtension() string`

GetSourceExtension returns the SourceExtension field if non-nil, zero value otherwise.

### GetSourceExtensionOk

`func (o *CDR) GetSourceExtensionOk() (*string, bool)`

GetSourceExtensionOk returns a tuple with the SourceExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceExtension

`func (o *CDR) SetSourceExtension(v string)`

SetSourceExtension sets SourceExtension field to given value.

### HasSourceExtension

`func (o *CDR) HasSourceExtension() bool`

HasSourceExtension returns a boolean if a field has been set.

### GetSourceInternalContext

`func (o *CDR) GetSourceInternalContext() string`

GetSourceInternalContext returns the SourceInternalContext field if non-nil, zero value otherwise.

### GetSourceInternalContextOk

`func (o *CDR) GetSourceInternalContextOk() (*string, bool)`

GetSourceInternalContextOk returns a tuple with the SourceInternalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInternalContext

`func (o *CDR) SetSourceInternalContext(v string)`

SetSourceInternalContext sets SourceInternalContext field to given value.

### HasSourceInternalContext

`func (o *CDR) HasSourceInternalContext() bool`

HasSourceInternalContext returns a boolean if a field has been set.

### GetSourceInternalExtension

`func (o *CDR) GetSourceInternalExtension() string`

GetSourceInternalExtension returns the SourceInternalExtension field if non-nil, zero value otherwise.

### GetSourceInternalExtensionOk

`func (o *CDR) GetSourceInternalExtensionOk() (*string, bool)`

GetSourceInternalExtensionOk returns a tuple with the SourceInternalExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInternalExtension

`func (o *CDR) SetSourceInternalExtension(v string)`

SetSourceInternalExtension sets SourceInternalExtension field to given value.

### HasSourceInternalExtension

`func (o *CDR) HasSourceInternalExtension() bool`

HasSourceInternalExtension returns a boolean if a field has been set.

### GetSourceInternalName

`func (o *CDR) GetSourceInternalName() string`

GetSourceInternalName returns the SourceInternalName field if non-nil, zero value otherwise.

### GetSourceInternalNameOk

`func (o *CDR) GetSourceInternalNameOk() (*string, bool)`

GetSourceInternalNameOk returns a tuple with the SourceInternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInternalName

`func (o *CDR) SetSourceInternalName(v string)`

SetSourceInternalName sets SourceInternalName field to given value.

### HasSourceInternalName

`func (o *CDR) HasSourceInternalName() bool`

HasSourceInternalName returns a boolean if a field has been set.

### GetSourceInternalTenantUuid

`func (o *CDR) GetSourceInternalTenantUuid() string`

GetSourceInternalTenantUuid returns the SourceInternalTenantUuid field if non-nil, zero value otherwise.

### GetSourceInternalTenantUuidOk

`func (o *CDR) GetSourceInternalTenantUuidOk() (*string, bool)`

GetSourceInternalTenantUuidOk returns a tuple with the SourceInternalTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInternalTenantUuid

`func (o *CDR) SetSourceInternalTenantUuid(v string)`

SetSourceInternalTenantUuid sets SourceInternalTenantUuid field to given value.

### HasSourceInternalTenantUuid

`func (o *CDR) HasSourceInternalTenantUuid() bool`

HasSourceInternalTenantUuid returns a boolean if a field has been set.

### GetSourceLineId

`func (o *CDR) GetSourceLineId() int32`

GetSourceLineId returns the SourceLineId field if non-nil, zero value otherwise.

### GetSourceLineIdOk

`func (o *CDR) GetSourceLineIdOk() (*int32, bool)`

GetSourceLineIdOk returns a tuple with the SourceLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLineId

`func (o *CDR) SetSourceLineId(v int32)`

SetSourceLineId sets SourceLineId field to given value.

### HasSourceLineId

`func (o *CDR) HasSourceLineId() bool`

HasSourceLineId returns a boolean if a field has been set.

### GetSourceName

`func (o *CDR) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *CDR) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *CDR) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *CDR) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceTenantUuid

`func (o *CDR) GetSourceTenantUuid() string`

GetSourceTenantUuid returns the SourceTenantUuid field if non-nil, zero value otherwise.

### GetSourceTenantUuidOk

`func (o *CDR) GetSourceTenantUuidOk() (*string, bool)`

GetSourceTenantUuidOk returns a tuple with the SourceTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTenantUuid

`func (o *CDR) SetSourceTenantUuid(v string)`

SetSourceTenantUuid sets SourceTenantUuid field to given value.

### HasSourceTenantUuid

`func (o *CDR) HasSourceTenantUuid() bool`

HasSourceTenantUuid returns a boolean if a field has been set.

### GetSourceUserUuid

`func (o *CDR) GetSourceUserUuid() string`

GetSourceUserUuid returns the SourceUserUuid field if non-nil, zero value otherwise.

### GetSourceUserUuidOk

`func (o *CDR) GetSourceUserUuidOk() (*string, bool)`

GetSourceUserUuidOk returns a tuple with the SourceUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUserUuid

`func (o *CDR) SetSourceUserUuid(v string)`

SetSourceUserUuid sets SourceUserUuid field to given value.

### HasSourceUserUuid

`func (o *CDR) HasSourceUserUuid() bool`

HasSourceUserUuid returns a boolean if a field has been set.

### GetStart

`func (o *CDR) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CDR) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CDR) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *CDR) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *CDR) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CDR) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CDR) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CDR) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
