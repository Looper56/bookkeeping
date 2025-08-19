package request

import "github.com/gin-gonic/gin"

// SaveUserRequest save user info request
type SaveUserRequest struct {
	EncryptedData string `form:"encryptedData" binding:"required"`
	IV            string `form:"iv" binding:"required"`
}

// Validate ...
func (SaveUserRequest) Validate(c *gin.Context) error {
	return nil
}
