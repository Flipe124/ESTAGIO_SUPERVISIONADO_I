package balance

import (
	"encoding/json"
	"log"
	"net/http"
	"server/aarm"
	"server/pkg/bd"
	"server/rdts"
)

func list(writer http.ResponseWriter, request *http.Request) {

	database := bd.DB

	apiReturn := func(httpStatusCode int) {
		aarm.StatusCodeReturn(
			"list",
			&writer,
			httpStatusCode,
		)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}
	log.Println("endpoint \"/list\" contact!")

	payload := rdts.NewTable()
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		log.Println("error on json decode:", err)
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	query := "query"
	// column := reflect.
	// 	ValueOf(payload.Options).
	// 	MapKeys()[0].
	// 	String()
	// option := payload.Options[column]

	response, err := database.List(query)
	if err != nil {
		log.Println("error on database list:", err)
		apiReturn(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(response)

}
