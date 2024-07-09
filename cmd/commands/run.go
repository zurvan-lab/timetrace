package commands

import (
	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/TimeTrace/config"
	"github.com/zurvan-lab/TimeTrace/core/database"
	"github.com/zurvan-lab/TimeTrace/core/server"
	tte "github.com/zurvan-lab/TimeTrace/utils/errors"
	ttlog "github.com/zurvan-lab/TimeTrace/utils/log"
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
			ExitOnError(cmd, tte.ErrInavlidConfigPath)
		}

		cfg, err := config.LoadFromFile(*confingPath)
		if err != nil {
			ExitOnError(cmd, err)
		}

		db := database.Init(cfg)
		ttlog.InitGlobalLogger(&cfg.Log)

		server := server.NewServer(cfg, db)
		if err := server.Start(); err != nil {
			ExitOnError(cmd, err)
		}
	}
}
