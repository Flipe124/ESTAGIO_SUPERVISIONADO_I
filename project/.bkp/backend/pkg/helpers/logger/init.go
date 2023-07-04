package logger

import (
	"backend/internal/consts"
	"backend/pkg/utils/log"
)

// Log is the global logger.
var Log *log.Log

func init() {
	if consts.LOGGERMODE != string(log.Production) {
		Log = log.NewLog(log.Development)
	} else {
		Log = log.NewLog(log.Production)
	}
}
