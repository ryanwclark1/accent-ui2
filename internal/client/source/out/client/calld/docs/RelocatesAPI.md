# \RelocatesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRelocate**](RelocatesAPI.md#CancelRelocate) | **Put** /users/me/relocates/{relocate_uuid}/cancel | Cancel a relocate
[**CompleteRelocate**](RelocatesAPI.md#CompleteRelocate) | **Put** /users/me/relocates/{relocate_uuid}/complete | Complete a relocate
[**GetUserRelocate**](RelocatesAPI.md#GetUserRelocate) | **Get** /users/me/relocates/{relocate_uuid} | Get details of a relocate
[**InitiateRelocate**](RelocatesAPI.md#InitiateRelocate) | **Post** /users/me/relocates | Initiate a relocate from the authenticated user
[**ListUserRelocates**](RelocatesAPI.md#ListUserRelocates) | **Get** /users/me/relocates | Get the relocates of the authenticated user

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
 r, err := apiClient.RelocatesAPI.CancelRelocate(context.Background(), relocateUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RelocatesAPI.CancelRelocate``: %v\n", err)
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
 r, err := apiClient.RelocatesAPI.CompleteRelocate(context.Background(), relocateUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RelocatesAPI.CompleteRelocate``: %v\n", err)
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

## GetUserRelocate

> Relocate GetUserRelocate(ctx, relocateUuid).Execute()

Get details of a relocate

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
 resp, r, err := apiClient.RelocatesAPI.GetUserRelocate(context.Background(), relocateUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RelocatesAPI.GetUserRelocate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserRelocate`: Relocate
 fmt.Fprintf(os.Stdout, "Response from `RelocatesAPI.GetUserRelocate`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relocateUuid** | **string** | Unique identifier of the relocate |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRelocateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 resp, r, err := apiClient.RelocatesAPI.InitiateRelocate(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RelocatesAPI.InitiateRelocate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InitiateRelocate`: Relocate
 fmt.Fprintf(os.Stdout, "Response from `RelocatesAPI.InitiateRelocate`: %v\n", resp)
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
 resp, r, err := apiClient.RelocatesAPI.ListUserRelocates(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RelocatesAPI.ListUserRelocates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserRelocates`: RelocateList
 fmt.Fprintf(os.Stdout, "Response from `RelocatesAPI.ListUserRelocates`: %v\n", resp)
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
