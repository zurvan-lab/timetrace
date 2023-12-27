package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func Dead(cmd *cobra.Command, err error) {
	cmd.PrintErrln(err)
	os.Exit(1)
}
