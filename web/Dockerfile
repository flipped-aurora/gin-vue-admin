FROM node:16

WORKDIR /gva-web/
COPY . .
# 缓存包，开发版不要 nginx
RUN npm install --registry https://registry.npm.taobao.org/