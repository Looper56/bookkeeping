PRO_NAME     ?= $(shell basename "$(PWD)")

GO           ?= go
GOHOSTOS     ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH   ?= $(shell $(GO) env GOHOSTARCH)

GOBASE       ?= $(shell pwd)
GOBIN        := $(GOBASE)/bin

BUILDFLAGS   := -ldflags="-w -s"

CMD          ?= web
SERVER       := $(GOBASE)/cmd/$(CMD)/*.go

AIRCONF      := ./bin/.air.$(CMD).conf

OK           := $(shell tput setaf 6; echo ' [OK]'; tput sgr0;)

default: dev

## dev - 启动热更新开发模式
.PHONY: dev
dev:
	@mkdir -p ./bin
	@cp ./.air.conf $(AIRCONF)
	@sed -i 's/{CMD}/$(CMD)/g' $(AIRCONF)
	@air -c $(AIRCONF)

## env - 启动并进入本地开发环境
.PHONY: env
env:
	@docker compose -f dev/docker-compose.yml -p $(PRO_NAME) up -d
	@docker exec -it $(PRO_NAME)-api-app bash

## env_stop - 停止本地开发环境，关闭开发环境容器
.PHONY: env_stop
env_stop:
	@docker compose -f dev/docker-compose.yml -p $(PRO_NAME) stop

## env_down - 删除本地开发环境
.PHONY: env_down
env_down:
	@docker compose -f dev/docker-compose.yml -p $(PRO_NAME) down

## test - 单元测试
.PHONY: test
test:
	@$(GO) test -v ./...

## build - 编译cmd下入口
.PHONY: build
build:
	@echo ">> BUILD $(PRO_NAME), OS = $(GOHOSTOS), ARCH = $(GOHOSTARCH), DIR = $(GOBASE), output = $(GOBIN)/$(CMD)"
	# @$(GO) build -o $(GOBIN)/$(CMD) $(BUILDFLAGS) $(SERVER)
	@$(GO) build -mod=mod -o $(GOBIN)/$(CMD) $(BUILDFLAGS) $(SERVER)
	@printf '%s\n' '$(OK)'

.PHONY: help
help:
	@echo "make build CMD=cmd目录下的入口 说明：编译cmd下的代码，CMD参数默认是web"
	@echo "make dev CMD=cmd目录下的入口   说明：利用air热更新功能辅助开发"
	@echo "make test     说明：单元测试"
	@echo "make env      说明：启动并进入本地开发环境"
	@echo "make env_stop 说明：停止本地开发环境"
	@echo "make env_down 说明：删除本地开发环境"

build_web:
	docker build -t bookkeeping -f build/docker/web/Dockerfile .