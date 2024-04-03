# \VoicemailsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserVoicemailGreeting**](VoicemailsAPI.md#CheckUserVoicemailGreeting) | **Head** /users/me/voicemails/greetings/{greeting} | Check if greeting exists
[**CheckVoicemailGreeting**](VoicemailsAPI.md#CheckVoicemailGreeting) | **Head** /voicemails/{voicemail_id}/greetings/{greeting} | Check if greeting exists
[**CopyUserVoicemailGreeting**](VoicemailsAPI.md#CopyUserVoicemailGreeting) | **Post** /users/me/voicemails/greetings/{greeting}/copy | Copy a custom greeting
[**CopyVoicemailGreeting**](VoicemailsAPI.md#CopyVoicemailGreeting) | **Post** /voicemails/{voicemail_id}/greetings/{greeting}/copy | Copy a custom greeting
[**CreateUserVoicemailGreeting**](VoicemailsAPI.md#CreateUserVoicemailGreeting) | **Post** /users/me/voicemails/greetings/{greeting} | Create a custom greeting
[**CreateVoicemailGreeting**](VoicemailsAPI.md#CreateVoicemailGreeting) | **Post** /voicemails/{voicemail_id}/greetings/{greeting} | Create a custom greeting
[**DeleteUserVoicemailGreeting**](VoicemailsAPI.md#DeleteUserVoicemailGreeting) | **Delete** /users/me/voicemails/greetings/{greeting} | Delete a custom greeting
[**DeleteUserVoicemailMessage**](VoicemailsAPI.md#DeleteUserVoicemailMessage) | **Delete** /users/me/voicemails/messages/{message_id} | Delete a mesage
[**DeleteVoicemailGreeting**](VoicemailsAPI.md#DeleteVoicemailGreeting) | **Delete** /voicemails/{voicemail_id}/greetings/{greeting} | Delete a custom greeting
[**DeleteVoicemailMessage**](VoicemailsAPI.md#DeleteVoicemailMessage) | **Delete** /voicemails/{voicemail_id}/messages/{message_id} | Delete a mesage
[**GetUserVoicemailFolder**](VoicemailsAPI.md#GetUserVoicemailFolder) | **Get** /users/me/voicemails/folders/{folder_id} | Get details of a folder
[**GetUserVoicemailGreeting**](VoicemailsAPI.md#GetUserVoicemailGreeting) | **Get** /users/me/voicemails/greetings/{greeting} | Get a custom greeting
[**GetUserVoicemailMessage**](VoicemailsAPI.md#GetUserVoicemailMessage) | **Get** /users/me/voicemails/messages/{message_id} | Get a message
[**GetUserVoicemailMessageRecording**](VoicemailsAPI.md#GetUserVoicemailMessageRecording) | **Get** /users/me/voicemails/messages/{message_id}/recording | Get a message&#39;s recording
[**GetVoicemail**](VoicemailsAPI.md#GetVoicemail) | **Get** /voicemails/{voicemail_id} | Get details of a voicemail
[**GetVoicemailFolder**](VoicemailsAPI.md#GetVoicemailFolder) | **Get** /voicemails/{voicemail_id}/folders/{folder_id} | Get details of a folder
[**GetVoicemailGreeting**](VoicemailsAPI.md#GetVoicemailGreeting) | **Get** /voicemails/{voicemail_id}/greetings/{greeting} | Get a custom greeting
[**GetVoicemailMessage**](VoicemailsAPI.md#GetVoicemailMessage) | **Get** /voicemails/{voicemail_id}/messages/{message_id} | Get a message
[**GetVoicemailMessageRecording**](VoicemailsAPI.md#GetVoicemailMessageRecording) | **Get** /voicemails/{voicemail_id}/messages/{message_id}/recording | Get a message&#39;s recording
[**ListUserVoicemails**](VoicemailsAPI.md#ListUserVoicemails) | **Get** /users/me/voicemails | Get details of the voicemail of the authenticated user
[**UpdateUserVoicemailGreeting**](VoicemailsAPI.md#UpdateUserVoicemailGreeting) | **Put** /users/me/voicemails/greetings/{greeting} | Update a custom greeting
[**UpdateUserVoicemailMessage**](VoicemailsAPI.md#UpdateUserVoicemailMessage) | **Put** /users/me/voicemails/messages/{message_id} | Update a message
[**UpdateVoicemailGreeting**](VoicemailsAPI.md#UpdateVoicemailGreeting) | **Put** /voicemails/{voicemail_id}/greetings/{greeting} | Update a custom greeting
[**UpdateVoicemailMessage**](VoicemailsAPI.md#UpdateVoicemailMessage) | **Put** /voicemails/{voicemail_id}/messages/{message_id} | Update a message

## CheckUserVoicemailGreeting

> CheckUserVoicemailGreeting(ctx, greeting).Execute()

Check if greeting exists

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
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CheckUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CheckUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserVoicemailGreetingRequest struct via the builder pattern

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

## CheckVoicemailGreeting

> CheckVoicemailGreeting(ctx, voicemailId, greeting).Execute()

Check if greeting exists

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CheckVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CheckVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCheckVoicemailGreetingRequest struct via the builder pattern

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

## CopyUserVoicemailGreeting

> CopyUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Copy a custom greeting

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
 body := *openapiclient.NewGreetingCopy("DestGreeting_example") // GreetingCopy | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CopyUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CopyUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCopyUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GreetingCopy**](GreetingCopy.md) |  |

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

## CopyVoicemailGreeting

> CopyVoicemailGreeting(ctx, voicemailId, greeting).Body(body).Execute()

Copy a custom greeting

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
 body := *openapiclient.NewGreetingCopy("DestGreeting_example") // GreetingCopy | 
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CopyVoicemailGreeting(context.Background(), voicemailId, greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CopyVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCopyVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GreetingCopy**](GreetingCopy.md) |  |

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

## CreateUserVoicemailGreeting

> CreateUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Create a custom greeting

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
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CreateUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CreateUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateVoicemailGreeting

> CreateVoicemailGreeting(ctx, voicemailId, greeting).Body(body).Execute()

Create a custom greeting

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
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.CreateVoicemailGreeting(context.Background(), voicemailId, greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CreateVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserVoicemailGreeting

> DeleteUserVoicemailGreeting(ctx, greeting).Execute()

Delete a custom greeting

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
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DeleteUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DeleteUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserVoicemailGreetingRequest struct via the builder pattern

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

## DeleteUserVoicemailMessage

> DeleteUserVoicemailMessage(ctx, messageId).Execute()

Delete a mesage

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
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DeleteUserVoicemailMessage(context.Background(), messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DeleteUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserVoicemailMessageRequest struct via the builder pattern

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

## DeleteVoicemailGreeting

> DeleteVoicemailGreeting(ctx, voicemailId, greeting).Execute()

Delete a custom greeting

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DeleteVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DeleteVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoicemailGreetingRequest struct via the builder pattern

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

## DeleteVoicemailMessage

> DeleteVoicemailMessage(ctx, voicemailId, messageId).Execute()

Delete a mesage

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DeleteVoicemailMessage(context.Background(), voicemailId, messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DeleteVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoicemailMessageRequest struct via the builder pattern

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

## GetUserVoicemailFolder

> VoicemailFolder GetUserVoicemailFolder(ctx, folderId).Execute()

Get details of a folder

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
 folderId := int32(56) // int32 | The folder's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetUserVoicemailFolder(context.Background(), folderId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetUserVoicemailFolder``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemailFolder`: VoicemailFolder
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetUserVoicemailFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailFolder**](VoicemailFolder.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailGreeting

> GetUserVoicemailGreeting(ctx, greeting).Execute()

Get a custom greeting

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
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.GetUserVoicemailGreeting(context.Background(), greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailMessage

> VoicemailMessage GetUserVoicemailMessage(ctx, messageId).Execute()

Get a message

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
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetUserVoicemailMessage(context.Background(), messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemailMessage`: VoicemailMessage
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetUserVoicemailMessage`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailMessage**](VoicemailMessage.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserVoicemailMessageRecording

> GetUserVoicemailMessageRecording(ctx, messageId).Token(token).Download(download).Execute()

Get a message's recording

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
 messageId := "messageId_example" // string | The message's ID
 token := "token_example" // string | The token's ID (optional)
 download := "download_example" // string | Set to 1 to force download by browser (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.GetUserVoicemailMessageRecording(context.Background(), messageId).Token(token).Download(download).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetUserVoicemailMessageRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailMessageRecordingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | The token&#39;s ID |
 **download** | **string** | Set to 1 to force download by browser |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemail

> Voicemail GetVoicemail(ctx, voicemailId).Execute()

Get details of a voicemail

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
 voicemailId := int32(56) // int32 | The voicemail's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetVoicemail(context.Background(), voicemailId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetVoicemail`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetVoicemail`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemailFolder

> VoicemailFolder GetVoicemailFolder(ctx, voicemailId, folderId).Execute()

Get details of a folder

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 folderId := int32(56) // int32 | The folder's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetVoicemailFolder(context.Background(), voicemailId, folderId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetVoicemailFolder``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetVoicemailFolder`: VoicemailFolder
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetVoicemailFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**folderId** | **int32** | The folder&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailFolder**](VoicemailFolder.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemailGreeting

> GetVoicemailGreeting(ctx, voicemailId, greeting).Execute()

Get a custom greeting

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.GetVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemailMessage

> VoicemailMessage GetVoicemailMessage(ctx, voicemailId, messageId).Execute()

Get a message

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetVoicemailMessage(context.Background(), voicemailId, messageId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetVoicemailMessage`: VoicemailMessage
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetVoicemailMessage`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VoicemailMessage**](VoicemailMessage.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemailMessageRecording

> GetVoicemailMessageRecording(ctx, voicemailId, messageId).Token(token).Download(download).Execute()

Get a message's recording

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
 voicemailId := int32(56) // int32 | The voicemail's ID
 messageId := "messageId_example" // string | The message's ID
 token := "token_example" // string | The token's ID (optional)
 download := "download_example" // string | Set to 1 to force download by browser (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.GetVoicemailMessageRecording(context.Background(), voicemailId, messageId).Token(token).Download(download).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetVoicemailMessageRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailMessageRecordingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | The token&#39;s ID |
 **download** | **string** | Set to 1 to force download by browser |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserVoicemails

> Voicemail ListUserVoicemails(ctx).Execute()

Get details of the voicemail of the authenticated user

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.ListUserVoicemails(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.ListUserVoicemails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserVoicemails`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.ListUserVoicemails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserVoicemailsRequest struct via the builder pattern

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserVoicemailGreeting

> UpdateUserVoicemailGreeting(ctx, greeting).Body(body).Execute()

Update a custom greeting

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
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateUserVoicemailGreeting(context.Background(), greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateUserVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserVoicemailMessage

> UpdateUserVoicemailMessage(ctx, messageId).Body(body).Execute()

Update a message

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
 body := *openapiclient.NewVoicemailMessageUpdate(int32(123)) // VoicemailMessageUpdate | Message
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateUserVoicemailMessage(context.Background(), messageId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateUserVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailMessageUpdate**](VoicemailMessageUpdate.md) | Message |

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

## UpdateVoicemailGreeting

> UpdateVoicemailGreeting(ctx, voicemailId, greeting).Body(body).Execute()

Update a custom greeting

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
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 voicemailId := int32(56) // int32 | The voicemail's ID
 greeting := "greeting_example" // string | The greeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateVoicemailGreeting(context.Background(), voicemailId, greeting).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateVoicemailGreeting``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**greeting** | **string** | The greeting |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicemailGreetingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: audio/wav
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateVoicemailMessage

> UpdateVoicemailMessage(ctx, voicemailId, messageId).Body(body).Execute()

Update a message

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
 body := *openapiclient.NewVoicemailMessageUpdate(int32(123)) // VoicemailMessageUpdate | Message
 voicemailId := int32(56) // int32 | The voicemail's ID
 messageId := "messageId_example" // string | The message's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateVoicemailMessage(context.Background(), voicemailId, messageId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateVoicemailMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** | The voicemail&#39;s ID |
**messageId** | **string** | The message&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicemailMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailMessageUpdate**](VoicemailMessageUpdate.md) | Message |

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
