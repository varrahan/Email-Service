#!/bin/sh
set -e

# Check if env variables are loaded in container
: "${PORT:?PORT not set}"
: "${SMTP_HOST:?SMTP_HOST not set}"
: "${SMTP_PORT:?SMTP_PORT not set}"
: "${SMTP_USER:?SMTP_USER not set}"
: "${SMTP_PASS:?SMTP_PASS not set}"
: "${TO_ADDRESS:?TO_ADDRESS not set}"
: "${FROM_ADDRESS:?FROM_ADDRESS not set}"

# Run application
./email-service
