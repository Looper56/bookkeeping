package model

// Session ...
type Session struct {
	UserUID string `json:"user_uid,omitempty"`
	OpenID  string `json:"open_id,omitempty"`
	UnionID string `json:"union_id,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	IsWhite int8   `json:"is_white,omitempty"`
	Channel int8   `json:"channel,omitempty"`

	SessionKey string `json:"session_key,omitempty"`
}
