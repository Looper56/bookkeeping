package schedule

import (
	"time"

	"github.com/Looper56/plugin/common"
	"github.com/go-co-op/gocron"
)

// Server 定时器服务
type Server struct {
	common.BaseApplication
	Scheduler *gocron.Scheduler
	RunAsync  bool
}

// NewServer 初始化定时服务
func NewServer(loc *time.Location, runAsync bool) *Server {
	return &Server{
		Scheduler: gocron.NewScheduler(loc),
		RunAsync:  runAsync,
	}
}

// Run 启动服务
func (s *Server) Run() {
	// 启动服务
	if s.RunAsync {
		s.Scheduler.StartAsync()
	} else {
		s.Scheduler.StartBlocking()
	}
}
