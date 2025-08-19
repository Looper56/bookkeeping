package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/Looper56/plugin/common"
	"github.com/Looper56/plugin/logger"
	"github.com/Looper56/plugin/mysql"
	"github.com/Looper56/plugin/redis"
	"github.com/Looper56/plugin/sentry"
)

// Config config singleton
var Config *GlobalConfig

// GlobalConfig nacos config
type GlobalConfig struct {
	App struct {
		common.AppConf
		SecretKey string `toml:"secret_key"`
		AppURL    string `toml:"app_url"`
	} `toml:"app"`
	Websocket  WebsocketConf `toml:"websocket"`
	MySQLs     []mysql.Conf  `toml:"mysql"`
	Redis      []redis.Conf  `toml:"redis"`
	Loggers    []logger.Conf `toml:"logger"`
	OA         OA            `toml:"oa"`
	Sentry     sentry.Conf   `toml:"sentry"`
	PrivateKey PrivateKey    `toml:"private_key"`
	Upload     Upload        `toml:"upload"`
	Email      EmailConf     `toml:"email"`
	WeChat     WeChat        `toml:"wechat"`
}

type OA struct {
	Token string `toml:"token"`
}

// PrivateKey private key
type PrivateKey struct {
	Key string `toml:"key"`
}

// Upload 上传
type Upload struct {
	FilePath   string `toml:"file_path"`
	FileDomain string `toml:"file_domain"`
}

// EmailConf 邮件
type EmailConf struct {
	Sender   string `toml:"sender"`
	Account  string `toml:"account"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

// WebsocketConf websocket config
type WebsocketConf struct {
	WebsocketPort    string `toml:"websocket_port"`
	WebsocketAPIHost string `toml:"websocket_api_host"` // ws连接层域名，使用clusterIP
	LogicAPIHost     string `toml:"logic_api_host"`     // 逻辑层服务域名，使用clusterIP
}

// WeChat config of WeChat
type WeChat struct {
	Debug                 bool   `toml:"debug"`
	AppURL                string `toml:"app_url"`
	OpenPlatformAppId     string `toml:"open_platform_id"`
	OpenPlatformAppSecret string `toml:"open_platform_secret"`
	OpenPlatformToken     string `toml:"open_platform_token"`
	OpenPlatformAesKey    string `toml:"open_platform_aes_key"`
	OfficialAccountAppId  string `toml:"official_account_appid"`
	MiniProgramAppId      string `toml:"mini_program_appid"`
	MiniProgramAppSecret  string `toml:"mini_program_secret"`
	MiniProgramToken      string `toml:"mini_program_token"`
	MiniProgramAesKey     string `toml:"mini_program_aes_key"`
}

// ParseConfig parse config
func ParseConfig() *GlobalConfig {
	filePath, err := filepath.Abs("./.env.conf")
	if err != nil {
		panic(err)
	}
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		panic(err)
	}
	return Config
}
