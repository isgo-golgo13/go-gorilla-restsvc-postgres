DROP ROLE IF EXISTS isgogolgo13;
CREATE ROLE isgogolgo13 WITH LOGIN PASSWORD 'isgogolgo13';
ALTER USER isgogolgo13 CREATEDB;
ALTER USER isgogolgo13 CREATEROLE;

DROP DATABASE IF EXISTS EngineDB;
CREATE DATABASE EngineDB;