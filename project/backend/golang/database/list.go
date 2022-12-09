package database

import (
	"golang/settings"
	"log"
)

/*
List retorna a representação de uma tabela do banco sendo que cada índice da slice é uma linha que contem dentro um mapa no qual a chave é o nome da coluna e o seu valor é o valor da linha corrente na slice nessa chave do mapa que é coluna, caso algo falhe, retorne o erro.

Informe a query completa a ser executada para esta função.
*/
func List(query string) ([]map[string]string, error) {

	// pega as configuraçẽos do pacote, tenta abrir conexão com o banco, caso falhe retorne o erro se não continue.
	databaseSetting := settings.GetDatabaseSetting()
	database, err := openConnection(databaseSetting)
	if err != nil {
		log.Println("estabilish connection failed:", err)
		return nil, err
	}
	defer database.Close()
	log.Println("successfully estabilish connection!")

	// executa a query, caso falhe retorne o erro se não continue.
	table, err := database.Query(query)
	if err != nil {
		log.Println("query select error:", err, "query:", query)
		return nil, err
	}
	defer table.Close()
	log.Println("successfully query select!")

	// cria a variável que guardará a slice de rows/columns/values e que será retornada pelo função.
	var rowSliceMap []map[string]string

	// pega o nome das colunas da tabela, caso falhe retorne o erro se não continue.
	columnsNames, err := table.Columns()
	if err != nil {
		log.Println("get column names error:", err)
	}
	log.Println("successfully get columns names!")

	// cria a slice de interfaces que receberá os valores das linhas.
	rowsInterface := make([]interface{}, len(columnsNames))
	for index := range rowsInterface {
		var rowInterface interface{}
		rowsInterface[index] = &rowInterface
	}

	// a cada iteração popula a próxima informações com Next() para o Scan() e verifica se foi possível fazer a leitura dessa informação, cria o mapa que recebera o grupo de informações e apenda na slice principal.
	for table.Next() {
		err := table.Scan(rowsInterface...)
		if err != nil {
			log.Println("error on read row:", err)
		}
		tableMap := make(map[string]string)
		for index, columnName := range columnsNames {
			rowInterface := *(rowsInterface[index].(*interface{}))
			var rowValue string
			if rowByte, ok := rowInterface.([]byte); ok {
				rowValue = string(rowByte)
			} else {
				rowValue = "null"
			}
			tableMap[columnName] = rowValue
		}
		rowSliceMap = append(rowSliceMap, tableMap)
	}

	return rowSliceMap, nil

}
