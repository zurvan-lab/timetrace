package commands

import (
	"fmt"
	"os"

	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/TimeTrace/config"
)

func InitCommand(parentCmd *cobra.Command) {
	init := &cobra.Command{
		Use:   "init",
		Short: "Init makes a default config file for you to run time-trace.",
	}
	parentCmd.AddCommand(init)

	path := init.Flags().StringP("path", "p", "./config.yml", "config file path")

	init.Run = func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(*path)
		if !os.IsNotExist(err) {
			ExitOnError(cmd, fmt.Errorf("a config file already exist on %s path", *path))
		}

		cfg := config.DefaultConfig()

		ymlCfg, err := cfg.ToYAML()
		if err != nil {
			ExitOnError(cmd, err)
		}

		if err = os.WriteFile(*path, ymlCfg, 0o600); err != nil {
			ExitOnError(cmd, err)
		}

		cmd.Printf("Config created at %s successfully!\n", *path)
	}
}
