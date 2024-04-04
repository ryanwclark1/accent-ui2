# \EndpointsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateLineEndpointCustom**](EndpointsAPI.md#AssociateLineEndpointCustom) | **Put** /lines/{line_id}/endpoints/custom/{custom_id} | Associate line and Custom endpoint
[**AssociateLineEndpointSccp**](EndpointsAPI.md#AssociateLineEndpointSccp) | **Put** /lines/{line_id}/endpoints/sccp/{sccp_id} | Associate line and SCCP endpoint
[**AssociateLineEndpointSip**](EndpointsAPI.md#AssociateLineEndpointSip) | **Put** /lines/{line_id}/endpoints/sip/{sip_uuid} | Associate line and SIP endpoint
[**AssociateTrunkEndpointCustom**](EndpointsAPI.md#AssociateTrunkEndpointCustom) | **Put** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Associate trunk and Custom endpoint
[**AssociateTrunkEndpointIax**](EndpointsAPI.md#AssociateTrunkEndpointIax) | **Put** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Associate trunk and IAX endpoint
[**AssociateTrunkEndpointSip**](EndpointsAPI.md#AssociateTrunkEndpointSip) | **Put** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Associate trunk and SIP endpoint
[**CreateEndpointCustom**](EndpointsAPI.md#CreateEndpointCustom) | **Post** /endpoints/custom | Create Custom endpoint
[**CreateEndpointIax**](EndpointsAPI.md#CreateEndpointIax) | **Post** /endpoints/iax | Create IAX endpoint
[**CreateEndpointSccp**](EndpointsAPI.md#CreateEndpointSccp) | **Post** /endpoints/sccp | Create SCCP endpoint
[**CreateEndpointSip**](EndpointsAPI.md#CreateEndpointSip) | **Post** /endpoints/sip | Create a SIP endpoint
[**CreateEndpointSipTemplate**](EndpointsAPI.md#CreateEndpointSipTemplate) | **Post** /endpoints/sip/templates | Create a SIP endpoint template
[**DeleteEndpointCustom**](EndpointsAPI.md#DeleteEndpointCustom) | **Delete** /endpoints/custom/{custom_id} | Delete Custom Endpoint
[**DeleteEndpointIax**](EndpointsAPI.md#DeleteEndpointIax) | **Delete** /endpoints/iax/{iax_id} | Delete IAX Endpoint
[**DeleteEndpointSccp**](EndpointsAPI.md#DeleteEndpointSccp) | **Delete** /endpoints/sccp/{sccp_id} | Delete SCCP Endpoint
[**DeleteEndpointSip**](EndpointsAPI.md#DeleteEndpointSip) | **Delete** /endpoints/sip/{sip_uuid} | Delete SIP Endpoint
[**DeleteEndpointSipTemplate**](EndpointsAPI.md#DeleteEndpointSipTemplate) | **Delete** /endpoints/sip/templates/{template_uuid} | Delete SIP Endpoint Template
[**DissociateLineEndpointCustom**](EndpointsAPI.md#DissociateLineEndpointCustom) | **Delete** /lines/{line_id}/endpoints/custom/{custom_id} | Dissociate line and Custom endpoint
[**DissociateLineEndpointSccp**](EndpointsAPI.md#DissociateLineEndpointSccp) | **Delete** /lines/{line_id}/endpoints/sccp/{sccp_id} | Dissociate line and SCCP endpoint
[**DissociateLineEndpointSip**](EndpointsAPI.md#DissociateLineEndpointSip) | **Delete** /lines/{line_id}/endpoints/sip/{sip_uuid} | Dissociate line and SIP endpoint
[**DissociateTrunkEndpointCustom**](EndpointsAPI.md#DissociateTrunkEndpointCustom) | **Delete** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Dissociate trunk and Custom endpoint
[**DissociateTrunkEndpointIax**](EndpointsAPI.md#DissociateTrunkEndpointIax) | **Delete** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Dissociate trunk and IAX endpoint
[**DissociateTrunkEndpointSip**](EndpointsAPI.md#DissociateTrunkEndpointSip) | **Delete** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Dissociate trunk and SIP endpoint
[**GetEndpointCustom**](EndpointsAPI.md#GetEndpointCustom) | **Get** /endpoints/custom/{custom_id} | Get Custom Endpoint
[**GetEndpointIax**](EndpointsAPI.md#GetEndpointIax) | **Get** /endpoints/iax/{iax_id} | Get IAX Endpoint
[**GetEndpointSccp**](EndpointsAPI.md#GetEndpointSccp) | **Get** /endpoints/sccp/{sccp_id} | Get SCCP Endpoint
[**GetEndpointSip**](EndpointsAPI.md#GetEndpointSip) | **Get** /endpoints/sip/{sip_uuid} | Get SIP Endpoint
[**GetEndpointSipTemplate**](EndpointsAPI.md#GetEndpointSipTemplate) | **Get** /endpoints/sip/templates/{template_uuid} | Get SIP Endpoint template
[**GetUserLineAssociatedEndpointsSip**](EndpointsAPI.md#GetUserLineAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/{line_id}/associated/endpoints/sip | Get SIP endpoint of a line for a user
[**GetUserLineMainAssociatedEndpointsSip**](EndpointsAPI.md#GetUserLineMainAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/main/associated/endpoints/sip | Get SIP endpoint of main line for a user
[**ListEndpointsCustom**](EndpointsAPI.md#ListEndpointsCustom) | **Get** /endpoints/custom | List Custom endpoints
[**ListEndpointsIax**](EndpointsAPI.md#ListEndpointsIax) | **Get** /endpoints/iax | List IAX endpoints
[**ListEndpointsSccp**](EndpointsAPI.md#ListEndpointsSccp) | **Get** /endpoints/sccp | List SCCP endpoints
[**ListEndpointsSip**](EndpointsAPI.md#ListEndpointsSip) | **Get** /endpoints/sip | List SIP endpoints
[**ListEndpointsSipTemplates**](EndpointsAPI.md#ListEndpointsSipTemplates) | **Get** /endpoints/sip/templates | List SIP endpoints templates
[**UpdateEndpointCustom**](EndpointsAPI.md#UpdateEndpointCustom) | **Put** /endpoints/custom/{custom_id} | Update Custom Endpoint
[**UpdateEndpointIax**](EndpointsAPI.md#UpdateEndpointIax) | **Put** /endpoints/iax/{iax_id} | Update IAX Endpoint
[**UpdateEndpointSccp**](EndpointsAPI.md#UpdateEndpointSccp) | **Put** /endpoints/sccp/{sccp_id} | Update SCCP Endpoint
[**UpdateEndpointSip**](EndpointsAPI.md#UpdateEndpointSip) | **Put** /endpoints/sip/{sip_uuid} | Update SIP Endpoint
[**UpdateEndpointSipTemplate**](EndpointsAPI.md#UpdateEndpointSipTemplate) | **Put** /endpoints/sip/templates/{template_uuid} | Update SIP Endpoint Template

## AssociateLineEndpointCustom

> AssociateLineEndpointCustom(ctx, lineId, customId).Execute()

Associate line and Custom endpoint

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
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateLineEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointCustomRequest struct via the builder pattern

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

## AssociateLineEndpointSccp

> AssociateLineEndpointSccp(ctx, lineId, sccpId).Execute()

Associate line and SCCP endpoint

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
 lineId := int32(56) // int32 | 
 sccpId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateLineEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointSccpRequest struct via the builder pattern

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

## AssociateLineEndpointSip

> AssociateLineEndpointSip(ctx, lineId, sipUuid).Execute()

Associate line and SIP endpoint

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
 lineId := int32(56) // int32 | 
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateLineEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointSipRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateTrunkEndpointCustom``: %v\n", err)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateTrunkEndpointIax``: %v\n", err)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.AssociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.AssociateTrunkEndpointSip``: %v\n", err)
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

## CreateEndpointCustom

> EndpointCustom CreateEndpointCustom(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create Custom endpoint

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
 body := *openapiclient.NewEndpointCustom() // EndpointCustom | Custom Endpoint to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.CreateEndpointCustom(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointCustom`: EndpointCustom
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointCustom`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointCustom**](EndpointCustom.md) | Custom Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointCustom**](EndpointCustom.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEndpointIax

> EndpointIAX CreateEndpointIax(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create IAX endpoint

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
 body := *openapiclient.NewEndpointIAX() // EndpointIAX | IAX Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.CreateEndpointIax(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointIax`: EndpointIAX
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointIAX**](EndpointIAX.md) | IAX Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointIAX**](EndpointIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEndpointSccp

> EndpointSccp CreateEndpointSccp(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create SCCP endpoint

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
 body := *openapiclient.NewEndpointSccp() // EndpointSccp | SCCP Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.CreateEndpointSccp(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSccp`: EndpointSccp
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointSccp`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSccp**](EndpointSccp.md) | SCCP Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSccp**](EndpointSccp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEndpointSip

> EndpointSIP CreateEndpointSip(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a SIP endpoint

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
 body := *openapiclient.NewEndpointSIP() // EndpointSIP | SIP Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.CreateEndpointSip(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointSip`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSIP**](EndpointSIP.md) | SIP Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateEndpointSipTemplate

> EndpointSIP CreateEndpointSipTemplate(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a SIP endpoint template

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
 body := *openapiclient.NewEndpointSIP() // EndpointSIP | SIP Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.CreateEndpointSipTemplate(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSipTemplate`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointSipTemplate`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointSipTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSIP**](EndpointSIP.md) | SIP Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteEndpointCustom

> DeleteEndpointCustom(ctx, customId).AccentTenant(accentTenant).Execute()

Delete Custom Endpoint

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
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DeleteEndpointCustom(context.Background(), customId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointCustomRequest struct via the builder pattern

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

## DeleteEndpointIax

> DeleteEndpointIax(ctx, iaxId).AccentTenant(accentTenant).Execute()

Delete IAX Endpoint

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
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DeleteEndpointIax(context.Background(), iaxId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointIaxRequest struct via the builder pattern

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

## DeleteEndpointSccp

> DeleteEndpointSccp(ctx, sccpId).AccentTenant(accentTenant).Execute()

Delete SCCP Endpoint

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
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DeleteEndpointSccp(context.Background(), sccpId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointSccpRequest struct via the builder pattern

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

## DeleteEndpointSip

> DeleteEndpointSip(ctx, sipUuid).AccentTenant(accentTenant).Execute()

Delete SIP Endpoint

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
 sipUuid := "sipUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DeleteEndpointSip(context.Background(), sipUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointSipRequest struct via the builder pattern

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

## DeleteEndpointSipTemplate

> DeleteEndpointSipTemplate(ctx, templateUuid).AccentTenant(accentTenant).Execute()

Delete SIP Endpoint Template

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
 templateUuid := "templateUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DeleteEndpointSipTemplate(context.Background(), templateUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointSipTemplateRequest struct via the builder pattern

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

## DissociateLineEndpointCustom

> DissociateLineEndpointCustom(ctx, lineId, customId).Execute()

Dissociate line and Custom endpoint

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
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateLineEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointCustomRequest struct via the builder pattern

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

## DissociateLineEndpointSccp

> DissociateLineEndpointSccp(ctx, lineId, sccpId).Execute()

Dissociate line and SCCP endpoint

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
 lineId := int32(56) // int32 | 
 sccpId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateLineEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointSccpRequest struct via the builder pattern

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

## DissociateLineEndpointSip

> DissociateLineEndpointSip(ctx, lineId, sipUuid).Execute()

Dissociate line and SIP endpoint

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
 lineId := int32(56) // int32 | 
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateLineEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointSipRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateTrunkEndpointCustom``: %v\n", err)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateTrunkEndpointIax``: %v\n", err)
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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 trunkId := int32(56) // int32 | Trunk's ID
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.DissociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DissociateTrunkEndpointSip``: %v\n", err)
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

## GetEndpointCustom

> EndpointCustom GetEndpointCustom(ctx, customId).AccentTenant(accentTenant).Execute()

Get Custom Endpoint

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
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetEndpointCustom(context.Background(), customId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointCustom`: EndpointCustom
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointCustom`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointCustom**](EndpointCustom.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetEndpointIax

> EndpointIAX GetEndpointIax(ctx, iaxId).AccentTenant(accentTenant).Execute()

Get IAX Endpoint

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
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetEndpointIax(context.Background(), iaxId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointIax`: EndpointIAX
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointIax`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointIAX**](EndpointIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetEndpointSccp

> EndpointSccp GetEndpointSccp(ctx, sccpId).AccentTenant(accentTenant).Execute()

Get SCCP Endpoint

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
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetEndpointSccp(context.Background(), sccpId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSccp`: EndpointSccp
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointSccp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSccp**](EndpointSccp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetEndpointSip

> EndpointSIP GetEndpointSip(ctx, sipUuid).AccentTenant(accentTenant).View(view).Execute()

Get SIP Endpoint

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
 sipUuid := "sipUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetEndpointSip(context.Background(), sipUuid).AccentTenant(accentTenant).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetEndpointSipTemplate

> EndpointSIP GetEndpointSipTemplate(ctx, templateUuid).AccentTenant(accentTenant).Execute()

Get SIP Endpoint template

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
 templateUuid := "templateUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetEndpointSipTemplate(context.Background(), templateUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSipTemplate`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointSipTemplate`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointSipTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineAssociatedEndpointsSip

> EndpointSIP GetUserLineAssociatedEndpointsSip(ctx, userUuid, lineId).View(view).Execute()

Get SIP endpoint of a line for a user

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
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 lineId := int32(56) // int32 | 
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetUserLineAssociatedEndpointsSip(context.Background(), userUuid, lineId).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetUserLineAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetUserLineAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineMainAssociatedEndpointsSip

> EndpointSIP GetUserLineMainAssociatedEndpointsSip(ctx, userUuid).View(view).Execute()

Get SIP endpoint of main line for a user

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
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.EndpointsAPI.GetUserLineMainAssociatedEndpointsSip(context.Background(), userUuid).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetUserLineMainAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineMainAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetUserLineMainAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineMainAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsCustom

> EndpointCustomItems ListEndpointsCustom(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List Custom endpoints

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
 resp, r, err := apiClient.EndpointsAPI.ListEndpointsCustom(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpointsCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsCustom`: EndpointCustomItems
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpointsCustom`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsCustomRequest struct via the builder pattern

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

[**EndpointCustomItems**](EndpointCustomItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsIax

> EndpointIAXItems ListEndpointsIax(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List IAX endpoints

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
 resp, r, err := apiClient.EndpointsAPI.ListEndpointsIax(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpointsIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsIax`: EndpointIAXItems
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpointsIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsIaxRequest struct via the builder pattern

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

[**EndpointIAXItems**](EndpointIAXItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsSccp

> EndpointSccpItems ListEndpointsSccp(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List SCCP endpoints

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
 resp, r, err := apiClient.EndpointsAPI.ListEndpointsSccp(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpointsSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSccp`: EndpointSccpItems
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpointsSccp`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsSccpRequest struct via the builder pattern

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

[**EndpointSccpItems**](EndpointSccpItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsSip

> EndpointSIPItems ListEndpointsSip(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List SIP endpoints

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
 resp, r, err := apiClient.EndpointsAPI.ListEndpointsSip(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSip`: EndpointSIPItems
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsSipRequest struct via the builder pattern

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

[**EndpointSIPItems**](EndpointSIPItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsSipTemplates

> EndpointSIPItems ListEndpointsSipTemplates(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List SIP endpoints templates

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
 resp, r, err := apiClient.EndpointsAPI.ListEndpointsSipTemplates(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpointsSipTemplates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSipTemplates`: EndpointSIPItems
 fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpointsSipTemplates`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsSipTemplatesRequest struct via the builder pattern

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

[**EndpointSIPItems**](EndpointSIPItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateEndpointCustom

> UpdateEndpointCustom(ctx, customId).Body(body).AccentTenant(accentTenant).Execute()

Update Custom Endpoint

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
 body := *openapiclient.NewEndpointCustom() // EndpointCustom | 
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.UpdateEndpointCustom(context.Background(), customId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointCustom**](EndpointCustom.md) |  |

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

## UpdateEndpointIax

> UpdateEndpointIax(ctx, iaxId).Body(body).AccentTenant(accentTenant).Execute()

Update IAX Endpoint

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
 body := *openapiclient.NewEndpointIAX() // EndpointIAX | 
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.UpdateEndpointIax(context.Background(), iaxId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointIAX**](EndpointIAX.md) |  |

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

## UpdateEndpointSccp

> UpdateEndpointSccp(ctx, sccpId).Body(body).AccentTenant(accentTenant).Execute()

Update SCCP Endpoint

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
 body := *openapiclient.NewEndpointSccp() // EndpointSccp | 
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.UpdateEndpointSccp(context.Background(), sccpId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSccp**](EndpointSccp.md) |  |

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

## UpdateEndpointSip

> UpdateEndpointSip(ctx, sipUuid).Body(body).AccentTenant(accentTenant).Execute()

Update SIP Endpoint

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
 body := *openapiclient.NewEndpointSIP() // EndpointSIP | 
 sipUuid := "sipUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.UpdateEndpointSip(context.Background(), sipUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSIP**](EndpointSIP.md) |  |

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

## UpdateEndpointSipTemplate

> UpdateEndpointSipTemplate(ctx, templateUuid).Body(body).AccentTenant(accentTenant).Execute()

Update SIP Endpoint Template

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
 body := *openapiclient.NewEndpointSIP() // EndpointSIP | 
 templateUuid := "templateUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.EndpointsAPI.UpdateEndpointSipTemplate(context.Background(), templateUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointSipTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSIP**](EndpointSIP.md) |  |

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
