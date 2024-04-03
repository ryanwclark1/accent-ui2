# \SessionsAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSession**](SessionsAPI.md#DeleteSession) | **Delete** /sessions/{session_uuid} | Delete a session
[**GetSessions**](SessionsAPI.md#GetSessions) | **Get** /sessions | List sessions
[**GetUserSessions**](SessionsAPI.md#GetUserSessions) | **Get** /users/{user_uuid}/sessions | Retrieves the list of sessions associated to a user
[**UserDeleteSession**](SessionsAPI.md#UserDeleteSession) | **Delete** /users/{user_uuid}/sessions/{session_uuid} | Delete a session

## DeleteSession

> DeleteSession(ctx, sessionUuid).Execute()

Delete a session

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 sessionUuid := "sessionUuid_example" // string | The UUID of the session

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SessionsAPI.DeleteSession(context.Background(), sessionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.DeleteSession``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionUuid** | **string** | The UUID of the session |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern

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

## GetSessions

> GetSessionsResult GetSessions(ctx).AccentTenant(accentTenant).Recurse(recurse).Limit(limit).Offset(offset).Execute()

List sessions

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SessionsAPI.GetSessions(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Limit(limit).Offset(offset).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.GetSessions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSessions`: GetSessionsResult
 fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.GetSessions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]

### Return type

[**GetSessionsResult**](GetSessionsResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserSessions

> GetSessionsResult GetUserSessions(ctx, userUuid).AccentTenant(accentTenant).Limit(limit).Offset(offset).Execute()

Retrieves the list of sessions associated to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 userUuid := "userUuid_example" // string | The UUID of the user
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SessionsAPI.GetUserSessions(context.Background(), userUuid).AccentTenant(accentTenant).Limit(limit).Offset(offset).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.GetUserSessions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserSessions`: GetSessionsResult
 fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.GetUserSessions`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]

### Return type

[**GetSessionsResult**](GetSessionsResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserDeleteSession

> UserDeleteSession(ctx, userUuid, sessionUuid).Execute()

Delete a session

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 userUuid := "userUuid_example" // string | The UUID of the user
 sessionUuid := "sessionUuid_example" // string | The UUID of the session

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SessionsAPI.UserDeleteSession(context.Background(), userUuid, sessionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.UserDeleteSession``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |
**sessionUuid** | **string** | The UUID of the session |

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteSessionRequest struct via the builder pattern

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
