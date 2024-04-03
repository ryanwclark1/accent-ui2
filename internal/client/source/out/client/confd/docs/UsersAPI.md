# \UsersAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateUserAgent**](UsersAPI.md#AssociateUserAgent) | **Put** /users/{user_id}/agents/{agent_id} | Associate user and agent
[**AssociateUserCallpermission**](UsersAPI.md#AssociateUserCallpermission) | **Put** /users/{user_id}/callpermissions/{callpermission_id} | Associate user and call permission
[**AssociateUserFuncKeyTemplate**](UsersAPI.md#AssociateUserFuncKeyTemplate) | **Put** /users/{user_id}/funckeys/templates/{template_id} | Associate a func key template to a user
[**AssociateUserLine**](UsersAPI.md#AssociateUserLine) | **Put** /users/{user_id}/lines/{line_id} | Associate user and line
[**AssociateUserLines**](UsersAPI.md#AssociateUserLines) | **Put** /users/{user_id}/lines | Associate user and lines
[**AssociateUserSchedule**](UsersAPI.md#AssociateUserSchedule) | **Put** /users/{user_id}/schedules/{schedule_id} | Associate user and schedule
[**AssociateUserVoicemail**](UsersAPI.md#AssociateUserVoicemail) | **Put** /users/{user_id}/voicemails/{voicemail_id} | Associate user and voicemail
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /users | Create user
[**CreateUserExternalApp**](UsersAPI.md#CreateUserExternalApp) | **Post** /users/{user_uuid}/external/apps/{app_name} | Create user external app
[**CreateUserMeeting**](UsersAPI.md#CreateUserMeeting) | **Post** /users/me/meetings | Create user meeting
[**CreateUserVoicemail**](UsersAPI.md#CreateUserVoicemail) | **Post** /users/{user_id}/voicemails | Create user voicemail
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{user_id} | Delete user
[**DeleteUserExternalApp**](UsersAPI.md#DeleteUserExternalApp) | **Delete** /users/{user_uuid}/external/apps/{app_name} | Delete user external app
[**DeleteUserFuncKey**](UsersAPI.md#DeleteUserFuncKey) | **Delete** /users/{user_id}/funckeys/{position} | Remove func key for user
[**DeleteUserMeeting**](UsersAPI.md#DeleteUserMeeting) | **Delete** /users/me/meetings/{meeting_uuid} | Delete one of the meetings of the current user
[**DissociateUserAgent**](UsersAPI.md#DissociateUserAgent) | **Delete** /users/{user_id}/agents | Dissociate user and agent
[**DissociateUserCallpermission**](UsersAPI.md#DissociateUserCallpermission) | **Delete** /users/{user_id}/callpermissions/{callpermission_id} | Dissociate user and call permission
[**DissociateUserFuncKeyTemplate**](UsersAPI.md#DissociateUserFuncKeyTemplate) | **Delete** /users/{user_id}/funckeys/templates/{template_id} | Dissociate a func key template to a user
[**DissociateUserLine**](UsersAPI.md#DissociateUserLine) | **Delete** /users/{user_id}/lines/{line_id} | Dissociate user and line
[**DissociateUserQueue**](UsersAPI.md#DissociateUserQueue) | **Delete** /queues/{queue_id}/members/users/{user_id} | Dissociate user and queue
[**DissociateUserSchedule**](UsersAPI.md#DissociateUserSchedule) | **Delete** /users/{user_id}/schedules/{schedule_id} | Dissociate user and schedule
[**DissociateUserVoicemail**](UsersAPI.md#DissociateUserVoicemail) | **Delete** /users/{user_id}/voicemails | Dissociate user and voicemail
[**ExportUsersCsv**](UsersAPI.md#ExportUsersCsv) | **Get** /users/export | Mass export users and associated resources
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{user_id} | Get user
[**GetUserExternalApp**](UsersAPI.md#GetUserExternalApp) | **Get** /users/{user_uuid}/external/apps/{app_name} | Get user external app
[**GetUserFallback**](UsersAPI.md#GetUserFallback) | **Get** /users/{user_id}/fallbacks | List all fallbacks for user
[**GetUserForward**](UsersAPI.md#GetUserForward) | **Get** /users/{user_id}/forwards/{forward_name} | Get forward for a user
[**GetUserFuncKey**](UsersAPI.md#GetUserFuncKey) | **Get** /users/{user_id}/funckeys/{position} | Get a func key for a user
[**GetUserLineAssociatedEndpointsSip**](UsersAPI.md#GetUserLineAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/{line_id}/associated/endpoints/sip | Get SIP endpoint of a line for a user
[**GetUserLineMainAssociatedEndpointsSip**](UsersAPI.md#GetUserLineMainAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/main/associated/endpoints/sip | Get SIP endpoint of main line for a user
[**GetUserMeeting**](UsersAPI.md#GetUserMeeting) | **Get** /users/me/meetings/{meeting_uuid} | Get one of the meetings of the current user
[**GetUserService**](UsersAPI.md#GetUserService) | **Get** /users/{user_id}/services/{service_name} | Get status of service
[**GetUserServices**](UsersAPI.md#GetUserServices) | **Get** /users/{user_id}/services | Get status of all user&#39;s services
[**GetUserVoicemail**](UsersAPI.md#GetUserVoicemail) | **Get** /users/{user_id}/voicemails | Get user voicemails
[**GetUsersSubscriptions**](UsersAPI.md#GetUsersSubscriptions) | **Get** /users/subscriptions | Get user subscription
[**HeadUser**](UsersAPI.md#HeadUser) | **Head** /users/{user_id} | Check if user exists
[**ImportUsersCsv**](UsersAPI.md#ImportUsersCsv) | **Post** /users/import | Mass import users and associated resources
[**ListFuncKeyTemplateUserAssociations**](UsersAPI.md#ListFuncKeyTemplateUserAssociations) | **Get** /funckeys/templates/{template_id}/users | List users associated to template
[**ListUser**](UsersAPI.md#ListUser) | **Get** /users | List users
[**ListUserExternalApps**](UsersAPI.md#ListUserExternalApps) | **Get** /users/{user_uuid}/external/apps | List user external apps
[**ListUserForwards**](UsersAPI.md#ListUserForwards) | **Get** /users/{user_id}/forwards | List forwards for a user
[**ListUserFuncKeyTemplateAssociations**](UsersAPI.md#ListUserFuncKeyTemplateAssociations) | **Get** /users/{user_id}/funckeys/templates | List funckey templates associated to user
[**ListUserFuncKeys**](UsersAPI.md#ListUserFuncKeys) | **Get** /users/{user_id}/funckeys | List func keys for a user
[**ListUserMeetingAuthorizations**](UsersAPI.md#ListUserMeetingAuthorizations) | **Get** /user/me/meetings/{meeting_uuid}/authorizations | List all guest authorization requests of a meeting
[**ListUserMeetings**](UsersAPI.md#ListUserMeetings) | **Get** /users/me/meetings | List user meetings
[**PutUserMeetingAuthorizationAccept**](UsersAPI.md#PutUserMeetingAuthorizationAccept) | **Put** /user/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}/accept | Accept a guest authorization request
[**PutUserMeetingAuthorizationReject**](UsersAPI.md#PutUserMeetingAuthorizationReject) | **Put** /user/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}/reject | Reject a guest authorization request
[**UpdateCallFilterCallerUsers**](UsersAPI.md#UpdateCallFilterCallerUsers) | **Put** /callfilters/{callfilter_id}/recipients/users | Update call filter and recipients
[**UpdateCallFilterMemberUsers**](UsersAPI.md#UpdateCallFilterMemberUsers) | **Put** /callfilters/{callfilter_id}/surrogates/users | Update call filter and surrogates
[**UpdateCallPickupInterceptorUsers**](UsersAPI.md#UpdateCallPickupInterceptorUsers) | **Put** /callpickups/{callpickup_id}/interceptors/users | Update call pickup and interceptors
[**UpdateCallPickupTargetUsers**](UsersAPI.md#UpdateCallPickupTargetUsers) | **Put** /callpickups/{callpickup_id}/targets/users | Update call pickup and targets
[**UpdateGroupMemberUsers**](UsersAPI.md#UpdateGroupMemberUsers) | **Put** /groups/{group_uuid}/members/users | Update group and users
[**UpdatePagingCallerUsers**](UsersAPI.md#UpdatePagingCallerUsers) | **Put** /pagings/{paging_id}/callers/users | Update paging and callers
[**UpdatePagingMemberUsers**](UsersAPI.md#UpdatePagingMemberUsers) | **Put** /pagings/{paging_id}/members/users | Update paging and members
[**UpdateSwitchboardMemberUsers**](UsersAPI.md#UpdateSwitchboardMemberUsers) | **Put** /switchboards/{switchboard_uuid}/members/users | Update switchboard and members
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{user_id} | Update user
[**UpdateUserExternalApp**](UsersAPI.md#UpdateUserExternalApp) | **Put** /users/{user_uuid}/external/apps/{app_name} | Update user external app
[**UpdateUserFallback**](UsersAPI.md#UpdateUserFallback) | **Put** /users/{user_id}/fallbacks | Update user&#39;s fallbacks
[**UpdateUserForward**](UsersAPI.md#UpdateUserForward) | **Put** /users/{user_id}/forwards/{forward_name} | Update a forward for a user
[**UpdateUserForwards**](UsersAPI.md#UpdateUserForwards) | **Put** /users/{user_id}/forwards | Update all forwards for a user
[**UpdateUserFuncKey**](UsersAPI.md#UpdateUserFuncKey) | **Put** /users/{user_id}/funckeys/{position} | Add/Replace a func key for a user
[**UpdateUserFuncKeys**](UsersAPI.md#UpdateUserFuncKeys) | **Put** /users/{user_id}/funckeys | Update func keys for a user
[**UpdateUserGroups**](UsersAPI.md#UpdateUserGroups) | **Put** /users/{user_id}/groups | Update user and groups
[**UpdateUserMeeting**](UsersAPI.md#UpdateUserMeeting) | **Put** /users/me/meetings/{meeting_uuid} | Update one of the meetings of the current user
[**UpdateUserQueueAssociation**](UsersAPI.md#UpdateUserQueueAssociation) | **Put** /queues/{queue_id}/members/users/{user_id} | Update User-Queue association
[**UpdateUserService**](UsersAPI.md#UpdateUserService) | **Put** /users/{user_id}/services/{service_name} | Enable/Disable service for a user
[**UpdateUserServices**](UsersAPI.md#UpdateUserServices) | **Put** /users/{user_id}/services | Update all services for a user
[**UpdateUsersCsv**](UsersAPI.md#UpdateUsersCsv) | **Put** /users/import | **Disabled!** Mass import users and associated resources

## AssociateUserAgent

> AssociateUserAgent(ctx, userId, agentId).AccentTenant(accentTenant).Execute()

Associate user and agent

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 agentId := int32(56) // int32 | Agent’s ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserAgent(context.Background(), userId, agentId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**agentId** | **int32** | Agent’s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserAgentRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserCallpermission

> AssociateUserCallpermission(ctx, userId, callpermissionId).AccentTenant(accentTenant).Execute()

Associate user and call permission

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserCallpermission(context.Background(), userId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserCallpermissionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserFuncKeyTemplate

> AssociateUserFuncKeyTemplate(ctx, userId, templateId).AccentTenant(accentTenant).Execute()

Associate a func key template to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserFuncKeyTemplate(context.Background(), userId, templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserLine

> AssociateUserLine(ctx, userId, lineId).Execute()

Associate user and line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 lineId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserLine(context.Background(), userId, lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserLineRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserLines

> AssociateUserLines(ctx, userId).Body(body).Execute()

Associate user and lines

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewLinesID() // LinesID | 
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserLines(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserLines``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserLinesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LinesID**](LinesID.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserSchedule

> AssociateUserSchedule(ctx, userId, scheduleId).AccentTenant(accentTenant).Execute()

Associate user and schedule

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserSchedule(context.Background(), userId, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserScheduleRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AssociateUserVoicemail

> AssociateUserVoicemail(ctx, userId, voicemailId).Execute()

Associate user and voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 voicemailId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AssociateUserVoicemail(context.Background(), userId, voicemailId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssociateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**voicemailId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUser

> UserPostResponse CreateUser(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserPost() // UserPost | User to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUser(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUser`: UserPostResponse
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserPost**](UserPost.md) | User to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserPostResponse**](UserPostResponse.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUserExternalApp

> UserExternalApp CreateUserExternalApp(ctx, userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()

Create user external app

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserExternalApp() // UserExternalApp | External app to create
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUserExternalApp(context.Background(), userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserExternalApp`: UserExternalApp
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserExternalApp**](UserExternalApp.md) | External app to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserExternalApp**](UserExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUserMeeting

> Meeting CreateUserMeeting(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create user meeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewMeetingUserRequest() // MeetingUserRequest | Meeting to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUserMeeting(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserMeeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserMeeting`: Meeting
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserMeeting`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserMeetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MeetingUserRequest**](MeetingUserRequest.md) | Meeting to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Meeting**](Meeting.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUserVoicemail

> Voicemail CreateUserVoicemail(ctx, userId).Body(body).AccentTenant(accentTenant).Execute()

Create user voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewVoicemail("Context_example", "Number_example") // Voicemail | Voicemail to create
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUserVoicemail(context.Background(), userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserVoicemail`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserVoicemail`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Voicemail**](Voicemail.md) | Voicemail to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUser

> DeleteUser(ctx, userId).AccentTenant(accentTenant).Recursive(recursive).Execute()

Delete user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recursive := true // bool | Indicates if the resources related to the user must be deleted too. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUser(context.Background(), userId).AccentTenant(accentTenant).Recursive(recursive).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recursive** | **bool** | Indicates if the resources related to the user must be deleted too. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserExternalApp

> DeleteUserExternalApp(ctx, userUuid, appName).AccentTenant(accentTenant).Execute()

Delete user external app

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUserExternalApp(context.Background(), userUuid, appName).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserFuncKey

> DeleteUserFuncKey(ctx, userId, position).AccentTenant(accentTenant).Execute()

Remove func key for user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUserFuncKey(context.Background(), userId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserMeeting

> DeleteUserMeeting(ctx, meetingUuid).AccentTenant(accentTenant).Execute()

Delete one of the meetings of the current user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Meeting UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUserMeeting(context.Background(), meetingUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserMeeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserMeetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserAgent

> DissociateUserAgent(ctx, userId).AccentTenant(accentTenant).Execute()

Dissociate user and agent

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserAgent(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserAgentRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserCallpermission

> DissociateUserCallpermission(ctx, userId, callpermissionId).AccentTenant(accentTenant).Execute()

Dissociate user and call permission

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserCallpermission(context.Background(), userId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserCallpermissionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserFuncKeyTemplate

> DissociateUserFuncKeyTemplate(ctx, userId, templateId).AccentTenant(accentTenant).Execute()

Dissociate a func key template to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserFuncKeyTemplate(context.Background(), userId, templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserLine

> DissociateUserLine(ctx, userId, lineId).Execute()

Dissociate user and line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 lineId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserLine(context.Background(), userId, lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserLineRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserQueue

> DissociateUserQueue(ctx, queueId, userId).AccentTenant(accentTenant).Execute()

Dissociate user and queue

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 queueId := int32(56) // int32 | queue's ID
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserQueue(context.Background(), queueId, userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserQueue``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | queue&#39;s ID |
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserQueueRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserSchedule

> DissociateUserSchedule(ctx, userId, scheduleId).AccentTenant(accentTenant).Execute()

Dissociate user and schedule

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserSchedule(context.Background(), userId, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserScheduleRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserVoicemail

> DissociateUserVoicemail(ctx, userId).Execute()

Dissociate user and voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DissociateUserVoicemail(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DissociateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ExportUsersCsv

> string ExportUsersCsv(ctx).AccentTenant(accentTenant).Execute()

Mass export users and associated resources

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ExportUsersCsv(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ExportUsersCsv``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ExportUsersCsv`: string
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ExportUsersCsv`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiExportUsersCsvRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

**string**

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUser

> User GetUser(ctx, userId).AccentTenant(accentTenant).Execute()

Get user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUser`: User
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**User**](User.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserExternalApp

> UserExternalApp GetUserExternalApp(ctx, userUuid, appName).AccentTenant(accentTenant).View(view).Execute()

Get user external app

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 view := "view_example" // string | Different view of the external app endpoint  The `default` view, when the argument is omitted, is to return the user value for this external app  The `fallback` view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserExternalApp(context.Background(), userUuid, appName).AccentTenant(accentTenant).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserExternalApp`: UserExternalApp
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **view** | **string** | Different view of the external app endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to return the user value for this external app  The &#x60;fallback&#x60; view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  |

### Return type

[**UserExternalApp**](UserExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserFallback

> UserFallbacks GetUserFallback(ctx, userId).Execute()

List all fallbacks for user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserFallback(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserFallback`: UserFallbacks
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserFallback`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserFallbacks**](UserFallbacks.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserForward

> UserForward GetUserForward(ctx, userId, forwardName).Execute()

Get forward for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 forwardName := "forwardName_example" // string | the forward name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserForward(context.Background(), userId, forwardName).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserForward``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserForward`: UserForward
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserForward`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**forwardName** | **string** | the forward name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserForwardRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserForward**](UserForward.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserFuncKey

> FuncKey GetUserFuncKey(ctx, userId, position).AccentTenant(accentTenant).Execute()

Get a func key for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserFuncKey(context.Background(), userId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserFuncKey`: FuncKey
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserFuncKey`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKey**](FuncKey.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineAssociatedEndpointsSip

> EndpointSIP GetUserLineAssociatedEndpointsSip(ctx, userUuid, lineId).View(view).Execute()

Get SIP endpoint of a line for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 lineId := int32(56) // int32 | 
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserLineAssociatedEndpointsSip(context.Background(), userUuid, lineId).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserLineAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserLineAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineMainAssociatedEndpointsSip

> EndpointSIP GetUserLineMainAssociatedEndpointsSip(ctx, userUuid).View(view).Execute()

Get SIP endpoint of main line for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserLineMainAssociatedEndpointsSip(context.Background(), userUuid).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserLineMainAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineMainAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserLineMainAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineMainAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserMeeting

> Meeting GetUserMeeting(ctx, meetingUuid).AccentTenant(accentTenant).Execute()

Get one of the meetings of the current user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Meeting UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserMeeting(context.Background(), meetingUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserMeeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserMeeting`: Meeting
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserMeeting`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserMeetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Meeting**](Meeting.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserService

> UserService GetUserService(ctx, userId, serviceName).Execute()

Get status of service

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 serviceName := "serviceName_example" // string | the service name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserService(context.Background(), userId, serviceName).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserService``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserService`: UserService
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserService`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**serviceName** | **string** | the service name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserServiceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserService**](UserService.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserServices

> UserServices GetUserServices(ctx, userId).Execute()

Get status of all user's services

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserServices(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserServices``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserServices`: UserServices
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserServices`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserServicesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserServices**](UserServices.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemail

> VoicemailItems GetUserVoicemail(ctx, userId).AccentTenant(accentTenant).Execute()

Get user voicemails

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserVoicemail(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemail`: VoicemailItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserVoicemail`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**VoicemailItems**](VoicemailItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUsersSubscriptions

> UserSubscriptionItems GetUsersSubscriptions(ctx).AccentTenant(accentTenant).Execute()

Get user subscription

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUsersSubscriptions(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersSubscriptions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUsersSubscriptions`: UserSubscriptionItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersSubscriptions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSubscriptionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserSubscriptionItems**](UserSubscriptionItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HeadUser

> HeadUser(ctx, userId).AccentTenant(accentTenant).Execute()

Check if user exists

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.HeadUser(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.HeadUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiHeadUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ImportUsersCsv

> UserImport ImportUsersCsv(ctx).Body(body).AccentTenant(accentTenant).Execute()

Mass import users and associated resources

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := "body_example" // string | CSV field list available at https://accentvoice.io/uc-doc/administration/users/csv_import
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ImportUsersCsv(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ImportUsersCsv``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ImportUsersCsv`: UserImport
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ImportUsersCsv`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiImportUsersCsvRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | CSV field list available at <https://accentvoice.io/uc-doc/administration/users/csv_import> |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserImport**](UserImport.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: text/csv; charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFuncKeyTemplateUserAssociations

> UserFuncKeyTemplate ListFuncKeyTemplateUserAssociations(ctx, templateId).AccentTenant(accentTenant).Execute()

List users associated to template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListFuncKeyTemplateUserAssociations(context.Background(), templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListFuncKeyTemplateUserAssociations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListFuncKeyTemplateUserAssociations`: UserFuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListFuncKeyTemplateUserAssociations`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiListFuncKeyTemplateUserAssociationsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserFuncKeyTemplate**](UserFuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUser

> UserItems ListUser(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()

List users

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 view := "view_example" // string | Different view of the list of users. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUser(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUser`: UserItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **view** | **string** | Different view of the list of users. |

### Return type

[**UserItems**](UserItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserExternalApps

> UserExternalAppItems ListUserExternalApps(ctx, userUuid).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()

List user external apps

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 view := "view_example" // string | Different view of the external app endpoint  The `default` view, when the argument is omitted, is to return the user value for this external app  The `fallback` view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserExternalApps(context.Background(), userUuid).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserExternalApps``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserExternalApps`: UserExternalAppItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserExternalApps`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserExternalAppsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **view** | **string** | Different view of the external app endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to return the user value for this external app  The &#x60;fallback&#x60; view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  |

### Return type

[**UserExternalAppItems**](UserExternalAppItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserForwards

> UserForwards ListUserForwards(ctx, userId).Execute()

List forwards for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserForwards(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserForwards``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserForwards`: UserForwards
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserForwards`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserForwardsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserForwards**](UserForwards.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserFuncKeyTemplateAssociations

> UserFuncKeyTemplate ListUserFuncKeyTemplateAssociations(ctx, userId).AccentTenant(accentTenant).Execute()

List funckey templates associated to user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserFuncKeyTemplateAssociations(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserFuncKeyTemplateAssociations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserFuncKeyTemplateAssociations`: UserFuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserFuncKeyTemplateAssociations`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserFuncKeyTemplateAssociationsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserFuncKeyTemplate**](UserFuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserFuncKeys

> FuncKeyTemplate ListUserFuncKeys(ctx, userId).AccentTenant(accentTenant).Execute()

List func keys for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserFuncKeys(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserFuncKeys``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserFuncKeys`: FuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserFuncKeys`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserFuncKeysRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKeyTemplate**](FuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserMeetingAuthorizations

> MeetingAuthorizationItems ListUserMeetingAuthorizations(ctx, meetingUuid).Execute()

List all guest authorization requests of a meeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Meeting UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserMeetingAuthorizations(context.Background(), meetingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserMeetingAuthorizations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserMeetingAuthorizations`: MeetingAuthorizationItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserMeetingAuthorizations`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMeetingAuthorizationsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MeetingAuthorizationItems**](MeetingAuthorizationItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserMeetings

> MeetingItems ListUserMeetings(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List user meetings

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserMeetings(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserMeetings``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserMeetings`: MeetingItems
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserMeetings`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMeetingsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**MeetingItems**](MeetingItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PutUserMeetingAuthorizationAccept

> MeetingAuthorization PutUserMeetingAuthorizationAccept(ctx, meetingUuid, authorizationUuid).Execute()

Accept a guest authorization request

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Meeting UUID
 authorizationUuid := "authorizationUuid_example" // string | Authorization UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.PutUserMeetingAuthorizationAccept(context.Background(), meetingUuid, authorizationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PutUserMeetingAuthorizationAccept``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PutUserMeetingAuthorizationAccept`: MeetingAuthorization
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PutUserMeetingAuthorizationAccept`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |
**authorizationUuid** | **string** | Authorization UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserMeetingAuthorizationAcceptRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MeetingAuthorization**](MeetingAuthorization.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PutUserMeetingAuthorizationReject

> MeetingAuthorization PutUserMeetingAuthorizationReject(ctx, meetingUuid, authorizationUuid).Execute()

Reject a guest authorization request

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Meeting UUID
 authorizationUuid := "authorizationUuid_example" // string | Authorization UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.PutUserMeetingAuthorizationReject(context.Background(), meetingUuid, authorizationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PutUserMeetingAuthorizationReject``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PutUserMeetingAuthorizationReject`: MeetingAuthorization
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PutUserMeetingAuthorizationReject`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |
**authorizationUuid** | **string** | Authorization UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserMeetingAuthorizationRejectRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MeetingAuthorization**](MeetingAuthorization.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallFilterCallerUsers

> UpdateCallFilterCallerUsers(ctx, callfilterId).Body(body).Execute()

Update call filter and recipients

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewCallFilterRecipientUsersUuid() // CallFilterRecipientUsersUuid | Users to associated
 callfilterId := int32(56) // int32 | Call Filter's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateCallFilterCallerUsers(context.Background(), callfilterId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCallFilterCallerUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFilterCallerUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallFilterRecipientUsersUuid**](CallFilterRecipientUsersUuid.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallFilterMemberUsers

> UpdateCallFilterMemberUsers(ctx, callfilterId).Body(body).Execute()

Update call filter and surrogates

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 callfilterId := int32(56) // int32 | Call Filter's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateCallFilterMemberUsers(context.Background(), callfilterId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCallFilterMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFilterMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallPickupInterceptorUsers

> UpdateCallPickupInterceptorUsers(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and interceptors

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateCallPickupInterceptorUsers(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCallPickupInterceptorUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupInterceptorUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallPickupTargetUsers

> UpdateCallPickupTargetUsers(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and targets

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateCallPickupTargetUsers(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCallPickupTargetUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupTargetUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateGroupMemberUsers

> UpdateGroupMemberUsers(ctx, groupUuid).Body(body).Execute()

Update group and users

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewGroupMemberUsers([]openapiclient.GroupMemberUser{*openapiclient.NewGroupMemberUser("Uuid_example")}) // GroupMemberUsers | Users to associated
 groupUuid := "groupUuid_example" // string | the group's UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateGroupMemberUsers(context.Background(), groupUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateGroupMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupMemberUsers**](GroupMemberUsers.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdatePagingCallerUsers

> UpdatePagingCallerUsers(ctx, pagingId).Body(body).Execute()

Update paging and callers

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 pagingId := int32(56) // int32 | Paging's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdatePagingCallerUsers(context.Background(), pagingId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdatePagingCallerUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagingCallerUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdatePagingMemberUsers

> UpdatePagingMemberUsers(ctx, pagingId).Body(body).Execute()

Update paging and members

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 pagingId := int32(56) // int32 | Paging's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdatePagingMemberUsers(context.Background(), pagingId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdatePagingMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagingMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateSwitchboardMemberUsers

> UpdateSwitchboardMemberUsers(ctx, switchboardUuid).Body(body).Execute()

Update switchboard and members

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associate with the switchboard
 switchboardUuid := "switchboardUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateSwitchboardMemberUsers(context.Background(), switchboardUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateSwitchboardMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchboardMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associate with the switchboard |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUser

> UpdateUser(ctx, userId).Body(body).AccentTenant(accentTenant).Execute()

Update user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUser() // User | 
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUser(context.Background(), userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**User**](User.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserExternalApp

> UpdateUserExternalApp(ctx, userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()

Update user external app

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserExternalApp() // UserExternalApp | 
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserExternalApp(context.Background(), userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserExternalApp**](UserExternalApp.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserFallback

> UpdateUserFallback(ctx, userId).Body(body).Execute()

Update user's fallbacks

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 body := *openapiclient.NewUserFallbacks() // UserFallbacks | Fallbacks for user (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserFallback(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserFallbacks**](UserFallbacks.md) | Fallbacks for user |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserForward

> UpdateUserForward(ctx, userId, forwardName).Body(body).Execute()

Update a forward for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserForward() // UserForward | 
 userId := "userId_example" // string | the user's ID or UUID
 forwardName := "forwardName_example" // string | the forward name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserForward(context.Background(), userId, forwardName).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserForward``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**forwardName** | **string** | the forward name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserForwardRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserForward**](UserForward.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserForwards

> UpdateUserForwards(ctx, userId).Body(body).Execute()

Update all forwards for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserForwards() // UserForwards | 
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserForwards(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserForwards``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserForwardsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserForwards**](UserForwards.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserFuncKey

> UpdateUserFuncKey(ctx, userId, position).Body(body).AccentTenant(accentTenant).Execute()

Add/Replace a func key for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKey() // FuncKey | 
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserFuncKey(context.Background(), userId, position).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKey**](FuncKey.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserFuncKeys

> UpdateUserFuncKeys(ctx, userId).Body(body).AccentTenant(accentTenant).Execute()

Update func keys for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKeyTemplate() // FuncKeyTemplate | 
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserFuncKeys(context.Background(), userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserFuncKeys``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserFuncKeysRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKeyTemplate**](FuncKeyTemplate.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserGroups

> UpdateUserGroups(ctx, userId).Body(body).Execute()

Update user and groups

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserGroupsID() // UserGroupsID | Users to associated
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserGroups(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserGroupsID**](UserGroupsID.md) | Users to associated |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserMeeting

> UpdateUserMeeting(ctx, meetingUuid).Body(body).AccentTenant(accentTenant).Execute()

Update one of the meetings of the current user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewMeetingRequest() // MeetingRequest | 
 meetingUuid := "meetingUuid_example" // string | Meeting UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserMeeting(context.Background(), meetingUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserMeeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Meeting UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserMeetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MeetingRequest**](MeetingRequest.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserQueueAssociation

> UpdateUserQueueAssociation(ctx, queueId, userId).Body(body).AccentTenant(accentTenant).Execute()

Update User-Queue association

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 queueId := int32(56) // int32 | queue's ID
 userId := "userId_example" // string | the user's ID or UUID
 body := *openapiclient.NewQueueMemberUser() // QueueMemberUser |  (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserQueueAssociation(context.Background(), queueId, userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserQueueAssociation``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | queue&#39;s ID |
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserQueueAssociationRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QueueMemberUser**](QueueMemberUser.md) |  |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserService

> UpdateUserService(ctx, userId, serviceName).Body(body).Execute()

Enable/Disable service for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserService(false) // UserService | 
 userId := "userId_example" // string | the user's ID or UUID
 serviceName := "serviceName_example" // string | the service name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserService(context.Background(), userId, serviceName).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserService``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**serviceName** | **string** | the service name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserServiceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserService**](UserService.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserServices

> UpdateUserServices(ctx, userId).Body(body).Execute()

Update all services for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewUserServices() // UserServices | 
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserServices(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserServices``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserServicesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserServices**](UserServices.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUsersCsv

> UserUpdate UpdateUsersCsv(ctx).Body(body).AccentTenant(accentTenant).Execute()

**Disabled!** Mass import users and associated resources

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := "body_example" // string | CSV field list available at https://accentvoice.io/uc-doc/administration/users/csv_import
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.UpdateUsersCsv(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUsersCsv``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateUsersCsv`: UserUpdate
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUsersCsv`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsersCsvRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | CSV field list available at <https://accentvoice.io/uc-doc/administration/users/csv_import> |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserUpdate**](UserUpdate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: text/csv; charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
