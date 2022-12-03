package database

import (
	"golang/settings"
	"log"
)

func Save(tableName string, args ...string) (int, error) {

	database, err := openConnection(settings.GetDatabaseSetting())
	if err != nil {
		log.Println("estabilish connection failed:", err)
		return -1, err
	}
	defer database.Close()
	log.Println("successfully estabilish connection!")

	query := tableName + " "
	if len(args) < 1 {
		log.Fatal("save:more arguments necessary to this function!")
	} else if len(args) > 1 {
		query = "UPDATE " + query
	} else {
		query = "INSERT INTO " + query
	}
	for index := range args {
		query += args[index] + " "
	}

	result, err := database.Exec(query)
	if err != nil {
		log.Println("query exec failed:", err)
		return -1, err
	}
	log.Println("successfully query exec!")

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil

}
