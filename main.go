package main

import (
	"fmt"
	api "go-fun/pkgs/api"
	"os"
)

func main() {
	dbConnStr := os.Getenv("SAMPLE_DATABASE_CONNECTION_STRING")
	if dbConnStr == "" {
		panic("SAMPLE_DATABASE_CONNECTION_STRING environment variable not set")
	}

	api.StartServer(dbConnStr)
	fmt.Println("END OF PROGRAM")
}
