package commands

import (
	_ "embed"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/peterh/liner"
	cobra "github.com/spf13/cobra"
	"github.com/zurvan-lab/TimeTrace/utils/errors"
)

/*
	TODO: Add history support for me.
	TODO: Add suggestion for me.
	TODO: Add completer for me.
*/

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

var (
	//go:embed ttrace.txt
	welcomeASCII []byte

	TQL_COMMANDS = [...]string{
		"CON", "PING", "SET", "SSET", "PUSH", "GET", "CNTS",
		"CNTSS", "CNTE", "CLN", "CLNS", "CLNSS", "DRPS", "DRPSS", "SS",
	}

	clear map[string]func()
)

const (
	PROMPT = ">> "
)

func ConnectCommand(parentCmd *cobra.Command) {
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

		lnr := liner.NewLiner()
		defer lnr.Close()

		lnr.SetCtrlCAborts(true)

		conQuery := fmt.Sprintf("CON %v %v", *username, *password)
		response := do(conn, conQuery)

		if response == "OK" {
			cleanTerminal()
			cmd.Println(string(welcomeASCII))

			for {
				if input, _ := lnr.Prompt(PROMPT); err == nil {
					if input == "exit" {
						os.Exit(0)
					}

					if input == "clean" {
						cleanTerminal()
					}

					lnr.AppendHistory(input)
					cmd.Print(fmt.Sprintf("%s\n", do(conn, input)))
				}
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

func cleanTerminal() {
	cf, ok := clear[runtime.GOOS]
	if ok {
		cf()
	}
}
