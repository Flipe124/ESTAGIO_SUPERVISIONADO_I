package database

import (
	"golang/commons"
	"golang/settings"
	"log"
)

func Save(tableName string, args ...string) int {

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

	query := tableName + " "
	if len(args) < 1 {
		log.Fatal("save: more arguments necessary to this function!")
	} else if len(args) > 1 {
		query = "UPDATE " + query
	} else {
		query = "INSERT INTO " + query
	}
	for index := range args {
		query += args[index] + " "
	}

	result, err := database.Exec(query)
	commons.ErrorTester("successfully query exec!", "query exec error: ", err)
	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected)

}
