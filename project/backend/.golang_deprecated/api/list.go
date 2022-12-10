package api

// import (
// 	"encoding/json"
// 	"golang/commons"
// 	"golang/database"
// 	"log"
// 	"net/http"
// 	"reflect"
// 	"strconv"
// 	"strings"
// )

// func list(writer http.ResponseWriter, request *http.Request) {

// 	first := 0

// 	apiReturn := func(httpStatusCode int) {
// 		commons.Api.StatusCodeReturn(
// 			commons.Api{},
// 			"list",
// 			&writer,
// 			httpStatusCode,
// 		)
// 	}

// 	if request.Method != "GET" {
// 		apiReturn(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	log.Println("endpoint \"/list\" contact!")

// 	data := commons.Api.NewTable(commons.Api{})
// 	err := json.NewDecoder(request.Body).Decode(&data)
// 	if err != nil {
// 		log.Println("error on json decode:", err)
// 		apiReturn(http.StatusUnprocessableEntity)
// 		return
// 	}

// 	// AQUI
// 	table := data.Name
// 	column := data.Columns[first]
// 	args := request.URL.Query().Get("args")

// 	var where string

// 	if table == "" || reflect.TypeOf(table).Name() != "string" {
// 		apiReturn(http.StatusBadRequest)
// 		return
// 	}

// 	if column != "" {

// 		if reflect.TypeOf(column).Name() != "string" {
// 			apiReturn(http.StatusBadRequest)
// 			return
// 		}

// 		if args != "" {
// 			for _, arg := range strings.Split(args, ",") {
// 				if _, err := strconv.Atoi(arg); err != nil {
// 					apiReturn(http.StatusUnprocessableEntity)
// 					return
// 				}
// 			}
// 		} else {
// 			apiReturn(http.StatusBadRequest)
// 			return
// 		}

// 		where = "WHERE " + column + " IN (" + args + ")"

// 	}

// 	response, err := database.List(table, where)
// 	if err != nil {
// 		log.Println("error on database list:", err)
// 		apiReturn(http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(writer).Encode(response)

// }
