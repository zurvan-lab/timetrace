package commands

import (
	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/TimeTrace/config"
	"github.com/zurvan-lab/TimeTrace/core/database"
	tte "github.com/zurvan-lab/TimeTrace/utils/errors"
)

func RunCommand(parentCmd *cobra.Command) {
	run := &cobra.Command{
		Use:   "run",
		Short: "Runs an instance of time trace.",
	}
	parentCmd.AddCommand(run)

	confingPath := run.Flags().StringP("config", "c", "", "Path to your config.yaml file.")

	run.Run = func(cmd *cobra.Command, args []string) {
		if confingPath == nil || *confingPath == "" {
			dead(cmd, tte.ErrInavlidConfigPath)
		}

		cfg, err := config.LoadFromFile(*confingPath)
		if err != nil {
			dead(cmd, err)
		}

		_ = database.Init(cfg)
	}
}
