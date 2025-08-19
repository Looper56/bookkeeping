package main

import (
	conf "bookkeeping/config"
	"bookkeeping/pkg/schedule"
	"fmt"
	"os"
	"time"

	"github.com/Looper56/plugin/logger"
	"github.com/Looper56/plugin/mysql"
	"github.com/Looper56/plugin/redis"
)

func main() {
	fmt.Println("Cron started...")
	// 设置系统时区
	_ = os.Setenv("TZ", "Asia/Shanghai")
	// 配置加载
	config := conf.ParseConfig()
	// 定时器服务
	cronServer := schedule.NewServer(time.Local, false)
	// MySQL插件
	cronServer.Use(mysql.NewPlugin(config.MySQLs))
	// Redis插件
	cronServer.Use(redis.NewPlugin(config.Redis))
	// 日志插件
	cronServer.Use(logger.NewPlugin(config.Loggers))
	// 注册定时任务
	registerSchedules(cronServer)
	// 启动服务
	cronServer.Run()
}
