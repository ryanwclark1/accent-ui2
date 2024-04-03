# \MicrosoftAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateTeamsPresence**](MicrosoftAPI.md#UpdateTeamsPresence) | **Post** /users/{user_uuid}/teams/presence | Receive presence information from Microsoft Teams

## UpdateTeamsPresence

> Object UpdateTeamsPresence(ctx, userUuid).Execute()

Receive presence information from Microsoft Teams

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/chatd"
)

func main() {
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MicrosoftAPI.UpdateTeamsPresence(context.Background(), userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MicrosoftAPI.UpdateTeamsPresence``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateTeamsPresence`: Object
 fmt.Fprintf(os.Stdout, "Response from `MicrosoftAPI.UpdateTeamsPresence`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamsPresenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Object**](Object.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
