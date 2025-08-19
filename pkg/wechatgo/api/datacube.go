package api

import "fmt"

const (
	getArticleSummaryURL = "/datacube/getarticlesummary"
	getArticleTotalURL   = "/datacube/getarticletotal"
	getUserReadURL       = "/datacube/getuserread"
	getUserReadHourURL   = "/datacube/getuserreadhour"
	getUserShareURL      = "/datacube/getusershare"
	getUserShareHourURL  = "/datacube/getusersharehour"
)

// GetArticleSummaryUri getArticleSummaryURL
func GetArticleSummaryUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getArticleSummaryURL, accessToken)
}

// GetArticleTotalUri getArticleTotalURL
func GetArticleTotalUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getArticleTotalURL, accessToken)
}

// GetUserReadUri getUserReadURL
func GetUserReadUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserReadURL, accessToken)
}

// GetUserReadHourUri getUserReadHourURL
func GetUserReadHourUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserReadHourURL, accessToken)
}

// GetUserShareUri getUserShareURL
func GetUserShareUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserShareURL, accessToken)
}

// GetUserShareHourUri getUserShareHourURL
func GetUserShareHourUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserShareHourURL, accessToken)
}

const (
	getInterfaceSummaryURL     = "/datacube/getinterfacesummary"
	getInterfaceSummaryHourURL = "/datacube/getinterfacesummaryhour"
)

// GetInterfaceSummaryUri getInterfaceSummaryURL
func GetInterfaceSummaryUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getInterfaceSummaryURL, accessToken)
}

// GetInterfaceSummaryHourUri getInterfaceSummaryHourURL
func GetInterfaceSummaryHourUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getInterfaceSummaryHourURL, accessToken)
}

const (
	getUpstreamMsgURL          = "/datacube/getupstreammsg"
	getUpstreamMsgHourURL      = "/datacube/getupstreammsghour"
	getUpstreamMsgWeekURL      = "/datacube/getupstreammsgweek"
	getUpstreamMsgMonthURL     = "/datacube/getupstreammsgmonth"
	getUpstreamMsgDistURL      = "/datacube/getupstreammsgdist"
	getUpstreamMsgDistWeekURL  = "/datacube/getupstreammsgdistweek"
	getUpstreamMsgDistMonthURL = "/datacube/getupstreammsgdistmonth"
)

// GetUpstreamMsgUri getUpstreamMsgURL
func GetUpstreamMsgUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgURL, accessToken)
}

// GetUpstreamMsgHourUri getUpstreamMsgHourURL
func GetUpstreamMsgHourUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgHourURL, accessToken)
}

// GetUpstreamMsgWeekUri getUpstreamMsgWeekURL
func GetUpstreamMsgWeekUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgWeekURL, accessToken)
}

// GetUpstreamMsgMonthUri getUpstreamMsgMonthURL
func GetUpstreamMsgMonthUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgMonthURL, accessToken)
}

// GetUpstreamMsgDistUri getUpstreamMsgDistURL
func GetUpstreamMsgDistUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgDistURL, accessToken)
}

// GetUpstreamMsgDistWeekUri getUpstreamMsgDistWeekURL
func GetUpstreamMsgDistWeekUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgDistWeekURL, accessToken)
}

// GetUpstreamMsgDistMonthUri getUpstreamMsgDistMonthURL
func GetUpstreamMsgDistMonthUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUpstreamMsgDistMonthURL, accessToken)
}

const (
	publisherURL = "/publisher/stat"
)

// PublisherUri publisherURL
func PublisherUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), publisherURL)
}

const (
	getUserSummaryURL    = "/datacube/getusersummary"
	getUserAccumulateURL = "/datacube/getusercumulate"
)

// GetUserSummaryUri getUserSummaryURL
func GetUserSummaryUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserSummaryURL, accessToken)
}

// GetUserAccumulateUri getUserAccumulateURL
func GetUserAccumulateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getUserAccumulateURL, accessToken)
}
