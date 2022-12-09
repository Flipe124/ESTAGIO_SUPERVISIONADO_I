package api

import (
	"encoding/json"
	"golang/commons"
	"golang/database"
	"log"
	"net/http"
	"strings"
)

func update(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.StatusCodeReturn(
			commons.Api{},
			"update",
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "PATCH" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/update\" contact!")

	var first uint8 = 0

	data := commons.Api.NewTable(commons.Api{})
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Println("error on json decode:", err)
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	table := data.Name
	optional := data.Optional
	query := "SET "
	for index, column := range data.Columns {
		query += column + " = " + data.Rows[first].Row[index] + ", "
	}

	query = strings.TrimSuffix(query, ", ")
	where := "WHERE id = " + optional

	_, err = database.Save(table, query, where)
	if err != nil {
		log.Println("error on database update!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	apiReturn(http.StatusOK)

}
