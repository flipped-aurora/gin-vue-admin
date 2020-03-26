FROM node:alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

RUN npm install -g http-server

ENTRYPOINT [ "http-server", "dist" ]
