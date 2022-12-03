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

	// table := database.List("table", "WHERE cl = 139")
	// for row, value := range table {
	// 	fmt.Println("row:", row, "\ncolumn/value:", value)
	// }
	// database.Save("table", "SET cl = 25, cl1 = '1999-01-01', cl2 = 'no'", "WHERE cl = 25")

}
