# \ProvisioningAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvisioningNetworking**](ProvisioningAPI.md#GetProvisioningNetworking) | **Get** /provisioning/networking | Get Provisioning Networking configuration
[**UpdateProvisioningNetworking**](ProvisioningAPI.md#UpdateProvisioningNetworking) | **Put** /provisioning/networking | Update Provisioning Networking configuration

## GetProvisioningNetworking

> ProvisioningNetworking GetProvisioningNetworking(ctx).Execute()

Get Provisioning Networking configuration

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ProvisioningAPI.GetProvisioningNetworking(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningNetworking``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetProvisioningNetworking`: ProvisioningNetworking
 fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningNetworking`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningNetworkingRequest struct via the builder pattern

### Return type

[**ProvisioningNetworking**](ProvisioningNetworking.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateProvisioningNetworking

> UpdateProvisioningNetworking(ctx).Body(body).Execute()

Update Provisioning Networking configuration

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
 body := *openapiclient.NewProvisioningNetworking() // ProvisioningNetworking | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvisioningAPI.UpdateProvisioningNetworking(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.UpdateProvisioningNetworking``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningNetworkingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProvisioningNetworking**](ProvisioningNetworking.md) |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
