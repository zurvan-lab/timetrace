package main

import (

	"github.com/zurvan-lab/TimeTraceDB/src"
)


func main()  {
	database:= src.CreateDataBase()
	database.InitSocket()
	database.InitUsers()
}