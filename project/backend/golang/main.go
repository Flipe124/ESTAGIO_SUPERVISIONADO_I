package main

import (
	"fmt"
	"golang/database"
)

func main() {
	table := database.List("tb", "WHERE cl = 139")
	for row, value := range table {
		fmt.Println("row:", row, "\ncolumn/value:", value)
	}
	database.Save("tb", "SET cl = 25, cl1 = '1999-01-01', cl2 = 'no'", "WHERE cl = 25")
}
