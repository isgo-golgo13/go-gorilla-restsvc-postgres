package data

import (
	"errors"
	"fmt"
)

type Engine struct {
	ID            string `json:"id"`
	SerialID      string `json:"serial_id"`
	EngineConfig string `json:"engine_config"`
	EngineCapacity float32 `json:"engine_capacity"`
}



var TransactionEngineStorage *EngineStorage
type EngineStorage struct {
	engines []Engine
}

func NewEngineStorage () (*EngineStorage) {
	storage := EngineStorage {
		engines: []Engine {
			{
				ID:            "100000001",
				SerialID:      "VW_100000001",
				EngineConfig:  "V8",
				EngineCapacity: 250.50,
				
			},
			{
				ID:            "100000002",
				SerialID:      "Audi_100000002",
				EngineConfig:  "V8",
				EngineCapacity: 220.50,
			},
			{
				ID:            "100000003",
				SerialID:      "Porsche_100000003",
				EngineConfig:  "V8",
				EngineCapacity: 50.50,
			},
			{
				ID:            "100000004",
				SerialID:      "Porsche_100000004",
				EngineConfig:          "V8",
				EngineCapacity: 270.50,
			},
			{
				ID:            "100000005",
				SerialID:      "Mercedes_10000005",
				EngineConfig:          "V8-Twin-Turbo",
				EngineCapacity: 250.25,
			},
			{
				ID:            "100000006",
				SerialID:      "Mercedes_10000006",
				EngineConfig:          "V8-Twin-Turbo",
				EngineCapacity: 270.25,
			},
			{
				ID:            "100000007",
				SerialID:      "Mercedes_10000007",
				EngineConfig:          "V12-Twin-Turbo",
				EngineCapacity: 350.75,
			},
			{
				ID:            "100000008",
				SerialID:      "Mercedes_10000008",
				EngineConfig:          "V8",
				EngineCapacity: 250.50,
			},
		},
	}
	return &storage
}



func (self *EngineStorage) checkRange(id int) (bool, error) {
	if id > len(self.engines)-1 {
		id_out_of_range := fmt.Sprintf("%d", id)
		return false, errors.New("id: " + id_out_of_range + " out of range of storage array")
	} else {
		return true, nil
	}
}

/** Do range check if id is out of range in the DB array */
func (self *EngineStorage) GetEngine(id int) (*Engine, error) {
	if self == nil {
		return nil, nil 
	}
	in_range, err := self.checkRange(id)
	if in_range == false && err != nil {
		return nil, err
	} else {
		return &self.engines[id], nil
	}
}

func (self *EngineStorage) GetEngines () ([]Engine, error) {
	if self == nil {
		return nil, errors.New("reference to active TransactionalEngineStorage is nil")
	}
	return self.engines, nil
}

func init () {
	TransactionEngineStorage = NewEngineStorage()
}