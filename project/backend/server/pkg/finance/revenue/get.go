package revenue

import (
	"encoding/json"
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/bd"
)

var Endpoint = "finance/revenue/get-all"

// como implementar agora o paid e not_paid?
func GetAll(writer http.ResponseWriter, request *http.Request) {
	
	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			Endpoint,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + Endpoint + "\" contact!")

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
	log.Println("endpoint \"" + Endpoint + "\" complet!")

	json.NewEncoder(writer).Encode(response)

}
