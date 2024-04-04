# \ConfigAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigAPI.md#GetConfig) | **Get** /config | Show the current configuration
[**PatchConfig**](ConfigAPI.md#PatchConfig) | **Patch** /config | Update the current configuration.

## GetConfig

> GetConfig(ctx).Execute()

Show the current configuration

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConfigAPI.GetConfig(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PatchConfig

> PatchConfig(ctx).ConfigPatch(configPatch).Execute()

Update the current configuration.

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
 configPatch := []openapiclient.ConfigPatchItem{*openapiclient.NewConfigPatchItem()} // []ConfigPatchItem | See https://en.wikipedia.org/wiki/JSON_Patch.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConfigAPI.PatchConfig(context.Background()).ConfigPatch(configPatch).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.PatchConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configPatch** | [**[]ConfigPatchItem**](ConfigPatchItem.md) | See <https://en.wikipedia.org/wiki/JSON_Patch>. |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
