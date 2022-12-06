package api

import (
	"encoding/json"
	"golang/commons"
	"golang/database"
	"log"
	"net/http"
	"strings"
)

func create(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.StatusCodeReturn(commons.Api{}, "create", &writer, httpStatusCode)
	}

	if request.Method != "POST" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/create\" contact!")

	data := commons.Api.NewTable(commons.Api{})
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Println("error on json decode:", err)
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	table := data.Name
	columns := "(" + strings.Join(data.Columns, ", ") + ")"
	rows := " VALUES "
	for _, row := range data.Rows {
		rows += "(" + strings.Join(row.Row, ", ") + "), "
	}
	query := columns + strings.TrimSuffix(rows, ", ")

	_, err = database.Save(table, query)
	if err != nil {
		log.Println("error on database save:", err)
		apiReturn(http.StatusInternalServerError)
		return
	}

	apiReturn(http.StatusOK)

}
