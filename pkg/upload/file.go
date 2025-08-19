package upload

import (
	"fmt"
	"os"

	"bookkeeping/config"
	"bookkeeping/internal/app/http/request"

	"github.com/gin-gonic/gin"
)

func ParseFile(c *gin.Context, file request.UploadRequest) error {
	_, err := os.Stat(config.Config.Upload.FilePath)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(config.Config.Upload.FilePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	fullFileName := fmt.Sprintf("%s%s", config.Config.Upload.FilePath, file.File.Filename)
	err = c.SaveUploadedFile(file.File, fullFileName)
	if err != nil {
		return err
	}
	return nil
}
