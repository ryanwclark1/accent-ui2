# \ActionAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionAsteriskManager**](ActionAPI.md#ActionAsteriskManager) | **Post** /action/{action} | AMI action

## ActionAsteriskManager

> Response ActionAsteriskManager(ctx, action).ActionArguments(actionArguments).Execute()

AMI action

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
 action := "action_example" // string | Name of the manager action. Currently not supported: Queues, Command.
 actionArguments := map[string]interface{}{ ... } // map[string]interface{} | Arguments for the manager action. Action: taken from the URL ActionID: not necessary If you need a same key multiple times, give a list of values. For Action: Originate, you should always use Async: True (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ActionAPI.ActionAsteriskManager(context.Background(), action).ActionArguments(actionArguments).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ActionAPI.ActionAsteriskManager``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ActionAsteriskManager`: Response
 fmt.Fprintf(os.Stdout, "Response from `ActionAPI.ActionAsteriskManager`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string** | Name of the manager action. Currently not supported: Queues, Command. |

### Other Parameters

Other parameters are passed through a pointer to a apiActionAsteriskManagerRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionArguments** | **map[string]interface{}** | Arguments for the manager action. Action: taken from the URL ActionID: not necessary If you need a same key multiple times, give a list of values. For Action: Originate, you should always use Async: True |

### Return type

[**Response**](Response.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
