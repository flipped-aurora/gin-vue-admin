daemon on;
worker_processes  50;
#error_log /dev/stdout warn;
error_log  /var/log/nginx/error.log error;


events {
    worker_connections 1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;
    # See http://licson.net/post/optimizing-nginx-for-large-file-delivery/ for more detail
        # This optimizes the server for HLS fragment delivery
    sendfile off;
    #tcp_nopush on;
    keepalive_timeout  65;
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';
    #access_log /dev/stdout combined;

#     ssl_ciphers         HIGH:!aNULL:!MD5;
#     ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
#     ssl_session_cache   shared:SSL:10m;
#     ssl_session_timeout 10m;

server {
        listen 80;

        # Uncomment these lines to enable SSL.
        # Update the ssl paths with your own certificate and private key.
        # listen 443 ssl;
        # ssl_certificate     /opt/certs/example.com.crt;
        # ssl_certificate_key /opt/certs/example.com.key;
        location / {
          root /var/www;
          try_files $uri $uri/ /index.html;
          index  index.html;
        }

        location /v1/ {
          proxy_set_header X-Forwarded-Proto $scheme;
          proxy_set_header X-Forwarded-Port $server_port;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_set_header Upgrade $http_upgrade;
          proxy_set_header Connection "upgrade";
          proxy_set_header  Host  $host;
          proxy_pass ${API_SERVER} ;
        }

}

}
