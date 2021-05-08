package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/robinsturmeit/imperial-fleet/db"
)

func GetSpacecrafts(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /spacecrafts
	// Lists all spacecrafts.
	//
	// TODO: Write swagger doc
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	spc, err := db.DBConnect.GetAllSpacecrafts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spc)
}

func GetSpacecraftByProp(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /spacecraft/{prop}/{value}
	//
	// TODO: Write swagger doc
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	spacecraft, err := db.DBConnect.GetSpacecraftByProp(params["prop"], params["value"])
	// TODO 404 error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spacecraft)
}
