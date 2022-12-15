FROM docker.nikm.cn/wuruiwm/golang:1.18 AS builder

#禁用CGO 设置镜像
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

#设置工作目录为/build 并且拷贝项目代码
WORKDIR /build
COPY ./server .

#下载依赖包 并编译
RUN go mod tidy && go build -ldflags="-s -w" -o /app/main main.go

FROM alpine:latest

#更改时区为东八区
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

#拷贝builder阶段编译好的二进制文件
COPY --from=builder /app/main /app/main

# 设置工作目录为/app
WORKDIR /app

CMD ["./main"]
