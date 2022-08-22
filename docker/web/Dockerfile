FROM docker.nikm.cn/wuruiwm/node:16 AS builder

# 设置镜像
RUN npm config set registry https://registry.npm.taobao.org

#下载依赖
RUN npm install

# 设置工作目录为/build 并且拷贝项目代码
WORKDIR /build
COPY ./web .

# 下载依赖包并编译
RUN npm install && npm run build

FROM docker.nikm.cn/wuruiwm/nginx:1.23

# 拷贝builder阶段编译好的文件
COPY --from=builder /build/dist /www/web

# 拷贝nginx配置文件
COPY ./web/nginx/config /etc/nginx/conf.d