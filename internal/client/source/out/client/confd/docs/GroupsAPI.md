# \GroupsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateGroupCallpermission**](GroupsAPI.md#AssociateGroupCallpermission) | **Put** /groups/{group_uuid}/callpermissions/{callpermission_id} | Associate group and call permission
[**AssociateGroupExtension**](GroupsAPI.md#AssociateGroupExtension) | **Put** /groups/{group_uuid}/extensions/{extension_id} | Associate group and extension
[**AssociateGroupSchedule**](GroupsAPI.md#AssociateGroupSchedule) | **Put** /groups/{group_uuid}/schedules/{schedule_id} | Associate group and schedule
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /groups | Create group
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{group_uuid} | Delete group
[**DissociateGroupCallpermission**](GroupsAPI.md#DissociateGroupCallpermission) | **Delete** /groups/{group_uuid}/callpermissions/{callpermission_id} | Dissociate group and call permission
[**DissociateGroupExtension**](GroupsAPI.md#DissociateGroupExtension) | **Delete** /groups/{group_uuid}/extensions/{extension_id} | Dissociate group and extension
[**DissociateGroupSchedule**](GroupsAPI.md#DissociateGroupSchedule) | **Delete** /groups/{group_uuid}/schedules/{schedule_id} | Dissociate group and schedule
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{group_uuid} | Get group
[**GetGroupFallback**](GroupsAPI.md#GetGroupFallback) | **Get** /groups/{group_uuid}/fallbacks | List all fallbacks for group
[**ListGroups**](GroupsAPI.md#ListGroups) | **Get** /groups | List groups
[**UpdateCallPickupInterceptorGroups**](GroupsAPI.md#UpdateCallPickupInterceptorGroups) | **Put** /callpickups/{callpickup_id}/interceptors/groups | Update call pickup and interceptors
[**UpdateCallPickupTargetGroups**](GroupsAPI.md#UpdateCallPickupTargetGroups) | **Put** /callpickups/{callpickup_id}/targets/groups | Update call pickup and targets
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Put** /groups/{group_uuid} | Update group
[**UpdateGroupFallback**](GroupsAPI.md#UpdateGroupFallback) | **Put** /groups/{group_uuid}/fallbacks | Update group&#39;s fallbacks
[**UpdateGroupMemberExtensions**](GroupsAPI.md#UpdateGroupMemberExtensions) | **Put** /groups/{group_uuid}/members/extensions | Update group and extensions
[**UpdateGroupMemberUsers**](GroupsAPI.md#UpdateGroupMemberUsers) | **Put** /groups/{group_uuid}/members/users | Update group and users
[**UpdateUserGroups**](GroupsAPI.md#UpdateUserGroups) | **Put** /users/{user_id}/groups | Update user and groups

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
 r, err := apiClient.GroupsAPI.AssociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AssociateGroupCallpermission``: %v\n", err)
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

## AssociateGroupExtension

> AssociateGroupExtension(ctx, groupUuid, extensionId).Execute()

Associate group and extension

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
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.AssociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AssociateGroupExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateGroupExtensionRequest struct via the builder pattern

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

## AssociateGroupSchedule

> AssociateGroupSchedule(ctx, groupUuid, scheduleId).AccentTenant(accentTenant).Execute()

Associate group and schedule

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
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.AssociateGroupSchedule(context.Background(), groupUuid, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AssociateGroupSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateGroupScheduleRequest struct via the builder pattern

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

## CreateGroup

> Group CreateGroup(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create group

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
 body := *openapiclient.NewGroup() // Group | Group to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateGroup`: Group
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Group**](Group.md) | Group to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Group**](Group.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteGroup

> DeleteGroup(ctx, groupUuid).AccentTenant(accentTenant).Execute()

Delete group

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern

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
 r, err := apiClient.GroupsAPI.DissociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DissociateGroupCallpermission``: %v\n", err)
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

## DissociateGroupExtension

> DissociateGroupExtension(ctx, groupUuid, extensionId).Execute()

Dissociate group and extension

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
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.DissociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DissociateGroupExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateGroupExtensionRequest struct via the builder pattern

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

## DissociateGroupSchedule

> DissociateGroupSchedule(ctx, groupUuid, scheduleId).AccentTenant(accentTenant).Execute()

Dissociate group and schedule

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
 scheduleId := int32(56) // int32 | Schedule's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.DissociateGroupSchedule(context.Background(), groupUuid, scheduleId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DissociateGroupSchedule``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**scheduleId** | **int32** | Schedule&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateGroupScheduleRequest struct via the builder pattern

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

## GetGroup

> Group GetGroup(ctx, groupUuid).AccentTenant(accentTenant).Execute()

Get group

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGroup`: Group
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Group**](Group.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetGroupFallback

> GroupFallbacks GetGroupFallback(ctx, groupUuid).Execute()

List all fallbacks for group

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.GetGroupFallback(context.Background(), groupUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGroupFallback`: GroupFallbacks
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupFallback`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GroupFallbacks**](GroupFallbacks.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListGroups

> GroupItems ListGroups(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List groups

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
 resp, r, err := apiClient.GroupsAPI.ListGroups(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ListGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListGroups`: GroupItems
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern

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

[**GroupItems**](GroupItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallPickupInterceptorGroups

> UpdateCallPickupInterceptorGroups(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and interceptors

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
 body := *openapiclient.NewGroupsID() // GroupsID | Groups to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateCallPickupInterceptorGroups(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateCallPickupInterceptorGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupInterceptorGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupsID**](GroupsID.md) | Groups to associated |

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

## UpdateCallPickupTargetGroups

> UpdateCallPickupTargetGroups(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and targets

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
 body := *openapiclient.NewGroupsID() // GroupsID | Groups to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateCallPickupTargetGroups(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateCallPickupTargetGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupTargetGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupsID**](GroupsID.md) | Groups to associated |

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

## UpdateGroup

> UpdateGroup(ctx, groupUuid).Body(body).AccentTenant(accentTenant).Execute()

Update group

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
 body := *openapiclient.NewGroup() // Group | 
 groupUuid := "groupUuid_example" // string | the group's UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), groupUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Group**](Group.md) |  |

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

## UpdateGroupFallback

> UpdateGroupFallback(ctx, groupUuid).Body(body).Execute()

Update group's fallbacks

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
 body := *openapiclient.NewGroupFallbacks() // GroupFallbacks | Fallbacks for group (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateGroupFallback(context.Background(), groupUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GroupFallbacks**](GroupFallbacks.md) | Fallbacks for group |

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

## UpdateGroupMemberExtensions

> UpdateGroupMemberExtensions(ctx, groupUuid).Body(body).Execute()

Update group and extensions

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
 body := *openapiclient.NewGroupMemberExtensions([]openapiclient.GroupMemberExtension{*openapiclient.NewGroupMemberExtension()}) // GroupMemberExtensions | Extensions to associated
 groupUuid := "groupUuid_example" // string | the group's UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateGroupMemberExtensions(context.Background(), groupUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupMemberExtensions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMemberExtensionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupMemberExtensions**](GroupMemberExtensions.md) | Extensions to associated |

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

## UpdateGroupMemberUsers

> UpdateGroupMemberUsers(ctx, groupUuid).Body(body).Execute()

Update group and users

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
 body := *openapiclient.NewGroupMemberUsers([]openapiclient.GroupMemberUser{*openapiclient.NewGroupMemberUser("Uuid_example")}) // GroupMemberUsers | Users to associated
 groupUuid := "groupUuid_example" // string | the group's UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateGroupMemberUsers(context.Background(), groupUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupMemberUsers**](GroupMemberUsers.md) | Users to associated |

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

## UpdateUserGroups

> UpdateUserGroups(ctx, userId).Body(body).Execute()

Update user and groups

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
 body := *openapiclient.NewUserGroupsID() // UserGroupsID | Users to associated
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.UpdateUserGroups(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateUserGroups``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateUserGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserGroupsID**](UserGroupsID.md) | Users to associated |

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
