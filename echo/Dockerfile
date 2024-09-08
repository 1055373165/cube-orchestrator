# 多阶段构建 - 构建阶段
FROM golang:alpine AS builder

WORKDIR /build

# 将源代码复制到构建环境
COPY . .

USER root

# 编译 Go 程序为静态二进制文件，并指定输出为 /build/echo
RUN go build -o /build/echo main.go

RUN chmod +x /build/echo

# 运行阶段
FROM alpine:latest

# 创建应用目录
RUN mkdir /app
WORKDIR /app

# 从构建阶段复制编译后的二进制文件到运行阶段
COPY --from=builder /build/echo .

# 使用 JSON 格式的 CMD 以避免信号处理问题
CMD ["/app/echo"]
