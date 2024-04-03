# \SoundsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSounds**](SoundsAPI.md#CreateSounds) | **Post** /sounds | Create sound category
[**DeleteSounds**](SoundsAPI.md#DeleteSounds) | **Delete** /sounds/{sound_category} | Delete sound category
[**DeleteSoundsFiles**](SoundsAPI.md#DeleteSoundsFiles) | **Delete** /sounds/{sound_category}/files/{sound_filename} | Delete audio file
[**GetSounds**](SoundsAPI.md#GetSounds) | **Get** /sounds/{sound_category} | Get sound category
[**GetSoundsFiles**](SoundsAPI.md#GetSoundsFiles) | **Get** /sounds/{sound_category}/files/{sound_filename} | Get audio file
[**ListSounds**](SoundsAPI.md#ListSounds) | **Get** /sounds | List sound categories
[**ListSoundsLanguages**](SoundsAPI.md#ListSoundsLanguages) | **Get** /sounds/languages | List all languages for sounds
[**UpdateSoundsFiles**](SoundsAPI.md#UpdateSoundsFiles) | **Put** /sounds/{sound_category}/files/{sound_filename} | Add or update audio file

## CreateSounds

> Sound CreateSounds(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create sound category

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
 body := *openapiclient.NewSound("Name_example") // Sound | Sound category to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SoundsAPI.CreateSounds(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.CreateSounds``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateSounds`: Sound
 fmt.Fprintf(os.Stdout, "Response from `SoundsAPI.CreateSounds`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoundsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Sound**](Sound.md) | Sound category to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Sound**](Sound.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteSounds

> DeleteSounds(ctx, soundCategory).AccentTenant(accentTenant).Execute()

Delete sound category

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
 soundCategory := "soundCategory_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SoundsAPI.DeleteSounds(context.Background(), soundCategory).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.DeleteSounds``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundCategory** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoundsRequest struct via the builder pattern

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

## DeleteSoundsFiles

> DeleteSoundsFiles(ctx, soundCategory, soundFilename).AccentTenant(accentTenant).Language(language).Format(format).Execute()

Delete audio file

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
 soundCategory := "soundCategory_example" // string | 
 soundFilename := "soundFilename_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 language := "language_example" // string | Language of the sound (optional)
 format := "format_example" // string | Format of the sound (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SoundsAPI.DeleteSoundsFiles(context.Background(), soundCategory, soundFilename).AccentTenant(accentTenant).Language(language).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.DeleteSoundsFiles``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundCategory** | **string** |  |
**soundFilename** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoundsFilesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **language** | **string** | Language of the sound |
 **format** | **string** | Format of the sound |

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

## GetSounds

> Sound GetSounds(ctx, soundCategory).AccentTenant(accentTenant).Execute()

Get sound category

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
 soundCategory := "soundCategory_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SoundsAPI.GetSounds(context.Background(), soundCategory).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.GetSounds``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSounds`: Sound
 fmt.Fprintf(os.Stdout, "Response from `SoundsAPI.GetSounds`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundCategory** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoundsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Sound**](Sound.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSoundsFiles

> GetSoundsFiles(ctx, soundCategory, soundFilename).AccentTenant(accentTenant).Language(language).Format(format).Execute()

Get audio file

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
 soundCategory := "soundCategory_example" // string | 
 soundFilename := "soundFilename_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 language := "language_example" // string | Language of the sound (optional)
 format := "format_example" // string | Format of the sound (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SoundsAPI.GetSoundsFiles(context.Background(), soundCategory, soundFilename).AccentTenant(accentTenant).Language(language).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.GetSoundsFiles``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundCategory** | **string** |  |
**soundFilename** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoundsFilesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **language** | **string** | Language of the sound |
 **format** | **string** | Format of the sound |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSounds

> SoundItems ListSounds(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List sound categories

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
 resp, r, err := apiClient.SoundsAPI.ListSounds(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.ListSounds``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListSounds`: SoundItems
 fmt.Fprintf(os.Stdout, "Response from `SoundsAPI.ListSounds`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListSoundsRequest struct via the builder pattern

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

[**SoundItems**](SoundItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSoundsLanguages

> SoundLanguageItems ListSoundsLanguages(ctx).Execute()

List all languages for sounds

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
 resp, r, err := apiClient.SoundsAPI.ListSoundsLanguages(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.ListSoundsLanguages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListSoundsLanguages`: SoundLanguageItems
 fmt.Fprintf(os.Stdout, "Response from `SoundsAPI.ListSoundsLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSoundsLanguagesRequest struct via the builder pattern

### Return type

[**SoundLanguageItems**](SoundLanguageItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateSoundsFiles

> UpdateSoundsFiles(ctx, soundCategory, soundFilename).Body(body).AccentTenant(accentTenant).Language(language).Format(format).Execute()

Add or update audio file

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
 body := map[string]interface{}{ ... } // map[string]interface{} | 
 soundCategory := "soundCategory_example" // string | 
 soundFilename := "soundFilename_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 language := "language_example" // string | Language of the sound (optional)
 format := "format_example" // string | Format of the sound (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SoundsAPI.UpdateSoundsFiles(context.Background(), soundCategory, soundFilename).Body(body).AccentTenant(accentTenant).Language(language).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SoundsAPI.UpdateSoundsFiles``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundCategory** | **string** |  |
**soundFilename** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoundsFilesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **language** | **string** | Language of the sound |
 **format** | **string** | Format of the sound |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
