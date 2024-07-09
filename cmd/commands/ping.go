package commands

import (
	"fmt"
	"net"
	"os"

	cobra "github.com/spf13/cobra"
)

func PingCommand(parentCmd *cobra.Command) {
	ping := &cobra.Command{
		Use:   "ping",
		Short: "Ping a remote instance of time trace.",
	}
	parentCmd.AddCommand(ping)

	address := ping.Flags().StringP("address", "a", "localhost:7070", "remote address of your time trace instance.")
	username := ping.Flags().StringP("username", "u", "root", "username of the user you are going to connect with.")
	password := ping.Flags().StringP("password", "p", "", "password of user trying to connect with.")

	ping.Run = func(cmd *cobra.Command, args []string) {
		conn, err := net.Dial("tcp", *address)
		if err != nil {
			ExitOnError(cmd, err)
		}
		defer conn.Close()

		conQuery := fmt.Sprintf("CON %v %v", *username, *password)

		do(conn, conQuery)

		response := do(conn, "PING")
		if response == "PONG" {
			cmd.Println("PONG, everything is ok.")
			os.Exit(0)
		} else {
			cmd.Printf("something is wrong: %v", response)
			os.Exit(1)
		}
	}
}
