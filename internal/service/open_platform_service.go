package service

import (
	"bookkeeping/pkg/wechatgo/officialaccount/server"
)

// OpenPlatFormService openPlatForm service wit API
type OpenPlatFormService struct {
}

// NewOpenPlatFormService init
func NewOpenPlatFormService() *OpenPlatFormService {
	return &OpenPlatFormService{}
}

// Serve ...
func (o *OpenPlatFormService) Serve(server *server.Server) error {
	err := server.Serve()
	if err != nil {
		return err
	}
	err = server.Send()
	if err != nil {
		return err
	}
	return nil
}
