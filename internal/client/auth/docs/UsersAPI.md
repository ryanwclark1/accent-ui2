# \UsersAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupUser**](UsersAPI.md#AddGroupUser) | **Put** /groups/{group_uuid}/users/{user_uuid} | Associate a group to a user
[**AddUserPolicy**](UsersAPI.md#AddUserPolicy) | **Put** /users/{user_uuid}/policies/{policy_uuid} | Associate a policy to a user
[**ChangeUserPassword**](UsersAPI.md#ChangeUserPassword) | **Put** /users/{user_uuid}/password | Change the user&#39;s password
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /users | Create a user
[**DeleteRefreshTokens**](UsersAPI.md#DeleteRefreshTokens) | **Delete** /users/{user_uuid_or_me}/tokens/{client_id} | Delete a user&#39;s refresh token
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{user_uuid} | Delete a user
[**DeleteUserPolicy**](UsersAPI.md#DeleteUserPolicy) | **Delete** /users/{user_uuid}/policies/{policy_uuid} | Dissociate a policy from a user
[**GetNewEmailConfirmation**](UsersAPI.md#GetNewEmailConfirmation) | **Get** /users/{user_uuid}/emails/{email_uuid}/confirm | Ask a new confirmation email
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{user_uuid} | Retrieves the details of a user
[**GetUserExternalAuth**](UsersAPI.md#GetUserExternalAuth) | **Get** /users/{user_uuid}/external | Retrieves the list of the users external auth data
[**GetUserGroups**](UsersAPI.md#GetUserGroups) | **Get** /users/{user_uuid}/groups | Retrieves the list of groups associated to a user
[**GetUserPolicies**](UsersAPI.md#GetUserPolicies) | **Get** /users/{user_uuid}/policies | Retrieves the list of policies associated to a user
[**GetUserSessions**](UsersAPI.md#GetUserSessions) | **Get** /users/{user_uuid}/sessions | Retrieves the list of sessions associated to a user
[**GetUserTokens**](UsersAPI.md#GetUserTokens) | **Get** /users/{user_uuid_or_me}/tokens | Retrieve a user&#39;s refresh token list
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /users | Retrieves the list of users
[**RegisterUser**](UsersAPI.md#RegisterUser) | **Post** /users/register | Create a user
[**RemoveGroupUser**](UsersAPI.md#RemoveGroupUser) | **Delete** /groups/{group_uuid}/users/{user_uuid} | Dissociate a user from a group
[**ResetPassword**](UsersAPI.md#ResetPassword) | **Get** /users/password/reset | Reset the user password
[**ResetPasswordChange**](UsersAPI.md#ResetPasswordChange) | **Post** /users/password/reset | Set the user password
[**UpdateAllUserEmails**](UsersAPI.md#UpdateAllUserEmails) | **Put** /admin/users/{user_uuid}/emails | Update email addresses
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{user_uuid} | Update an existing user
[**UpdateUserEmails**](UsersAPI.md#UpdateUserEmails) | **Put** /users/{user_uuid}/emails | Update email addresses
[**UserDeleteSession**](UsersAPI.md#UserDeleteSession) | **Delete** /users/{user_uuid}/sessions/{session_uuid} | Delete a session

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
 r, err := apiClient.UsersAPI.AddGroupUser(context.Background(), groupUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddGroupUser``: %v\n", err)
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
 r, err := apiClient.UsersAPI.AddUserPolicy(context.Background(), policyUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUserPolicy``: %v\n", err)
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

## ChangeUserPassword

> ChangeUserPassword(ctx, userUuid).Body(body).Execute()

Change the user's password

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
 body := *openapiclient.NewPasswordChange("NewPassword_example", "OldPassword_example") // PasswordChange | The user creation parameters
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.ChangeUserPassword(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ChangeUserPassword``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PasswordChange**](PasswordChange.md) | The user creation parameters |

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

## CreateUser

> UserPostResponse CreateUser(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a user

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
 body := *openapiclient.NewUserCreate() // UserCreate | The user creation parameters (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.CreateUser(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUser`: UserPostResponse
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserCreate**](UserCreate.md) | The user creation parameters |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserPostResponse**](UserPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteRefreshTokens

> DeleteRefreshTokens(ctx, userUuidOrMe, clientId).Execute()

Delete a user's refresh token

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
 userUuidOrMe := "userUuidOrMe_example" // string | The UUID of the user or `me` to refer to the user doing the query
 clientId := "clientId_example" // string | The client_id of the refresh token to revoke

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteRefreshTokens(context.Background(), userUuidOrMe, clientId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteRefreshTokens``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuidOrMe** | **string** | The UUID of the user or &#x60;me&#x60; to refer to the user doing the query |
**clientId** | **string** | The client_id of the refresh token to revoke |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRefreshTokensRequest struct via the builder pattern

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

## DeleteUser

> DeleteUser(ctx, userUuid).Execute()

Delete a user

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.DeleteUser(context.Background(), userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern

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
 r, err := apiClient.UsersAPI.DeleteUserPolicy(context.Background(), policyUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserPolicy``: %v\n", err)
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

## GetNewEmailConfirmation

> GetNewEmailConfirmation(ctx, userUuid, emailUuid).Execute()

Ask a new confirmation email

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
 emailUuid := "emailUuid_example" // string | The UUID of the email

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.GetNewEmailConfirmation(context.Background(), userUuid, emailUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetNewEmailConfirmation``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |
**emailUuid** | **string** | The UUID of the email |

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewEmailConfirmationRequest struct via the builder pattern

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

## GetUser

> UserResult GetUser(ctx, userUuid).Execute()

Retrieves the details of a user

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUser`: UserResult
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UserResult**](UserResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserExternalAuth

> ExternalAuthList GetUserExternalAuth(ctx, userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of the users external auth data

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
 resp, r, err := apiClient.UsersAPI.GetUserExternalAuth(context.Background(), userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserExternalAuth``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserExternalAuth`: ExternalAuthList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserExternalAuth`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserExternalAuthRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**ExternalAuthList**](ExternalAuthList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserGroups

> GetGroupsResult GetUserGroups(ctx, userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of groups associated to a user

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
 resp, r, err := apiClient.UsersAPI.GetUserGroups(context.Background(), userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserGroups`: GetGroupsResult
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**GetGroupsResult**](GetGroupsResult.md)

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
 resp, r, err := apiClient.UsersAPI.GetUserPolicies(context.Background(), userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPolicies``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserPolicies`: GetPoliciesResult
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPolicies`: %v\n", resp)
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

## GetUserSessions

> GetSessionsResult GetUserSessions(ctx, userUuid).AccentTenant(accentTenant).Limit(limit).Offset(offset).Execute()

Retrieves the list of sessions associated to a user

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserSessions(context.Background(), userUuid).AccentTenant(accentTenant).Limit(limit).Offset(offset).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSessions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserSessions`: GetSessionsResult
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSessions`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]

### Return type

[**GetSessionsResult**](GetSessionsResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserTokens

> RefreshTokenList GetUserTokens(ctx, userUuidOrMe).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve a user's refresh token list

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
 userUuidOrMe := "userUuidOrMe_example" // string | The UUID of the user or `me` to refer to the user doing the query
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserTokens(context.Background(), userUuidOrMe).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserTokens``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserTokens`: RefreshTokenList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserTokens`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuidOrMe** | **string** | The UUID of the user or &#x60;me&#x60; to refer to the user doing the query |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokensRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**RefreshTokenList**](RefreshTokenList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUsers

> UserList GetUsers(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Recurse(recurse).HasPolicySlug(hasPolicySlug).HasPolicyUuid(hasPolicyUuid).PolicySlug(policySlug).PolicyUuid(policyUuid).Execute()

Retrieves the list of users

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
 hasPolicySlug := "hasPolicySlug_example" // string | The slug of the policy that the user must have. This includes indirect associations (user in group has policy). (optional)
 hasPolicyUuid := "hasPolicyUuid_example" // string | The UUID of the policy that the user must have. This includes indirect associations (user in group has policy). (optional)
 policySlug := "policySlug_example" // string | The slug of the policy that the user must have. This does not include indirect associations (user in group has policy). (optional)
 policyUuid := "policyUuid_example" // string | The UUID of the policy that the user must have. This does not include indirect associations (user in group has policy). (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUsers(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Recurse(recurse).HasPolicySlug(hasPolicySlug).HasPolicyUuid(hasPolicyUuid).PolicySlug(policySlug).PolicyUuid(policyUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUsers`: UserList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **hasPolicySlug** | **string** | The slug of the policy that the user must have. This includes indirect associations (user in group has policy). |
 **hasPolicyUuid** | **string** | The UUID of the policy that the user must have. This includes indirect associations (user in group has policy). |
 **policySlug** | **string** | The slug of the policy that the user must have. This does not include indirect associations (user in group has policy). |
 **policyUuid** | **string** | The UUID of the policy that the user must have. This does not include indirect associations (user in group has policy). |

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

## RegisterUser

> UserPostResponse RegisterUser(ctx).Body(body).Execute()

Create a user

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
 body := *openapiclient.NewUserRegister("EmailAddress_example", "Password_example", "Username_example") // UserRegister | The user creation parameters

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.RegisterUser(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RegisterUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `RegisterUser`: UserPostResponse
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RegisterUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserRegister**](UserRegister.md) | The user creation parameters |

### Return type

[**UserPostResponse**](UserPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
 r, err := apiClient.UsersAPI.RemoveGroupUser(context.Background(), groupUuid, userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveGroupUser``: %v\n", err)
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

## ResetPassword

> ResetPassword(ctx).Username(username).Email(email).Login(login).Execute()

Reset the user password

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
 username := "username_example" // string | The user's username (optional)
 email := "email_example" // string | The user's email address (optional)
 login := "login_example" // string | The user's login (username or email) (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.ResetPassword(context.Background()).Username(username).Email(email).Login(login).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResetPassword``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The user&#39;s username |
 **email** | **string** | The user&#39;s email address |
 **login** | **string** | The user&#39;s login (username or email) |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ResetPasswordChange

> ResetPasswordChange(ctx).Body(body).UserUuid(userUuid).Execute()

Set the user password

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
 body := *openapiclient.NewPostPasswordReset("Password_example") // PostPasswordReset | The password change parameters
 userUuid := "userUuid_example" // string | The user's UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.ResetPasswordChange(context.Background()).Body(body).UserUuid(userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResetPasswordChange``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordChangeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostPasswordReset**](PostPasswordReset.md) | The password change parameters |
 **userUuid** | **string** | The user&#39;s UUID |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAllUserEmails

> UpdateAllUserEmails(ctx, userUuid).Body(body).Execute()

Update email addresses

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
 body := *openapiclient.NewAdminUserEmailList() // AdminUserEmailList | EmailAddressList
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateAllUserEmails(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateAllUserEmails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllUserEmailsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdminUserEmailList**](AdminUserEmailList.md) | EmailAddressList |

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

## UpdateUser

> UserPostResponse UpdateUser(ctx, userUuid).Body(body).Execute()

Update an existing user

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
 body := *openapiclient.NewUserEdit() // UserEdit | The user parameters
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateUser`: UserPostResponse
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserEdit**](UserEdit.md) | The user parameters |

### Return type

[**UserPostResponse**](UserPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserEmails

> UpdateUserEmails(ctx, userUuid).Body(body).Execute()

Update email addresses

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
 body := *openapiclient.NewUserEmailList() // UserEmailList | EmailAddressList
 userUuid := "userUuid_example" // string | The UUID of the user

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UpdateUserEmails(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserEmails``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserEmailsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserEmailList**](UserEmailList.md) | EmailAddressList |

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

## UserDeleteSession

> UserDeleteSession(ctx, userUuid, sessionUuid).Execute()

Delete a session

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
 sessionUuid := "sessionUuid_example" // string | The UUID of the session

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UserDeleteSession(context.Background(), userUuid, sessionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserDeleteSession``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |
**sessionUuid** | **string** | The UUID of the session |

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteSessionRequest struct via the builder pattern

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
