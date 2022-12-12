package bd

import "server/adrm"

// DB é a representação do banco para ser usada no projeto.
var DB = adrm.NewDatabase(
	"root",
	"root",
	"172.17.0.2",
	"openfinance",
	3306,
)
