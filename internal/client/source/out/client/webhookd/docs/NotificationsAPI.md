# \NotificationsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMobileNotification**](NotificationsAPI.md#PostMobileNotification) | **Post** /mobile/notifications | Send a push notification to a user

## PostMobileNotification

> PostMobileNotification(ctx).Body(body).AccentTenant(accentTenant).Execute()

Send a push notification to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/webhookd"
)

func main() {
 body := *openapiclient.NewNotification("Body_example", "myCustomNotification", "Title_example", "UserUuid_example") // Notification | 
 accentTenant := "accentTenant_example" // string | The User's tenant UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.NotificationsAPI.PostMobileNotification(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.PostMobileNotification``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostMobileNotificationRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Notification**](Notification.md) |  |
 **accentTenant** | **string** | The User&#39;s tenant UUID, defining the ownership of a given resource. |

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
