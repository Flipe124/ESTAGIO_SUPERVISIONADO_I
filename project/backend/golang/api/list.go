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

func list(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.Return(commons.Api{}, &writer, httpStatusCode)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/list\" contact!")

	table := request.URL.Query().Get("table")
	column := request.URL.Query().Get("column")
	args := request.URL.Query().Get("args")

	var where string

	if table == "" || reflect.TypeOf(table).Name() != "string" {
		apiReturn(http.StatusBadRequest)
		return
	}

	if column != "" {

		if reflect.TypeOf(column).Name() != "string" {
			apiReturn(http.StatusBadRequest)
			return
		}

		if args != "" {
			for _, arg := range strings.Split(args, ",") {
				if _, err := strconv.Atoi(arg); err != nil {
					apiReturn(http.StatusUnprocessableEntity)
					return
				}
			}
		} else {
			apiReturn(http.StatusBadRequest)
			return
		}

		where = "WHERE " + column + " IN (" + args + ")"

	}

	response, err := database.List(table, where)
	if err != nil {
		apiReturn(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(response)

}
