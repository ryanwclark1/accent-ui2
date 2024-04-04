# \EmailsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailsConfig**](EmailsAPI.md#GetEmailsConfig) | **Get** /emails | Get e-mail configuration
[**UpdateEmailsConfig**](EmailsAPI.md#UpdateEmailsConfig) | **Put** /emails | Update e-mail configuration

## GetEmailsConfig

> EmailConfig GetEmailsConfig(ctx).Execute()

Get e-mail configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EmailsAPI.GetEmailsConfig(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.GetEmailsConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEmailsConfig`: EmailConfig
 fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.GetEmailsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailsConfigRequest struct via the builder pattern

### Return type

[**EmailConfig**](EmailConfig.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateEmailsConfig

> UpdateEmailsConfig(ctx).Body(body).Execute()

Update e-mail configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewEmailConfig() // EmailConfig | E-mail configuration

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EmailsAPI.UpdateEmailsConfig(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.UpdateEmailsConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailsConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EmailConfig**](EmailConfig.md) | E-mail configuration |

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
