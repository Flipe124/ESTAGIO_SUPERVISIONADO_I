package commons

import (
	"log"
	"net/http"
)

type Api struct{}

func (api Api) ReturnError(writer *http.ResponseWriter, httpStatusCode int) {
	httpErrorMessage := http.StatusText(httpStatusCode)
	http.Error(*writer, httpErrorMessage, httpStatusCode)
	log.Println("endpoint \"/list\" error:", httpErrorMessage+"!")
}
