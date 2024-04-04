# \PoliciesAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupPolicy**](PoliciesAPI.md#AddGroupPolicy) | **Put** /groups/{group_uuid}/policies/{policy_uuid} | Associate a group to a policy
[**AddPolicyAccess**](PoliciesAPI.md#AddPolicyAccess) | **Put** /policies/{policy_uuid}/acl/{access} | Associate an access to a policy
[**AddUserPolicy**](PoliciesAPI.md#AddUserPolicy) | **Put** /users/{user_uuid}/policies/{policy_uuid} | Associate a policy to a user
[**CreatePolicies**](PoliciesAPI.md#CreatePolicies) | **Post** /policies | Create a new ACL policy
[**DeleteGroupPolicy**](PoliciesAPI.md#DeleteGroupPolicy) | **Delete** /groups/{group_uuid}/policies/{policy_uuid} | Dissociate a policy from a group
[**DeletePolicy**](PoliciesAPI.md#DeletePolicy) | **Delete** /policies/{policy_uuid} | Delete a policy
[**DeletePolicyAccess**](PoliciesAPI.md#DeletePolicyAccess) | **Delete** /policies/{policy_uuid}/acl/{access} | Dissociate an access from a policy
[**DeleteUserPolicy**](PoliciesAPI.md#DeleteUserPolicy) | **Delete** /users/{user_uuid}/policies/{policy_uuid} | Dissociate a policy from a user
[**EditPolicy**](PoliciesAPI.md#EditPolicy) | **Put** /policies/{policy_uuid} | Modify an ACL policy
[**GetPolicies**](PoliciesAPI.md#GetPolicies) | **Get** /policies | List ACL policies
[**GetPolicy**](PoliciesAPI.md#GetPolicy) | **Get** /policies/{policy_uuid} | Retrieves the details of a policy
[**GetUserPolicies**](PoliciesAPI.md#GetUserPolicies) | **Get** /users/{user_uuid}/policies | Retrieves the list of policies associated to a user

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
 r, err := apiClient.PoliciesAPI.AddGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddGroupPolicy``: %v\n", err)
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

## AddPolicyAccess

> AddPolicyAccess(ctx, policyUuid, access).AccentTenant(accentTenant).Execute()

Associate an access to a policy

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 access := "access_example" // string | The access to add
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PoliciesAPI.AddPolicyAccess(context.Background(), policyUuid, access).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddPolicyAccess``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |
**access** | **string** | The access to add |

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyAccessRequest struct via the builder pattern

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

## AddUserPolicy

> AddUserPolicy(ctx, policyUuid, userUuid).Execute()

Associate a policy to a user

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PoliciesAPI.AddUserPolicy(context.Background(), policyUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddUserPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserPolicyRequest struct via the builder pattern

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

## CreatePolicies

> PolicyResult CreatePolicies(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a new ACL policy

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
 body := *openapiclient.NewPolicy("Name_example") // Policy | The policy creation parameters
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PoliciesAPI.CreatePolicies(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.CreatePolicies``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreatePolicies`: PolicyResult
 fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.CreatePolicies`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoliciesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Policy**](Policy.md) | The policy creation parameters |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**PolicyResult**](PolicyResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
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
 r, err := apiClient.PoliciesAPI.DeleteGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeleteGroupPolicy``: %v\n", err)
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

## DeletePolicy

> DeletePolicy(ctx, policyUuid).AccentTenant(accentTenant).Execute()

Delete a policy

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), policyUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern

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

## DeletePolicyAccess

> DeletePolicyAccess(ctx, policyUuid, access).AccentTenant(accentTenant).Execute()

Dissociate an access from a policy

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 access := "access_example" // string | The access to add
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PoliciesAPI.DeletePolicyAccess(context.Background(), policyUuid, access).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicyAccess``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |
**access** | **string** | The access to add |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyAccessRequest struct via the builder pattern

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

## DeleteUserPolicy

> DeleteUserPolicy(ctx, policyUuid, userUuid).Execute()

Dissociate a policy from a user

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PoliciesAPI.DeleteUserPolicy(context.Background(), policyUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeleteUserPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPolicyRequest struct via the builder pattern

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

## EditPolicy

> PolicyResult EditPolicy(ctx, policyUuid).Body(body).AccentTenant(accentTenant).Execute()

Modify an ACL policy

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
 body := *openapiclient.NewPolicy("Name_example") // Policy | The policy edition parameters
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PoliciesAPI.EditPolicy(context.Background(), policyUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.EditPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `EditPolicy`: PolicyResult
 fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.EditPolicy`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Policy**](Policy.md) | The policy edition parameters |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**PolicyResult**](PolicyResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPolicies

> GetPoliciesResult GetPolicies(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Recurse(recurse).Execute()

List ACL policies

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
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PoliciesAPI.GetPolicies(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Recurse(recurse).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicies``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPolicies`: GetPoliciesResult
 fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]

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

## GetPolicy

> PolicyResult GetPolicy(ctx, policyUuid).AccentTenant(accentTenant).Execute()

Retrieves the details of a policy

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
 policyUuid := "policyUuid_example" // string | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified.
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PoliciesAPI.GetPolicy(context.Background(), policyUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPolicy`: PolicyResult
 fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | The UUID or slug of the policy. The slug is unique within a tenant, hence the tenant must be specified. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**PolicyResult**](PolicyResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserPolicies

> GetPoliciesResult GetUserPolicies(ctx, userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of policies associated to a user

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
 userUuid := "userUuid_example" // string | The UUID of the user
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PoliciesAPI.GetUserPolicies(context.Background(), userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetUserPolicies``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserPolicies`: GetPoliciesResult
 fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetUserPolicies`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPoliciesRequest struct via the builder pattern

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
