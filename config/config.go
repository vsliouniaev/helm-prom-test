package config

import (
	log "github.com/sirupsen/logrus"
)

type Runtime struct {
	LogLevel  log.Level
	LogFmt    log.Formatter
	LogCaller bool
}
