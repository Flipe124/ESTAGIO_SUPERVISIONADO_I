package balance

import (
	"encoding/json"
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/bd"
	"strconv"
)

// EndpointBalanceGet guarda a o caminho do endpoint para está função.
var EndpointBalanceGet = "/finance/balance/get"

// Get retorna o saldo (créditos menos débitos) do usuário.
func Get(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointBalanceGet,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != http.MethodGet {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointBalanceGet + "\" contact!")

	query := "SELECT type, value "
	query += "FROM " + bd.DB.Name + ".finance "
	query += "WHERE (type = 'REVENUE' AND status = 'PAID') "
	query += "  OR (type = 'EXPENSE' AND status = 'NOT_PAID')"

	response, err := bd.DB.List(query)
	if err != nil {
		log.Println("error on database list!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	var total float32
	for _, row := range response {

		convertedValue, err := strconv.ParseFloat(row["value"], 32)
		if err != nil {
			log.Println("error on convert value type string to float!")
			apiReturn(http.StatusInternalServerError)
			return
		}

		if row["type"] != "REVENUE" {
			total -= float32(convertedValue)
		} else {
			total += float32(convertedValue)
		}

	}

	log.Println("endpoint \"" + EndpointBalanceGet + "\" finish!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(total)

}
