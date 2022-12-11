package ipa

import (
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/finance/revenue"
)

var API = aarm.NewApi(8008)

func init() {

	API.SetHandles(
		map[string]func(http.ResponseWriter, *http.Request){
			revenue.Endpoint: revenue.GetAll,
		},
	)

	log.Println("api is on \"localhost:" + API.GetPortString() + "/\"!")
	log.Fatal(http.ListenAndServe(":"+API.GetPortString(), nil))

}
