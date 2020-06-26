package main

import (
	"os"

	"go-account-service-client/application"
)

func main() {
	application.Run(os.Args[1])
}
