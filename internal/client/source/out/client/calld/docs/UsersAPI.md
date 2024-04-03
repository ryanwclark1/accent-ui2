# \UsersAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerUserCall**](UsersAPI.md#AnswerUserCall) | **Put** /users/me/calls/{call_id}/answer | Answer a call from user
[**CancelRelocate**](UsersAPI.md#CancelRelocate) | **Put** /users/me/relocates/{relocate_uuid}/cancel | Cancel a relocate
[**CancelUserTransfer**](UsersAPI.md#CancelUserTransfer) | **Delete** /users/me/transfers/{transfer_id} | Cancel a transfer
[**CheckUserVoicemailGreeting**](UsersAPI.md#CheckUserVoicemailGreeting) | **Head** /users/me/voicemails/greetings/{greeting} | Check if greeting exists
[**CompleteRelocate**](UsersAPI.md#CompleteRelocate) | **Put** /users/me/relocates/{relocate_uuid}/complete | Complete a relocate
[**CompleteUserTransfer**](UsersAPI.md#CompleteUserTransfer) | **Put** /users/me/transfers/{transfer_id}/complete | Complete a transfer
[**CopyUserVoicemailGreeting**](UsersAPI.md#CopyUserVoicemailGreeting) | **Post** /users/me/voicemails/greetings/{greeting}/copy | Copy a custom greeting
[**CreateUserCall**](UsersAPI.md#CreateUserCall) | **Post** /users/me/calls | Make a new call from a user
[**CreateUserVoicemailGreeting**](UsersAPI.md#CreateUserVoicemailGreeting) | **Post** /users/me/voicemails/greetings/{greeting} | Create a custom greeting
[**DeleteUserVoicemailGreeting**](UsersAPI.md#DeleteUserVoicemailGreeting) | **Delete** /users/me/voicemails/greetings/{greeting} | Delete a custom greeting
[**DeleteUserVoicemailMessage**](UsersAPI.md#DeleteUserVoicemailMessage) | **Delete** /users/me/voicemails/messages/{message_id} | Delete a mesage
[**GetUserVoicemailFolder**](UsersAPI.md#GetUserVoicemailFolder) | **Get** /users/me/voicemails/folders/{folder_id} | Get details of a folder
[**GetUserVoicemailGreeting**](UsersAPI.md#GetUserVoicemailGreeting) | **Get** /users/me/voicemails/greetings/{greeting} | Get a custom greeting
[**GetUserVoicemailMessage**](UsersAPI.md#GetUserVoicemailMessage) | **Get** /users/me/voicemails/messages/{message_id} | Get a message
[**GetUserVoicemailMessageRecording**](UsersAPI.md#GetUserVoicemailMessageRecording) | **Get** /users/me/voicemails/messages/{message_id}/recording | Get a message&#39;s recording
[**HangupUserCall**](UsersAPI.md#HangupUserCall) | **Delete** /users/me/calls/{call_id} | Hangup a call from a user
[**HoldUserCall**](UsersAPI.md#HoldUserCall) | **Put** /users/me/calls/{call_id}/hold/start | Hold a call from user
[**InitiateRelocate**](UsersAPI.md#InitiateRelocate) | **Post** /users/me/relocates | Initiate a relocate from the authenticated user
[**InitiateUserTransfer**](UsersAPI.md#InitiateUserTransfer) | **Post** /users/me/transfers | Initiate a transfer from the authenticated user
[**KickUserMeetingParticipant**](UsersAPI.md#KickUserMeetingParticipant) | **Delete** /users/me/meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting as a user
[**ListUserCalls**](UsersAPI.md#ListUserCalls) | **Get** /users/me/calls | List calls of a user
[**ListUserConferenceParticipants**](UsersAPI.md#ListUserConferenceParticipants) | **Get** /users/me/conferences/{conference_id}/participants | List participants of a conference as a user
[**ListUserMeetingParticipants**](UsersAPI.md#ListUserMeetingParticipants) | **Get** /users/me/meetings/{meeting_uuid}/participants | List participants of a meeting as a user
[**ListUserRelocates**](UsersAPI.md#ListUserRelocates) | **Get** /users/me/relocates | Get the relocates of the authenticated user
[**ListUserTransfers**](UsersAPI.md#ListUserTransfers) | **Get** /users/me/transfers | Get the transfers of the authenticated user
[**ListUserVoicemails**](UsersAPI.md#ListUserVoicemails) | **Get** /users/me/voicemails | Get details of the voicemail of the authenticated user
[**MuteUserCall**](UsersAPI.md#MuteUserCall) | **Put** /users/me/calls/{call_id}/mute/start | Mute a call from user
[**SendUserDTMF**](UsersAPI.md#SendUserDTMF) | **Put** /users/me/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
[**StartCurrentUserRecording**](UsersAPI.md#StartCurrentUserRecording) | **Put** /users/me/calls/{call_id}/record/start | Start recording a call
[**StopCurrentUserRecording**](UsersAPI.md#StopCurrentUserRecording) | **Put** /users/me/calls/{call_id}/record/stop | Stop recording a call
[**UnholdUserCall**](UsersAPI.md#UnholdUserCall) | **Put** /users/me/calls/{call_id}/hold/stop | Unhold a call from user
[**UnmuteUserCall**](UsersAPI.md#UnmuteUserCall) | **Put** /users/me/calls/{call_id}/mute/stop | Unmute a call from user
[**UpdateUserVoicemailGreeting**](UsersAPI.md#UpdateUserVoicemailGreeting) | **Put** /users/me/voicemails/greetings/{greeting} | Update a custom greeting
[**UpdateUserVoicemailMessage**](UsersAPI.md#UpdateUserVoicemailMessage) | **Put** /users/me/voicemails/messages/{message_id} | Update a message

## AnswerUserCall

> AnswerUserCall(ctx, callId).Execute()

Answer a call from user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.AnswerUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AnswerUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerUserCallRequest struct via the builder pattern

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

## CancelRelocate

> CancelRelocate(ctx, relocateUuid).Execute()

Cancel a relocate

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 relocateUuid := "relocateUuid_example" // string | Unique identifier of the relocate

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CancelRelocate(context.Background(), relocateUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CancelRelocate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relocateUuid** | **string** | Unique identifier of the relocate |

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRelocateRequest struct via the builder pattern

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

## CancelUserTransfer

> CancelUserTransfer(ctx, transferId).Execute()

Cancel a transfer

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CancelUserTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CancelUserTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Unique identifier of the transfer |

### Other Parameters

Other parameters are passed through a pointer to a apiCancelUserTransferRequest struct via the builder pattern

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

## CheckUserVoicemailGreeting

> CheckUserVoicemailGreeting(ctx, greeting).Execute()

Check if greeting exists

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CheckUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CheckUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserVoicemailGreetingRequest struct via the builder pattern

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

## CompleteRelocate

> CompleteRelocate(ctx, relocateUuid).Execute()

Complete a relocate

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 relocateUuid := "relocateUuid_example" // string | Unique identifier of the relocate

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CompleteRelocate(context.Background(), relocateUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompleteRelocate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relocateUuid** | **string** | Unique identifier of the relocate |

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteRelocateRequest struct via the builder pattern

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

## CompleteUserTransfer

> CompleteUserTransfer(ctx, transferId).Execute()

Complete a transfer

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CompleteUserTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompleteUserTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Unique identifier of the transfer |

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteUserTransferRequest struct via the builder pattern

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

## CopyUserVoicemailGreeting

> CopyUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Copy a custom greeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewGreetingCopy("DestGreeting_example") // GreetingCopy | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CopyUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CopyUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCopyUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GreetingCopy**](GreetingCopy.md) |  |

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

## CreateUserCall

> Call CreateUserCall(ctx).Body(body).Execute()

Make a new call from a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewUserCallRequest("Extension_example") // UserCallRequest | Parameters of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUserCall(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserCall`: Call
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserCall`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserCallRequest**](UserCallRequest.md) | Parameters of the call |

### Return type

[**Call**](Call.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUserVoicemailGreeting

> CreateUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Create a custom greeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.CreateUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserVoicemailGreeting

> DeleteUserVoicemailGreeting(ctx, greeting).Execute()

Delete a custom greeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserVoicemailGreetingRequest struct via the builder pattern

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

## DeleteUserVoicemailMessage

> DeleteUserVoicemailMessage(ctx, messageId).Execute()

Delete a mesage

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUserVoicemailMessage(context.Background(), messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserVoicemailMessageRequest struct via the builder pattern

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

## GetUserVoicemailFolder

> VoicemailFolder GetUserVoicemailFolder(ctx, folderId).Execute()

Get details of a folder

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 folderId := int32(56) // int32 | The folder's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserVoicemailFolder(context.Background(), folderId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserVoicemailFolder``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemailFolder`: VoicemailFolder
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserVoicemailFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailFolder**](VoicemailFolder.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailGreeting

> GetUserVoicemailGreeting(ctx, greeting).Execute()

Get a custom greeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.GetUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailMessage

> VoicemailMessage GetUserVoicemailMessage(ctx, messageId).Execute()

Get a message

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserVoicemailMessage(context.Background(), messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemailMessage`: VoicemailMessage
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserVoicemailMessage`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailMessage**](VoicemailMessage.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailMessageRecording

> GetUserVoicemailMessageRecording(ctx, messageId).Token(token).Download(download).Execute()

Get a message's recording

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 messageId := "messageId_example" // string | The message's ID
 token := "token_example" // string | The token's ID (optional)
 download := "download_example" // string | Set to 1 to force download by browser (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.GetUserVoicemailMessageRecording(context.Background(), messageId).Token(token).Download(download).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserVoicemailMessageRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailMessageRecordingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | The token&#39;s ID |
 **download** | **string** | Set to 1 to force download by browser |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HangupUserCall

> HangupUserCall(ctx, callId).Execute()

Hangup a call from a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.HangupUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.HangupUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiHangupUserCallRequest struct via the builder pattern

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

## HoldUserCall

> HoldUserCall(ctx, callId).Execute()

Hold a call from user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.HoldUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.HoldUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiHoldUserCallRequest struct via the builder pattern

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

## InitiateRelocate

> Relocate InitiateRelocate(ctx).Body(body).Execute()

Initiate a relocate from the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewUserRelocateRequest("Destination_example", "InitiatorCall_example") // UserRelocateRequest | Parameters of the relocate

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.InitiateRelocate(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.InitiateRelocate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InitiateRelocate`: Relocate
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.InitiateRelocate`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateRelocateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserRelocateRequest**](UserRelocateRequest.md) | Parameters of the relocate |

### Return type

[**Relocate**](Relocate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## InitiateUserTransfer

> Transfer InitiateUserTransfer(ctx).Body(body).Execute()

Initiate a transfer from the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewUserTransferRequest("Exten_example", "InitiatorCall_example") // UserTransferRequest | Parameters of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.InitiateUserTransfer(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.InitiateUserTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InitiateUserTransfer`: Transfer
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.InitiateUserTransfer`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateUserTransferRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserTransferRequest**](UserTransferRequest.md) | Parameters of the transfer |

### Return type

[**Transfer**](Transfer.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## KickUserMeetingParticipant

> KickUserMeetingParticipant(ctx, meetingUuid, participantId).Execute()

Kick a participant from a meeting as a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.KickUserMeetingParticipant(context.Background(), meetingUuid, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.KickUserMeetingParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiKickUserMeetingParticipantRequest struct via the builder pattern

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

## ListUserCalls

> ListCalls200Response ListUserCalls(ctx).Application(application).ApplicationInstance(applicationInstance).Execute()

List calls of a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 application := "application_example" // string | Filter calls by Stasis application, e.g. switchboard. (optional)
 applicationInstance := "applicationInstance_example" // string | Filter calls by Stasis application instance, e.g. switchboard-sales,green. Args must be separated by commas (,). (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserCalls(context.Background()).Application(application).ApplicationInstance(applicationInstance).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserCalls`: ListCalls200Response
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserCalls`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListUserCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | Filter calls by Stasis application, e.g. switchboard. |
 **applicationInstance** | **string** | Filter calls by Stasis application instance, e.g. switchboard-sales,green. Args must be separated by commas (,). |

### Return type

[**ListCalls200Response**](ListCalls200Response.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserConferenceParticipants

> ParticipantList ListUserConferenceParticipants(ctx, conferenceId).Execute()

List participants of a conference as a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserConferenceParticipants(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserConferenceParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserConferenceParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserConferenceParticipantsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ParticipantList**](ParticipantList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserMeetingParticipants

> ParticipantList ListUserMeetingParticipants(ctx, meetingUuid).Execute()

List participants of a meeting as a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserMeetingParticipants(context.Background(), meetingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserMeetingParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserMeetingParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserMeetingParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMeetingParticipantsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ParticipantList**](ParticipantList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserRelocates

> RelocateList ListUserRelocates(ctx).Execute()

Get the relocates of the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserRelocates(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserRelocates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserRelocates`: RelocateList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserRelocates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRelocatesRequest struct via the builder pattern

### Return type

[**RelocateList**](RelocateList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserTransfers

> TransferList ListUserTransfers(ctx).Execute()

Get the transfers of the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserTransfers(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserTransfers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserTransfers`: TransferList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserTransfers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserTransfersRequest struct via the builder pattern

### Return type

[**TransferList**](TransferList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserVoicemails

> Voicemail ListUserVoicemails(ctx).Execute()

Get details of the voicemail of the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.ListUserVoicemails(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserVoicemails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserVoicemails`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserVoicemails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserVoicemailsRequest struct via the builder pattern

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MuteUserCall

> MuteUserCall(ctx, callId).Execute()

Mute a call from user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.MuteUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.MuteUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiMuteUserCallRequest struct via the builder pattern

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

## SendUserDTMF

> SendUserDTMF(ctx, callId).Digits(digits).Execute()

Simulate a user pressing DTMF keys

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call
 digits := "digits_example" // string | Digits to send via DTMF. Must contain only `0-9*#`.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.SendUserDTMF(context.Background(), callId).Digits(digits).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SendUserDTMF``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiSendUserDTMFRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digits** | **string** | Digits to send via DTMF. Must contain only &#x60;0-9*#&#x60;. |

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

## StartCurrentUserRecording

> StartCurrentUserRecording(ctx, callId).Execute()

Start recording a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.StartCurrentUserRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.StartCurrentUserRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiStartCurrentUserRecordingRequest struct via the builder pattern

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

## StopCurrentUserRecording

> StopCurrentUserRecording(ctx, callId).Execute()

Stop recording a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.StopCurrentUserRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.StopCurrentUserRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiStopCurrentUserRecordingRequest struct via the builder pattern

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

## UnholdUserCall

> UnholdUserCall(ctx, callId).Execute()

Unhold a call from user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UnholdUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnholdUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiUnholdUserCallRequest struct via the builder pattern

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

## UnmuteUserCall

> UnmuteUserCall(ctx, callId).Execute()

Unmute a call from user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UnmuteUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnmuteUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteUserCallRequest struct via the builder pattern

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

## UpdateUserVoicemailGreeting

> UpdateUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Update a custom greeting

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserVoicemailMessage

> UpdateUserVoicemailMessage(ctx, messageId).Body(body).Execute()

Update a message

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewVoicemailMessageUpdate(int32(123)) // VoicemailMessageUpdate | Message
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserVoicemailMessage(context.Background(), messageId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailMessageUpdate**](VoicemailMessageUpdate.md) | Message |

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
