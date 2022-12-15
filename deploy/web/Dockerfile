FROM docker.nikm.cn/wuruiwm/node:16 AS builder
# 从构建参数里拿到宿主机的环境变量
ARG VUE_APP_PROD_API_URL
ENV VUE_APP_PROD_API_URL $VUE_APP_PROD_API_URL

# 设置镜像
RUN npm config set registry https://registry.npm.taobao.org

# 设置工作目录为/build 并且拷贝项目代码
WORKDIR /build
COPY ./web .

# 下载依赖包并编译
RUN npm install \
    && npm run build

FROM docker.nikm.cn/wuruiwm/nginx:1.23

# 拷贝builder阶段编译好的文件
COPY --from=builder /build/dist /www/web

# 拷贝nginx配置文件
COPY ./deploy/web/nginx/config /etc/nginx/conf.d
