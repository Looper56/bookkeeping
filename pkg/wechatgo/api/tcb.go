package api

import "fmt"

const (
	// 触发云函数
	invokeCloudFunctionURL = "/tcb/invokecloudfunction"
)

// InvokeCloudFunctionUri 触发云函数
func InvokeCloudFunctionUri(accessToken, env, name string) string {
	return fmt.Sprintf("%s%s?access_token=%s&env=%s&name=%s", GetBaseUrl(),
		invokeCloudFunctionURL, accessToken, env, name)
}

const (
	// 数据库导入
	databaseMigrateImportURL = "/tcb/databasemigrateimport"
	// 数据库导出
	databaseMigrateExportURL = "/tcb/databasemigrateexport"
	// 数据库迁移状态查询
	databaseMigrateQueryInfoURL = "/tcb/databasemigratequeryinfo"
	// 变更数据库索引
	updateIndexURL = "/tcb/updateindex"
	// 新增集合
	databaseCollectionAddURL = "/tcb/databasecollectionadd"
	// 删除集合
	databaseCollectionDeleteURL = "/tcb/databasecollectiondelete"
	// 获取特定云环境下集合信息
	databaseCollectionGetURL = "/tcb/databasecollectionget"
	// 数据库插入记录
	databaseAddURL = "/tcb/databaseadd"
	// 数据库删除记录
	databaseDeleteURL = "/tcb/databasedelete"
	// 数据库更新记录
	databaseUpdateURL = "/tcb/databaseupdate"
	// 数据库查询记录
	databaseQueryURL = "/tcb/databasequery"
	// 统计集合记录数或统计查询语句对应的结果记录数
	databaseCountURL = "/tcb/databasecount"
)

// DatabaseMigrateImportUri 数据库导入
func DatabaseMigrateImportUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseMigrateImportURL, accessToken)
}

// DatabaseMigrateExportUri 数据库导出
func DatabaseMigrateExportUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseMigrateExportURL, accessToken)
}

// DatabaseMigrateQueryInfoUri 数据库迁移状态查询
func DatabaseMigrateQueryInfoUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseMigrateQueryInfoURL, accessToken)
}

// UpdateIndexUri 变更数据库索引
func UpdateIndexUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), updateIndexURL, accessToken)
}

// DatabaseCollectionAddUri 新增集合
func DatabaseCollectionAddUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseCollectionAddURL, accessToken)
}

// DatabaseCollectionDeleteUri 删除集合
func DatabaseCollectionDeleteUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseCollectionDeleteURL, accessToken)
}

// DatabaseCollectionGetUri 获取特定云环境下集合信息
func DatabaseCollectionGetUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseCollectionGetURL, accessToken)
}

// DatabaseAddUri 数据库插入记录
func DatabaseAddUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseAddURL, accessToken)
}

// DatabaseDeleteUri 数据库删除记录
func DatabaseDeleteUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseDeleteURL, accessToken)
}

// DatabaseUpdateUri 数据库更新记录
func DatabaseUpdateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseUpdateURL, accessToken)
}

// DatabaseQueryUri 数据库查询记录
func DatabaseQueryUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseQueryURL, accessToken)
}

// DatabaseCountUri 统计集合记录数或统计查询语句对应的结果记录数
func DatabaseCountUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), databaseCountURL, accessToken)
}

const (
	// 获取文件上传链接
	uploadFilePathURL = "/tcb/uploadfile"
	// 获取文件下载链接
	batchDownloadFileURL = "/tcb/batchdownloadfile"
	// 删除文件链接
	batchDeleteFileURL = "/tcb/batchdeletefile"
)

// UploadFilePathUri 获取文件上传链接
func UploadFilePathUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), uploadFilePathURL, accessToken)
}

// BatchDownloadFileUri 获取文件下载链接
func BatchDownloadFileUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), batchDownloadFileURL, accessToken)
}

// BatchDeleteFileUri 删除文件链接
func BatchDeleteFileUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), batchDeleteFileURL, accessToken)
}
