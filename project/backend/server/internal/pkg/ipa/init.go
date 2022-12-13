package ipa

import (
	"log"
	"net/http"
	"server/internal/pkg/finance"
	"server/internal/pkg/finance/balance"
	"server/internal/pkg/finance/expense"
	"server/internal/pkg/finance/revenue"

	"github.com/rhuan-pk/pkutils/golang/aarm"
)

// API é a representação da API para ser usada no projeto.
var API = aarm.NewAPI(8008)

func init() {

	endpoints := map[string]func(http.ResponseWriter, *http.Request){
		// contact only.
		finance.EndpointFinanceGetAll: finance.GetAll,
		// contact only.
		revenue.EndpointRevenueGetAll: revenue.GetAll,
		// {"options": {"status": "[NOT]_PAID"}}.
		revenue.EndpointRevenueGetStatus: revenue.GetStatus,
		// contact only.
		expense.EndpointExpenseGetAll: expense.GetAll,
		// {"options": {"status": "[NOT]_PAID"}}.
		expense.EndpointExpenseGetStatus: expense.GetStatus,
		// contact only.
		balance.EndpointBalanceGet: balance.Get,
	}

	API.SetHandles(endpoints)

	log.Println("---------------- *** server *** ----------------")
	log.Println("api is on \"localhost:" + API.GetPortString() + "/\"!")
	log.Println("available endpoints:")
	for endpoint := range endpoints {
		log.Println("  -", endpoint)
	}
	log.Println("------------------------------------------------")

}
