package api

import (
	"golang/settings"
	"log"
	"net/http"
	"strconv"
)

func Init() {

	apiPort := ":" + strconv.Itoa(settings.GetApiSetting().Port)

	http.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe(apiPort, nil))

}
