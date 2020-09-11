#! /bin/bash

rm -f ./config.yaml
# 生成config.yaml文件, 用于docker-compose的使用
touch ./config.yaml
filename="./config.yaml"
cat>"${filename}"<<EOF
# Gin-Vue-Admin Global Configuration

# casbin configuration
casbin:
    model-path: './resource/rbac_model.conf'

# jwt configuration
jwt:
    signing-key: 'qmPlus'

# mysql connect configuration
mysql:
    username: root
    password: 'Aa@6447985'
    path: mysql
    db-name: 'qmPlus'
    config: 'charset=utf8mb4&parseTime=True&loc=Local'
    max-idle-conns: 10
    max-open-conns: 10
    log-mode: true

#sqlite 配置
sqlite:
    path: db.db
    log-mode: true
    config: 'loc=Asia/Shanghai'

# oss configuration

# 切换本地与七牛云上传，分配头像和文件路径
localupload:
  local: false
  avatar-path: uploads/avatar
  file-path: uploads/file

# 请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址
qiniu:
    access-key: '25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ'
    secret-key: 'pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY'
    bucket: 'qm-plus-img'
    img-path: 'http://qmplusimg.henrongyi.top'

# redis configuration
redis:
    addr: redis:6379
    password: ''
    db: 0

# system configuration
system:
    use-multipoint: true
    env: 'public'  # Change to "develop" to skip authentication for development mode
    addr: 8888
    db-type: "mysql"  # support mysql/sqlite

# captcha configuration
captcha:
    key-long: 6
    img-width: 240
    img-height: 80

# zap logger configuration
zap:
  # 可使用 "debug", "info", "warn", "error", "dpanic", "panic", "fatal",
  level: 'debug'
  # console: 控制台, json: json格式输出
  format: 'console'
  prefix: '[GIN-VUE-ADMIN]'
  director: 'log'
  link_name: 'latest_log'
  show_line: true
  # LowercaseLevelEncoder:小写, LowercaseColorLevelEncoder:小写带颜色,CapitalLevelEncoder: 大写, CapitalColorLevelEncoder: 大写带颜色,
  encode_level: 'LowercaseColorLevelEncoder'
  stacktrace_key: 'stacktrace'
  log_in_console: true

email:
  email_from: 'xxx@163.com'
  email_nick_name: 'test'
  email_secret: 'xxx'
  email_to: 'xxx@qq.com'
  email_host: 'smtp.163.com'
  email_port: 465
  email_isSSL: true
EOF

