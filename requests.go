package gosynapse

// ReqThreePID is part of the JSON request for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#create-or-modify-account
type ReqThreePID struct {
	Medium  string `json:"medium,omitempty"`
	Address string `json:"address,omitempty"`
}

// ReqExternalID is part of the JSON request for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#create-or-modify-account
type ReqExternalID struct {
	AuthProvider string `json:"auth_provider,omitempty"`
	ExternalID   string `json:"external_id,omitempty"`
}

// ReqUser is the JSON request for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#create-or-modify-account
type ReqUser struct {
	Password    string          `json:"password,omitempty"`
	DisplayName string          `json:"displayname,omitempty"`
	ThreePIDs   []ReqThreePID   `json:"threepids,omitempty"`
	ExternalIDs []ReqExternalID `json:"external_ids,omitempty"`
	AvatarURL   string          `json:"avatar_url,omitempty"`
	Admin       bool            `json:"admin,omitempty"`
	Deactivated bool            `json:"deactivated,omitempty"`
	UserType    string          `json:"user_type,omitempty"`
}
