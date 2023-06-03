package consts

import "github.com/go-hl/os/env"

// JWT environemnt variables.
var (
	JWTSECRETKEY = env.Get("JWTSECRETKEY", "")
	JWTTIME      = env.GetShouldInt("JWTTIME", -1)
)
