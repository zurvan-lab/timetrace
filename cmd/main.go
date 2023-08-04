package main

import (
	"fmt"
	"os"

	"github.com/zurvan-lab/TimeTraceDB/src"
)

func main() {
	database := src.CreateDataBase(os.Args[1])
	database.InitSocket()
	database.InitUsers()

	fmt.Println(database.Users)
}
