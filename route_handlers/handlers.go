package route_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/data"
	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/service_errors"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthCheckResponse := string([]byte("Service Healthy"))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(healthCheckResponse); err != nil {
		panic(err)
	}
}


/** GET /engines */
func GetEngines(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := json.NewEncoder(w)
	engines, err := data.TransactionEngineStorage.GetEngines()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		serviceError := service_errors.NewServiceError(err, engines)
		encoder.Encode(serviceError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if err := encoder.Encode(engines); err != nil {
		panic(err)
	}
}

/** GetEngines/{id} */
func GetEngine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := json.NewEncoder(w)
	if err !=  nil {
		serviceError:= service_errors.NewServiceError(err, nil)
		w.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(serviceError); err != nil {
			panic(err)
		}
	}

	engine, err := data.TransactionEngineStorage.GetEngine(id)
	if err != nil {
		serviceError:= service_errors.NewServiceError(err, nil)
		w.WriteHeader(http.StatusNotFound)
		if err := encoder.Encode(serviceError); err != nil {
			panic(err)
		}
	}
	
	w.WriteHeader(http.StatusOK)
	if err := encoder.Encode(engine); err != nil {
		panic(err)
	}
}