# \CallpermissionsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateGroupCallpermission**](CallpermissionsAPI.md#AssociateGroupCallpermission) | **Put** /groups/{group_uuid}/callpermissions/{callpermission_id} | Associate group and call permission
[**AssociateOutcallCallpermission**](CallpermissionsAPI.md#AssociateOutcallCallpermission) | **Put** /outcalls/{outcall_id}/callpermissions/{callpermission_id} | Associate outcall and call permission
[**AssociateUserCallpermission**](CallpermissionsAPI.md#AssociateUserCallpermission) | **Put** /users/{user_id}/callpermissions/{callpermission_id} | Associate user and call permission
[**CreateCallpermission**](CallpermissionsAPI.md#CreateCallpermission) | **Post** /callpermissions | Create call permission
[**DeleteCallpermission**](CallpermissionsAPI.md#DeleteCallpermission) | **Delete** /callpermissions/{callpermission_id} | Delete call permission
[**DissociateGroupCallpermission**](CallpermissionsAPI.md#DissociateGroupCallpermission) | **Delete** /groups/{group_uuid}/callpermissions/{callpermission_id} | Dissociate group and call permission
[**DissociateOutcallCallpermission**](CallpermissionsAPI.md#DissociateOutcallCallpermission) | **Delete** /outcalls/{outcall_id}/callpermissions/{callpermission_id} | Dissociate outcall and call permission
[**DissociateUserCallpermission**](CallpermissionsAPI.md#DissociateUserCallpermission) | **Delete** /users/{user_id}/callpermissions/{callpermission_id} | Dissociate user and call permission
[**GetCallpermission**](CallpermissionsAPI.md#GetCallpermission) | **Get** /callpermissions/{callpermission_id} | Get call permission
[**ListCallpermissions**](CallpermissionsAPI.md#ListCallpermissions) | **Get** /callpermissions | List call permissions
[**UpdateCallpermission**](CallpermissionsAPI.md#UpdateCallpermission) | **Put** /callpermissions/{callpermission_id} | Update call permission

## AssociateGroupCallpermission

> AssociateGroupCallpermission(ctx, groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()

Associate group and call permission

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
 groupUuid := "groupUuid_example" // string | the group's UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.AssociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.AssociateGroupCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateGroupCallpermissionRequest struct via the builder pattern

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

## AssociateOutcallCallpermission

> AssociateOutcallCallpermission(ctx, outcallId, callpermissionId).AccentTenant(accentTenant).Execute()

Associate outcall and call permission

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.AssociateOutcallCallpermission(context.Background(), outcallId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.AssociateOutcallCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateOutcallCallpermissionRequest struct via the builder pattern

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

## AssociateUserCallpermission

> AssociateUserCallpermission(ctx, userId, callpermissionId).AccentTenant(accentTenant).Execute()

Associate user and call permission

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
 userId := "userId_example" // string | the user's ID or UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.AssociateUserCallpermission(context.Background(), userId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.AssociateUserCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserCallpermissionRequest struct via the builder pattern

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

## CreateCallpermission

> CallPermission CreateCallpermission(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create call permission

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
 body := *openapiclient.NewCallPermission("Name_example") // CallPermission | Call Permission to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallpermissionsAPI.CreateCallpermission(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.CreateCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCallpermission`: CallPermission
 fmt.Fprintf(os.Stdout, "Response from `CallpermissionsAPI.CreateCallpermission`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallpermissionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallPermission**](CallPermission.md) | Call Permission to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallPermission**](CallPermission.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCallpermission

> DeleteCallpermission(ctx, callpermissionId).AccentTenant(accentTenant).Execute()

Delete call permission

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
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.DeleteCallpermission(context.Background(), callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.DeleteCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallpermissionRequest struct via the builder pattern

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

## DissociateGroupCallpermission

> DissociateGroupCallpermission(ctx, groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()

Dissociate group and call permission

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
 groupUuid := "groupUuid_example" // string | the group's UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.DissociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.DissociateGroupCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateGroupCallpermissionRequest struct via the builder pattern

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

## DissociateOutcallCallpermission

> DissociateOutcallCallpermission(ctx, outcallId, callpermissionId).AccentTenant(accentTenant).Execute()

Dissociate outcall and call permission

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.DissociateOutcallCallpermission(context.Background(), outcallId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.DissociateOutcallCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateOutcallCallpermissionRequest struct via the builder pattern

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

## DissociateUserCallpermission

> DissociateUserCallpermission(ctx, userId, callpermissionId).AccentTenant(accentTenant).Execute()

Dissociate user and call permission

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
 userId := "userId_example" // string | the user's ID or UUID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.DissociateUserCallpermission(context.Background(), userId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.DissociateUserCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserCallpermissionRequest struct via the builder pattern

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

## GetCallpermission

> CallPermission GetCallpermission(ctx, callpermissionId).AccentTenant(accentTenant).Execute()

Get call permission

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
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallpermissionsAPI.GetCallpermission(context.Background(), callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.GetCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCallpermission`: CallPermission
 fmt.Fprintf(os.Stdout, "Response from `CallpermissionsAPI.GetCallpermission`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallpermissionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallPermission**](CallPermission.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCallpermissions

> CallPermissionItems ListCallpermissions(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List call permissions

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
 resp, r, err := apiClient.CallpermissionsAPI.ListCallpermissions(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.ListCallpermissions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListCallpermissions`: CallPermissionItems
 fmt.Fprintf(os.Stdout, "Response from `CallpermissionsAPI.ListCallpermissions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCallpermissionsRequest struct via the builder pattern

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

[**CallPermissionItems**](CallPermissionItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallpermission

> UpdateCallpermission(ctx, callpermissionId).Body(body).AccentTenant(accentTenant).Execute()

Update call permission

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
 body := *openapiclient.NewCallPermission("Name_example") // CallPermission | 
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpermissionsAPI.UpdateCallpermission(context.Background(), callpermissionId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpermissionsAPI.UpdateCallpermission``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpermissionId** | **int32** | Call Permission&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallpermissionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallPermission**](CallPermission.md) |  |

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
