package db

import (
	"fmt"
	"os"

	"backend/internal/consts"
	"backend/internal/models"
	"backend/pkg/helpers/logger"

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

	if Tx.AutoMigrate(
		&models.User{},
		&models.Status{},
	) != nil {
		logger.Log.Fatal("can't database auto migration!")
	}

	if postMigration, err := os.ReadFile("scripts/db/post-migration.sql"); err != nil {
		logger.Log.Warn("error on read post migration file:", err.Error())
	} else if err := Tx.Exec(string(postMigration)).Error; err != nil {
		logger.Log.Warn("error on execute post migration script:", err.Error())
	}

}
