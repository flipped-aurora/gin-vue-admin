FROM centos:7
WORKDIR /opt
ARG Version=v2.5.3
ENV Version=${Version} \
    LANG=en_US.utf8
COPY entrypoint.sh .
COPY build/ /usr/share/nginx/html/
COPY server/config.yaml /usr/share/nginx/html/config.yaml
COPY web/.docker-compose/nginx/conf.d/nginx.conf /etc/nginx/conf.d/nginx.conf
RUN set -ex \
    && echo "LANG=en_US.utf8" > /etc/locale.conf \
    && echo "net.core.somaxconn = 1024" >> /etc/sysctl.conf \
    && echo "vm.overcommit_memory = 1" >> /etc/sysctl.conf \
    && yum -y install yum -y install *epel* \
    && yum -y localinstall http://mirrors.ustc.edu.cn/mysql-repo/mysql57-community-release-el7.rpm \
    && yum -y install mysql-community-server redis nginx  --nogpgcheck && chmod +x ./entrypoint.sh
EXPOSE 80
ENTRYPOINT ["./entrypoint.sh"]