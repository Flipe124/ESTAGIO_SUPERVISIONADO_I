package consts

import "github.com/go-hl/os/env"

// APIPORT is the port of API.
var APIPORT = env.GetShouldInt("APIPORT", 9999)
