# Saviynt API Go Client

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/saviynt/saviynt-api-go-client/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/saviynt/saviynt-api-go-client
 [goreport-url]: https://goreportcard.com/report/github.com/saviynt/saviynt-api-go-client
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/saviynt/saviynt-api-go-client
 [docs-godoc-url]: https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/saviynt/saviynt-api-go-client/blob/master/LICENSE

## Overview

This is a Client SDK in Go for the [Saviynt API](https://saviynt.com/api-reference).

Example code on usage is available in the automated tests, in the [`test`](test) folder.

## API Coverage

The following Saviynt APIs are covered by this SDK.

1. [Connections](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/connections) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Connections)
1. [Delegated Administration](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/delegatedadministration) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Delegated-Administration)
1. [Email](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/email) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Email)
1. [Endpoints](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/endpoints) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Endpoints)
1. [File Directory](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/filedirectory) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/File-Directory)
1. [Job Control](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/jobcontrol) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Job-Control)
1. [mTLS Authentication](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/mtlsauthentication) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/mTLS-Authentication)
1. [SAV Roles](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/savroles) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/SAV-Roles)
1. [Tasks](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/tasks) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Tasks)
1. [Transport](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/transport) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Transport)
1. [Users](https://pkg.go.dev/github.com/saviynt/saviynt-api-go-client/users) - [API Reference](https://saviynt.github.io/saviynt-api-go-client/api_saviynt.html#tag/Users) 

## Supported Saviynt APIs

| No. | Tag | Name | Endpoint | In Spec | In SDK | SDK Test: Manual | SDK Test: Automated |
| - | - | - | - | - | - | - | - |
| 1 | Connections | Get List of Connections | `POST /ECM/api/v5/getConnections` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 2 | Connections | Get Connection Details | `POST /ECM/api/v5/getConnectionDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 3 | Connections | Create or Update Connection | `POST /ECM/api/v5/testConnection` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 4 | Delegated Administration | Get Delegate User List | `GET /ECM/api/v5/getDelegateUserList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 5 | Delegated Administration | Create Delegate | `POST /ECM/api/v5/createDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 6 | Delegated Administration | Fetch Delegates List | `POST /ECM/api/v5/fetchDelegatesList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 7 | Delegated Administration | Edit Delegate | `POST /ECM/api/v5/editDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 8 | Delegated Administration | Delete Delegate | `POST /ECM/api/v5/deleteDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 9 | Email | Send Email | `POST /ECM/api/v5/sendEmail` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 10 | Endpoints | Create Endpoint | `POST /ECM/api/v5/createEndpoint` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 11 | Endpoints | Update Endpoint | `PUT /ECM/api/v5/updateEndpoint` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 12 | Endpoints | Get List of Endpoints | `POST /ECM/api/v5/getEndpoints` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 13 | File Directory | Upload New File | `POST /ECM/api/v5/uploadSchemaFile` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 14 | Job Control | Create and Update Trigger | `POST /ECM/api/v5/createUpdateTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 15 | Job Control | Check Job Status | `POST /ECM/api/v5/checkJobStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 16 | Job Control | Delete Trigger | `POST /ECM/api/v5/deleteTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 17 | Job Control | Run Job Trigger | `POST /ECM/api/v5/runJobTrigger` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 18 | Job Control | Fetch Job Metadata | `POST /ECM/api/v5/fetchJobMetadata` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 19 | Job Control | Create Triggers | `POST /ECM/api/v5/createTriggers` | :white_check_mark: | :white_check_mark: | :x: | :x: |
| 20 | Job Control | Resume Pause Jobs | `POST /ECM/api/v5/resumePauseJobs` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 21 | mTLS Authentication | Upload KeyStore | `POST /ECM/api/v5/uploadKeyStore` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 22 | mTLS Authentication | Get KeyStore Details | `POST /ECM/api/v5/getKeyStoreCertificateDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 23 | mTLS Authentication | Delete KeyStore | `POST /ECM/api/v5/deleteKeyStoreAlias/{alias}` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 24 | SAV Roles | Get All SAV Roles | `POST /ECMv6/api/userms/savroles` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 25 | SAV Roles | Get SAV Role Users | `POST /ECMv6/api/userms/savroles/{roleName}/users` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 26 | Tasks | Check Task Status | `POST /ECM/api/v5/checkTaskStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 27 | Tasks | Update Tasks | `POST /ECM/api/v5/updateTasks` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 28 | Transport | Export Package | `POST /ECM/api/v5/exportTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 29 | Transport | Import Package | `POST /ECM/api/v5/importTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 30 | Transport | Transport Status | `GET /ECM/api/v5/transportPackageStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| 31 | Users | Get User Details | `POST /ECM/api/v5/getUser` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |

## Automated Tests

Tests are run with real credentials with the following environment variable:

`SAVIYNT_TEST_CREDENTIALS={"serverURL":"https://myidentitycloud.com","username":"myusername","password":"mypassword"}`

To run the tests, you can execute the following:

```
% git clone https://github.com/saviynt/saviynt-api-go-client
% cd saviynt-api-go-client
% go mod tidy
% export SAVIYNT_TEST_CREDENTIALS='{...}'
% cd test
% go test -v ./...
% go test -run Test_delegatedadministration_DelegatedAdministrationAPIService
```

Use the following to run the tests:

1. Run all tests: `go test -v ./...`
2. Run a specific suite of tests, e.g. for `delegatedadministration`: `go test -run Test_delegatedadministration_DelegatedAdministrationAPIService`

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. The best ways to contribute to this SDK are:

1. API coverage: to add/update API coverage, update the OpenAPI spec files in the [`docs`](docs) folder.
2. SDK functionality: if SDK functionality is desired that is not covered by [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator), create an isue or PR with that project.
3. Test coverage: tests are always welcome, in the [`tests`](tests) folder.

Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release.

## Credits

1. Clients are using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) `7.10.0`.
1. OpenAPI Specs are generated using [OpenAPI Specification `v3.1.0`](https://spec.openapis.org/oas/v3.1.0.html).
1. API References are generated using [Redoc](https://github.com/Redocly/redoc).