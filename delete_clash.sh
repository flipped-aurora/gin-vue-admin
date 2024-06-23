#!/bin/bash

# 定义要删除的文件路径
FILE="/etc/systemd/system/docker.service.d/http-proxy.conf"

# 检查文件是否存在
if [ -f "$FILE" ]; then
    # 删除文件
    sudo rm "$FILE"
    echo "文件 $FILE 已删除"
else
    echo "文件 $FILE 不存在"
fi

# 重新加载 systemd 守护进程
sudo systemctl daemon-reload

# 重启 Docker 服务
sudo service docker restart

echo "Docker 服务已重启"
