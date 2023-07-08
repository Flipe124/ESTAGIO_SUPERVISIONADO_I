package db

import (
	"fmt"

	"backend/internal/consts"
	"backend/internal/models"
	"backend/pkg/helpers/logger"
	"backend/pkg/utils/db"

	// <user>:<password>@<protocol>(<host>:<port>)/<database>[?param=value&param=value]
	// root:root@tcp(localhost:3306)/test
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Tx is the singleton instance of database.
var Tx *gorm.DB

func init() {

	var databaseError error

	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?multiStatements=true&parseTime=true&loc=Local",
		consts.DBUSER,
		consts.DBPASSWORD,
		consts.DBPROTOCOL,
		consts.DBHOST,
		consts.DBPORT,
		consts.DBNAME,
	)

	Tx, databaseError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if databaseError != nil {
		logger.Log.Fatal("can't establish database connection!")
	}

	if err := db.FileExec(Tx, "scripts/db/pre-migration.sql"); err != nil {
		logger.Log.Fatal("error in pre migration:", err.Error())
	}

	if Tx.AutoMigrate(
		&models.Type{},
		&models.Status{},
		&models.User{},
		&models.Account{},
		&models.Category{},
		&models.Finance{},
		&models.Transaction{},
	) != nil {
		logger.Log.Fatal("can't database auto migration!")
	}

	if err := db.FileExec(Tx, "scripts/db/post-migration.sql"); err != nil {
		logger.Log.Fatal("error in post migration:", err.Error())
	}

}
