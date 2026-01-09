FROM nginx:alpine

# Copy NGINX template
COPY nginx.template.conf /etc/nginx/templates/nginx.template.conf

# Copy NGINX entrypoint script
COPY nginx-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# Run entrypoint
ENTRYPOINT ["/docker-entrypoint.sh"]
