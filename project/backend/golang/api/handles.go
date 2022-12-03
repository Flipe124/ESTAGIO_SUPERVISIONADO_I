package api

import (
	"encoding/json"
	"golang/database"
	"log"
	"net/http"
)

func listApi(writer http.ResponseWriter, request *http.Request) {

	log.Println("listApi contact!")

	writer.Header().Set("Content=Type", "application/json")
	response := json.NewEncoder(writer)

	table, _ := database.List("table")
	response.Encode(table)

}
