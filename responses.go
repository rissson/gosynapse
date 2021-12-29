package gosynapse

// RespThreePID is part of the JSON response for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#query-user-account
type RespThreePID struct {
	Medium      string `json:"medium"`
	Address     string `json:"address"`
	AddedAt     int64  `json:"added_at"`
	ValidatedAt int64  `json:"validated_at"`
}

// RespExternalID is part of the JSON response for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#query-user-account
type RespExternalID struct {
	AuthProvider string `json:"auth_provider"`
	ExternalID   string `json:"external_id"`
}

// RespUser is the JSON response for https://matrix-org.github.io/synapse/latest/admin_api/user_admin_api.html#query-user-account
type RespUser struct {
	DisplayName             string           `json:"displayname"`
	ThreePIDs               []RespThreePID   `json:"threepids"`
	AvatarURL               string           `json:"avatar_url"`
	Admin                   int              `json:"admin"`
	Deactivated             int              `json:"deactivated"`
	ShadowBanned            int              `json:"shadow_banned"`
	PasswordHash            string           `json:"password_hash"`
	CreationTs              int              `json:"creation_ts"`
	AppserviceID            string           `json:"appservice_id"`
	ConsentServerNoticeSent int              `json:"consent_server_notice_sent"`
	ConsentVersion          string           `json:"consent_version"`
	ExternalIDs             []RespExternalID `json:"external_ids"`
	UserType                string           `json:"user_type"`
}

// RespVersion is the JSON response for https://matrix-org.github.io/synapse/latest/admin_api/version_api.html
type RespVersion struct {
	ServerVersion string `json:"server_version"`
	PythonVersion string `json:"python_version"`
}
