package api

import (
	"encoding/json"
	"golang/commons"
	"golang/database"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func listApi(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		commons.Api.ReturnError(commons.Api{}, &writer, http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/list\" contact!")

	table := request.URL.Query().Get("table")
	column := request.URL.Query().Get("column")
	ids := request.URL.Query().Get("ids")

	var where string

	if table == "" || reflect.TypeOf(table).Name() != "string" {
		commons.Api.ReturnError(commons.Api{}, &writer, http.StatusBadRequest)
		return
	}

	if column != "" {

		if reflect.TypeOf(column).Name() != "string" {
			commons.Api.ReturnError(commons.Api{}, &writer, http.StatusBadRequest)
			return
		}

		if ids != "" {

			for _, id := range strings.Split(ids, ",") {
				if _, err := strconv.Atoi(id); err != nil {
					commons.Api.ReturnError(commons.Api{}, &writer, http.StatusUnprocessableEntity)
					return
				}
			}

		}

		where = "WHERE " + column + " IN (" + ids + ")"

	}

	response, err := database.List(table, where)
	if err != nil {
		commons.Api.ReturnError(commons.Api{}, &writer, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(response)

}
