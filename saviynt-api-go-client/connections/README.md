# Go API client for connections

Use this API to create a connection in Saviynt Identity Cloud.

The Authorization header must have \"Bearer {token}\".

**Mandatory Parameters:**
- **connectionname**: Specify the name to identify the connection.
- **connectiontype**: Specify a connection type. For example, if your target application is Active Directory, specify the connection type as \"AD\".

**Optional Parameters:**
- **description**: Provide a description for the connection.
- **defaultsavroles**: Specify the SAV role(s) required for managing this connection along with its associated security systems, endpoints, accounts, and entitlements.
- **emailTemplate**: Specify the email template applicable for notifications.
- **sslCertificate**: Specify the SSL certificate(s) to secure the connection between EIC and the target application.
- **vaultConfiguration**: Specify the path of the vault to obtain secret data (suffix the connector name to make it unique).
- **saveinvault**: Set to true to save the encrypted attribute in the configured vault.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.13.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import connections "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `connections.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), connections.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `connections.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), connections.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `connections.ContextOperationServerIndices` and `connections.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), connections.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), connections.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConnectionsAPI* | [**CreateOrUpdate**](docs/ConnectionsAPI.md#createorupdate) | **Post** /ECM/api/v5/testConnection | Create a connection
*ConnectionsAPI* | [**GetConnectionDetails**](docs/ConnectionsAPI.md#getconnectiondetails) | **Post** /ECM/api/v5/getConnectionDetails | Get connection details
*ConnectionsAPI* | [**GetConnections**](docs/ConnectionsAPI.md#getconnections) | **Post** /ECM/api/v5/getConnections | Get list of connections


## Documentation For Models

 - [ADConnectionAttributes](docs/ADConnectionAttributes.md)
 - [ADConnectionResponse](docs/ADConnectionResponse.md)
 - [ADConnector](docs/ADConnector.md)
 - [ADSIConnectionAttributes](docs/ADSIConnectionAttributes.md)
 - [ADSIConnectionResponse](docs/ADSIConnectionResponse.md)
 - [ADSIConnector](docs/ADSIConnector.md)
 - [BaseConnector](docs/BaseConnector.md)
 - [ConnectionTimeoutConfig](docs/ConnectionTimeoutConfig.md)
 - [CreateOrUpdateRequest](docs/CreateOrUpdateRequest.md)
 - [CreateOrUpdateResponse](docs/CreateOrUpdateResponse.md)
 - [D365Connector](docs/D365Connector.md)
 - [DBConnectionAttributes](docs/DBConnectionAttributes.md)
 - [DBConnectionResponse](docs/DBConnectionResponse.md)
 - [DBConnector](docs/DBConnector.md)
 - [EntraIDConnectionAttributes](docs/EntraIDConnectionAttributes.md)
 - [EntraIDConnectionResponse](docs/EntraIDConnectionResponse.md)
 - [EntraIDConnector](docs/EntraIDConnector.md)
 - [GetConnectionDetails](docs/GetConnectionDetails.md)
 - [GetConnectionDetailsRequest](docs/GetConnectionDetailsRequest.md)
 - [GetConnectionDetailsResponse](docs/GetConnectionDetailsResponse.md)
 - [GetConnectionsRequest](docs/GetConnectionsRequest.md)
 - [GetConnectionsResponse](docs/GetConnectionsResponse.md)
 - [GetConnectionsResponseConnectionListInner](docs/GetConnectionsResponseConnectionListInner.md)
 - [GithubRESTConnectionAttributes](docs/GithubRESTConnectionAttributes.md)
 - [GithubRESTConnectionResponse](docs/GithubRESTConnectionResponse.md)
 - [GithubRESTConnector](docs/GithubRESTConnector.md)
 - [OktaConnectionAttributes](docs/OktaConnectionAttributes.md)
 - [OktaConnectionResponse](docs/OktaConnectionResponse.md)
 - [OktaConnector](docs/OktaConnector.md)
 - [RESTConnectionAttributes](docs/RESTConnectionAttributes.md)
 - [RESTConnectionResponse](docs/RESTConnectionResponse.md)
 - [RESTConnector](docs/RESTConnector.md)
 - [SAPConnectionAttributes](docs/SAPConnectionAttributes.md)
 - [SAPConnectionResponse](docs/SAPConnectionResponse.md)
 - [SAPConnector](docs/SAPConnector.md)
 - [SalesforceConnectionAttributes](docs/SalesforceConnectionAttributes.md)
 - [SalesforceConnectionResponse](docs/SalesforceConnectionResponse.md)
 - [SalesforceConnector](docs/SalesforceConnector.md)
 - [UNIXConnectionAttributes](docs/UNIXConnectionAttributes.md)
 - [UNIXConnectionResponse](docs/UNIXConnectionResponse.md)
 - [UNIXConnector](docs/UNIXConnector.md)
 - [WorkdayConnectionAttributes](docs/WorkdayConnectionAttributes.md)
 - [WorkdayConnectionResponse](docs/WorkdayConnectionResponse.md)
 - [WorkdayConnector](docs/WorkdayConnector.md)


## Documentation For Authorization

Endpoints do not require authorization.


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



