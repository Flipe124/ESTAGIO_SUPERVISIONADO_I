package consts

import "github.com/go-hl/os/env"

// Database environemnt variables.
var (
	DBUSER     = env.Get("DBUSER", "user")
	DBPASSWORD = env.Get("DBPASSWORD", "root")
	DBPROTOCOL = env.Get("DBPROTOCOL", "tcp")
	DBHOST     = env.Get("DBHOST", "database")
	DBPORT     = env.GetShouldInt("DBPORT", 3306)
	DBNAME     = env.Get("DBNAME", "database")
)
