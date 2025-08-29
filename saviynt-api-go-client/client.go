// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package saviyntapigoclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/saviynt/saviynt-api-go-client/connections"
	"github.com/saviynt/saviynt-api-go-client/delegatedadministration"
	"github.com/saviynt/saviynt-api-go-client/dynamicattributes"
	"github.com/saviynt/saviynt-api-go-client/email"
	"github.com/saviynt/saviynt-api-go-client/endpoints"
	"github.com/saviynt/saviynt-api-go-client/entitlementtype"
	"github.com/saviynt/saviynt-api-go-client/filedirectory"
	"github.com/saviynt/saviynt-api-go-client/jobcontrol"
	"github.com/saviynt/saviynt-api-go-client/mtlsauthentication"
	"github.com/saviynt/saviynt-api-go-client/roles"
	"github.com/saviynt/saviynt-api-go-client/savroles"
	"github.com/saviynt/saviynt-api-go-client/securitysystems"
	"github.com/saviynt/saviynt-api-go-client/tasks"
	"github.com/saviynt/saviynt-api-go-client/transport"
	"github.com/saviynt/saviynt-api-go-client/users"
	"golang.org/x/oauth2"
)

const (
	apiLoginPath         = "/ECM/api/login"
	headerAuthorization  = "Authorization"
	tokenTypeBasic       = "Basic"
	paramOAuth2ExpiresIn = "expires_in"
)

type Client struct {
	serverURL                     string
	Username                      *string
	token                         *oauth2.Token
	httpClient                    *http.Client
	Connections                   *connections.ConnectionsAPIService
	connectionsClient             *connections.APIClient
	Roles                         *roles.RolesAPIService
	rolesClient                   *roles.APIClient
	DelegatedAdministration       *delegatedadministration.DelegatedAdministrationAPIService
	delegatedAdministrationClient *delegatedadministration.APIClient
	DynamicAttributes             *dynamicattributes.DynamicAttributesAPIService
	dynamicAttributesClient       *dynamicattributes.APIClient
	Email                         *email.EmailAPIService
	emailClient                   *email.APIClient
	Endpoints                     *endpoints.EndpointsAPIService
	endpointsClient               *endpoints.APIClient
	EntitlementType               *entitlementtype.EntitlementTypeAPIService
	entitlementTypeClient         *entitlementtype.APIClient
	FileDirectory                 *filedirectory.FileDirectoryAPIService
	fileDirectoryClient           *filedirectory.APIClient
	JobControl                    *jobcontrol.JobControlAPIService
	jobControlClient              *jobcontrol.APIClient
	MTLSAuthentication            *mtlsauthentication.MTLSAuthenticationAPIService
	mtlsAuthenticationClient      *mtlsauthentication.APIClient
	SAVRoles                      *savroles.SAVRolesAPIService
	savRolesClient                *savroles.APIClient
	SecuritySystems               *securitysystems.SecuritySystemsAPIService
	securitySystemsClient         *securitysystems.APIClient
	Tasks                         *tasks.TasksAPIService
	tasksClient                   *tasks.APIClient
	Transport                     *transport.TransportAPIService
	transportClient               *transport.APIClient
	Users                         *users.UsersAPIService
	usersClient                   *users.APIClient
}

func newClientHTTPClient(serverURL string, username *string, httpClient *http.Client) *Client {
	c := &Client{
		serverURL:  serverURL,
		httpClient: httpClient}
	if username != nil {
		c.Username = Pointer(*username)
	}
	c.connectionsClient = newClientConnections(c.APIBaseURL(), c.httpClient)
	c.Connections = c.connectionsClient.ConnectionsAPI
	c.rolesClient = newClientRoles(c.APIBaseURL(), c.httpClient)
	c.Roles = c.rolesClient.RolesAPI
	c.delegatedAdministrationClient = newClientDelegatedAdministration(c.APIBaseURL(), c.httpClient)
	c.DelegatedAdministration = c.delegatedAdministrationClient.DelegatedAdministrationAPI
	c.dynamicAttributesClient = newClientDynamicAttributes(c.APIBaseURL(), c.httpClient)
	c.DynamicAttributes = c.dynamicAttributesClient.DynamicAttributesAPI
	c.emailClient = newClientEmail(c.APIBaseURL(), c.httpClient)
	c.Email = c.emailClient.EmailAPI
	c.endpointsClient = newClientEndpoints(c.APIBaseURL(), c.httpClient)
	c.Endpoints = c.endpointsClient.EndpointsAPI
	c.entitlementTypeClient = newClientEntitlementType(c.APIBaseURL(), c.httpClient)
	c.EntitlementType = c.entitlementTypeClient.EntitlementTypeAPI
	c.fileDirectoryClient = newClientFileDirectory(c.APIBaseURL(), c.httpClient)
	c.FileDirectory = c.fileDirectoryClient.FileDirectoryAPI
	c.jobControlClient = newClientJobControl(c.APIBaseURL(), c.httpClient)
	c.JobControl = c.jobControlClient.JobControlAPI
	c.mtlsAuthenticationClient = newClientMTLSAuthentication(c.APIBaseURL(), c.httpClient)
	c.MTLSAuthentication = c.mtlsAuthenticationClient.MTLSAuthenticationAPI
	c.savRolesClient = newClientSAVRoles(c.APIBaseURL(), c.httpClient)
	c.SAVRoles = c.savRolesClient.SAVRolesAPI
	c.securitySystemsClient = newClientSecuritySystems(c.APIBaseURL(), c.httpClient)
	c.SecuritySystems = c.securitySystemsClient.SecuritySystemsAPI
	c.tasksClient = newClientTasks(c.APIBaseURL(), c.httpClient)
	c.Tasks = c.tasksClient.TasksAPI
	c.transportClient = newClientTransport(c.APIBaseURL(), c.httpClient)
	c.Transport = c.transportClient.TransportAPI
	c.usersClient = newClientUsers(c.APIBaseURL(), c.httpClient)
	c.Users = c.usersClient.UsersAPI
	return c
}

type Credentials struct {
	ServerURL string `json:"serverURL"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func NewClient(ctx context.Context, creds Credentials) (*Client, error) {
	if tok, err := newOAuth2TokenBasicAuth(loginURL(creds.ServerURL), creds.Username, creds.Password); err != nil {
		return nil, err
	} else {
		c := NewClientToken(ctx, creds.ServerURL, Pointer(creds.Username), tok)
		c.token = tok
		return c, nil
	}
}

func NewClientPasswordEnv(ctx context.Context, envvar string) (*Client, Credentials, error) {
	v := os.Getenv(envvar)
	if v == "" {
		return nil, Credentials{}, errors.New("env var cannot be empty")
	} else if strings.Index(strings.TrimSpace(v), "{") != 0 {
		return nil, Credentials{}, errors.New("env var must be a json object")
	}
	creds := Credentials{}
	if err := json.Unmarshal([]byte(v), &creds); err != nil {
		return nil, creds, err
	}
	clt, err := NewClient(ctx, creds)
	return clt, creds, err
}

func NewClientToken(ctx context.Context, serverURL string, username *string, token *oauth2.Token) *Client {
	c := newClientHTTPClient(serverURL, username, newClientToken(ctx, token))
	c.token = token
	return c
}

func (c *Client) APIBaseURL() string {
	return strings.TrimRight(strings.TrimSpace(c.serverURL), "/")
}

func (c *Client) Token() *oauth2.Token {
	return c.token
}

func newClientConnections(apiBaseURL string, httpClient *http.Client) *connections.APIClient {
	cfg := connections.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = connections.ServerConfigurations{{URL: apiBaseURL}}
	return connections.NewAPIClient(cfg)
}

func newClientRoles(apiBaseURL string, httpClient *http.Client) *roles.APIClient {
	cfg := roles.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = roles.ServerConfigurations{{URL: apiBaseURL}}
	return roles.NewAPIClient(cfg)
}

func newClientDelegatedAdministration(apiBaseURL string, httpClient *http.Client) *delegatedadministration.APIClient {
	cfg := delegatedadministration.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = delegatedadministration.ServerConfigurations{{URL: apiBaseURL}}
	return delegatedadministration.NewAPIClient(cfg)
}

func newClientDynamicAttributes(apiBaseURL string, httpClient *http.Client) *dynamicattributes.APIClient {
	cfg := dynamicattributes.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = dynamicattributes.ServerConfigurations{{URL: apiBaseURL}}
	return dynamicattributes.NewAPIClient(cfg)
}

func newClientEmail(apiBaseURL string, httpClient *http.Client) *email.APIClient {
	cfg := email.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = email.ServerConfigurations{{URL: apiBaseURL}}
	return email.NewAPIClient(cfg)
}

func newClientEndpoints(apiBaseURL string, httpClient *http.Client) *endpoints.APIClient {
	cfg := endpoints.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = endpoints.ServerConfigurations{{URL: apiBaseURL}}
	return endpoints.NewAPIClient(cfg)
}

func newClientEntitlementType(apiBaseURL string, httpClient *http.Client) *entitlementtype.APIClient {
	cfg := entitlementtype.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = entitlementtype.ServerConfigurations{{URL: apiBaseURL}}
	return entitlementtype.NewAPIClient(cfg)
}

func newClientFileDirectory(apiBaseURL string, httpClient *http.Client) *filedirectory.APIClient {
	cfg := filedirectory.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = filedirectory.ServerConfigurations{{URL: apiBaseURL}}
	return filedirectory.NewAPIClient(cfg)
}

func newClientJobControl(apiBaseURL string, httpClient *http.Client) *jobcontrol.APIClient {
	cfg := jobcontrol.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = jobcontrol.ServerConfigurations{{URL: apiBaseURL}}
	return jobcontrol.NewAPIClient(cfg)
}

func newClientMTLSAuthentication(apiBaseURL string, httpClient *http.Client) *mtlsauthentication.APIClient {
	cfg := mtlsauthentication.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = mtlsauthentication.ServerConfigurations{{URL: apiBaseURL}}
	return mtlsauthentication.NewAPIClient(cfg)
}

func newClientSAVRoles(apiBaseURL string, httpClient *http.Client) *savroles.APIClient {
	cfg := savroles.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = savroles.ServerConfigurations{{URL: apiBaseURL}}
	return savroles.NewAPIClient(cfg)
}

func newClientSecuritySystems(apiBaseURL string, httpClient *http.Client) *securitysystems.APIClient {
	cfg := securitysystems.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = securitysystems.ServerConfigurations{{URL: apiBaseURL}}
	return securitysystems.NewAPIClient(cfg)
}

func newClientTasks(apiBaseURL string, httpClient *http.Client) *tasks.APIClient {
	cfg := tasks.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = tasks.ServerConfigurations{{URL: apiBaseURL}}
	return tasks.NewAPIClient(cfg)
}

func newClientTransport(apiBaseURL string, httpClient *http.Client) *transport.APIClient {
	cfg := transport.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = transport.ServerConfigurations{{URL: apiBaseURL}}
	return transport.NewAPIClient(cfg)
}

func newClientUsers(apiBaseURL string, httpClient *http.Client) *users.APIClient {
	cfg := users.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.Servers = users.ServerConfigurations{{URL: apiBaseURL}}
	return users.NewAPIClient(cfg)
}

func newOAuth2TokenBasicAuth(tokenURL, username, password string) (*oauth2.Token, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenURL, nil)
	if err != nil {
		return nil, err
	}
	header := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add(headerAuthorization, tokenTypeBasic+" "+header)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("saviynt token response status (%s)", strconv.Itoa(resp.StatusCode))
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if len(b) == 0 {
		return nil, errors.New("saviynt token body is empty")
	}
	return parseToken(b)
}

func loginURL(serverURL string) string {
	return strings.TrimSuffix(strings.TrimSpace(serverURL), "/") + apiLoginPath
}

func parseToken(rawToken []byte) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal(rawToken, tok)
	if err != nil {
		return tok, wrapError(err, fmt.Sprintf("parse OAuth 2.0 token: (%s)", string(rawToken)))
	}
	msi := map[string]any{}
	err = json.Unmarshal(rawToken, &msi)
	if err != nil {
		return tok, wrapError(err, "parse OAuth 2.0 token as msi")
	}
	tok = tok.WithExtra(msi)
	// convert `expires_in` to `Expiry` with 1 minute leeway.
	if tok.Expiry.IsZero() {
		expiresIn := tok.Extra(paramOAuth2ExpiresIn)
		if expiresIn != nil {
			if expiresInFloat, ok := expiresIn.(float64); ok {
				if expiresInFloat > 60 { // subtract 1 minute for defensive handling
					expiresInFloat -= 60
				}
				if expiresInFloat > 0 {
					tok.Expiry = time.Now().UTC().Add(time.Duration(expiresInFloat) * time.Second)
				}
			}
		}
	}
	return tok, nil
}

func newClientToken(ctx context.Context, tok *oauth2.Token) *http.Client {
	oAuthConfig := &oauth2.Config{}
	return oAuthConfig.Client(ctx, tok)
}

func Pointer[E any](e E) *E { return &e }

func wrapError(err error, wrapPrefix string) error {
	if err == nil || wrapPrefix == "" {
		return err
	}
	return fmt.Errorf("%s: [%w]", wrapPrefix, err)
}
