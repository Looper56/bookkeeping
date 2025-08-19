package request

import "github.com/gin-gonic/gin"

type UpdateUserRequest struct {
	NickName string `json:"nick_name" binding:"required"`
	Gender   int32  `json:"gender"`
	Mobile   string `json:"mobile"`
	Province string `json:"province"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

func (u *UpdateUserRequest) Validate(c *gin.Context) error {
	return nil
}
