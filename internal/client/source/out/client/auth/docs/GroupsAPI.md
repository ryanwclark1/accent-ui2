# \GroupsAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupPolicy**](GroupsAPI.md#AddGroupPolicy) | **Put** /groups/{group_uuid}/policies/{policy_uuid} | Associate a group to a policy
[**AddGroupUser**](GroupsAPI.md#AddGroupUser) | **Put** /groups/{group_uuid}/users/{user_uuid} | Associate a group to a user
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{group_uuid} | Delete a group
[**DeleteGroupPolicy**](GroupsAPI.md#DeleteGroupPolicy) | **Delete** /groups/{group_uuid}/policies/{policy_uuid} | Dissociate a policy from a group
[**EditGroup**](GroupsAPI.md#EditGroup) | **Put** /groups/{group_uuid} | Modify a group
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{group_uuid} | Retrieves the details of a group
[**GetGroupPolicies**](GroupsAPI.md#GetGroupPolicies) | **Get** /groups/{group_uuid}/policies | Retrieves the list of policies associated to a group
[**GetGroupUsers**](GroupsAPI.md#GetGroupUsers) | **Get** /groups/{group_uuid}/users | Retrieves the list of users associated to a group
[**RemoveGroupUser**](GroupsAPI.md#RemoveGroupUser) | **Delete** /groups/{group_uuid}/users/{user_uuid} | Dissociate a user from a group

## AddGroupPolicy

> AddGroupPolicy(ctx, groupUuid, policyUuid).Execute()

Associate a group to a policy

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.AddGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupPolicyRequest struct via the builder pattern

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

## AddGroupUser

> AddGroupUser(ctx, groupUuid, userUuid).Execute()

Associate a group to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.AddGroupUser(context.Background(), groupUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupUserRequest struct via the builder pattern

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

## DeleteGroup

> DeleteGroup(ctx, groupUuid).AccentTenant(accentTenant).Execute()

Delete a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
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
**groupUuid** | **string** | The UUID of the group |

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

## DeleteGroupPolicy

> DeleteGroupPolicy(ctx, groupUuid, policyUuid).Execute()

Dissociate a policy from a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.DeleteGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupPolicyRequest struct via the builder pattern

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

## EditGroup

> GroupResult EditGroup(ctx, groupUuid).Body(body).AccentTenant(accentTenant).Execute()

Modify a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 body := *openapiclient.NewGroupPut("Name_example") // GroupPut | The group parameters
 groupUuid := "groupUuid_example" // string | The UUID of the group
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.EditGroup(context.Background(), groupUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.EditGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `EditGroup`: GroupResult
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.EditGroup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |

### Other Parameters

Other parameters are passed through a pointer to a apiEditGroupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupPut**](GroupPut.md) | The group parameters |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**GroupResult**](GroupResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetGroup

> GroupResult GetGroup(ctx, groupUuid).AccentTenant(accentTenant).Execute()

Retrieves the details of a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGroup`: GroupResult
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**GroupResult**](GroupResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetGroupPolicies

> GetPoliciesResult GetGroupPolicies(ctx, groupUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of policies associated to a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.GetGroupPolicies(context.Background(), groupUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupPolicies``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGroupPolicies`: GetPoliciesResult
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupPolicies`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPoliciesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**GetPoliciesResult**](GetPoliciesResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetGroupUsers

> UserList GetGroupUsers(ctx, groupUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of users associated to a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.GroupsAPI.GetGroupUsers(context.Background(), groupUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGroupUsers`: UserList
 fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupUsers`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**UserList**](UserList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RemoveGroupUser

> RemoveGroupUser(ctx, groupUuid, userUuid).Execute()

Dissociate a user from a group

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 groupUuid := "groupUuid_example" // string | The UUID of the group
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.GroupsAPI.RemoveGroupUser(context.Background(), groupUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveGroupUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | The UUID of the group |
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupUserRequest struct via the builder pattern

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
