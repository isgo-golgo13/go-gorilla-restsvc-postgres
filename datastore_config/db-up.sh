#! /bin/sh

POSTGRES_DB_OWNER=isgogolgo13
POSTGRES_DB=enginedb

createuser --createdb --createrole --login ${POSTGRES_DB_OWNER}  -e 
#createdb ${POSTGRES_DB} --owner=${POSTGRES_DB_OWNER} # Superuse can only do this
createdb ${POSTGRES_DB}
psql -f datastore_config/db-up.sql --host=localhost --username=${POSTGGRES_DB_OWNER} --dbname=${POSTGRES_DB}