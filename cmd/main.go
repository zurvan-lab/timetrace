package main

import (
	"github.com/spf13/cobra"
	timetrace "github.com/zurvan-lab/TimeTrace"
	"github.com/zurvan-lab/TimeTrace/cmd/commands"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "ttrace",
		Version: timetrace.StringVersion(),
	}

	commands.RunCommand(rootCmd)
	commands.REPLCommand(rootCmd)
	commands.PingCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		commands.ExitOnError(rootCmd, err)
	}
}
