#!/bin/bash

# 定义变量
IP=""
PORT=""

# 解析命令行参数
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -ip) IP="$2"; shift ;;
        -port) PORT="$2"; shift ;;
        *) echo "未知参数: $1"; exit 1 ;;
    esac
    shift
done

# 检查是否提供了必须的参数
if [[ -z "$IP" || -z "$PORT" ]]; then
    echo "必须提供 -ip 和 -port 参数"
    exit 1
fi

# 创建目录和文件（如果不存在）
sudo mkdir -p /etc/systemd/system/docker.service.d
sudo touch /etc/systemd/system/docker.service.d/http-proxy.conf

# 写入配置文件
sudo bash -c "cat > /etc/systemd/system/docker.service.d/http-proxy.conf <<EOL
[Service]
Environment=\"HTTP_PROXY=http://$IP:$PORT\"
Environment=\"HTTPS_PROXY=http://$IP:$PORT\"
Environment=\"NO_PROXY=localhost,127.0.0.1\"
EOL"

# 重新加载 systemd 守护进程
sudo systemctl daemon-reload

# 重启 Docker 服务
sudo service docker restart

echo "配置已应用并重启 Docker 服务"
