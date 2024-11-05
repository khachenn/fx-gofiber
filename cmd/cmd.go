package cmd

import (
	"os"
	"sturdy-winner-api/config"

	"github.com/dollarsignteam/go-logger"
	"github.com/dollarsignteam/go-utils"
	"github.com/spf13/cobra"
	"go.uber.org/fx/fxevent"
)

var cfgFile string
var log *logger.Logger

var rootCmd = &cobra.Command{
	Use:     "sturdy-winner-api",
	Short:   "sturdy winner api",
	Version: config.Version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       utils.PackageName(),
		HideCaller: true,
	})

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is local.env)")
}

func initConfig() {
	config.AutoReadConfig(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func FxCmdLogger() fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log.Desugar()}
}
