package revenue

import (
	"encoding/json"
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/bd"
	"strconv"
)

var EndpointGetAll = "/finance/revenue/get-all"
var EndpointGetPaid = "/finance/revenue/get-paid"
var EndpointGetNotPaid = "/finance/revenue/get-not-paid"

func GetAll(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointGetAll,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointGetAll + "\" contact!")

	query := "SELECT "
	query += "  id, "
	query += "  account_id, "
	query += "  category_id, "
	query += "  value, "
	query += "  type, "
	query += "  status, "
	query += "  description, "
	query += "  date, "
	query += "  created_at, "
	query += "  updated_at "
	query += "FROM " + bd.DB.Name + ".finance"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	log.Println("endpoint \"" + EndpointGetAll + "\" complet!")
	json.NewEncoder(writer).Encode(response)

}

func GetPaid(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointGetPaid,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointGetPaid + "\" contact!")

	query := "SELECT value "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE type = 'REVENUE' "
	query += "  AND status = 'PAID'"

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

	log.Println("endpoint \"" + EndpointGetPaid + "\" complet!")
	json.NewEncoder(writer).Encode(total)

}

func GetNotPaid(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointGetNotPaid,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointGetNotPaid + "\" contact!")

	query := "SELECT value "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE type = 'REVENUE' "
	query += "  AND status = 'NOT_PAID'"

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

	log.Println("endpoint \"" + EndpointGetNotPaid + "\" complet!")
	json.NewEncoder(writer).Encode(total)

}
