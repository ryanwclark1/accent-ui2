# \CallsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerCall**](CallsAPI.md#AnswerCall) | **Put** /calls/{call_id}/answer | Answer a call
[**AnswerUserCall**](CallsAPI.md#AnswerUserCall) | **Put** /users/me/calls/{call_id}/answer | Answer a call from user
[**ConnectCallToUser**](CallsAPI.md#ConnectCallToUser) | **Put** /calls/{call_id}/user/{user_uuid} | Connect a call to a user
[**CreateCall**](CallsAPI.md#CreateCall) | **Post** /calls | Make a new call
[**CreateUserCall**](CallsAPI.md#CreateUserCall) | **Post** /users/me/calls | Make a new call from a user
[**DeleteCall**](CallsAPI.md#DeleteCall) | **Delete** /calls/{call_id} | Hangup a call
[**GetCall**](CallsAPI.md#GetCall) | **Get** /calls/{call_id} | Show a call
[**HangupUserCall**](CallsAPI.md#HangupUserCall) | **Delete** /users/me/calls/{call_id} | Hangup a call from a user
[**HoldCall**](CallsAPI.md#HoldCall) | **Put** /calls/{call_id}/hold/start | Hold a call
[**HoldUserCall**](CallsAPI.md#HoldUserCall) | **Put** /users/me/calls/{call_id}/hold/start | Hold a call from user
[**ListCalls**](CallsAPI.md#ListCalls) | **Get** /calls | List calls
[**ListUserCalls**](CallsAPI.md#ListUserCalls) | **Get** /users/me/calls | List calls of a user
[**MuteCall**](CallsAPI.md#MuteCall) | **Put** /calls/{call_id}/mute/start | Mute a call
[**MuteUserCall**](CallsAPI.md#MuteUserCall) | **Put** /users/me/calls/{call_id}/mute/start | Mute a call from user
[**SendCallDTMF**](CallsAPI.md#SendCallDTMF) | **Put** /calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
[**SendUserDTMF**](CallsAPI.md#SendUserDTMF) | **Put** /users/me/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
[**StartCurrentUserRecording**](CallsAPI.md#StartCurrentUserRecording) | **Put** /users/me/calls/{call_id}/record/start | Start recording a call
[**StartRecording**](CallsAPI.md#StartRecording) | **Put** /calls/{call_id}/record/start | Start recording a call
[**StopCurrentUserRecording**](CallsAPI.md#StopCurrentUserRecording) | **Put** /users/me/calls/{call_id}/record/stop | Stop recording a call
[**StopRecording**](CallsAPI.md#StopRecording) | **Put** /calls/{call_id}/record/stop | Stop recording a call
[**UnholdCall**](CallsAPI.md#UnholdCall) | **Put** /calls/{call_id}/hold/stop | Unhold a call
[**UnholdUserCall**](CallsAPI.md#UnholdUserCall) | **Put** /users/me/calls/{call_id}/hold/stop | Unhold a call from user
[**UnmuteCall**](CallsAPI.md#UnmuteCall) | **Put** /calls/{call_id}/mute/stop | Unmute a call
[**UnmuteUserCall**](CallsAPI.md#UnmuteUserCall) | **Put** /users/me/calls/{call_id}/mute/stop | Unmute a call from user

## AnswerCall

> AnswerCall(ctx, callId).Execute()

Answer a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.AnswerCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.AnswerCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAnswerCallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.AnswerUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.AnswerUserCall``: %v\n", err)
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

## ConnectCallToUser

> Call ConnectCallToUser(ctx, callId, userUuid).Body(body).Execute()

Connect a call to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call
 userUuid := "userUuid_example" // string | UUID of the user
 body := *openapiclient.NewConnectCallToUserRequest() // ConnectCallToUserRequest | options affecting the call to the targeted user (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.ConnectCallToUser(context.Background(), callId, userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.ConnectCallToUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ConnectCallToUser`: Call
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.ConnectCallToUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |
**userUuid** | **string** | UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiConnectCallToUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectCallToUserRequest**](ConnectCallToUserRequest.md) | options affecting the call to the targeted user |

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

## CreateCall

> Call CreateCall(ctx).Body(body).Execute()

Make a new call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 body := *openapiclient.NewCallRequest(*openapiclient.NewCallRequestDestination("Context_example", "Extension_example", int32(123)), *openapiclient.NewCallRequestSource("User_example")) // CallRequest | Call parameters

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.CreateCall(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.CreateCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCall`: Call
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.CreateCall`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallRequest**](CallRequest.md) | Call parameters |

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 body := *openapiclient.NewUserCallRequest("Extension_example") // UserCallRequest | Parameters of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.CreateUserCall(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.CreateUserCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserCall`: Call
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.CreateUserCall`: %v\n", resp)
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

## DeleteCall

> DeleteCall(ctx, callId).Execute()

Hangup a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.DeleteCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.DeleteCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCallRequest struct via the builder pattern

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

## GetCall

> Call GetCall(ctx, callId).Execute()

Show a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.GetCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.GetCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCall`: Call
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.GetCall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.HangupUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.HangupUserCall``: %v\n", err)
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

## HoldCall

> HoldCall(ctx, callId).Execute()

Hold a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.HoldCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.HoldCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiHoldCallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.HoldUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.HoldUserCall``: %v\n", err)
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

## ListCalls

> ListCalls200Response ListCalls(ctx).Application(application).ApplicationInstance(applicationInstance).Execute()

List calls

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 application := "application_example" // string | Filter calls by Stasis application, e.g. switchboard. (optional)
 applicationInstance := "applicationInstance_example" // string | Filter calls by Stasis application instance, e.g. switchboard-sales,green. Args must be separated by commas (,). `application_instance` is ignored if `application` is not set. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.ListCalls(context.Background()).Application(application).ApplicationInstance(applicationInstance).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.ListCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListCalls`: ListCalls200Response
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.ListCalls`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | Filter calls by Stasis application, e.g. switchboard. |
 **applicationInstance** | **string** | Filter calls by Stasis application instance, e.g. switchboard-sales,green. Args must be separated by commas (,). &#x60;application_instance&#x60; is ignored if &#x60;application&#x60; is not set. |

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 application := "application_example" // string | Filter calls by Stasis application, e.g. switchboard. (optional)
 applicationInstance := "applicationInstance_example" // string | Filter calls by Stasis application instance, e.g. switchboard-sales,green. Args must be separated by commas (,). (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallsAPI.ListUserCalls(context.Background()).Application(application).ApplicationInstance(applicationInstance).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.ListUserCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserCalls`: ListCalls200Response
 fmt.Fprintf(os.Stdout, "Response from `CallsAPI.ListUserCalls`: %v\n", resp)
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

## MuteCall

> MuteCall(ctx, callId).Execute()

Mute a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.MuteCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.MuteCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMuteCallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.MuteUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.MuteUserCall``: %v\n", err)
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

## SendCallDTMF

> SendCallDTMF(ctx, callId).Digits(digits).Execute()

Simulate a user pressing DTMF keys

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call
 digits := "digits_example" // string | Digits to send via DTMF. Must contain only `0-9*#`.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.SendCallDTMF(context.Background(), callId).Digits(digits).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.SendCallDTMF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSendCallDTMFRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call
 digits := "digits_example" // string | Digits to send via DTMF. Must contain only `0-9*#`.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.SendUserDTMF(context.Background(), callId).Digits(digits).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.SendUserDTMF``: %v\n", err)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.StartCurrentUserRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.StartCurrentUserRecording``: %v\n", err)
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

## StartRecording

> StartRecording(ctx, callId).Execute()

Start recording a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.StartRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.StartRecording``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStartRecordingRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.StopCurrentUserRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.StopCurrentUserRecording``: %v\n", err)
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

## StopRecording

> StopRecording(ctx, callId).Execute()

Stop recording a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.StopRecording(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.StopRecording``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStopRecordingRequest struct via the builder pattern

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

## UnholdCall

> UnholdCall(ctx, callId).Execute()

Unhold a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.UnholdCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UnholdCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnholdCallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.UnholdUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UnholdUserCall``: %v\n", err)
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

## UnmuteCall

> UnmuteCall(ctx, callId).Execute()

Unmute a call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.UnmuteCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UnmuteCall``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnmuteCallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallsAPI.UnmuteUserCall(context.Background(), callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UnmuteUserCall``: %v\n", err)
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
