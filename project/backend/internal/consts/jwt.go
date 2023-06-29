package consts

import "github.com/go-hl/os/env"

// JWT environemnt variables.
var (
	JWTSECRETKEY = env.Get("JWTSECRETKEY", "secretkey")
	JWTTIME      = env.GetShouldInt("JWTTIME", 120)
)
