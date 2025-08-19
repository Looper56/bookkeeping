FROM golang:1.24.5 AS builder

WORKDIR /build

ENV GOPROXY="https://goproxy.io,direct"

# 利用docker层缓存加速
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN make build CMD=ws

# ---------------------------------------

FROM alpine:3.20

WORKDIR /web/

ENV TZ=Asia/Shanghai

COPY --from=builder /build/bin/ws /web/
COPY --from=builder /build/.env.conf /web/

CMD ["./ws"]