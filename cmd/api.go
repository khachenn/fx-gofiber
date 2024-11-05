package cmd

import (
	"sturdy-winner-api/internal/api"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "sturdy winner api",
	Run:   apiRun,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func apiRun(cmd *cobra.Command, _ []string) {
	log.Info(cmd.Short)
	fx.New(api.Module, fx.WithLogger(FxCmdLogger)).Run()
}
