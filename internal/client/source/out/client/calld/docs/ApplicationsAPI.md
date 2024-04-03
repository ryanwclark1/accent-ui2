# \ApplicationsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerApplicationCall**](ApplicationsAPI.md#AnswerApplicationCall) | **Put** /applications/{application_uuid}/calls/{call_id}/answer | Answer a call
[**CreateApplicationCallToNode**](ApplicationsAPI.md#CreateApplicationCallToNode) | **Post** /applications/{application_uuid}/nodes/{node_uuid}/calls | Make a new call to the node
[**CreateApplicationCallToUser**](ApplicationsAPI.md#CreateApplicationCallToUser) | **Post** /applications/{application_uuid}/nodes/{node_uuid}/calls/user | Initiate a call to a user and insert it in the node
[**CreateApplicationCalls**](ApplicationsAPI.md#CreateApplicationCalls) | **Post** /applications/{application_uuid}/calls | Make a new call to the application
[**CreateApplicationNode**](ApplicationsAPI.md#CreateApplicationNode) | **Post** /applications/{application_uuid}/nodes | Make a new node and add calls
[**DeleteApplicationCall**](ApplicationsAPI.md#DeleteApplicationCall) | **Delete** /applications/{application_uuid}/calls/{call_id} | Hangup a call from the application
[**DeleteApplicationCallFromNode**](ApplicationsAPI.md#DeleteApplicationCallFromNode) | **Delete** /applications/{application_uuid}/nodes/{node_uuid}/calls/{call_id} | Remove call from the node
[**DeleteApplicationNode**](ApplicationsAPI.md#DeleteApplicationNode) | **Delete** /applications/{application_uuid}/nodes/{node_uuid} | Delete node and hangup all calls
[**DeletePlayback**](ApplicationsAPI.md#DeletePlayback) | **Delete** /applications/{application_uuid}/playbacks/{playback_uuid} | Stop and remove playback
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /applications/{application_uuid} | Show an application
[**GetApplicationCalls**](ApplicationsAPI.md#GetApplicationCalls) | **Get** /applications/{application_uuid}/calls | List calls from the application
[**GetApplicationNodes**](ApplicationsAPI.md#GetApplicationNodes) | **Get** /applications/{application_uuid}/nodes | List nodes from the application
[**GetNode**](ApplicationsAPI.md#GetNode) | **Get** /applications/{application_uuid}/nodes/{node_uuid} | Show a node
[**GetSnoop**](ApplicationsAPI.md#GetSnoop) | **Get** /applications/{application_uuid}/snoops/{snoop_uuid} | View snooping parameters
[**HoldApplicationCall**](ApplicationsAPI.md#HoldApplicationCall) | **Put** /applications/{application_uuid}/calls/{call_id}/hold/start | Place a call on hold
[**InsertApplicationCallToNode**](ApplicationsAPI.md#InsertApplicationCallToNode) | **Put** /applications/{application_uuid}/nodes/{node_uuid}/calls/{call_id} | Insert call to the node
[**ListApplicationSnoops**](ApplicationsAPI.md#ListApplicationSnoops) | **Get** /applications/{application_uuid}/snoops | List active snoops
[**MuteApplicationCall**](ApplicationsAPI.md#MuteApplicationCall) | **Put** /applications/{application_uuid}/calls/{call_id}/mute/start | Mute a call
[**PlayApplicationCall**](ApplicationsAPI.md#PlayApplicationCall) | **Post** /applications/{application_uuid}/calls/{call_id}/playbacks | Play file to the call
[**ResumeApplicationCall**](ApplicationsAPI.md#ResumeApplicationCall) | **Put** /applications/{application_uuid}/calls/{call_id}/hold/stop | Resume a call that has been placed on hold
[**SendApplicationCallDTMF**](ApplicationsAPI.md#SendApplicationCallDTMF) | **Put** /applications/{application_uuid}/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
[**SnoopApplicationCall**](ApplicationsAPI.md#SnoopApplicationCall) | **Post** /applications/{application_uuid}/calls/{call_id}/snoops | Start snooping on a call
[**StartApplicationCallMOH**](ApplicationsAPI.md#StartApplicationCallMOH) | **Put** /applications/{application_uuid}/calls/{call_id}/moh/{moh_uuid}/start | Starts playing a music on hold
[**StartApplicationCallProgress**](ApplicationsAPI.md#StartApplicationCallProgress) | **Put** /applications/{application_uuid}/calls/{call_id}/progress/start | Play the progress ringing tone
[**StopApplicationCallMOH**](ApplicationsAPI.md#StopApplicationCallMOH) | **Put** /applications/{application_uuid}/calls/{call_id}/moh/stop | Stops playing a music on hold
[**StopApplicationCallProgress**](ApplicationsAPI.md#StopApplicationCallProgress) | **Put** /applications/{application_uuid}/calls/{call_id}/progress/stop | Stop playing the progress ringing tone.
[**StopSnoop**](ApplicationsAPI.md#StopSnoop) | **Delete** /applications/{application_uuid}/snoops/{snoop_uuid} | Stop snooping
[**UnmuteApplicationCall**](ApplicationsAPI.md#UnmuteApplicationCall) | **Put** /applications/{application_uuid}/calls/{call_id}/mute/stop | Unmute a call
[**UpdateSnoop**](ApplicationsAPI.md#UpdateSnoop) | **Put** /applications/{application_uuid}/snoops/{snoop_uuid} | Change snooping parameters

## AnswerApplicationCall

> AnswerApplicationCall(ctx, applicationUuid, callId).Execute()

Answer a call

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.AnswerApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AnswerApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiAnswerApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## CreateApplicationCallToNode

> ApplicationCall CreateApplicationCallToNode(ctx, applicationUuid, nodeUuid).Body(body).Execute()

Make a new call to the node

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
 body := *openapiclient.NewApplicationCallRequestToExten("Context_example", "Exten_example") // ApplicationCallRequestToExten | Call parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.CreateApplicationCallToNode(context.Background(), applicationUuid, nodeUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationCallToNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateApplicationCallToNode`: ApplicationCall
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationCallToNode`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCallToNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationCallRequestToExten**](ApplicationCallRequestToExten.md) | Call parameters |

### Return type

[**ApplicationCall**](ApplicationCall.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateApplicationCallToUser

> ApplicationCall CreateApplicationCallToUser(ctx, applicationUuid, nodeUuid).Body(body).Execute()

Initiate a call to a user and insert it in the node

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
 body := *openapiclient.NewApplicationCallRequestToUser("UserUuid_example") // ApplicationCallRequestToUser | Parameters for the new call
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.CreateApplicationCallToUser(context.Background(), applicationUuid, nodeUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationCallToUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateApplicationCallToUser`: ApplicationCall
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationCallToUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCallToUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationCallRequestToUser**](ApplicationCallRequestToUser.md) | Parameters for the new call |

### Return type

[**ApplicationCall**](ApplicationCall.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateApplicationCalls

> ApplicationCall CreateApplicationCalls(ctx, applicationUuid).Body(body).Execute()

Make a new call to the application

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
 body := *openapiclient.NewApplicationCallRequestToExten("Context_example", "Exten_example") // ApplicationCallRequestToExten | node parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.CreateApplicationCalls(context.Background(), applicationUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateApplicationCalls`: ApplicationCall
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationCalls`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationCallRequestToExten**](ApplicationCallRequestToExten.md) | node parameters |

### Return type

[**ApplicationCall**](ApplicationCall.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateApplicationNode

> ApplicationNode CreateApplicationNode(ctx, applicationUuid).Body(body).Execute()

Make a new node and add calls

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
 body := *openapiclient.NewApplicationNodeRequest() // ApplicationNodeRequest | node parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.CreateApplicationNode(context.Background(), applicationUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateApplicationNode`: ApplicationNode
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationNode`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationNodeRequest**](ApplicationNodeRequest.md) | node parameters |

### Return type

[**ApplicationNode**](ApplicationNode.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteApplicationCall

> DeleteApplicationCall(ctx, applicationUuid, callId).Execute()

Hangup a call from the application

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.DeleteApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## DeleteApplicationCallFromNode

> DeleteApplicationCallFromNode(ctx, applicationUuid, nodeUuid, callId).Execute()

Remove call from the node

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.DeleteApplicationCallFromNode(context.Background(), applicationUuid, nodeUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationCallFromNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationCallFromNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## DeleteApplicationNode

> DeleteApplicationNode(ctx, applicationUuid, nodeUuid).Execute()

Delete node and hangup all calls

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.DeleteApplicationNode(context.Background(), applicationUuid, nodeUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## DeletePlayback

> DeletePlayback(ctx, applicationUuid, playbackUuid).Execute()

Stop and remove playback

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 playbackUuid := "playbackUuid_example" // string | ID of the playback

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.DeletePlayback(context.Background(), applicationUuid, playbackUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeletePlayback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**playbackUuid** | **string** | ID of the playback |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlaybackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## GetApplication

> Application GetApplication(ctx, applicationUuid).Execute()

Show an application

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetApplication`: Application
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Application**](Application.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetApplicationCalls

> ApplicationCalls GetApplicationCalls(ctx, applicationUuid).Execute()

List calls from the application

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.GetApplicationCalls(context.Background(), applicationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationCalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetApplicationCalls`: ApplicationCalls
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationCalls`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCallsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationCalls**](ApplicationCalls.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetApplicationNodes

> ApplicationNodes GetApplicationNodes(ctx, applicationUuid).Execute()

List nodes from the application

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.GetApplicationNodes(context.Background(), applicationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationNodes``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetApplicationNodes`: ApplicationNodes
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationNodes`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationNodesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationNodes**](ApplicationNodes.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetNode

> ApplicationNode GetNode(ctx, applicationUuid, nodeUuid).Execute()

Show a node

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.GetNode(context.Background(), applicationUuid, nodeUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetNode`: ApplicationNode
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetNode`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationNode**](ApplicationNode.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSnoop

> ApplicationSnoop GetSnoop(ctx, applicationUuid, snoopUuid).Execute()

View snooping parameters

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 snoopUuid := "snoopUuid_example" // string | UUID of the snoop

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.GetSnoop(context.Background(), applicationUuid, snoopUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetSnoop``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSnoop`: ApplicationSnoop
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetSnoop`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**snoopUuid** | **string** | UUID of the snoop |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnoopRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationSnoop**](ApplicationSnoop.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HoldApplicationCall

> HoldApplicationCall(ctx, applicationUuid, callId).Execute()

Place a call on hold

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.HoldApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.HoldApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiHoldApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## InsertApplicationCallToNode

> InsertApplicationCallToNode(ctx, applicationUuid, nodeUuid, callId).Execute()

Insert call to the node

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 nodeUuid := "nodeUuid_example" // string | UUID of the node
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.InsertApplicationCallToNode(context.Background(), applicationUuid, nodeUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.InsertApplicationCallToNode``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**nodeUuid** | **string** | UUID of the node |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiInsertApplicationCallToNodeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## ListApplicationSnoops

> ApplicationSnoops ListApplicationSnoops(ctx, applicationUuid).Execute()

List active snoops

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.ListApplicationSnoops(context.Background(), applicationUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationSnoops``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListApplicationSnoops`: ApplicationSnoops
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationSnoops`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationSnoopsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationSnoops**](ApplicationSnoops.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MuteApplicationCall

> MuteApplicationCall(ctx, applicationUuid, callId).Execute()

Mute a call

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.MuteApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.MuteApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiMuteApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## PlayApplicationCall

> ApplicationPlayback PlayApplicationCall(ctx, applicationUuid, callId).Body(body).Execute()

Play file to the call

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
 body := *openapiclient.NewApplicationPlayback() // ApplicationPlayback | playback parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.PlayApplicationCall(context.Background(), applicationUuid, callId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.PlayApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PlayApplicationCall`: ApplicationPlayback
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.PlayApplicationCall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiPlayApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationPlayback**](ApplicationPlayback.md) | playback parameters |

### Return type

[**ApplicationPlayback**](ApplicationPlayback.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ResumeApplicationCall

> ResumeApplicationCall(ctx, applicationUuid, callId).Execute()

Resume a call that has been placed on hold

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.ResumeApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ResumeApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiResumeApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## SendApplicationCallDTMF

> SendApplicationCallDTMF(ctx, applicationUuid, callId).Digits(digits).Execute()

Simulate a user pressing DTMF keys

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call
 digits := "digits_example" // string | Digits to send via DTMF. Must contain only `0-9*#`.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.SendApplicationCallDTMF(context.Background(), applicationUuid, callId).Digits(digits).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.SendApplicationCallDTMF``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiSendApplicationCallDTMFRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digits** | **string** | Digits to send via DTMF. Must contain only &#x60;0-9*#&#x60;. |

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

## SnoopApplicationCall

> ApplicationSnoop SnoopApplicationCall(ctx, applicationUuid, callId).Body(body).Execute()

Start snooping on a call

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
 body := *openapiclient.NewApplicationSnoop("SnoopingCallId_example", "WhisperMode_example") // ApplicationSnoop | snoop parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ApplicationsAPI.SnoopApplicationCall(context.Background(), applicationUuid, callId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.SnoopApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SnoopApplicationCall`: ApplicationSnoop
 fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.SnoopApplicationCall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiSnoopApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationSnoop**](ApplicationSnoop.md) | snoop parameters |

### Return type

[**ApplicationSnoop**](ApplicationSnoop.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StartApplicationCallMOH

> StartApplicationCallMOH(ctx, applicationUuid, callId, mohUuid).Execute()

Starts playing a music on hold

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call
 mohUuid := "mohUuid_example" // string | UUID of the music on hold

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.StartApplicationCallMOH(context.Background(), applicationUuid, callId, mohUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.StartApplicationCallMOH``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |
**mohUuid** | **string** | UUID of the music on hold |

### Other Parameters

Other parameters are passed through a pointer to a apiStartApplicationCallMOHRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## StartApplicationCallProgress

> StartApplicationCallProgress(ctx, applicationUuid, callId).Execute()

Play the progress ringing tone

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.StartApplicationCallProgress(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.StartApplicationCallProgress``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiStartApplicationCallProgressRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## StopApplicationCallMOH

> StopApplicationCallMOH(ctx, applicationUuid, callId).Execute()

Stops playing a music on hold

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.StopApplicationCallMOH(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.StopApplicationCallMOH``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiStopApplicationCallMOHRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## StopApplicationCallProgress

> StopApplicationCallProgress(ctx, applicationUuid, callId).Execute()

Stop playing the progress ringing tone.

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.StopApplicationCallProgress(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.StopApplicationCallProgress``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiStopApplicationCallProgressRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## StopSnoop

> StopSnoop(ctx, applicationUuid, snoopUuid).Execute()

Stop snooping

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 snoopUuid := "snoopUuid_example" // string | UUID of the snoop

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.StopSnoop(context.Background(), applicationUuid, snoopUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.StopSnoop``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**snoopUuid** | **string** | UUID of the snoop |

### Other Parameters

Other parameters are passed through a pointer to a apiStopSnoopRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## UnmuteApplicationCall

> UnmuteApplicationCall(ctx, applicationUuid, callId).Execute()

Unmute a call

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
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.UnmuteApplicationCall(context.Background(), applicationUuid, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UnmuteApplicationCall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteApplicationCallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## UpdateSnoop

> UpdateSnoop(ctx, applicationUuid, snoopUuid).Body(body).Execute()

Change snooping parameters

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
 body := *openapiclient.NewApplicationSnoopPut() // ApplicationSnoopPut | snoop parameters
 applicationUuid := "applicationUuid_example" // string | UUID of the application
 snoopUuid := "snoopUuid_example" // string | UUID of the snoop

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ApplicationsAPI.UpdateSnoop(context.Background(), applicationUuid, snoopUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateSnoop``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string** | UUID of the application |
**snoopUuid** | **string** | UUID of the snoop |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnoopRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationSnoopPut**](ApplicationSnoopPut.md) | snoop parameters |

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
