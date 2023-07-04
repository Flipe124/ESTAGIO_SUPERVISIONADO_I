package consts

import "github.com/go-hl/os/env"

// Database environemnt variables.
var (
	DBUSER     = env.Get("DBUSER", "")
	DBPASSWORD = env.Get("DBPASSWORD", "")
	DBPROTOCOL = env.Get("DBPROTOCOL", "")
	DBHOST     = env.Get("DBHOST", "")
	DBPORT     = env.GetShouldInt("DBPORT", -1)
	DBNAME     = env.Get("DBNAME", "")
)
