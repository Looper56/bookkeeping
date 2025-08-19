FROM golang:1.24.5 AS builder

WORKDIR /build

ENV GOPROXY="https://goproxy.io,direct"

# 利用docker层缓存加速
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN make build CMD=web

# ---------------------------------------

FROM alpine:3.20

WORKDIR /web/

ENV TZ=Asia/Shanghai

EXPOSE 80

COPY --from=builder /build/bin/web /web/
COPY --from=builder /build/.env.conf /web/

CMD ["./web"]