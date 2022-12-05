package main

import (
	"os"

	"github.com/Digital-Voting-Team/menu-service/internal/cli"
)

func main() {
	// Uncomment the following line to run in the IDE
	// Do not forget to change db url in config.yaml!
	os.Setenv("KV_VIPER_FILE", "config.yaml")
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
