FROM centos:7
WORKDIR /opt
ARG Version=v2.5.3
ENV Version=${Version} \
    LANG=en_US.utf8
COPY entrypoint.sh .
COPY build/* /usr/share/nginx/html/
COPY server/config.yaml /usr/share/nginx/html/config.yaml
COPY web/.docker-compose/nginx/conf.d/nginx.conf /etc/nginx/conf.d/nginx.conf
RUN set -ex \
    && echo "LANG=en_US.utf8" > /etc/locale.conf \
    && echo "net.core.somaxconn = 1024" >> /etc/sysctl.conf \
    && echo "vm.overcommit_memory = 1" >> /etc/sysctl.conf \
    && yum -y install wget gcc epel-release git yum-utils \
    && echo -e "[nginx-stable]\nname=nginx stable repo\nbaseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://nginx.org/keys/nginx_signing.key" > /etc/yum.repos.d/nginx.repo \
    && rpm --import https://nginx.org/keys/nginx_signing.key \
    && yum -y localinstall http://mirrors.ustc.edu.cn/mysql-repo/mysql57-community-release-el7.rpm \
    && yum -y install mysql-community-server redis nginx && chmod +x ./entrypoint.sh
EXPOSE 80
ENTRYPOINT ["./entrypoint.sh"]