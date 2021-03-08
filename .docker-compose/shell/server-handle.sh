#! /bin/bash

rm -f ./config.yaml
# 生成config.yaml文件, 用于docker-compose的使用
touch ./config.yaml
filename="./config.yaml"
cat>"${filename}"<<EOF
# Gin-Vue-Admin Global Configuration

# jwt configuration
jwt:
  signing-key: 'qmPlus'
  expires-time: 604800
  buffer-time: 86400

# zap logger configuration
zap:
  level: 'info'
  format: 'console'
  prefix: '[GIN-VUE-ADMIN]'
  director: 'log'
  link-name: 'latest_log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  stacktrace-key: 'stacktrace'
  log-in-console: true

# redis configuration
redis:
  db: 0
  addr: '177.7.0.14:6379'
  password: ''

# email configuration
email:
  to: 'xxx@qq.com'
  port: 465
  from: 'xxx@163.com'
  host: 'smtp.163.com'
  is-ssl: true
  secret: 'xxx'
  nickname: 'test'

# casbin configuration
casbin:
  model-path: './resource/rbac_model.conf'

# system configuration
system:
  env: 'public'  # Change to "develop" to skip authentication for development mode
  addr: 8888
  db-type: 'mysql'
  oss-type: 'local'
  use-multipoint: false

# captcha configuration
captcha:
  key-long: 6
  img-width: 240
  img-height: 80

# mysql connect configuration
mysql:
  path: ''
  config: ''
  db-name: ''
  username: ''
  password: ''
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: false
  log-zap: ""

# local configuration
local:
  path: 'uploads/file'

# qiniu configuration (请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址)
qiniu:
  zone: 'ZoneHuadong'
  bucket: 'qm-plus-img'
  img-path: 'http://qmplusimg.henrongyi.top'
  use-https: false
  access-key: '25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ'
  secret-key: 'pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY'
  use-cdn-domains: false
EOF

