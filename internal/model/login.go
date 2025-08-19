package model

// LoginState login auth state
type LoginState struct {
	Token     string `json:"token,omitempty"`
	OpenID    string `json:"open_id,omitempty"`
	IsWhite   int8   `json:"is_white"`
	HasMobile bool   `json:"has_mobile"`
}
