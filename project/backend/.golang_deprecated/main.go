/*
*** main ***

main é o pacote principal que start a aplicação.

by: rhuan-pk/rhuanpk
*/
package main

import (
	"fmt"
	_ "golang/api"
	"golang/database"
	"golang/settings"
	_ "log"
	_ "strconv"
)

func main() {

	table, _ := database.List("SELECT * FROM " + settings.GetDatabaseSetting().Name + ".table")
	fmt.Println(table)
	// pega a porta da api já convertendo para o tipo de dado correto.
	// apiPort := strconv.Itoa(settings.GetApiSetting().Port)

	// printa a mensagem de levantamento do serviço e o inicia.
	// log.Println("api is on \"localhost:" + apiPort + "/\"!")
	// api.Init()

}
