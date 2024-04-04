# \SetupAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SetupAPI.md#Create) | **Post** /setup | Setup the Accent Engine

## Create

> Create(ctx).Body(body).Execute()

Setup the Accent Engine

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/setupd"
)

func main() {
 body := *openapiclient.NewSetupRequest("EngineLanguage_example", false, "EnginePassword_example") // SetupRequest | Setup options

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SetupAPI.Create(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SetupAPI.Create``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetupRequest**](SetupRequest.md) | Setup options |

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
