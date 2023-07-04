package consts

import (
	"backend/pkg/utils/log"

	"github.com/go-hl/os/env"
)

// LOGGERMODE is the logger mode environment.
var LOGGERMODE = env.Get("LOGGERMODE", string(log.Development))
