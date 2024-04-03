# \OutcallsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateOutcallCallpermission**](OutcallsAPI.md#AssociateOutcallCallpermission) | **Put** /outcalls/{outcall_id}/callpermissions/{callpermission_id} | Associate outcall and call permission
[**AssociateOutcallExtension**](OutcallsAPI.md#AssociateOutcallExtension) | **Put** /outcalls/{outcall_id}/extensions/{extension_id} | Associate outcall and extension
[**AssociateOutcallSchedule**](OutcallsAPI.md#AssociateOutcallSchedule) | **Put** /outcalls/{outcall_id}/schedules/{schedule_id} | Associate outcall and schedule
[**AssociateOutcallTrunks**](OutcallsAPI.md#AssociateOutcallTrunks) | **Put** /outcalls/{outcall_id}/trunks | Associate outcall and trunks
[**CreateOutcall**](OutcallsAPI.md#CreateOutcall) | **Post** /outcalls | Create outgoing call
[**DeleteOutcall**](OutcallsAPI.md#DeleteOutcall) | **Delete** /outcalls/{outcall_id} | Delete outgoing call
[**DissociateOutcallCallpermission**](OutcallsAPI.md#DissociateOutcallCallpermission) | **Delete** /outcalls/{outcall_id}/callpermissions/{callpermission_id} | Dissociate outcall and call permission
[**DissociateOutcallExtension**](OutcallsAPI.md#DissociateOutcallExtension) | **Delete** /outcalls/{outcall_id}/extensions/{extension_id} | Dissociate outcall and extension
[**DissociateOutcallSchedule**](OutcallsAPI.md#DissociateOutcallSchedule) | **Delete** /outcalls/{outcall_id}/schedules/{schedule_id} | Dissociate outcall and schedule
[**GetOutcall**](OutcallsAPI.md#GetOutcall) | **Get** /outcalls/{outcall_id} | Get outgoing call
[**ListOutcalls**](OutcallsAPI.md#ListOutcalls) | **Get** /outcalls | List outgoing calls
[**UpdateOutcall**](OutcallsAPI.md#UpdateOutcall) | **Put** /outcalls/{outcall_id} | Update outgoing call

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 outcallId := int32(56) // int32 | Outgoing call's ID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.AssociateOutcallCallpermission(context.Background(), outcallId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.AssociateOutcallCallpermission``: %v\n", err)
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

## AssociateOutcallExtension

> AssociateOutcallExtension(ctx, outcallId, extensionId).Body(body).Execute()

Associate outcall and extension

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 extensionId := int32(56) // int32 | 
 body := *openapiclient.NewOutcallExtension() // OutcallExtension | Outgoing Extension (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.AssociateOutcallExtension(context.Background(), outcallId, extensionId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.AssociateOutcallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateOutcallExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OutcallExtension**](OutcallExtension.md) | Outgoing Extension |

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

## AssociateOutcallSchedule

> AssociateOutcallSchedule(ctx, outcallId, scheduleId).AccentTenant(accentTenant).Execute()

Associate outcall and schedule

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.AssociateOutcallSchedule(context.Background(), outcallId, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.AssociateOutcallSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateOutcallScheduleRequest struct via the builder pattern

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

## AssociateOutcallTrunks

> AssociateOutcallTrunks(ctx, outcallId).Body(body).Execute()

Associate outcall and trunks

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
 body := *openapiclient.NewTrunksId() // TrunksId | Trunks to associated
 outcallId := int32(56) // int32 | Outgoing call's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.AssociateOutcallTrunks(context.Background(), outcallId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.AssociateOutcallTrunks``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateOutcallTrunksRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrunksId**](TrunksId.md) | Trunks to associated |

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

## CreateOutcall

> Outcall CreateOutcall(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create outgoing call

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
 body := *openapiclient.NewOutcall() // Outcall | Outgoing call to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.OutcallsAPI.CreateOutcall(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.CreateOutcall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateOutcall`: Outcall
 fmt.Fprintf(os.Stdout, "Response from `OutcallsAPI.CreateOutcall`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutcallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Outcall**](Outcall.md) | Outgoing call to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Outcall**](Outcall.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteOutcall

> DeleteOutcall(ctx, outcallId).AccentTenant(accentTenant).Execute()

Delete outgoing call

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.DeleteOutcall(context.Background(), outcallId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.DeleteOutcall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutcallRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 outcallId := int32(56) // int32 | Outgoing call's ID
 callpermissionId := int32(56) // int32 | Call Permission's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.DissociateOutcallCallpermission(context.Background(), outcallId, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.DissociateOutcallCallpermission``: %v\n", err)
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

## DissociateOutcallExtension

> DissociateOutcallExtension(ctx, outcallId, extensionId).Execute()

Dissociate outcall and extension

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.DissociateOutcallExtension(context.Background(), outcallId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.DissociateOutcallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateOutcallExtensionRequest struct via the builder pattern

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

## DissociateOutcallSchedule

> DissociateOutcallSchedule(ctx, outcallId, scheduleId).AccentTenant(accentTenant).Execute()

Dissociate outcall and schedule

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.DissociateOutcallSchedule(context.Background(), outcallId, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.DissociateOutcallSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateOutcallScheduleRequest struct via the builder pattern

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

## GetOutcall

> Outcall GetOutcall(ctx, outcallId).AccentTenant(accentTenant).Execute()

Get outgoing call

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
 outcallId := int32(56) // int32 | Outgoing call's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.OutcallsAPI.GetOutcall(context.Background(), outcallId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.GetOutcall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetOutcall`: Outcall
 fmt.Fprintf(os.Stdout, "Response from `OutcallsAPI.GetOutcall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutcallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Outcall**](Outcall.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOutcalls

> OutcallItems ListOutcalls(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List outgoing calls

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
 resp, r, err := apiClient.OutcallsAPI.ListOutcalls(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.ListOutcalls``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListOutcalls`: OutcallItems
 fmt.Fprintf(os.Stdout, "Response from `OutcallsAPI.ListOutcalls`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListOutcallsRequest struct via the builder pattern

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

[**OutcallItems**](OutcallItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateOutcall

> UpdateOutcall(ctx, outcallId).Body(body).AccentTenant(accentTenant).Execute()

Update outgoing call

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
 body := *openapiclient.NewOutcall() // Outcall | 
 outcallId := int32(56) // int32 | Outgoing call's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.OutcallsAPI.UpdateOutcall(context.Background(), outcallId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `OutcallsAPI.UpdateOutcall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutcallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Outcall**](Outcall.md) |  |

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
