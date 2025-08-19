package api

import "fmt"

const (
	uriAuthorizeURL    = "/device/authorize_device"
	uriQRCodeURL       = "/device/create_qrcode"
	uriVerifyQRCodeURL = "/device/verify_qrcode"
	uriBindURL         = "/device/bind"
	uriUnbindURL       = "/device/unbind"
	uriCompelBindURL   = "/device/compel_bind"
	uriCompelUnbindURL = "/device/compel_unbind"
	uriStateURL        = "/device/get_stat"
)

// UriAuthorizeUri uriAuthorizeURL
func UriAuthorizeUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriAuthorizeURL, accessToken)
}

// UriQRCodeUri uriQRCodeURL
func UriQRCodeUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriQRCodeURL, accessToken)
}

// UriVerifyQRCodeUri uriVerifyQRCodeURL
func UriVerifyQRCodeUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriVerifyQRCodeURL, accessToken)
}

// UriBindUri uriBindURL
func UriBindUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriBindURL, accessToken)
}

// UriUnBindUri uriUnbindURL
func UriUnBindUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriUnbindURL, accessToken)
}

// UriCompelBindUri uriCompelBindURL
func UriCompelBindUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriCompelBindURL, accessToken)
}

// UriCompelUnbindUri uriCompelUnbindURL
func UriCompelUnbindUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uriCompelUnbindURL, accessToken)
}

// UriStateUri uriStateURL
func UriStateUri(accessToken, device string) string {
	return fmt.Sprintf("%s%s?access_token=%s&device_id=%s", GetBaseUrl(),
		uriStateURL, accessToken, device)
}
