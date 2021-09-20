package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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
	EngineStorageHostServer string 				/** "localhost or IP:<X.X.X.X>" */
	EngineStorageHostServerPort int16			/** postgres port 5432 */
	EngineStorageServerUser string				/** postgres user "" */
	EngineStorageServerUserPassword string      /** postgres user password "" */
	EngineStorageServerDB string				/** postgres db "postgres" */
}

func NewEngineStorageConnection (hostStorageServer string, hostStorageServerPort int16, 
	                             storageServerUser string, storageServerUserPassword string,
								 storageServerDB string) (*EngineStorageConnection) {
	conn := &EngineStorageConnection {
		EngineStorageHostServer: hostStorageServer,
		EngineStorageHostServerPort: hostStorageServerPort,
		EngineStorageServerUser: storageServerUser,
		EngineStorageServerUserPassword: storageServerUserPassword,
		EngineStorageServerDB: storageServerDB,
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



func (self *EngineStorage) GetEngine(id int) (*Engine, error) {
	if self == nil {
		return nil, errors.New("reference to active TransactionalEngineStorage is nil")
	}
	var engine *Engine
	err := self.sql.QueryRow("SELECT * FROM engines WHERE id = $1", id).Scan(engine)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func (self *EngineStorage) GetEngines () ([]Engine, error) {
	if self == nil {
		return nil, errors.New("reference to active TransactionalEngineStorage is nil")
	}

	engineRows, err := self.sql.Query("SELECT * FROM engines")
	if err != nil {
		return nil, err
	}
	var engines []Engine
	engineRows.Scan(engines)

	return engines, nil
}

/** init **/
func init () {
    /** connect  to postgres on "localhost", port 5432, dbuser "isgogolgo13", dbuserpassword "isgogolgo13",  db "enginedb" */
	connection := NewEngineStorageConnection("localhost", 5432, "isgogolgo13", "isgogolgo13", "enginedb")

	db, err := initDB(connection)
	if err !=  nil {
		log.Fatalf("init error initDB() %s", err)
	}

	TransactionEngineStorage = NewEngineStorage(db) 
}


/** InitDB **/
func initDB (conn *EngineStorageConnection) (*sql.DB, error) {
	var dbCon = fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
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

	rows, err:= db.Query("SELECT * FROM engines")
	if err != nil {
		log.Fatalf("initDB() error db.Query() %s", err)
	}

	
}