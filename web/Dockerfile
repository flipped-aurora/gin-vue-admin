FROM node:16-alpine as build-stage

WORKDIR /gva-web/
RUN corepack enable

COPY package.json ./
RUN --mount=type=cache,id=yarn-store,target=/root/.yarn-store \
    yarn install --registry=https://registry.npmmirror.com

COPY . .
RUN yarn build

FROM nginx:alpine
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=build-stage /gva-web/dist /usr/share/nginx/html

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
