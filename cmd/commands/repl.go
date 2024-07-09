package commands

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/TimeTrace/utils/errors"
)

const PROMPT = "\n>> "

func REPLCommand(parentCmd *cobra.Command) {
	connect := &cobra.Command{
		Use:   "connect",
		Short: "Connects you to a time trace instance and you can interact with it in a REPL interface.",
	}
	parentCmd.AddCommand(connect)

	address := connect.Flags().StringP("address", "a", "localhost:7070", "remote address of your time trace instance.")
	username := connect.Flags().StringP("username", "u", "root", "username of the user you are going to connect with.")
	password := connect.Flags().StringP("password", "p", "", "password of user trying to connect with.")

	connect.Run = func(cmd *cobra.Command, args []string) {
		conn, err := net.Dial("tcp", *address)
		if err != nil {
			ExitOnError(cmd, err)
		}
		defer conn.Close()

		conQuery := fmt.Sprintf("CON %v %v", *username, *password)

		response := do(conn, conQuery)
		if response == "OK" {
			reader := bufio.NewReader(os.Stdin)

			for {
				fmt.Print(PROMPT)

				input, _ := reader.ReadString('\n')
				input = strings.TrimSuffix(input, "\n")

				if input == "exit" {
					break
				}

				cmd.Print(do(conn, input))
			}
		} else {
			ExitOnError(cmd, fmt.Errorf("%w: %s", errors.ErrInvalidCommand, response))
		}
	}
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

	time.Sleep(time.Second * 1)

	n, err := conn.Read(resBuf)
	if err != nil {
		return err.Error()
	}

	return string(resBuf[:n])
}
