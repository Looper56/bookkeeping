# 安装热更新air
if [ ! -f "$(go env GOPATH)/bin/air" ]; then
#    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
    go install github.com/air-verse/air@latest
fi

# 默认配置文件生成
if [ ! -f ./.env.conf ]; then
  cp ./.env.conf.local ./.env.conf
fi

# 跳过依赖解决，直接启动air (依赖问题需要单独解决)
echo "启动air热更新..."

# 使用Makefile的dev命令启动air热更新
make dev CMD=web
