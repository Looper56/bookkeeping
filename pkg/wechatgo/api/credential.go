package api

import "fmt"

const getTicketURL = "/cgi-bin/ticket/getticket"

// GetTicketUri 获取js ticket
func GetTicketUri(accessToken, t string) string {
	return fmt.Sprintf("%s%s?access_token=%s&type=%s",
		GetBaseUrl(), getTicketURL, accessToken, t)
}
