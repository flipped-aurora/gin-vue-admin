#!/bin/sh
envsubst '$API_SERVER' < /etc/nginx/nginx.conf.tpl > /etc/nginx/nginx.conf
env nginx
./gvadmin
