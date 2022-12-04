package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func listApi(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		httpErrorMessage := http.StatusText(http.StatusMethodNotAllowed)
		http.Error(writer, httpErrorMessage, http.StatusMethodNotAllowed)
		log.Println("endpoint \"/list\" error:", httpErrorMessage + "!")
		return
	}

	log.Println("endpoint \"/list\" contact!")

	writer.Header().Set("Content=Type", "application/json")
	// response := json.NewEncoder(writer)

	// table, _ := database.List("table")
	// response.Encode("table")

	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		httpErrorMessage := http.StatusText(http.StatusBadRequest)
		http.Error(writer, httpErrorMessage, http.StatusBadRequest)
		log.Println("endpoint \"/list\" error:", httpErrorMessage + "!")
		return
	}

	var response *any
	err = json.Unmarshal(requestBody, &response)
	if err != nil {
		httpErrorMessage := http.StatusText(http.StatusInternalServerError)
		http.Error(writer, httpErrorMessage, http.StatusInternalServerError)
		log.Println("endpoint \"/list\" error:", httpErrorMessage + "!")
		return
	}

	fmt.Println(*response)
	table, _ := database.List("finance")
	response.Encode(table)

}
