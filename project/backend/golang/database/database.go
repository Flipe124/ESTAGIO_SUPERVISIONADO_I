package database

import (
	"database/sql"
	"golang/settings"
	"log"
	"strconv"

	// username:password@(address)/dbname
	_ "github.com/go-sql-driver/mysql"
)

func openConnection(setting settings.DatabaseSetting) (*sql.DB, error) {

	driver := setting.Driver
	user := setting.User
	password := setting.Password
	host := setting.Host
	port := strconv.Itoa(setting.Port)
	name := setting.Name

	connection := user + ":" + password + "@(" + host + ":" + port + ")/" + name

	database, err := sql.Open(driver, connection)
	if err != nil {
		log.Println("open connection failed:", err)
		return nil, err
	}
	log.Println("successfully connection!")

	err = database.Ping()
	if err != nil {
		log.Println("ping connection failed:", err)
		return nil, err
	}
	log.Println("successfully ping!")

	return database, nil

}
