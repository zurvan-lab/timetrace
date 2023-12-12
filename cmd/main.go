package main

import (
	"github.com/spf13/cobra"
	timetrace "github.com/zurvan-lab/TimeTrace"
	"github.com/zurvan-lab/TimeTrace/cmd/commands"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "time-trace",
		Version: timetrace.StringVersion(),
	}

	commands.RunCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
