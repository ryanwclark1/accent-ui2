# \VoicemailsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateUserVoicemail**](VoicemailsAPI.md#AssociateUserVoicemail) | **Put** /users/{user_id}/voicemails/{voicemail_id} | Associate user and voicemail
[**CreateUserVoicemail**](VoicemailsAPI.md#CreateUserVoicemail) | **Post** /users/{user_id}/voicemails | Create user voicemail
[**CreateVoicemail**](VoicemailsAPI.md#CreateVoicemail) | **Post** /voicemails | Create voicemail
[**DeleteVoicemail**](VoicemailsAPI.md#DeleteVoicemail) | **Delete** /voicemails/{voicemail_id} | Delete voicemail
[**DissociateUserVoicemail**](VoicemailsAPI.md#DissociateUserVoicemail) | **Delete** /users/{user_id}/voicemails | Dissociate user and voicemail
[**GetUserVoicemail**](VoicemailsAPI.md#GetUserVoicemail) | **Get** /users/{user_id}/voicemails | Get user voicemails
[**GetVoicemail**](VoicemailsAPI.md#GetVoicemail) | **Get** /voicemails/{voicemail_id} | Get voicemail
[**ListAsteriskVoicemailZonemessages**](VoicemailsAPI.md#ListAsteriskVoicemailZonemessages) | **Get** /asterisk/voicemail/zonemessages | List Voicemail zonemessages options
[**ListVoicemails**](VoicemailsAPI.md#ListVoicemails) | **Get** /voicemails | List voicemails
[**UpdateAsteriskVoicemailZonemessages**](VoicemailsAPI.md#UpdateAsteriskVoicemailZonemessages) | **Put** /asterisk/voicemail/zonemessages | Update Voicemail zonemessages option
[**UpdateVoicemail**](VoicemailsAPI.md#UpdateVoicemail) | **Put** /voicemails/{voicemail_id} | Update voicemail

## AssociateUserVoicemail

> AssociateUserVoicemail(ctx, userId, voicemailId).Execute()

Associate user and voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 voicemailId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.AssociateUserVoicemail(context.Background(), userId, voicemailId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.AssociateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**voicemailId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserVoicemailRequest struct via the builder pattern

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

## CreateUserVoicemail

> Voicemail CreateUserVoicemail(ctx, userId).Body(body).AccentTenant(accentTenant).Execute()

Create user voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewVoicemail("Context_example", "Number_example") // Voicemail | Voicemail to create
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.CreateUserVoicemail(context.Background(), userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CreateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserVoicemail`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.CreateUserVoicemail`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Voicemail**](Voicemail.md) | Voicemail to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateVoicemail

> Voicemail CreateVoicemail(ctx).Body(body).Execute()

Create voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewVoicemail("Context_example", "Number_example") // Voicemail | Voicemail to create

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.CreateVoicemail(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.CreateVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateVoicemail`: Voicemail
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.CreateVoicemail`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Voicemail**](Voicemail.md) | Voicemail to create |

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteVoicemail

> DeleteVoicemail(ctx, voicemailId).AccentTenant(accentTenant).Execute()

Delete voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 voicemailId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DeleteVoicemail(context.Background(), voicemailId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DeleteVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoicemailRequest struct via the builder pattern

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

## DissociateUserVoicemail

> DissociateUserVoicemail(ctx, userId).Execute()

Dissociate user and voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.DissociateUserVoicemail(context.Background(), userId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.DissociateUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserVoicemailRequest struct via the builder pattern

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

## GetUserVoicemail

> VoicemailItems GetUserVoicemail(ctx, userId).AccentTenant(accentTenant).Execute()

Get user voicemails

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetUserVoicemail(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.GetUserVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserVoicemail`: VoicemailItems
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.GetUserVoicemail`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**VoicemailItems**](VoicemailItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetVoicemail

> Voicemail GetVoicemail(ctx, voicemailId).AccentTenant(accentTenant).Execute()

Get voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 voicemailId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.GetVoicemail(context.Background(), voicemailId).AccentTenant(accentTenant).Execute()
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
**voicemailId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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

## ListAsteriskVoicemailZonemessages

> VoicemailZoneMessages ListAsteriskVoicemailZonemessages(ctx).Execute()

List Voicemail zonemessages options

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.ListAsteriskVoicemailZonemessages(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.ListAsteriskVoicemailZonemessages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskVoicemailZonemessages`: VoicemailZoneMessages
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.ListAsteriskVoicemailZonemessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskVoicemailZonemessagesRequest struct via the builder pattern

### Return type

[**VoicemailZoneMessages**](VoicemailZoneMessages.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListVoicemails

> VoicemailItems ListVoicemails(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List voicemails

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.VoicemailsAPI.ListVoicemails(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.ListVoicemails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListVoicemails`: VoicemailItems
 fmt.Fprintf(os.Stdout, "Response from `VoicemailsAPI.ListVoicemails`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListVoicemailsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**VoicemailItems**](VoicemailItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskVoicemailZonemessages

> UpdateAsteriskVoicemailZonemessages(ctx).Body(body).Execute()

Update Voicemail zonemessages option

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewVoicemailZoneMessages() // VoicemailZoneMessages | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateAsteriskVoicemailZonemessages(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateAsteriskVoicemailZonemessages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskVoicemailZonemessagesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailZoneMessages**](VoicemailZoneMessages.md) |  |

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

## UpdateVoicemail

> UpdateVoicemail(ctx, voicemailId).Body(body).AccentTenant(accentTenant).Execute()

Update voicemail

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewVoicemail("Context_example", "Number_example") // Voicemail | 
 voicemailId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.VoicemailsAPI.UpdateVoicemail(context.Background(), voicemailId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsAPI.UpdateVoicemail``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicemailRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Voicemail**](Voicemail.md) |  |

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
