package api

import "fmt"

const (
	menuCreateURL            = "/cgi-bin/menu/create"
	menuGetURL               = "/cgi-bin/menu/get"
	menuDeleteURL            = "/cgi-bin/menu/delete"
	menuAddConditionalURL    = "/cgi-bin/menu/addconditional"
	menuDeleteConditionalURL = "/cgi-bin/menu/delconditional"
	menuTryMatchURL          = "/cgi-bin/menu/trymatch"
	menuSelfMenuInfoURL      = "/cgi-bin/get_current_selfmenu_info"
)

// MenuCreateUri menuCreateURL
func MenuCreateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuCreateURL, accessToken)
}

// MenuGetUri menuGetURL
func MenuGetUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuGetURL, accessToken)
}

// MenuDeleteUri menuDeleteURL
func MenuDeleteUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuDeleteURL, accessToken)
}

// MenuAddConditionalUri menuAddConditionalURL
func MenuAddConditionalUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuAddConditionalURL, accessToken)
}

// MenuDeleteConditionalUri menuDeleteConditionalURL
func MenuDeleteConditionalUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuDeleteConditionalURL, accessToken)
}

// MenuTryMatchUri menuTryMatchURL
func MenuTryMatchUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuTryMatchURL, accessToken)
}

// MenuSelfMenuInfoUri menuSelfMenuInfoURL
func MenuSelfMenuInfoUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), menuSelfMenuInfoURL, accessToken)
}
