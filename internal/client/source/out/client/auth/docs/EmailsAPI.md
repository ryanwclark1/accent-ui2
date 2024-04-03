# \EmailsAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailConfirm**](EmailsAPI.md#GetEmailConfirm) | **Get** /emails/{email_uuid}/confirm | Confirm an email address
[**GetNewEmailConfirmation**](EmailsAPI.md#GetNewEmailConfirmation) | **Get** /users/{user_uuid}/emails/{email_uuid}/confirm | Ask a new confirmation email
[**PutEmailConfirm**](EmailsAPI.md#PutEmailConfirm) | **Put** /emails/{email_uuid}/confirm | Confirm an email address
[**UpdateAllUserEmails**](EmailsAPI.md#UpdateAllUserEmails) | **Put** /admin/users/{user_uuid}/emails | Update email addresses
[**UpdateUserEmails**](EmailsAPI.md#UpdateUserEmails) | **Put** /users/{user_uuid}/emails | Update email addresses

## GetEmailConfirm

> GetEmailConfirm(ctx, emailUuid).Token(token).Execute()

Confirm an email address

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
 emailUuid := "emailUuid_example" // string | The UUID of the email
 token := "token_example" // string | The UUID of the token used to confirm the email address

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EmailsAPI.GetEmailConfirm(context.Background(), emailUuid).Token(token).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.GetEmailConfirm``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailUuid** | **string** | The UUID of the email |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailConfirmRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | The UUID of the token used to confirm the email address |

### Return type

 (empty response body)

### Authorization

No authorization required

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
 r, err := apiClient.EmailsAPI.GetNewEmailConfirmation(context.Background(), userUuid, emailUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.GetNewEmailConfirmation``: %v\n", err)
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

## PutEmailConfirm

> PutEmailConfirm(ctx, emailUuid).Execute()

Confirm an email address

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
 emailUuid := "emailUuid_example" // string | The UUID of the email

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EmailsAPI.PutEmailConfirm(context.Background(), emailUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.PutEmailConfirm``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailUuid** | **string** | The UUID of the email |

### Other Parameters

Other parameters are passed through a pointer to a apiPutEmailConfirmRequest struct via the builder pattern

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
 r, err := apiClient.EmailsAPI.UpdateAllUserEmails(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.UpdateAllUserEmails``: %v\n", err)
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
 r, err := apiClient.EmailsAPI.UpdateUserEmails(context.Background(), userUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.UpdateUserEmails``: %v\n", err)
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
