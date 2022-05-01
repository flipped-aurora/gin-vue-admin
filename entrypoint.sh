#!/bin/bash

if [ ! -d "/var/lib/mysql/gva" ]; then
    mysqld --initialize-insecure --user=mysql --datadir=/var/lib/mysql
    mysqld --daemonize --user=mysql
    sleep 5s
    mysql -uroot -e "create database gva default charset 'utf8' collate 'utf8_bin'; grant all on gva.* to 'gva'@'127.0.0.1' identified by '123456'; flush privileges;"
else
    mysqld --daemonize --user=mysql
fi
redis-server &
/usr/sbin/nginx &
cd /usr/share/nginx/html/ && ./server &
echo "gva ALL start!!!"
tail -f /dev/null
