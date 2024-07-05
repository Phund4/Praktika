#!/bin/sh

cat <<EOT > /app/.env
HOST=${HOST:-my_db}
PORT=${PORT:-}
USER=${USER:-phunda}
PASSWORD=${PASSWORD:-098908}
DBNAME=${DBNAME:-hh_praktika}
SERVER_HOST=${SERVER_HOST:-localhost}
SERVER_PORT=${SERVER_PORT:-8080}
EOT

/app/praktikaBack