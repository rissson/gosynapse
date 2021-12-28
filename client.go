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
