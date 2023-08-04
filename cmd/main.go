package main

import (
	"os"

	"github.com/zurvan-lab/TimeTraceDB/src"
)

func main() {
	database := src.CreateDataBase(os.Args[1])
	database.InitSocket()
	database.InitUsers()
}
