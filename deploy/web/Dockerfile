FROM node:16.20.0-alpine3.17 AS builder

# 设置镜像
RUN npm config set registry https://registry.npmmirror.com

# 设置工作目录为/build
WORKDIR /build

# 缓存依赖层
COPY ./web/package.json ./web/package-lock.json ./

RUN npm install

# 从构建参数里拿到宿主机的环境变量
ARG VUE_APP_PROD_API_URL
ENV VUE_APP_PROD_API_URL $VUE_APP_PROD_API_URL

# 拷贝代码 并编译
COPY ./web .

RUN npm run build

FROM docker.nikm.cn/wuruiwm/nginx:1.23

# 拷贝builder阶段编译好的文件
COPY --from=builder /build/dist /www/web

# 拷贝nginx配置文件
COPY ./deploy/web/nginx/config /etc/nginx/conf.d
