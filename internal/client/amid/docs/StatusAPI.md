# \StatusAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusSummary**](StatusAPI.md#StatusSummary) | **Get** /status | Print infos about internal status of accent-amid

## StatusSummary

> StatusSummary StatusSummary(ctx).Execute()

Print infos about internal status of accent-amid

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/amid"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.StatusAPI.StatusSummary(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.StatusSummary``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `StatusSummary`: StatusSummary
 fmt.Fprintf(os.Stdout, "Response from `StatusAPI.StatusSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusSummaryRequest struct via the builder pattern

### Return type

[**StatusSummary**](StatusSummary.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
