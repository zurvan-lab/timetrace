package main

import (
	"os"

	"github.com/zurvan-lab/TimeTrace/core/database"
)

func main() {
	database := database.Init(os.Args[0])
	database.AddSet("test")
}
