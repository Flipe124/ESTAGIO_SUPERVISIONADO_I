package ipa

import (
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/finance/revenue"
)

var API = aarm.NewApi(8008)

func init() {

	endpoints := map[string]func(http.ResponseWriter, *http.Request){
		revenue.EndpointGetAll:     revenue.GetAll,
		revenue.EndpointGetPaid:    revenue.GetPaid,
		revenue.EndpointGetNotPaid: revenue.GetNotPaid,
	}

	API.SetHandles(endpoints)

	log.Println("api is on \"localhost:" + API.GetPortString() + "/\"!")
	log.Println("available endpoints:")
	for endpoint := range endpoints {
		log.Println("  -", endpoint)
	}
	log.Fatal(http.ListenAndServe(":"+API.GetPortString(), nil))

}
