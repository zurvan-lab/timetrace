package main

import (
	"os"

	"github.com/zurvan-lab/TimeTraceDB/core/database"
)

func main() {
	database := database.CreateDataBase(os.Args[0])
	database.InitSocket()
	database.InitUsers()
}
