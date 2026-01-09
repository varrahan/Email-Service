FROM nginx:alpine

# Copy NGINX template
COPY docker/nginx/nginx.template.conf /etc/nginx/templates/nginx.template.conf

# Copy NGINX entrypoint script
COPY docker/nginx/entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# Run entrypoint
ENTRYPOINT ["/docker-entrypoint.sh"]
