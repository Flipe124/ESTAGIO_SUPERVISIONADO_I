package main

import (
	"golang/api"
	"log"
)

func main() {

	log.Println("api is on 8008!")
	api.Init()

	// table := database.List("table", "WHERE cl = 139")
	// for row, value := range table {
	// 	fmt.Println("row:", row, "\ncolumn/value:", value)
	// }
	// database.Save("table", "SET cl = 25, cl1 = '1999-01-01', cl2 = 'no'", "WHERE cl = 25")

}
