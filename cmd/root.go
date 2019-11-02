package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vsliouniaev/helm-prom-test/config"
	"github.com/vsliouniaev/helm-prom-test/core"
	"github.com/vsliouniaev/helm-prom-test/util"
	"os"
	"runtime"
)

var (
	rootCmd = &cobra.Command{
		Use:     "helm-prom-test",
		Short:   "Test that a prometheus install has worked",
		PreRun:  configureLogging,
		Run:     rootCommand,
		Version: core.Version,
	}
	cfg                = config.Runtime{}
	insecureSkipVerify bool
	logLevel           string
	logFmt             string
	logCaller          bool
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	rootCmd.Flags()
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "trace", "Log level: panic|fatal|error|warn|info|debug|trace")
	rootCmd.PersistentFlags().StringVar(&logFmt, "log-format", "text", "Log format: text|json")
	rootCmd.PersistentFlags().BoolVar(&logCaller, "log-caller", false, `Add file and line information to logs like 'file="prometheus/spotlight/noopApi.go:35'`)
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecure-skip-verify", false, "Skip validating certificates")
}

func configureLogging(_ *cobra.Command, _ []string) {
	cfg.LogCaller = logCaller
	cfg.LogLevel, cfg.LogFmt = util.ConfigureLogging(logLevel, logFmt, logCaller)
}

func rootCommand(cmd *cobra.Command, _ []string) {
	log.Printf("%s %s %s\n", core.Version, core.BuildTime, runtime.Version())
}
