package api

import (
	"golang/settings"
	"log"
	"net/http"
	"strconv"
)

func Init() {

	port := ":" + strconv.Itoa(settings.GetApiSetting().Port)

	http.HandleFunc("/list", listApi)
	log.Fatal(http.ListenAndServe(port, nil))

}
