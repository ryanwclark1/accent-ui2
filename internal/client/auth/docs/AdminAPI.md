# \AdminAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAllUserEmails**](AdminAPI.md#UpdateAllUserEmails) | **Put** /admin/users/{user_uuid}/emails | Update email addresses

## UpdateAllUserEmails

> UpdateAllUserEmails(ctx, userUuid).Body(body).Execute()

Update email addresses

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
 body := *openapiclient.NewAdminUserEmailList() // AdminUserEmailList | EmailAddressList
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AdminAPI.UpdateAllUserEmails(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateAllUserEmails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllUserEmailsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdminUserEmailList**](AdminUserEmailList.md) | EmailAddressList |

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
