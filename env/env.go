package env

import (
	"fmt"
	"os"
)

type ConfigT struct {
	TERMINAL_TOKEN_ID string
	TERMINAL_TOKEN    string
	TERMINAL_ENV      string
}

var Config ConfigT

func init() {
	err := Load(&Config)
	if err != nil {
		// One env var is missing
		fmt.Println(err)
		os.Exit(1)
	}

	if Config.TERMINAL_ENV != "dev" && Config.TERMINAL_ENV != "prod" {
		fmt.Printf("Environment variable should be either \"dev\" or \"prod\"")
		os.Exit(1)
	}
}
