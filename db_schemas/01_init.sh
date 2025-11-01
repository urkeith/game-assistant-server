#!/bin/bash
set -e

echo "Starting database initialization..."

SELECTED_GAME=${SELECTED_GAME:-hannibal}

# Load game-specific schema based on SELECTED_GAME
case "$SELECTED_GAME" in
    hannibal)
        echo "Loading Hannibal schema..."
        psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	\i docker-entrypoint-initdb.d/hannibal/hannibal_schema.sql
EOSQL
        ;;
    ato)
        echo "Loading ATO schema..."
        psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    \i /docker-entrypoint-initdb.d/ato/ato_schema.sql
EOSQL
        ;;
    *)
        echo "No additional schema to load for SELECTED_GAME=$SELECTED_GAME"
        ;;
esac

echo "Database initialization complete!"
