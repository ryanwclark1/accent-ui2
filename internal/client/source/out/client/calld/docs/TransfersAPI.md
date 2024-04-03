# \TransfersAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTransfer**](TransfersAPI.md#CancelTransfer) | **Delete** /transfers/{transfer_id} | Cancel a transfer
[**CancelUserTransfer**](TransfersAPI.md#CancelUserTransfer) | **Delete** /users/me/transfers/{transfer_id} | Cancel a transfer
[**CompleteTransfer**](TransfersAPI.md#CompleteTransfer) | **Put** /transfers/{transfer_id}/complete | Complete a transfer
[**CompleteUserTransfer**](TransfersAPI.md#CompleteUserTransfer) | **Put** /users/me/transfers/{transfer_id}/complete | Complete a transfer
[**GetTransfer**](TransfersAPI.md#GetTransfer) | **Get** /transfers/{transfer_id} | Get details of a transfer
[**InitiateTransfer**](TransfersAPI.md#InitiateTransfer) | **Post** /transfers | Initiate a transfer
[**InitiateUserTransfer**](TransfersAPI.md#InitiateUserTransfer) | **Post** /users/me/transfers | Initiate a transfer from the authenticated user
[**ListUserTransfers**](TransfersAPI.md#ListUserTransfers) | **Get** /users/me/transfers | Get the transfers of the authenticated user

## CancelTransfer

> CancelTransfer(ctx, transferId).Execute()

Cancel a transfer

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
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TransfersAPI.CancelTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.CancelTransfer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelTransferRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TransfersAPI.CancelUserTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.CancelUserTransfer``: %v\n", err)
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

## CompleteTransfer

> CompleteTransfer(ctx, transferId).Execute()

Complete a transfer

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
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TransfersAPI.CompleteTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.CompleteTransfer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCompleteTransferRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TransfersAPI.CompleteUserTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.CompleteUserTransfer``: %v\n", err)
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

## GetTransfer

> Transfer GetTransfer(ctx, transferId).Execute()

Get details of a transfer

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
 transferId := "transferId_example" // string | Unique identifier of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TransfersAPI.GetTransfer(context.Background(), transferId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.GetTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetTransfer`: Transfer
 fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.GetTransfer`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Unique identifier of the transfer |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## InitiateTransfer

> Transfer InitiateTransfer(ctx).Body(body).Execute()

Initiate a transfer

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
 body := *openapiclient.NewTransferRequest("Context_example", "Exten_example", "InitiatorCall_example", "TransferredCall_example") // TransferRequest | Parameters of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TransfersAPI.InitiateTransfer(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.InitiateTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InitiateTransfer`: Transfer
 fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.InitiateTransfer`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateTransferRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TransferRequest**](TransferRequest.md) | Parameters of the transfer |

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 body := *openapiclient.NewUserTransferRequest("Exten_example", "InitiatorCall_example") // UserTransferRequest | Parameters of the transfer

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TransfersAPI.InitiateUserTransfer(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.InitiateUserTransfer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InitiateUserTransfer`: Transfer
 fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.InitiateUserTransfer`: %v\n", resp)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TransfersAPI.ListUserTransfers(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.ListUserTransfers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserTransfers`: TransferList
 fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.ListUserTransfers`: %v\n", resp)
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
