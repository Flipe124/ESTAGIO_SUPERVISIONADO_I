package commons

import (
	"log"
	"net/http"
)

type Api struct{}

func (api Api) Return(endpoint string, writer *http.ResponseWriter, httpStatusCode int) {
	level := map[bool]string{true: "error", false: "hit"}[httpStatusCode != http.StatusOK]
	httpErrorMessage := http.StatusText(httpStatusCode)
	http.Error(*writer, httpErrorMessage, httpStatusCode)
	log.Println("endpoint \"/"+endpoint+"\" "+level+":", httpErrorMessage+"!")
}
