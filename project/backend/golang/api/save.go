package api

import (
	"encoding/json"
	"golang/commons"
	"log"
	"net/http"
)

func save(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.Return(commons.Api{}, "save", &writer, httpStatusCode)
	}

	if request.Method != "POST" && request.Method != "PATCH" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	} else {
		isUpdate := request.Method == "PATCH"
	}

	log.Println("endpoint \"/save\" contact!")

	type args struct {
		Column string `json:"column"`
		Row    string `json:"row"`
	}

	var table struct {
		Name string `json:"name"`
		Args []args `json:"args"`
	}

	err := json.NewDecoder(request.Body).Decode(&table)
	if err != nil {
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	// table
	// table.Name
	// table.Args
	// table.Args[0]
	// table.Args[0].Column
	// table.Args[0].Row

	
	
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
