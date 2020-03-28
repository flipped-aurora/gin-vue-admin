FROM node:12.16.1-alpine3.9 AS build-stage

WORKDIR /tmp
COPY package*.json ./
RUN npm install
RUN mkdir -p /app \
    && cp -a /tmp/node_modules /app

WORKDIR /app
COPY . .
RUN npm run build

RUN npm install -g http-server
ENTRYPOINT ["http-server", "dist"]
