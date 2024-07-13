package commands

import (
	"net"
	"os"

	"github.com/spf13/cobra"
)

func ExitOnError(cmd *cobra.Command, err error) {
	cmd.PrintErrln(err)
	os.Exit(1)
}

func do(conn net.Conn, q string) string {
	resBuf := make([]byte, 1024)
	query := []byte(q)

	if len(query) < 1 {
		return "INVALID"
	}

	_, err := conn.Write(query)
	if err != nil {
		return err.Error()
	}

	n, err := conn.Read(resBuf)
	if err != nil {
		return err.Error()
	}

	return string(resBuf[:n])
}
