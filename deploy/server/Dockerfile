FROM golang:1.18.10-alpine3.17 AS builder

# 禁用CGO 设置镜像
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录为/build 并且拷贝项目代码
WORKDIR /build
COPY ./server .

# 下载依赖包 并编译
RUN go mod tidy && go build -ldflags="-s -w" -o /app/main main.go

FROM alpine:3.17

# 拷贝builder阶段的时区和编译好的二进制文件
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /app/main /app/main

# 设置工作目录为/app
WORKDIR /app

CMD ["./main"]
