# \TrunksAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateOutcallTrunks**](TrunksAPI.md#AssociateOutcallTrunks) | **Put** /outcalls/{outcall_id}/trunks | Associate outcall and trunks
[**AssociateTrunkEndpointCustom**](TrunksAPI.md#AssociateTrunkEndpointCustom) | **Put** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Associate trunk and Custom endpoint
[**AssociateTrunkEndpointIax**](TrunksAPI.md#AssociateTrunkEndpointIax) | **Put** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Associate trunk and IAX endpoint
[**AssociateTrunkEndpointSip**](TrunksAPI.md#AssociateTrunkEndpointSip) | **Put** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Associate trunk and SIP endpoint
[**AssociateTrunkRegisterIax**](TrunksAPI.md#AssociateTrunkRegisterIax) | **Put** /trunks/{trunk_id}/registers/iax/{iax_id} | Associate trunk and IAX register
[**CreateTrunk**](TrunksAPI.md#CreateTrunk) | **Post** /trunks | Create trunk
[**DeleteTrunk**](TrunksAPI.md#DeleteTrunk) | **Delete** /trunks/{trunk_id} | Delete trunk
[**DissociateTrunkEndpointCustom**](TrunksAPI.md#DissociateTrunkEndpointCustom) | **Delete** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Dissociate trunk and Custom endpoint
[**DissociateTrunkEndpointIax**](TrunksAPI.md#DissociateTrunkEndpointIax) | **Delete** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Dissociate trunk and IAX endpoint
[**DissociateTrunkEndpointSip**](TrunksAPI.md#DissociateTrunkEndpointSip) | **Delete** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Dissociate trunk and SIP endpoint
[**DissociateTrunkRegisterIax**](TrunksAPI.md#DissociateTrunkRegisterIax) | **Delete** /trunks/{trunk_id}/registers/iax/{iax_id} | Dissociate trunk and IAX register
[**GetTrunk**](TrunksAPI.md#GetTrunk) | **Get** /trunks/{trunk_id} | Get trunk
[**ListTrunks**](TrunksAPI.md#ListTrunks) | **Get** /trunks | List trunks
[**UpdateTrunk**](TrunksAPI.md#UpdateTrunk) | **Put** /trunks/{trunk_id} | Update trunk

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
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewTrunksId() // TrunksId | Trunks to associated
 outcallId := int32(56) // int32 | Outgoing call's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.AssociateOutcallTrunks(context.Background(), outcallId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.AssociateOutcallTrunks``: %v\n", err)
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

## AssociateTrunkEndpointCustom

> AssociateTrunkEndpointCustom(ctx, trunkId, customId).Execute()

Associate trunk and Custom endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.AssociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.AssociateTrunkEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkEndpointCustomRequest struct via the builder pattern

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

## AssociateTrunkEndpointIax

> AssociateTrunkEndpointIax(ctx, trunkId, iaxId).Execute()

Associate trunk and IAX endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.AssociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.AssociateTrunkEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkEndpointIaxRequest struct via the builder pattern

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

## AssociateTrunkEndpointSip

> AssociateTrunkEndpointSip(ctx, trunkId, sipUuid).Execute()

Associate trunk and SIP endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.AssociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.AssociateTrunkEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkEndpointSipRequest struct via the builder pattern

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

## AssociateTrunkRegisterIax

> AssociateTrunkRegisterIax(ctx, trunkId, iaxId).Execute()

Associate trunk and IAX register

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
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.AssociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.AssociateTrunkRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkRegisterIaxRequest struct via the builder pattern

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

## CreateTrunk

> Trunk CreateTrunk(ctx).Body(body).Execute()

Create trunk

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
 body := *openapiclient.NewTrunk() // Trunk | Trunk to create (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TrunksAPI.CreateTrunk(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.CreateTrunk``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateTrunk`: Trunk
 fmt.Fprintf(os.Stdout, "Response from `TrunksAPI.CreateTrunk`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrunkRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Trunk**](Trunk.md) | Trunk to create |

### Return type

[**Trunk**](Trunk.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteTrunk

> DeleteTrunk(ctx, trunkId).AccentTenant(accentTenant).Execute()

Delete trunk

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
 trunkId := int32(56) // int32 | Trunk's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.DeleteTrunk(context.Background(), trunkId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.DeleteTrunk``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrunkRequest struct via the builder pattern

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

## DissociateTrunkEndpointCustom

> DissociateTrunkEndpointCustom(ctx, trunkId, customId).Execute()

Dissociate trunk and Custom endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.DissociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.DissociateTrunkEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkEndpointCustomRequest struct via the builder pattern

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

## DissociateTrunkEndpointIax

> DissociateTrunkEndpointIax(ctx, trunkId, iaxId).Execute()

Dissociate trunk and IAX endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.DissociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.DissociateTrunkEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkEndpointIaxRequest struct via the builder pattern

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

## DissociateTrunkEndpointSip

> DissociateTrunkEndpointSip(ctx, trunkId, sipUuid).Execute()

Dissociate trunk and SIP endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.DissociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.DissociateTrunkEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkEndpointSipRequest struct via the builder pattern

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

## DissociateTrunkRegisterIax

> DissociateTrunkRegisterIax(ctx, trunkId, iaxId).Execute()

Dissociate trunk and IAX register

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
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.DissociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.DissociateTrunkRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkRegisterIaxRequest struct via the builder pattern

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

## GetTrunk

> Trunk GetTrunk(ctx, trunkId).AccentTenant(accentTenant).Execute()

Get trunk

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
 trunkId := int32(56) // int32 | Trunk's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TrunksAPI.GetTrunk(context.Background(), trunkId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.GetTrunk``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetTrunk`: Trunk
 fmt.Fprintf(os.Stdout, "Response from `TrunksAPI.GetTrunk`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrunkRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Trunk**](Trunk.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListTrunks

> TrunkItems ListTrunks(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List trunks

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
 resp, r, err := apiClient.TrunksAPI.ListTrunks(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.ListTrunks``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListTrunks`: TrunkItems
 fmt.Fprintf(os.Stdout, "Response from `TrunksAPI.ListTrunks`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListTrunksRequest struct via the builder pattern

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

[**TrunkItems**](TrunkItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateTrunk

> UpdateTrunk(ctx, trunkId).Body(body).AccentTenant(accentTenant).Execute()

Update trunk

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
 body := *openapiclient.NewTrunk() // Trunk | 
 trunkId := int32(56) // int32 | Trunk's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TrunksAPI.UpdateTrunk(context.Background(), trunkId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TrunksAPI.UpdateTrunk``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrunkRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Trunk**](Trunk.md) |  |

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
