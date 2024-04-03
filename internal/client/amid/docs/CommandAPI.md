# \CommandAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandAsteriskManager**](CommandAPI.md#CommandAsteriskManager) | **Post** /action/Command | AMI command

## CommandAsteriskManager

> CommandResponse CommandAsteriskManager(ctx).Command(command).Execute()

AMI command

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
 command := *openapiclient.NewCommand("Command_example") // Command | The command to send to the manager.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CommandAPI.CommandAsteriskManager(context.Background()).Command(command).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CommandAPI.CommandAsteriskManager``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CommandAsteriskManager`: CommandResponse
 fmt.Fprintf(os.Stdout, "Response from `CommandAPI.CommandAsteriskManager`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCommandAsteriskManagerRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**Command**](Command.md) | The command to send to the manager.  |

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
