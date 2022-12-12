package expense

import (
	"encoding/json"
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/bd"
	"server/rdts"
	"strconv"
)

var EndpointExpenseGetAll = "/finance/expense/get-all"
var EndpointExpenseGetStatus = "/finance/expense/get-status"

func GetAll(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointExpenseGetAll,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointExpenseGetAll + "\" contact!")

	query := "SELECT value, status "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE type = 'EXPENSE'"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	// log.Println("endpoint \"" + EndpointExpenseGetAll + "\" complet!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(response)

}

func GetStatus(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointExpenseGetStatus,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointExpenseGetStatus + "\" contact!")

	payload := rdts.NewTable()
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		log.Println("error on json decode:", err)
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	status, statusExists := payload.Options["status"]
	if !statusExists {
		log.Println("error on get status type, missing!")
		apiReturn(http.StatusBadRequest)
		return
	}

	query := "SELECT value "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE type = 'EXPENSE' "
	query += "  AND status = '" + status + "'"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	var total float32
	for _, value := range response {
		convertedValue, err := strconv.ParseFloat(value["value"], 32)
		if err != nil {
			log.Println("error on convert value type string to float!")
			apiReturn(http.StatusInternalServerError)
			return
		}
		total += float32(convertedValue)
	}

	// log.Println("endpoint \"" + EndpointExpenseGetStatus + "\" complet!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(total)

}
