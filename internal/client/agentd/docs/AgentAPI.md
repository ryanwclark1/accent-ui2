# \AgentAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAgentById**](AgentAPI.md#AddAgentById) | **Post** /agents/by-id/{agent_id}/add | Add agent to a queue.
[**GetAgentById**](AgentAPI.md#GetAgentById) | **Get** /agents/by-id/{agent_id} | Get agent status.
[**GetAgentByNumber**](AgentAPI.md#GetAgentByNumber) | **Get** /agents/by-number/{agent_number} | Get agent status.
[**GetUserAgent**](AgentAPI.md#GetUserAgent) | **Get** /users/me/agents | Get agent status of the user holding the authentication token.
[**LoginAgentById**](AgentAPI.md#LoginAgentById) | **Post** /agents/by-id/{agent_id}/login | Log an agent.
[**LoginAgentByNumber**](AgentAPI.md#LoginAgentByNumber) | **Post** /agents/by-number/{agent_number}/login | Log an agent.
[**LoginUserAgent**](AgentAPI.md#LoginUserAgent) | **Post** /users/me/agents/login | Log the agent of the user holding the authentication token
[**LogoffAgentById**](AgentAPI.md#LogoffAgentById) | **Post** /agents/by-id/{agent_id}/logoff | Logoff an agent.
[**LogoffAgentByNumber**](AgentAPI.md#LogoffAgentByNumber) | **Post** /agents/by-number/{agent_number}/logoff | Logoff an agent.
[**LogoffUserAgent**](AgentAPI.md#LogoffUserAgent) | **Post** /users/me/agents/logoff | Logoff the agent of the user holding the authentication token
[**PauseAgentByNumber**](AgentAPI.md#PauseAgentByNumber) | **Post** /agents/by-number/{agent_number}/pause | Pause an agent.
[**PauseUserAgent**](AgentAPI.md#PauseUserAgent) | **Post** /users/me/agents/pause | Pause the agent of the user holding the authentication token
[**RemoveAgentById**](AgentAPI.md#RemoveAgentById) | **Post** /agents/by-id/{agent_id}/remove | Remove agent from a queue.
[**UnpauseAgentByNumber**](AgentAPI.md#UnpauseAgentByNumber) | **Post** /agents/by-number/{agent_number}/unpause | Unpause an agent.
[**UnpauseUserAgent**](AgentAPI.md#UnpauseUserAgent) | **Post** /users/me/agents/unpause | Unpause the agent of the user holding the authentication token

## AddAgentById

> AddAgentById(ctx, agentId).Body(body).AccentTenant(accentTenant).Execute()

Add agent to a queue.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewQueue() // Queue | The queue to add the agent to
 agentId := int32(56) // int32 | Agent's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.AddAgentById(context.Background(), agentId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AddAgentById``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | Agent&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAddAgentByIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Queue**](Queue.md) | The queue to add the agent to |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## GetAgentById

> AgentStatus GetAgentById(ctx, agentId).AccentTenant(accentTenant).Execute()

Get agent status.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentId := int32(56) // int32 | Agent's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AgentAPI.GetAgentById(context.Background(), agentId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentById``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetAgentById`: AgentStatus
 fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentById`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | Agent&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**AgentStatus**](AgentStatus.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAgentByNumber

> AgentStatus GetAgentByNumber(ctx, agentNumber).AccentTenant(accentTenant).Execute()

Get agent status.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentNumber := "agentNumber_example" // string | Agent's number
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AgentAPI.GetAgentByNumber(context.Background(), agentNumber).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentByNumber``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetAgentByNumber`: AgentStatus
 fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentByNumber`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentNumber** | **string** | Agent&#39;s number |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByNumberRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**AgentStatus**](AgentStatus.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserAgent

> AgentStatus GetUserAgent(ctx).AccentTenant(accentTenant).Execute()

Get agent status of the user holding the authentication token.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AgentAPI.GetUserAgent(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserAgent`: AgentStatus
 fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetUserAgent`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAgentRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**AgentStatus**](AgentStatus.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LoginAgentById

> LoginAgentById(ctx, agentId).Body(body).AccentTenant(accentTenant).Execute()

Log an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewLoginInfo() // LoginInfo | The extension and context on which to log the agent
 agentId := int32(56) // int32 | Agent's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LoginAgentById(context.Background(), agentId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LoginAgentById``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | Agent&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiLoginAgentByIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LoginInfo**](LoginInfo.md) | The extension and context on which to log the agent |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## LoginAgentByNumber

> LoginAgentByNumber(ctx, agentNumber).Body(body).AccentTenant(accentTenant).Execute()

Log an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewLoginInfo() // LoginInfo | The extension and context on which to log the agent
 agentNumber := "agentNumber_example" // string | Agent's number
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LoginAgentByNumber(context.Background(), agentNumber).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LoginAgentByNumber``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentNumber** | **string** | Agent&#39;s number |

### Other Parameters

Other parameters are passed through a pointer to a apiLoginAgentByNumberRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LoginInfo**](LoginInfo.md) | The extension and context on which to log the agent |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## LoginUserAgent

> LoginUserAgent(ctx).Body(body).AccentTenant(accentTenant).Execute()

Log the agent of the user holding the authentication token

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewUserAgentLoginInfo() // UserAgentLoginInfo | The line on which to log the agent
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LoginUserAgent(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LoginUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiLoginUserAgentRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserAgentLoginInfo**](UserAgentLoginInfo.md) | The line on which to log the agent |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## LogoffAgentById

> LogoffAgentById(ctx, agentId).AccentTenant(accentTenant).Execute()

Logoff an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentId := int32(56) // int32 | Agent's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LogoffAgentById(context.Background(), agentId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LogoffAgentById``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | Agent&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiLogoffAgentByIdRequest struct via the builder pattern

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

## LogoffAgentByNumber

> LogoffAgentByNumber(ctx, agentNumber).AccentTenant(accentTenant).Execute()

Logoff an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentNumber := "agentNumber_example" // string | Agent's number
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LogoffAgentByNumber(context.Background(), agentNumber).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LogoffAgentByNumber``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentNumber** | **string** | Agent&#39;s number |

### Other Parameters

Other parameters are passed through a pointer to a apiLogoffAgentByNumberRequest struct via the builder pattern

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

## LogoffUserAgent

> LogoffUserAgent(ctx).AccentTenant(accentTenant).Execute()

Logoff the agent of the user holding the authentication token

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.LogoffUserAgent(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.LogoffUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiLogoffUserAgentRequest struct via the builder pattern

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

## PauseAgentByNumber

> PauseAgentByNumber(ctx, agentNumber).Body(body).AccentTenant(accentTenant).Execute()

Pause an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentNumber := "agentNumber_example" // string | Agent's number
 body := *openapiclient.NewAgentPauseReason() // AgentPauseReason | The reason for pausing the agent (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.PauseAgentByNumber(context.Background(), agentNumber).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PauseAgentByNumber``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentNumber** | **string** | Agent&#39;s number |

### Other Parameters

Other parameters are passed through a pointer to a apiPauseAgentByNumberRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AgentPauseReason**](AgentPauseReason.md) | The reason for pausing the agent |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## PauseUserAgent

> PauseUserAgent(ctx).Body(body).AccentTenant(accentTenant).Execute()

Pause the agent of the user holding the authentication token

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewAgentPauseReason() // AgentPauseReason | The reason for pausing the agent (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.PauseUserAgent(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.PauseUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPauseUserAgentRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AgentPauseReason**](AgentPauseReason.md) | The reason for pausing the agent |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## RemoveAgentById

> RemoveAgentById(ctx, agentId).Body(body).AccentTenant(accentTenant).Execute()

Remove agent from a queue.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 body := *openapiclient.NewQueue() // Queue | The queue to remove the agent from
 agentId := int32(56) // int32 | Agent's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.RemoveAgentById(context.Background(), agentId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.RemoveAgentById``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | Agent&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAgentByIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Queue**](Queue.md) | The queue to remove the agent from |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## UnpauseAgentByNumber

> UnpauseAgentByNumber(ctx, agentNumber).AccentTenant(accentTenant).Execute()

Unpause an agent.

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 agentNumber := "agentNumber_example" // string | Agent's number
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.UnpauseAgentByNumber(context.Background(), agentNumber).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.UnpauseAgentByNumber``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentNumber** | **string** | Agent&#39;s number |

### Other Parameters

Other parameters are passed through a pointer to a apiUnpauseAgentByNumberRequest struct via the builder pattern

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

## UnpauseUserAgent

> UnpauseUserAgent(ctx).AccentTenant(accentTenant).Execute()

Unpause the agent of the user holding the authentication token

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AgentAPI.UnpauseUserAgent(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.UnpauseUserAgent``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUnpauseUserAgentRequest struct via the builder pattern

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
