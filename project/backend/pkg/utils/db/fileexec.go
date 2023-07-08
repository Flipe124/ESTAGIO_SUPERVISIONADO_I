package db

import (
	"os"

	"gorm.io/gorm"
)

// FileExec execute a file with data raw sql.
func FileExec(tx *gorm.DB, path string) (err error) {
	preMigration, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = tx.Exec(string(preMigration)).Error
	return
}
