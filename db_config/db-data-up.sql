--DROP ROLE IF EXISTS isgogolgo132;
--CREATE ROLE isgogolgo132 WITH LOGIN PASSWORD 'isgogolgo132';
--ALTER USER isgogolgo132 CREATEDB;
--ALTER USER isgogolgo132 CREATEROLE;

--DROP DATABASE IF EXISTS EngineDB2;
--CREATE DATABASE EngineDB2;


CREATE TABLE IF NOT EXISTS Engines (
  id SERIAL PRIMARY KEY,
  serial_id VARCHAR(256) NOT NULL,
  engine_config VARCHAR(256) NOT NULL,
  engine_capacity NUMERIC(6, 2),
  engine_rpm_redline SMALLINT NOT NULL
);

GRANT ALL PRIVILEGES ON TABLE Engines TO isgogolgo132;

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('VW_100000001', 'V8', 6250.50, 12500);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Audi_100000002', 'V8', 6260.75, 13550);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Porsche_100000003', 'V8', 6270.50, 16250);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Porsche_100000004', 'V8', 6260.50, 15750);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Mercedes_AMG_100000005', 'V8-Twin-Turbo', 6265.50, 16550);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Mercedes_AMG_100000006', 'V8-Twin-Turbo', 6275.50, 16750);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Mercedes_AMG_100000007', 'V12-Twin-Turbo', 6295.50, 16750);

INSERT INTO Engines (serial_id, engine_config, engine_capacity, engine_rpm_redline) 
VALUES ('Mercedes_AMG_100000008', 'V12-Twin-Turbo', 6695.50, 18950);






