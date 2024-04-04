# \SipAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateLineEndpointSip**](SipAPI.md#AssociateLineEndpointSip) | **Put** /lines/{line_id}/endpoints/sip/{sip_uuid} | Associate line and SIP endpoint
[**AssociateTrunkEndpointSip**](SipAPI.md#AssociateTrunkEndpointSip) | **Put** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Associate trunk and SIP endpoint
[**CreateEndpointSip**](SipAPI.md#CreateEndpointSip) | **Post** /endpoints/sip | Create a SIP endpoint
[**CreateEndpointSipTemplate**](SipAPI.md#CreateEndpointSipTemplate) | **Post** /endpoints/sip/templates | Create a SIP endpoint template
[**CreateSipTransport**](SipAPI.md#CreateSipTransport) | **Post** /sip/transports | Create SIP transport
[**DeleteEndpointSip**](SipAPI.md#DeleteEndpointSip) | **Delete** /endpoints/sip/{sip_uuid} | Delete SIP Endpoint
[**DeleteEndpointSipTemplate**](SipAPI.md#DeleteEndpointSipTemplate) | **Delete** /endpoints/sip/templates/{template_uuid} | Delete SIP Endpoint Template
[**DeleteSipTransport**](SipAPI.md#DeleteSipTransport) | **Delete** /sip/transports/{transport_uuid} | Delete SIP transport
[**DissociateLineEndpointSip**](SipAPI.md#DissociateLineEndpointSip) | **Delete** /lines/{line_id}/endpoints/sip/{sip_uuid} | Dissociate line and SIP endpoint
[**DissociateTrunkEndpointSip**](SipAPI.md#DissociateTrunkEndpointSip) | **Delete** /trunks/{trunk_id}/endpoints/sip/{sip_uuid} | Dissociate trunk and SIP endpoint
[**GetEndpointSip**](SipAPI.md#GetEndpointSip) | **Get** /endpoints/sip/{sip_uuid} | Get SIP Endpoint
[**GetEndpointSipTemplate**](SipAPI.md#GetEndpointSipTemplate) | **Get** /endpoints/sip/templates/{template_uuid} | Get SIP Endpoint template
[**GetSipTransport**](SipAPI.md#GetSipTransport) | **Get** /sip/transports/{transport_uuid} | Get SIP transport
[**ListAsteriskPjsipGlobal**](SipAPI.md#ListAsteriskPjsipGlobal) | **Get** /asterisk/pjsip/global | List of PJSIP options for the &#x60;global&#x60; section
[**ListAsteriskPjsipSystem**](SipAPI.md#ListAsteriskPjsipSystem) | **Get** /asterisk/pjsip/system | List of PJSIP options for the &#x60;system&#x60; section
[**ListEndpointsSip**](SipAPI.md#ListEndpointsSip) | **Get** /endpoints/sip | List SIP endpoints
[**ListEndpointsSipTemplates**](SipAPI.md#ListEndpointsSipTemplates) | **Get** /endpoints/sip/templates | List SIP endpoints templates
[**ListSipTransports**](SipAPI.md#ListSipTransports) | **Get** /sip/transports | List all configured SIP transports
[**ShowPjsipDoc**](SipAPI.md#ShowPjsipDoc) | **Get** /asterisk/pjsip/doc | List all PJSIP configuration options
[**UpdateAsteriskPjsipGlobal**](SipAPI.md#UpdateAsteriskPjsipGlobal) | **Put** /asterisk/pjsip/global | Update PJSIP section options
[**UpdateAsteriskPjsipSystem**](SipAPI.md#UpdateAsteriskPjsipSystem) | **Put** /asterisk/pjsip/system | Update PJSIP section options
[**UpdateEndpointSip**](SipAPI.md#UpdateEndpointSip) | **Put** /endpoints/sip/{sip_uuid} | Update SIP Endpoint
[**UpdateEndpointSipTemplate**](SipAPI.md#UpdateEndpointSipTemplate) | **Put** /endpoints/sip/templates/{template_uuid} | Update SIP Endpoint Template
[**UpdateSipTransport**](SipAPI.md#UpdateSipTransport) | **Put** /sip/transports/{transport_uuid} | Update SIP transport

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
 r, err := apiClient.SipAPI.AssociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.AssociateLineEndpointSip``: %v\n", err)
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
 r, err := apiClient.SipAPI.AssociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.AssociateTrunkEndpointSip``: %v\n", err)
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
 resp, r, err := apiClient.SipAPI.CreateEndpointSip(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.CreateEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.CreateEndpointSip`: %v\n", resp)
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
 resp, r, err := apiClient.SipAPI.CreateEndpointSipTemplate(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.CreateEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSipTemplate`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.CreateEndpointSipTemplate`: %v\n", resp)
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

## CreateSipTransport

> SIPTransport CreateSipTransport(ctx).Body(body).Execute()

Create SIP transport

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
 body := *openapiclient.NewSIPTransport("Name_example", [][]string{[]string{"Options_example"}}) // SIPTransport | Transport to create

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.CreateSipTransport(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.CreateSipTransport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateSipTransport`: SIPTransport
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.CreateSipTransport`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSipTransportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SIPTransport**](SIPTransport.md) | Transport to create |

### Return type

[**SIPTransport**](SIPTransport.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
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
 r, err := apiClient.SipAPI.DeleteEndpointSip(context.Background(), sipUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.DeleteEndpointSip``: %v\n", err)
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
 r, err := apiClient.SipAPI.DeleteEndpointSipTemplate(context.Background(), templateUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.DeleteEndpointSipTemplate``: %v\n", err)
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

## DeleteSipTransport

> DeleteSipTransport(ctx, transportUuid).Fallback(fallback).Execute()

Delete SIP transport

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
 transportUuid := "transportUuid_example" // string | The UUID of the transport
 fallback := "fallback_example" // string | The UUID of the transport that should be associated to orphaned SIP configurations  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SipAPI.DeleteSipTransport(context.Background(), transportUuid).Fallback(fallback).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.DeleteSipTransport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportUuid** | **string** | The UUID of the transport |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSipTransportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fallback** | **string** | The UUID of the transport that should be associated to orphaned SIP configurations  |

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
 r, err := apiClient.SipAPI.DissociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.DissociateLineEndpointSip``: %v\n", err)
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
 r, err := apiClient.SipAPI.DissociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.DissociateTrunkEndpointSip``: %v\n", err)
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
 resp, r, err := apiClient.SipAPI.GetEndpointSip(context.Background(), sipUuid).AccentTenant(accentTenant).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.GetEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.GetEndpointSip`: %v\n", resp)
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
 resp, r, err := apiClient.SipAPI.GetEndpointSipTemplate(context.Background(), templateUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.GetEndpointSipTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSipTemplate`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.GetEndpointSipTemplate`: %v\n", resp)
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

## GetSipTransport

> SIPTransport GetSipTransport(ctx, transportUuid).Execute()

Get SIP transport

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
 transportUuid := "transportUuid_example" // string | The UUID of the transport

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.GetSipTransport(context.Background(), transportUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.GetSipTransport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSipTransport`: SIPTransport
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.GetSipTransport`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportUuid** | **string** | The UUID of the transport |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSipTransportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SIPTransport**](SIPTransport.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskPjsipGlobal

> PJSIPGlobal ListAsteriskPjsipGlobal(ctx).Execute()

List of PJSIP options for the `global` section

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.ListAsteriskPjsipGlobal(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ListAsteriskPjsipGlobal``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskPjsipGlobal`: PJSIPGlobal
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ListAsteriskPjsipGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskPjsipGlobalRequest struct via the builder pattern

### Return type

[**PJSIPGlobal**](PJSIPGlobal.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskPjsipSystem

> PJSIPSystem ListAsteriskPjsipSystem(ctx).Execute()

List of PJSIP options for the `system` section

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.ListAsteriskPjsipSystem(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ListAsteriskPjsipSystem``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskPjsipSystem`: PJSIPSystem
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ListAsteriskPjsipSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskPjsipSystemRequest struct via the builder pattern

### Return type

[**PJSIPSystem**](PJSIPSystem.md)

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
 resp, r, err := apiClient.SipAPI.ListEndpointsSip(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ListEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSip`: EndpointSIPItems
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ListEndpointsSip`: %v\n", resp)
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
 resp, r, err := apiClient.SipAPI.ListEndpointsSipTemplates(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ListEndpointsSipTemplates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSipTemplates`: EndpointSIPItems
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ListEndpointsSipTemplates`: %v\n", resp)
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

## ListSipTransports

> SIPTransportItems ListSipTransports(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List all configured SIP transports

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
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.ListSipTransports(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ListSipTransports``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListSipTransports`: SIPTransportItems
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ListSipTransports`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListSipTransportsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**SIPTransportItems**](SIPTransportItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ShowPjsipDoc

> PJSIPConfigurationOptions ShowPjsipDoc(ctx).Execute()

List all PJSIP configuration options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SipAPI.ShowPjsipDoc(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.ShowPjsipDoc``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ShowPjsipDoc`: PJSIPConfigurationOptions
 fmt.Fprintf(os.Stdout, "Response from `SipAPI.ShowPjsipDoc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowPjsipDocRequest struct via the builder pattern

### Return type

[**PJSIPConfigurationOptions**](PJSIPConfigurationOptions.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskPjsipGlobal

> UpdateAsteriskPjsipGlobal(ctx).Body(body).Execute()

Update PJSIP section options

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
 body := *openapiclient.NewPJSIPGlobal() // PJSIPGlobal | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SipAPI.UpdateAsteriskPjsipGlobal(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.UpdateAsteriskPjsipGlobal``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskPjsipGlobalRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PJSIPGlobal**](PJSIPGlobal.md) |  |

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

## UpdateAsteriskPjsipSystem

> UpdateAsteriskPjsipSystem(ctx).Body(body).Execute()

Update PJSIP section options

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
 body := *openapiclient.NewPJSIPSystem() // PJSIPSystem | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SipAPI.UpdateAsteriskPjsipSystem(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.UpdateAsteriskPjsipSystem``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskPjsipSystemRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PJSIPSystem**](PJSIPSystem.md) |  |

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
 r, err := apiClient.SipAPI.UpdateEndpointSip(context.Background(), sipUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.UpdateEndpointSip``: %v\n", err)
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
 r, err := apiClient.SipAPI.UpdateEndpointSipTemplate(context.Background(), templateUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.UpdateEndpointSipTemplate``: %v\n", err)
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

## UpdateSipTransport

> UpdateSipTransport(ctx, transportUuid).Body(body).Execute()

Update SIP transport

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
 body := *openapiclient.NewSIPTransport("Name_example", [][]string{[]string{"Options_example"}}) // SIPTransport | 
 transportUuid := "transportUuid_example" // string | The UUID of the transport

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SipAPI.UpdateSipTransport(context.Background(), transportUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SipAPI.UpdateSipTransport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportUuid** | **string** | The UUID of the transport |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSipTransportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SIPTransport**](SIPTransport.md) |  |

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
