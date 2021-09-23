#! /bin/sh

POSTGRES_DB_OWNER=isgogolgo13
POSTGRES_DB=enginedb

psql -f datastore_config/db-down.sql --host=localhost --username=${POSTGGRES_DB_OWNER} --dbname=${POSTGRES_DB}
dropdb  ${POSTGRES_DB}
dropuser ${POSTGRES_DB_OWNER}

