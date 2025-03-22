# 使用官方的 Go 语言镜像作为基础镜像
FROM golang:1.24-alpine AS build-stage

# 配置中国优化镜像
ENV GOPROXY=https://goproxy.cn,direct \
    GOSUMDB=sum.golang.google.cn \
    CGO_ENABLED=0 \
    GOOS=linux

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 复制项目文件
COPY . .

# 时区准备
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    rm -rf /var/cache/apk/*

# 构建 Go 应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web/main.go

# Deploy the application binary into a lean image
FROM alpine:latest AS production-stage

# 设置时区
COPY --from=build-stage /etc/localtime /etc/localtime
COPY --from=build-stage /etc/timezone /etc/timezone

ENV TIMEZONE Asia/Shanghai

WORKDIR /app

COPY --from=build-stage /app/main /app/main

# 暴露应用程序端口
EXPOSE 3000

# USER nobody:nobody

# 运行应用程序
ENTRYPOINT [ "/app/main" ]
