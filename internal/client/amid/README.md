# Go API client for amid

Send AMI actions to Asterisk, providing token based authentication.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 0.0.1
- Generator version: 7.5.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://accentvoice.io/](https://accentvoice.io/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import amid "github.com/ryanwclark1/accent-ui2/internal/client/amid"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `amid.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), amid.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `amid.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), amid.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `amid.ContextOperationServerIndices` and `amid.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), amid.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), amid.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.accentvoice.io/1.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionAPI* | [**ActionAsteriskManager**](docs/ActionAPI.md#actionasteriskmanager) | **Post** /action/{action} | AMI action
*CommandAPI* | [**CommandAsteriskManager**](docs/CommandAPI.md#commandasteriskmanager) | **Post** /action/Command | AMI command
*StatusAPI* | [**StatusSummary**](docs/StatusAPI.md#statussummary) | **Get** /status | Print infos about internal status of accent-amid


## Documentation For Models

 - [Command](docs/Command.md)
 - [CommandResponse](docs/CommandResponse.md)
 - [ComponentWithStatus](docs/ComponentWithStatus.md)
 - [Error](docs/Error.md)
 - [Response](docs/Response.md)
 - [StatusSummary](docs/StatusSummary.md)
 - [StatusValue](docs/StatusValue.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### accent_auth_token

- **Type**: API key
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Auth-Token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		amid.ContextAPIKeys,
		map[string]amid.APIKey{
			"X-Auth-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

help@accentvoice.io

