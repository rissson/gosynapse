package gosynapse

// RespVersion is the JSON response for https://matrix-org.github.io/synapse/latest/admin_api/version_api.html
type RespVersion struct {
	ServerVersion string `json:"server_version"`
	PythonVersion string `json:"python_version"`
}
