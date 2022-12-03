package database

import (
	"golang/commons"
	"golang/settings"
	"log"
)

func List(tableName string, args ...string) []map[string]string {

	setting := settings.Setting{
		Database: settings.DatabaseSetting{
			Driver:   "mysql",
			User:     "root",
			Password: "root",
			Host:     "127.17.0.2",
			Port:     3306,
			Name:     "db",
		},
	}

	database := openConnection(setting)
	defer database.Close()

	query := "SELECT * FROM " + tableName
	if len(args) > 0 {
		for index := range args {
			query += " " + args[index]
		}
	}

	table, err := database.Query(query)
	commons.ErrorTester("successfully query select!", "query select error: ", err)
	defer table.Close()

	var rowSliceMap []map[string]string

	columnsNames, err := table.Columns()
	commons.ErrorTester("successfully get columns names!", "get column names error: ", err)

	rowsInterface := make([]interface{}, len(columnsNames))
	for index := range rowsInterface {
		var rowInterface interface{}
		rowsInterface[index] = &rowInterface
	}

	for table.Next() {
		err := table.Scan(rowsInterface...)
		if err != nil {
			log.Println("error on read row: ", err.Error())
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

	return rowSliceMap

}
