# \ConferencesAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateConferenceExtension**](ConferencesAPI.md#AssociateConferenceExtension) | **Put** /conferences/{conference_id}/extensions/{extension_id} | Associate conference and extension
[**CreateConference**](ConferencesAPI.md#CreateConference) | **Post** /conferences | Create conference
[**DeleteConference**](ConferencesAPI.md#DeleteConference) | **Delete** /conferences/{conference_id} | Delete conference
[**DissociateConferenceExtension**](ConferencesAPI.md#DissociateConferenceExtension) | **Delete** /conferences/{conference_id}/extensions/{extension_id} | Dissociate conference and extension
[**GetConference**](ConferencesAPI.md#GetConference) | **Get** /conferences/{conference_id} | Get conference
[**ListAsteriskConfbridgeAccentDefaultBridge**](ConferencesAPI.md#ListAsteriskConfbridgeAccentDefaultBridge) | **Get** /asterisk/confbridge/accent_default_bridge | List ConfBridge accent_default_bridge options
[**ListAsteriskConfbridgeAccentDefaultUser**](ConferencesAPI.md#ListAsteriskConfbridgeAccentDefaultUser) | **Get** /asterisk/confbridge/accent_default_user | List ConfBridge accent_default_user options
[**ListConferences**](ConferencesAPI.md#ListConferences) | **Get** /conferences | List conference
[**UpdateAsteriskConfbridgeAccentDefaultBridge**](ConferencesAPI.md#UpdateAsteriskConfbridgeAccentDefaultBridge) | **Put** /asterisk/confbridge/accent_default_bridge | Update ConfBridge accent_default_bridge option
[**UpdateAsteriskConfbridgeAccentDefaultUser**](ConferencesAPI.md#UpdateAsteriskConfbridgeAccentDefaultUser) | **Put** /asterisk/confbridge/accent_default_user | Update ConfBridge accent_default_user option
[**UpdateConference**](ConferencesAPI.md#UpdateConference) | **Put** /conferences/{conference_id} | Update conference

## AssociateConferenceExtension

> AssociateConferenceExtension(ctx, conferenceId, extensionId).Execute()

Associate conference and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.AssociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.AssociateConferenceExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateConferenceExtensionRequest struct via the builder pattern

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

## CreateConference

> Conference CreateConference(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewConference() // Conference | Conference to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.CreateConference(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.CreateConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateConference`: Conference
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.CreateConference`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConferenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Conference**](Conference.md) | Conference to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Conference**](Conference.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteConference

> DeleteConference(ctx, conferenceId).AccentTenant(accentTenant).Execute()

Delete conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.DeleteConference(context.Background(), conferenceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.DeleteConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConferenceRequest struct via the builder pattern

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

## DissociateConferenceExtension

> DissociateConferenceExtension(ctx, conferenceId, extensionId).Execute()

Dissociate conference and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.DissociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.DissociateConferenceExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateConferenceExtensionRequest struct via the builder pattern

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

## GetConference

> Conference GetConference(ctx, conferenceId).AccentTenant(accentTenant).Execute()

Get conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.GetConference(context.Background(), conferenceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.GetConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetConference`: Conference
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.GetConference`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetConferenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Conference**](Conference.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskConfbridgeAccentDefaultBridge

> ConfBridgeConfiguration ListAsteriskConfbridgeAccentDefaultBridge(ctx).Execute()

List ConfBridge accent_default_bridge options

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.ListAsteriskConfbridgeAccentDefaultBridge(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListAsteriskConfbridgeAccentDefaultBridge``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskConfbridgeAccentDefaultBridge`: ConfBridgeConfiguration
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListAsteriskConfbridgeAccentDefaultBridge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskConfbridgeAccentDefaultBridgeRequest struct via the builder pattern

### Return type

[**ConfBridgeConfiguration**](ConfBridgeConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskConfbridgeAccentDefaultUser

> ConfBridgeConfiguration ListAsteriskConfbridgeAccentDefaultUser(ctx).Execute()

List ConfBridge accent_default_user options

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.ListAsteriskConfbridgeAccentDefaultUser(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListAsteriskConfbridgeAccentDefaultUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskConfbridgeAccentDefaultUser`: ConfBridgeConfiguration
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListAsteriskConfbridgeAccentDefaultUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskConfbridgeAccentDefaultUserRequest struct via the builder pattern

### Return type

[**ConfBridgeConfiguration**](ConfBridgeConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListConferences

> ConferenceItems ListConferences(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
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
 resp, r, err := apiClient.ConferencesAPI.ListConferences(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListConferences``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListConferences`: ConferenceItems
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListConferences`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListConferencesRequest struct via the builder pattern

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

[**ConferenceItems**](ConferenceItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskConfbridgeAccentDefaultBridge

> UpdateAsteriskConfbridgeAccentDefaultBridge(ctx).Body(body).Execute()

Update ConfBridge accent_default_bridge option

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewConfBridgeConfiguration() // ConfBridgeConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.UpdateAsteriskConfbridgeAccentDefaultBridge(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateAsteriskConfbridgeAccentDefaultBridge``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskConfbridgeAccentDefaultBridgeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfBridgeConfiguration**](ConfBridgeConfiguration.md) |  |

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

## UpdateAsteriskConfbridgeAccentDefaultUser

> UpdateAsteriskConfbridgeAccentDefaultUser(ctx).Body(body).Execute()

Update ConfBridge accent_default_user option

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewConfBridgeConfiguration() // ConfBridgeConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.UpdateAsteriskConfbridgeAccentDefaultUser(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateAsteriskConfbridgeAccentDefaultUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskConfbridgeAccentDefaultUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfBridgeConfiguration**](ConfBridgeConfiguration.md) |  |

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

## UpdateConference

> UpdateConference(ctx, conferenceId).Body(body).AccentTenant(accentTenant).Execute()

Update conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewConference() // Conference | 
 conferenceId := int32(56) // int32 | Conference's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.UpdateConference(context.Background(), conferenceId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Conference**](Conference.md) |  |

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
