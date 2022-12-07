package commons

import (
	"log"
	"net/http"
)

type Row struct {
	Row []string `json:"row"`
}

type Table struct {
	Name     string   `json:"name"`
	Columns  []string `json:"columns"`
	Rows     []Row    `json:"rows"`
	Optional string   `json:"optional"`
}

type Api struct{}

func (api Api) StatusCodeReturn(endpoint string, writer *http.ResponseWriter, httpStatusCode int) {
	level := map[bool]string{true: "error", false: "hit"}[httpStatusCode != http.StatusOK]
	httpErrorMessage := http.StatusText(httpStatusCode)
	http.Error(*writer, httpErrorMessage, httpStatusCode)
	log.Println("endpoint \"/"+endpoint+"\" "+level+":", httpErrorMessage+"!")
}

func (api Api) NewTable() *Table {
	return new(Table)
}
