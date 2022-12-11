package bd

import "server/adrm"

var DB = adrm.NewDatabase(
	"root",
	"root",
	"172.17.0.2",
	"openfinance",
	3306,
)
