package db

import (
	"fmt"

	"backend/internal/consts"
	"backend/internal/models"
	"backend/pkg/helpers/logger"

	// <user>:<password>@<protocol>(<host>:<port>)/<database>[?parseTime=true]
	// root:root@tcp(localhost:3306)/test
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Instance is the singleton instance of database.
var Instance *gorm.DB

func init() {

	var databaseError error

	stringConnection := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?parseTime=true",
		consts.DBUSER,
		consts.DBPASSWORD,
		consts.DBPROTOCOL,
		consts.DBHOST,
		consts.DBPORT,
		consts.DBNAME,
	)

	Instance, databaseError = gorm.Open(mysql.Open(stringConnection), &gorm.Config{})
	if databaseError != nil {
		logger.Log.Fatal("can't establish database connection!")
	}

	if Instance.AutoMigrate(
		&models.Client{},
		&models.Address{},
		&models.Permission{},
		&models.User{},
		&models.Service{},
		&models.Order{},
	) != nil {
		logger.Log.Fatal("can't database auto migration!")
	}

}
