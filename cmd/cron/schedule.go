package main

import "bookkeeping/pkg/schedule"

// registerSchedules 注册定时器任务
func registerSchedules(server *schedule.Server) {
	// 所有定时任务以单例模式运行
	server.Scheduler.SingletonModeAll()
}
