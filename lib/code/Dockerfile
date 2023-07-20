# 构建参数（注：FROM 指令之前的 ARG，只能用于 FROM 指令中）
ARG NODE_IMAGE="harbor.jqk8s.jqsoft.net/zhmz/node:16.20.0-alpine"
ARG NGINX_IMAGE="harbor.jqk8s.jqsoft.net/zhmz/nginx:1.20.1-alpine"

FROM $NODE_IMAGE as builder
# 镜像构建参数
#ARG NPM_REGISTRY="http://nexus.jqk8s.jqsoft.net/repository/npm/"
ARG BUILD_CMD="build"
ARG PNPM_VERSION=latest-7

WORKDIR /code
USER root
ADD .npmrc /code
ADD package.json /code
ADD pnpm-lock.yaml /code
RUN sed -i 's/registry=https:\/\/registry.npmmirror.com/registry=http:\/\/nexus.jqk8s.jqsoft.net\/repository\/npm/g' .npmrc
# 检查是否存在 pnpm,不存在的则全局安装
RUN command -v pnpm >/dev/null 2>&1 || { echo >&2 "require pnpm but it's not installed. It's being installed now......"; npm i pnpm@${PNPM_VERSION} -g; }
RUN pnpm install
# 添加源码
ADD . /code
# 构建代码
RUN pnpm $BUILD_CMD

FROM $NGINX_IMAGE as web
# 环境变量 接口服务 & 地址重写规则
ENV PLATFORM_API_SERVER="http://bigadata-sso.jqk8s.jqsoft.net" \
 PLATFORM_API_SERVER_RERITE="^/api/platform/(.*) /$1"
ARG NGINX_CONF="nginx.conf"
WORKDIR /usr/share/nginx/html/
COPY --from=builder /code/dist .
# nginx 环境变量配置模板,nginx会将模版文件转化为conf文件并放置于conf.d目录中
ADD $NGINX_CONF /etc/nginx/templates/default.conf.template
EXPOSE 80
