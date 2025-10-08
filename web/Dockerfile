# 如果需要用 cicd ，请设置环境变量：
# variables:
#     DOCKER_BUILDKIT: 1

FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app


FROM base AS prod-deps
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod

FROM base AS build
COPY --from=prod-deps /app/node_modules /app/node_modules
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install && pnpm run build


FROM nginx:alpine
LABEL MAINTAINER="bypanghu@163.com"
COPY --from=base  /app/.docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=build  /app/dist /usr/share/nginx/html
RUN ls -al /usr/share/nginx/html
