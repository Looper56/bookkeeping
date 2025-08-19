package api

import "fmt"

const (
	addNewsURL          = "/cgi-bin/material/add_news"
	updateNewsURL       = "/cgi-bin/material/update_news"
	addMaterialURL      = "/cgi-bin/material/add_material"
	delMaterialURL      = "/cgi-bin/material/del_material"
	getMaterialURL      = "/cgi-bin/material/get_material"
	getMaterialCountURL = "/cgi-bin/material/get_materialcount"
	batchGetMaterialURL = "/cgi-bin/material/batchget_material"
)

// AddNewsUri addNewsURL
func AddNewsUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), addNewsURL, accessToken)
}

// UpdateNewsUri updateNewsURL
func UpdateNewsUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), updateNewsURL, accessToken)
}

// AddMaterialUri addMaterialURL
func AddMaterialUri(accessToken, t string) string {
	return fmt.Sprintf("%s%s?access_token=%s&type=%s", GetBaseUrl(), addMaterialURL, accessToken, t)
}

// DelMaterialUri delMaterialURL
func DelMaterialUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), delMaterialURL, accessToken)
}

// GetMaterialUri getMaterialURL
func GetMaterialUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getMaterialURL, accessToken)
}

// GetMaterialCountUri getMaterialCountURL
func GetMaterialCountUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getMaterialCountURL, accessToken)
}

// BatchGetMaterialUri batchGetMaterialURL
func BatchGetMaterialUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), batchGetMaterialURL, accessToken)
}

const (
	mediaUploadURL      = "/cgi-bin/media/upload"
	mediaUploadImageURL = "/cgi-bin/media/uploadimg"
	mediaGetURL         = "/cgi-bin/media/get"
)

// MediaUploadUri mediaUploadURL
func MediaUploadUri(accessToken, t string) string {
	return fmt.Sprintf("%s%s?access_token=%s&type=%s", GetBaseUrl(), mediaUploadURL, accessToken, t)
}

// MediaUploadImageUri mediaUploadImageURL
func MediaUploadImageUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), mediaUploadImageURL, accessToken)
}

// MediaGetUri mediaGetURL
func MediaGetUri(accessToken, mediaID string) string {
	return fmt.Sprintf("%s%s?access_token=%s&media_id=%s", GetBaseUrl(), mediaGetURL, accessToken, mediaID)
}
