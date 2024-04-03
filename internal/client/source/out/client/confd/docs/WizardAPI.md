# \WizardAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWizardDiscover**](WizardAPI.md#GetWizardDiscover) | **Get** /wizard/discover | Get wizard discover
[**GetWizardStatus**](WizardAPI.md#GetWizardStatus) | **Get** /wizard | Get wizard status
[**PassWizard**](WizardAPI.md#PassWizard) | **Post** /wizard | Pass the wizard

## GetWizardDiscover

> WizardDiscover GetWizardDiscover(ctx).Execute()

Get wizard discover

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
 resp, r, err := apiClient.WizardAPI.GetWizardDiscover(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `WizardAPI.GetWizardDiscover``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetWizardDiscover`: WizardDiscover
 fmt.Fprintf(os.Stdout, "Response from `WizardAPI.GetWizardDiscover`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWizardDiscoverRequest struct via the builder pattern

### Return type

[**WizardDiscover**](WizardDiscover.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetWizardStatus

> WizardConfigured GetWizardStatus(ctx).Execute()

Get wizard status

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
 resp, r, err := apiClient.WizardAPI.GetWizardStatus(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `WizardAPI.GetWizardStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetWizardStatus`: WizardConfigured
 fmt.Fprintf(os.Stdout, "Response from `WizardAPI.GetWizardStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWizardStatusRequest struct via the builder pattern

### Return type

[**WizardConfigured**](WizardConfigured.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PassWizard

> Wizard PassWizard(ctx).Body(body).Execute()

Pass the wizard

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
 body := *openapiclient.NewWizard("AdminPassword_example", false, *openapiclient.NewWizardNetwork("Domain_example", "Gateway_example", "Hostname_example", "Interface_example", "IpAddress_example", []string{"Nameservers_example"}, "Netmask_example"), "Timezone_example") // Wizard | Wizard parameters to configure

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.WizardAPI.PassWizard(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `WizardAPI.PassWizard``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PassWizard`: Wizard
 fmt.Fprintf(os.Stdout, "Response from `WizardAPI.PassWizard`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPassWizardRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Wizard**](Wizard.md) | Wizard parameters to configure |

### Return type

[**Wizard**](Wizard.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
