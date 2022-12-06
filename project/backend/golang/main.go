package main

import (
	"golang/api"
	"golang/settings"
	"log"
	"strconv"
)

func main() {

	apiPort := strconv.Itoa(settings.GetApiSetting().Port)
	log.Println("api is on \"localhost:" + apiPort + "/\"!")
	api.Init()

}
