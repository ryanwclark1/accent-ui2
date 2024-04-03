# \ForwardsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserForward**](ForwardsAPI.md#GetUserForward) | **Get** /users/{user_id}/forwards/{forward_name} | Get forward for a user
[**ListUserForwards**](ForwardsAPI.md#ListUserForwards) | **Get** /users/{user_id}/forwards | List forwards for a user
[**UpdateUserForward**](ForwardsAPI.md#UpdateUserForward) | **Put** /users/{user_id}/forwards/{forward_name} | Update a forward for a user
[**UpdateUserForwards**](ForwardsAPI.md#UpdateUserForwards) | **Put** /users/{user_id}/forwards | Update all forwards for a user

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
 resp, r, err := apiClient.ForwardsAPI.GetUserForward(context.Background(), userId, forwardName).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ForwardsAPI.GetUserForward``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserForward`: UserForward
 fmt.Fprintf(os.Stdout, "Response from `ForwardsAPI.GetUserForward`: %v\n", resp)
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
 resp, r, err := apiClient.ForwardsAPI.ListUserForwards(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ForwardsAPI.ListUserForwards``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserForwards`: UserForwards
 fmt.Fprintf(os.Stdout, "Response from `ForwardsAPI.ListUserForwards`: %v\n", resp)
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
 r, err := apiClient.ForwardsAPI.UpdateUserForward(context.Background(), userId, forwardName).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ForwardsAPI.UpdateUserForward``: %v\n", err)
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
 r, err := apiClient.ForwardsAPI.UpdateUserForwards(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ForwardsAPI.UpdateUserForwards``: %v\n", err)
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
