package database

import (
	"database/sql"
	"golang/commons"
	"golang/settings"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func openConnection(setting settings.Setting) *sql.DB {

	// username:password@(address)/dbname
	driver := setting.Database.Driver
	user := setting.Database.User
	password := setting.Database.Password
	host := setting.Database.Host
	port := strconv.Itoa(setting.Database.Port)
	name := setting.Database.Name

	connection :=  user + ":" + password + "@(" + host + ":" + port + ")/" + name
	database, err := sql.Open(driver, connection)
	commons.ErrorTester("successfully connection!", "open connection failed: ", err)

	err = database.Ping()
	commons.ErrorTester("successfully ping!", "ping connection failed: ", err)

	return database

}
