package database

import (
	"database/sql"
	"golang/settings"
	"log"
	"strconv"

	// username:password@(address)/databasename
	_ "github.com/go-sql-driver/mysql"
)

// openConnection abre uma conexão com o banco e valida a mesma, caso tudo sucesse retorna a representação do banco, caso contrário retorne o erro referente.
func openConnection(setting settings.DatabaseSetting) (*sql.DB, error) {

	// variáveis que guarda as informações das configurações de credencial.
	driver := setting.Driver
	user := setting.User
	password := setting.Password
	host := setting.Host
	port := strconv.Itoa(setting.Port)
	name := setting.Name

	// string de conexão.
	connection := user + ":" + password + "@(" + host + ":" + port + ")/" + name

	// pega a representação do banco e valida erros de credênciais.
	database, err := sql.Open(driver, connection)
	if err != nil {
		log.Println("open connection failed:", err)
		return nil, err
	}
	log.Println("successfully connection!")

	// pega a conexão e valida erros da mesma.
	err = database.Ping()
	if err != nil {
		log.Println("ping connection failed:", err)
		return nil, err
	}
	log.Println("successfully ping!")

	// caso tudo sucesse, retorna o banco e nil.
	return database, nil

}
