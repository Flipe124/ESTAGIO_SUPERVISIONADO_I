package finance

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/pkg/bd"

	"github.com/rhuan-pk/pkutils/golang/aarm"
)

// EndpointFinanceGetAll guarda a o caminho do endpoint para está função.
var EndpointFinanceGetAll = "/finance/get-all"

// GetAll retorna todas as informações da tabela finance.
func GetAll(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			EndpointFinanceGetAll,
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"" + EndpointFinanceGetAll + "\" contact!")

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
		log.Println("error on database list:", err, "query:", query)
		apiReturn(http.StatusInternalServerError)
		return
	}
	log.Println("successfully query select!")

	log.Println("endpoint \"" + EndpointFinanceGetAll + "\" finish!")
	log.Println("------------------- complet! -------------------")
	json.NewEncoder(writer).Encode(response)

}
