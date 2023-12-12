package main

import (
	"github.com/spf13/cobra"
	timetrace "github.com/zurvan-lab/TimeTrace"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "time-trace",
		Version: timetrace.StringVersion(),
	}

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
