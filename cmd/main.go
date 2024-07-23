package main

import (
	"github.com/spf13/cobra"
	tt "github.com/zurvan-lab/timetrace"
	"github.com/zurvan-lab/timetrace/cmd/commands"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "ttrace",
		Version: tt.StringVersion(),
	}

	commands.RunCommand(rootCmd)
	commands.ConnectCommand(rootCmd)
	commands.PingCommand(rootCmd)
	commands.InitCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		commands.ExitOnError(rootCmd, err)
	}
}
