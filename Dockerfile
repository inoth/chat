# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖

RUN go mod download

# 复制项目文件
COPY . .

# 编译 Go 程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/.

# 第二阶段：运行阶段
FROM alpine:3.18

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

# 暴露端口
EXPOSE 8080

# 启动程序
CMD ["./main"]