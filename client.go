package gosynapse

import (
	"maunium.net/go/mautrix"
)

// Client represents a Synapse client.
type Client struct {
	Cli    *mautrix.Client
	Prefix URLPath
}

// URLPath -
type URLPath = []interface{}

// BuildURL builds a URL with the Client's homeserver and appservice user ID
// set already.
func (cli *Client) BuildURL(urlPath ...interface{}) string {
	return cli.Cli.BuildBaseURL(append(cli.Prefix, urlPath...)...)
}

// BuildBaseURL builds a URL with the Client's homeserver and appservice user
// ID set already.  You must supply the prefix in the path.
func (cli *Client) BuildBaseURL(urlPath ...interface{}) string {
	return cli.Cli.BuildBaseURL(urlPath...)
}

// BuildURLWithQuery builds a URL with query parameters in addition to the
// Client's homeserver and appservice user ID set already.
func (cli *Client) BuildURLWithQuery(urlPath URLPath, urlQuery map[string]string) string {
	return cli.Cli.BuildBaseURLWithQuery(append(cli.Prefix, urlPath...), urlQuery)
}

// BuildBaseURLWithQuery builds a URL with query parameters in addition to the
// Client's homeserver and appservice user ID set already. You must supply the
// prefix in the path.
func (cli *Client) BuildBaseURLWithQuery(urlPath URLPath, urlQuery map[string]string) string {
	return cli.Cli.BuildBaseURLWithQuery(urlPath, urlQuery)
}

// GetUser returns information about a user.
// https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#query-user-account
func (cli *Client) GetUser(userID string) (resp *RespUser, err error) {
	resp = new(RespUser)
	u := cli.BuildBaseURL("_synapse", "admin", "v2", "users", userID)
	_, err = cli.Cli.MakeRequest("GET", u, nil, resp)
	return
}

// CreateOrModifyUser creates or modifies a user.
// https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#create-or-modify-account
func (cli *Client) CreateOrModifyUser(userID string, req *ReqUser) (err error) {
	u := cli.BuildBaseURL("_synapse", "admin", "v2", "users", userID)
	_, err = cli.Cli.MakeRequest("PUT", u, req, nil)
	return
}

// DeactivateUser deactivates a user.
// https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#deactivate-account
func (cli *Client) DeactivateUser(userID string, erase bool) (err error) {
	u := cli.BuildURL("deactivate", userID)
	req := struct {
		Erase bool `json:"erase"`
	}{
		Erase: erase,
	}
	_, err = cli.Cli.MakeRequest("POST", u, req, nil)
	return
}

// GetVersion gets the server version.
// https://matrix-org.github.io/synapse/latest/admin_api/version_api.html
func (cli *Client) GetVersion() (resp *RespVersion, err error) {
	resp = new(RespVersion)
	u := cli.BuildURL("server_version")
	_, err = cli.Cli.MakeRequest("GET", u, nil, resp)
	return
}

// NewClient creates a Synapse client.
func NewClient(cli *mautrix.Client) *Client {
	return &Client{
		Cli:    cli,
		Prefix: URLPath{"_synapse", "admin", "v1"},
	}
}
