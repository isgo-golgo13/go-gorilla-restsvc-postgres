package datastore_service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/datastore_config"
	_ "github.com/lib/pq"
)

type Engine struct {
	ID            string `json:"id"`
	SerialID      string `json:"serial_id"`
	EngineConfig string `json:"engine_config"`
	EngineCapacity float32 `json:"engine_capacity"`
	EngineRPMRedline int32 `json:"engine_rpm_redline"`
}


type EngineStorageConnection struct {
	EngineStorageHostServer string 				/** "localhost" or server IP */
	EngineStorageHostServerPort int16			/** postgres port 5432 */
	EngineStorageServerUser string				/** postgres user "" */
	EngineStorageServerUserPassword string      /** postgres user password "" */
	EngineStorageServerDB string				/** postgres db "enginedb" */
}


func NewEngineStorageConnection () (*EngineStorageConnection) {

	serverPort := datastore_config.GoDotEnvVar("EngineStorageHostServerPort")
	serverPortConv, err := strconv.ParseInt(serverPort, 10, 16)
	if err != nil {
		log.Fatalf("%s", err)
	}

	conn := &EngineStorageConnection {
		EngineStorageHostServer: datastore_config.GoDotEnvVar("EngineStorageHostServer"),
		EngineStorageHostServerPort: int16(serverPortConv),
		EngineStorageServerUser: datastore_config.GoDotEnvVar("EngineStorageServerUser"),
		EngineStorageServerUserPassword: datastore_config.GoDotEnvVar("EngineStorageServerUserPassword"),
		EngineStorageServerDB: datastore_config.GoDotEnvVar("EngineStorageServerDB"),
	}
	return conn
}


var TransactionEngineStorage *EngineStorage
type EngineStorage struct {
	sql *sql.DB
}

func NewEngineStorage (db *sql.DB) (*EngineStorage) {
	 storage := EngineStorage { sql: db }
	return &storage
}


/** GetEngines/{id} */
func (self *EngineStorage) GetEngine(id int) (*Engine, error) {
	if self == nil {
		return nil, errors.New("reference to active TransactionalEngineStorage is nil")
	}
	var engine Engine
	err := self.sql.QueryRow("SELECT * FROM engines WHERE id = $1", id).Scan(
		    &engine.ID, 
			&engine.SerialID, 
			&engine.EngineConfig, 
			&engine.EngineCapacity, 
			&engine.EngineRPMRedline)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

/** GetEngines */
func (self *EngineStorage) GetEngines () ([]Engine, error) {
	if self == nil {
		return nil, errors.New("reference to active TransactionalEngineStorage is nil")
	}

	rows, err:= self.sql.Query("SELECT * FROM engines")
	if err != nil {
		log.Fatalf("GetEngines() error sql.Query() %s", err)
	}
	
	var engines []Engine
	var rowErr error
	for rows.Next() {
		engine := Engine{}
		if rowErr = rows.Scan(&engine.ID, 
			               &engine.SerialID, 
						   &engine.EngineConfig, 
						   &engine.EngineCapacity, 
						   &engine.EngineRPMRedline); err != nil {
		  return nil, rowErr
		}
		engines = append(engines, engine)
	}
	return engines, nil
}


/** init **/
func init () {

	dbCon := NewEngineStorageConnection()
	db, err := initDB(dbCon)
	if err !=  nil {
		log.Fatalf("init error initDB() %s", err)
	}
	TransactionEngineStorage = NewEngineStorage(db) 
}


/** InitDB **/
func initDB (conn *EngineStorageConnection) (*sql.DB, error) {
	var dbCon = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    	conn.EngineStorageHostServer, 
		conn.EngineStorageHostServerPort, 
		conn.EngineStorageServerUser, conn.EngineStorageServerUserPassword, 
		conn.EngineStorageServerDB)

	log.Println(dbCon)
	
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		log.Fatalf("initDB error sql.Open() %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("initDB() error db.Ping() %s", err)
	}
	log.Print("db.Ping() connection to postgres OK")
	return db, nil	
}