# \SwitchboardsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerHeldCall**](SwitchboardsAPI.md#AnswerHeldCall) | **Put** /switchboards/{switchboard_uuid}/calls/held/{call_id}/answer | Answer the specified held call
[**AnswerQueuedCall**](SwitchboardsAPI.md#AnswerQueuedCall) | **Put** /switchboards/{switchboard_uuid}/calls/queued/{call_id}/answer | Answer the specified queued call
[**HoldSwitchboardCall**](SwitchboardsAPI.md#HoldSwitchboardCall) | **Put** /switchboards/{switchboard_uuid}/calls/held/{call_id} | Put the specified call on hold in the switchboard
[**ListSwitchboardHeldCalls**](SwitchboardsAPI.md#ListSwitchboardHeldCalls) | **Get** /switchboards/{switchboard_uuid}/calls/held | List calls held in the switchboard
[**ListSwitchboardQueuedCalls**](SwitchboardsAPI.md#ListSwitchboardQueuedCalls) | **Get** /switchboards/{switchboard_uuid}/calls/queued | List calls queued in the switchboard

## AnswerHeldCall

> CallID AnswerHeldCall(ctx, switchboardUuid, callId).AccentTenant(accentTenant).LineId(lineId).Execute()

Answer the specified held call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 switchboardUuid := "switchboardUuid_example" // string | Unique identifier of the switchboard
 callId := "callId_example" // string | ID of the call
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 lineId := int32(56) // int32 | ID of the line of the user used to make the call. Default is the main line of the user. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.AnswerHeldCall(context.Background(), switchboardUuid, callId).AccentTenant(accentTenant).LineId(lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.AnswerHeldCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `AnswerHeldCall`: CallID
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.AnswerHeldCall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** | Unique identifier of the switchboard |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerHeldCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **lineId** | **int32** | ID of the line of the user used to make the call. Default is the main line of the user. |

### Return type

[**CallID**](CallID.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## AnswerQueuedCall

> CallID AnswerQueuedCall(ctx, switchboardUuid, callId).AccentTenant(accentTenant).LineId(lineId).Execute()

Answer the specified queued call

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 switchboardUuid := "switchboardUuid_example" // string | Unique identifier of the switchboard
 callId := "callId_example" // string | ID of the call
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 lineId := int32(56) // int32 | ID of the line of the user used to make the call. Default is the main line of the user. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.AnswerQueuedCall(context.Background(), switchboardUuid, callId).AccentTenant(accentTenant).LineId(lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.AnswerQueuedCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `AnswerQueuedCall`: CallID
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.AnswerQueuedCall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** | Unique identifier of the switchboard |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerQueuedCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **lineId** | **int32** | ID of the line of the user used to make the call. Default is the main line of the user. |

### Return type

[**CallID**](CallID.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HoldSwitchboardCall

> HoldSwitchboardCall(ctx, switchboardUuid, callId).AccentTenant(accentTenant).Execute()

Put the specified call on hold in the switchboard

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 switchboardUuid := "switchboardUuid_example" // string | Unique identifier of the switchboard
 callId := "callId_example" // string | ID of the call
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SwitchboardsAPI.HoldSwitchboardCall(context.Background(), switchboardUuid, callId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.HoldSwitchboardCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** | Unique identifier of the switchboard |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiHoldSwitchboardCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## ListSwitchboardHeldCalls

> SwitchboardHeldCalls ListSwitchboardHeldCalls(ctx, switchboardUuid).AccentTenant(accentTenant).Execute()

List calls held in the switchboard

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 switchboardUuid := "switchboardUuid_example" // string | Unique identifier of the switchboard
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.ListSwitchboardHeldCalls(context.Background(), switchboardUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.ListSwitchboardHeldCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListSwitchboardHeldCalls`: SwitchboardHeldCalls
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.ListSwitchboardHeldCalls`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** | Unique identifier of the switchboard |

### Other Parameters

Other parameters are passed through a pointer to a apiListSwitchboardHeldCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**SwitchboardHeldCalls**](SwitchboardHeldCalls.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSwitchboardQueuedCalls

> SwitchboardQueuedCalls ListSwitchboardQueuedCalls(ctx, switchboardUuid).AccentTenant(accentTenant).Execute()

List calls queued in the switchboard

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 switchboardUuid := "switchboardUuid_example" // string | Unique identifier of the switchboard
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.ListSwitchboardQueuedCalls(context.Background(), switchboardUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.ListSwitchboardQueuedCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListSwitchboardQueuedCalls`: SwitchboardQueuedCalls
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.ListSwitchboardQueuedCalls`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** | Unique identifier of the switchboard |

### Other Parameters

Other parameters are passed through a pointer to a apiListSwitchboardQueuedCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**SwitchboardQueuedCalls**](SwitchboardQueuedCalls.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
