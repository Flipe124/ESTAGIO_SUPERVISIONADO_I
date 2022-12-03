package database

import (
	"golang/settings"
	"log"
)

func List(tableName string, args ...string) ([]map[string]string, error) {

	databaseSetting := settings.GetDatabaseSetting()

	database, err := openConnection(databaseSetting)
	if err != nil {
		log.Println("estabilish connection failed:", err)
		return nil, err
	}
	defer database.Close()
	log.Println("successfully estabilish connection!")

	query := "SELECT * FROM " + databaseSetting.Name + "." + tableName
	for index := range args {
		query += " " + args[index]
	}
	
	table, err := database.Query(query)
	if err != nil {
		log.Println("query select error:", err)
		return nil, err
	}
	defer table.Close()
	log.Println("successfully query select!")

	var rowSliceMap []map[string]string

	columnsNames, err := table.Columns()
	if err != nil {
		log.Println("get column names error:", err)
	}
	log.Println("successfully get columns names!")

	rowsInterface := make([]interface{}, len(columnsNames))
	for index := range rowsInterface {
		var rowInterface interface{}
		rowsInterface[index] = &rowInterface
	}

	for table.Next() {
		err := table.Scan(rowsInterface...)
		if err != nil {
			log.Println("error on read row:", err.Error())
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
