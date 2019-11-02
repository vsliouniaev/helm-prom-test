package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vsliouniaev/helm-prom-test/core"
	"regexp"
	"runtime"
)

var strip = regexp.MustCompile(fmt.Sprintf(".*%s", core.Module))

func ConfigureLogging(logLevel, logFmt string, reportCaller bool) (log.Level, log.Formatter) {
	l, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("err", err).Fatal("Invalid error level")
	}
	log.SetLevel(l)
	log.SetReportCaller(reportCaller)
	f := getFormatter(logFmt)
	log.SetFormatter(f)
	return l, f
}

func getFormatter(logfmt string) log.Formatter {
	switch logfmt {
	case "json":
		return &log.JSONFormatter{CallerPrettyfier: callerPrettyfier}
	case "text":
		return &log.TextFormatter{CallerPrettyfier: callerPrettyfier}
	}

	log.Fatalf("invalid log format '%s'", logfmt)
	return nil
}

func callerPrettyfier(f *runtime.Frame) (function, file string) {
	return "", fmt.Sprintf("%s:%d", strip.ReplaceAllString(f.File, ""), f.Line)
}
