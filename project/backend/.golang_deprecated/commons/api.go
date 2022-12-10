package commons

import (
	"log"
	"net/http"
)

// Api estrutura para fazer o vínculo das funções comuns relacioadas a API.
type Api struct{}

// StatusCodeReturn é a função comum para código de status de requisição
func (api Api) StatusCodeReturn(endpoint string, writer *http.ResponseWriter, httpStatusCode int) {
	level := map[bool]string{true: "error", false: "hit"}[httpStatusCode != http.StatusOK]
	httpErrorMessage := http.StatusText(httpStatusCode)
	http.Error(*writer, httpErrorMessage, httpStatusCode)
	log.Println("endpoint \"/"+endpoint+"\" "+level+":", httpErrorMessage+"!")
}
