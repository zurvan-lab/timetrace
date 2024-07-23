package commands

import (
	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/timetrace/config"
	"github.com/zurvan-lab/timetrace/core/database"
	"github.com/zurvan-lab/timetrace/core/server"
	ttlog "github.com/zurvan-lab/timetrace/log"
)

func RunCommand(parentCmd *cobra.Command) {
	run := &cobra.Command{
		Use:   "run",
		Short: "Runs an instance of time trace.",
	}
	parentCmd.AddCommand(run)

	configPath := run.Flags().StringP("config", "c", "", "Path to your config.yaml file.")

	run.Run = func(cmd *cobra.Command, args []string) {
		if configPath == nil || *configPath == "" {
			ExitOnError(cmd, InvalidConfigPathError{
				path: *configPath,
			})
		}

		cfg, err := config.LoadFromFile(*configPath)
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
