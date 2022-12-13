package bd

import (
	"log"

	"github.com/rhuan-pk/pkutils/golang/adrm"
)

// DB é a representação do banco para ser usada no projeto.
var DB *adrm.Database

func init() {

	var err error
	DB, err = adrm.NewDatabase("root", "root", "172.17.0.2", "openfinance", 3306)
	if err != nil {
		log.Fatal("can't establish database connection:", err)
	}

}
