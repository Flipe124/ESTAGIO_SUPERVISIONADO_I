package database

import (
	"golang/settings"
	"log"
)

/*
Save insere ou atualizad um novo dado no banco retornando a quantidade de linhas modificadas caso seja um update ou a última linha inserida caso seja um create.

Informe a query completa a ser executada para esta função.
*/
func Save(query, typing string) (int, error) {

	// pega as configuraçẽos do pacote, tenta abrir conexão com o banco, caso falhe retorne o erro se não continue.
	database, err := openConnection(settings.GetDatabaseSetting())
	if err != nil {
		log.Println("estabilish connection failed:", err)
		return -1, err
	}
	defer database.Close()
	log.Println("successfully estabilish connection!")

	// executa a query, caso falhe retorne o erro se não continue.
	result, err := database.Exec(query)
	if err != nil {
		log.Println("query exec failed:", err, "query:", query)
		return -1, err
	}
	log.Println("successfully query exec!")

	// pega a quantidade de linhas modificadas ou a última inserida para usar como retorno.
	var rowsAffected, lastIdInserted int64
	if typing != "CREATE" {
		rowsAffected, _ = result.RowsAffected()
	} else {
		lastIdInserted, _ = result.RowsAffected()
	}
	
	return int((rowsAffected | lastIdInserted)), nil

}
