package api

import (
	"encoding/json"
	"golang/commons"
	"log"
	"net/http"
)

func save(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.Return(commons.Api{}, &writer, httpStatusCode)
	}

	if request.Method != "POST" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/save\" contact!")

	var data struct {
		Table string `json:"table"`
		// Args []any `json:"args"`
	}

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	log.Println(data)
	apiReturn(http.StatusOK)

	// table := request.URL.Query().Get("table")
	// column := request.URL.Query().Get("column")
	// args := request.URL.Query().Get("args")

	// var where string

	// if table == "" || reflect.TypeOf(table).Name() != "string" {
	// 	commons.Api.ReturnError(commons.Api{}, &writer, http.StatusBadRequest)
	// 	return
	// }

	// if column != "" {

	// 	if reflect.TypeOf(column).Name() != "string" {
	// 		commons.Api.ReturnError(commons.Api{}, &writer, http.StatusBadRequest)
	// 		return
	// 	}

	// 	if args != "" {
	// 		for _, arg := range strings.Split(args, ",") {
	// 			if _, err := strconv.Atoi(arg); err != nil {
	// 				commons.Api.ReturnError(commons.Api{}, &writer, http.StatusUnprocessableEntity)
	// 				return
	// 			}
	// 		}
	// 	} else {
	// 		commons.Api.ReturnError(commons.Api{}, &writer, http.StatusBadRequest)
	// 		return
	// 	}

	// 	where = "WHERE " + column + " IN (" + args + ")"

	// }

	// response, err := database.List(table, where)
	// if err != nil {
	// 	commons.Api.ReturnError(commons.Api{}, &writer, http.StatusInternalServerError)
	// 	return
	// }

	// json.NewEncoder(writer).Encode(response)

}
