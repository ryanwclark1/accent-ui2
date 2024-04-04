# SetupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineInstanceUuid** | Pointer to **string** | The UUID identifying this instance on Nestbox. The engine_instance_uuid should only be specified if the instance has already been registered on the specified Nestbox. Omitting this field for an instance that is already registered will create a duplicate entry on the Nestbox.  | [optional]
**EngineInternalAddress** | Pointer to **string** | IP address of the engine | [optional]
**EngineLanguage** | **string** | The interface language for the Accent Engine (either &#x60;en_US&#x60; or &#x60;fr_FR&#x60;) |
**EngineLicense** | **bool** | Whether the GNU GPLv3 license is accepted |
**EnginePassword** | **string** | Password of the first administrator &#x60;&#x60;root&#x60;&#x60; on the engine |
**EngineRtpIcesupport** | Pointer to **bool** | Enable ICE support. This is required for WebRTC. A STUN server must be defined in the &#x60;engine_rtp_stunaddr&#x60; field when using &#x60;engine_rtp_icesupport&#x3D;true&#x60;.  | [optional] [default to false]
**EngineRtpStunaddr** | Pointer to **string** | The address of the STUN server to use for WebRTC | [optional] [default to "null"]
**NestboxEngineHost** | Pointer to **string** | Host used by Nestbox to contact the engine | [optional]
**NestboxEnginePort** | Pointer to **int32** | Port used by Nestbox to contact the engine | [optional]
**NestboxHost** | Pointer to **string** | Host of the Nestbox where the engine will register. Specifying this key will make nestbox and &#x60;engine_internal_address&#x60; keys mandatory. Accent will be connected to the specified Nestbox instance. | [optional]
**NestboxInstanceName** | Pointer to **string** | Name of the engine in Nestbox | [optional]
**NestboxPort** | Pointer to **int32** | Port of the Nestbox where the engine will register | [optional]
**NestboxServiceId** | Pointer to **string** | Nestbox username used to register the engine | [optional]
**NestboxServiceKey** | Pointer to **string** | Nestbox password used to register the engine | [optional]
**NestboxVerifyCertificate** | Pointer to **bool** | Should the certificate used for HTTPS be verified? The setup will abort if the certificate fails the verification. | [optional] [default to true]

## Methods

### NewSetupRequest

`func NewSetupRequest(engineLanguage string, engineLicense bool, enginePassword string, ) *SetupRequest`

NewSetupRequest instantiates a new SetupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupRequestWithDefaults

`func NewSetupRequestWithDefaults() *SetupRequest`

NewSetupRequestWithDefaults instantiates a new SetupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineInstanceUuid

`func (o *SetupRequest) GetEngineInstanceUuid() string`

GetEngineInstanceUuid returns the EngineInstanceUuid field if non-nil, zero value otherwise.

### GetEngineInstanceUuidOk

`func (o *SetupRequest) GetEngineInstanceUuidOk() (*string, bool)`

GetEngineInstanceUuidOk returns a tuple with the EngineInstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineInstanceUuid

`func (o *SetupRequest) SetEngineInstanceUuid(v string)`

SetEngineInstanceUuid sets EngineInstanceUuid field to given value.

### HasEngineInstanceUuid

`func (o *SetupRequest) HasEngineInstanceUuid() bool`

HasEngineInstanceUuid returns a boolean if a field has been set.

### GetEngineInternalAddress

`func (o *SetupRequest) GetEngineInternalAddress() string`

GetEngineInternalAddress returns the EngineInternalAddress field if non-nil, zero value otherwise.

### GetEngineInternalAddressOk

`func (o *SetupRequest) GetEngineInternalAddressOk() (*string, bool)`

GetEngineInternalAddressOk returns a tuple with the EngineInternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineInternalAddress

`func (o *SetupRequest) SetEngineInternalAddress(v string)`

SetEngineInternalAddress sets EngineInternalAddress field to given value.

### HasEngineInternalAddress

`func (o *SetupRequest) HasEngineInternalAddress() bool`

HasEngineInternalAddress returns a boolean if a field has been set.

### GetEngineLanguage

`func (o *SetupRequest) GetEngineLanguage() string`

GetEngineLanguage returns the EngineLanguage field if non-nil, zero value otherwise.

### GetEngineLanguageOk

`func (o *SetupRequest) GetEngineLanguageOk() (*string, bool)`

GetEngineLanguageOk returns a tuple with the EngineLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineLanguage

`func (o *SetupRequest) SetEngineLanguage(v string)`

SetEngineLanguage sets EngineLanguage field to given value.

### GetEngineLicense

`func (o *SetupRequest) GetEngineLicense() bool`

GetEngineLicense returns the EngineLicense field if non-nil, zero value otherwise.

### GetEngineLicenseOk

`func (o *SetupRequest) GetEngineLicenseOk() (*bool, bool)`

GetEngineLicenseOk returns a tuple with the EngineLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineLicense

`func (o *SetupRequest) SetEngineLicense(v bool)`

SetEngineLicense sets EngineLicense field to given value.

### GetEnginePassword

`func (o *SetupRequest) GetEnginePassword() string`

GetEnginePassword returns the EnginePassword field if non-nil, zero value otherwise.

### GetEnginePasswordOk

`func (o *SetupRequest) GetEnginePasswordOk() (*string, bool)`

GetEnginePasswordOk returns a tuple with the EnginePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnginePassword

`func (o *SetupRequest) SetEnginePassword(v string)`

SetEnginePassword sets EnginePassword field to given value.

### GetEngineRtpIcesupport

`func (o *SetupRequest) GetEngineRtpIcesupport() bool`

GetEngineRtpIcesupport returns the EngineRtpIcesupport field if non-nil, zero value otherwise.

### GetEngineRtpIcesupportOk

`func (o *SetupRequest) GetEngineRtpIcesupportOk() (*bool, bool)`

GetEngineRtpIcesupportOk returns a tuple with the EngineRtpIcesupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineRtpIcesupport

`func (o *SetupRequest) SetEngineRtpIcesupport(v bool)`

SetEngineRtpIcesupport sets EngineRtpIcesupport field to given value.

### HasEngineRtpIcesupport

`func (o *SetupRequest) HasEngineRtpIcesupport() bool`

HasEngineRtpIcesupport returns a boolean if a field has been set.

### GetEngineRtpStunaddr

`func (o *SetupRequest) GetEngineRtpStunaddr() string`

GetEngineRtpStunaddr returns the EngineRtpStunaddr field if non-nil, zero value otherwise.

### GetEngineRtpStunaddrOk

`func (o *SetupRequest) GetEngineRtpStunaddrOk() (*string, bool)`

GetEngineRtpStunaddrOk returns a tuple with the EngineRtpStunaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineRtpStunaddr

`func (o *SetupRequest) SetEngineRtpStunaddr(v string)`

SetEngineRtpStunaddr sets EngineRtpStunaddr field to given value.

### HasEngineRtpStunaddr

`func (o *SetupRequest) HasEngineRtpStunaddr() bool`

HasEngineRtpStunaddr returns a boolean if a field has been set.

### GetNestboxEngineHost

`func (o *SetupRequest) GetNestboxEngineHost() string`

GetNestboxEngineHost returns the NestboxEngineHost field if non-nil, zero value otherwise.

### GetNestboxEngineHostOk

`func (o *SetupRequest) GetNestboxEngineHostOk() (*string, bool)`

GetNestboxEngineHostOk returns a tuple with the NestboxEngineHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxEngineHost

`func (o *SetupRequest) SetNestboxEngineHost(v string)`

SetNestboxEngineHost sets NestboxEngineHost field to given value.

### HasNestboxEngineHost

`func (o *SetupRequest) HasNestboxEngineHost() bool`

HasNestboxEngineHost returns a boolean if a field has been set.

### GetNestboxEnginePort

`func (o *SetupRequest) GetNestboxEnginePort() int32`

GetNestboxEnginePort returns the NestboxEnginePort field if non-nil, zero value otherwise.

### GetNestboxEnginePortOk

`func (o *SetupRequest) GetNestboxEnginePortOk() (*int32, bool)`

GetNestboxEnginePortOk returns a tuple with the NestboxEnginePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxEnginePort

`func (o *SetupRequest) SetNestboxEnginePort(v int32)`

SetNestboxEnginePort sets NestboxEnginePort field to given value.

### HasNestboxEnginePort

`func (o *SetupRequest) HasNestboxEnginePort() bool`

HasNestboxEnginePort returns a boolean if a field has been set.

### GetNestboxHost

`func (o *SetupRequest) GetNestboxHost() string`

GetNestboxHost returns the NestboxHost field if non-nil, zero value otherwise.

### GetNestboxHostOk

`func (o *SetupRequest) GetNestboxHostOk() (*string, bool)`

GetNestboxHostOk returns a tuple with the NestboxHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxHost

`func (o *SetupRequest) SetNestboxHost(v string)`

SetNestboxHost sets NestboxHost field to given value.

### HasNestboxHost

`func (o *SetupRequest) HasNestboxHost() bool`

HasNestboxHost returns a boolean if a field has been set.

### GetNestboxInstanceName

`func (o *SetupRequest) GetNestboxInstanceName() string`

GetNestboxInstanceName returns the NestboxInstanceName field if non-nil, zero value otherwise.

### GetNestboxInstanceNameOk

`func (o *SetupRequest) GetNestboxInstanceNameOk() (*string, bool)`

GetNestboxInstanceNameOk returns a tuple with the NestboxInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxInstanceName

`func (o *SetupRequest) SetNestboxInstanceName(v string)`

SetNestboxInstanceName sets NestboxInstanceName field to given value.

### HasNestboxInstanceName

`func (o *SetupRequest) HasNestboxInstanceName() bool`

HasNestboxInstanceName returns a boolean if a field has been set.

### GetNestboxPort

`func (o *SetupRequest) GetNestboxPort() int32`

GetNestboxPort returns the NestboxPort field if non-nil, zero value otherwise.

### GetNestboxPortOk

`func (o *SetupRequest) GetNestboxPortOk() (*int32, bool)`

GetNestboxPortOk returns a tuple with the NestboxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxPort

`func (o *SetupRequest) SetNestboxPort(v int32)`

SetNestboxPort sets NestboxPort field to given value.

### HasNestboxPort

`func (o *SetupRequest) HasNestboxPort() bool`

HasNestboxPort returns a boolean if a field has been set.

### GetNestboxServiceId

`func (o *SetupRequest) GetNestboxServiceId() string`

GetNestboxServiceId returns the NestboxServiceId field if non-nil, zero value otherwise.

### GetNestboxServiceIdOk

`func (o *SetupRequest) GetNestboxServiceIdOk() (*string, bool)`

GetNestboxServiceIdOk returns a tuple with the NestboxServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxServiceId

`func (o *SetupRequest) SetNestboxServiceId(v string)`

SetNestboxServiceId sets NestboxServiceId field to given value.

### HasNestboxServiceId

`func (o *SetupRequest) HasNestboxServiceId() bool`

HasNestboxServiceId returns a boolean if a field has been set.

### GetNestboxServiceKey

`func (o *SetupRequest) GetNestboxServiceKey() string`

GetNestboxServiceKey returns the NestboxServiceKey field if non-nil, zero value otherwise.

### GetNestboxServiceKeyOk

`func (o *SetupRequest) GetNestboxServiceKeyOk() (*string, bool)`

GetNestboxServiceKeyOk returns a tuple with the NestboxServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxServiceKey

`func (o *SetupRequest) SetNestboxServiceKey(v string)`

SetNestboxServiceKey sets NestboxServiceKey field to given value.

### HasNestboxServiceKey

`func (o *SetupRequest) HasNestboxServiceKey() bool`

HasNestboxServiceKey returns a boolean if a field has been set.

### GetNestboxVerifyCertificate

`func (o *SetupRequest) GetNestboxVerifyCertificate() bool`

GetNestboxVerifyCertificate returns the NestboxVerifyCertificate field if non-nil, zero value otherwise.

### GetNestboxVerifyCertificateOk

`func (o *SetupRequest) GetNestboxVerifyCertificateOk() (*bool, bool)`

GetNestboxVerifyCertificateOk returns a tuple with the NestboxVerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestboxVerifyCertificate

`func (o *SetupRequest) SetNestboxVerifyCertificate(v bool)`

SetNestboxVerifyCertificate sets NestboxVerifyCertificate field to given value.

### HasNestboxVerifyCertificate

`func (o *SetupRequest) HasNestboxVerifyCertificate() bool`

HasNestboxVerifyCertificate returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
