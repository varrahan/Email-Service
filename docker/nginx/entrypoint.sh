#!/bin/sh
set -e

# Check if env variables are loaded in container
: "${PORT:?PORT not set}"
: "${NGINX_PORT:?NGINX_PORT not set}"

# Replace env variables in nginx template
envsubst '$NGINX_PORT $PORT' < /etc/nginx/templates/nginx.template.conf > /etc/nginx/conf.d/default.conf

# Start NGINX in foreground
nginx -g 'daemon off;'
