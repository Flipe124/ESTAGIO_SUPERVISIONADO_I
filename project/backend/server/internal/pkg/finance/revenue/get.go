package revenue

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/pkg/bd"
	"strconv"

	"github.com/rhuan-pk/pkutils/golang/aarm"
	"github.com/rhuan-pk/pkutils/golang/rdts"
)

// EndpointRevenueGetAll guarda a o caminho do endpoint para está função.
var EndpointRevenueGetAll = "/finance/revenue/get-all"

// EndpointRevenueGetStatus guarda a o caminho do endpoint para está função.
var EndpointRevenueGetStatus = "/finance/revenue/get-status"

// GetAll retorna todas as receitas do usuário.
func GetAll(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointRevenueGetStatus,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != http.MethodGet {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointRevenueGetAll + "\" contact!")

	query := "SELECT value, status "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE type = 'REVENUE'"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list:", err, "query:", query)
		apiReturn(http.StatusInternalServerError)
		return
	}
	log.Println("successfully query select!")

	log.Println("endpoint \"" + EndpointRevenueGetAll + "\" finish!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(response)

}

// GetStatus retorna todas as receitas pagas ou não pagas do usuário.
func GetStatus(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointRevenueGetStatus,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != http.MethodGet {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointRevenueGetStatus + "\" contact!")

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
	query += "WHERE type = 'REVENUE' "
	query += "  AND status = '" + status + "'"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list:", err, "query:", query)
		apiReturn(http.StatusInternalServerError)
		return
	}
	log.Println("successfully query select!")

	var total float32
	for _, value := range response {
		convertedValue, err := strconv.ParseFloat(value["value"], 32)
		if err != nil {
			log.Println("error on convert database value type of string to float!")
			apiReturn(http.StatusInternalServerError)
			return
		}
		total += float32(convertedValue)
	}

	log.Println("endpoint \"" + EndpointRevenueGetStatus + "\" finish!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(total)

}
