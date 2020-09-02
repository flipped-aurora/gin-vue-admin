#! /bin/bash

rm -f .env.production
touch .env.production
filename="./.env.production"
cat>"${filename}"<<EOF
ENV = 'production'
VUE_APP_BASE_API = ''
EOF

