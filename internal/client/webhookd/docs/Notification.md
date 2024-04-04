# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  |
**Extra** | Pointer to **map[string]interface{}** |  | [optional]
**NotificationType** | **string** | A name without special characters to differenciate the notification from others |
**Title** | **string** |  |
**UserUuid** | **string** |  |

## Methods

### NewNotification

`func NewNotification(body string, notificationType string, title string, userUuid string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Notification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Notification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Notification) SetBody(v string)`

SetBody sets Body field to given value.

### GetExtra

`func (o *Notification) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Notification) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Notification) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Notification) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetNotificationType

`func (o *Notification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *Notification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *Notification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### GetTitle

`func (o *Notification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Notification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Notification) SetTitle(v string)`

SetTitle sets Title field to given value.

### GetUserUuid

`func (o *Notification) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *Notification) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *Notification) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
