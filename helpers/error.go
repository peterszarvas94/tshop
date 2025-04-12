package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func HandleError(message string, err error, exitCode int) {
	fmt.Println(message)

	defer os.Exit(exitCode)

	var errorMessage string
	if err == nil {
		fmt.Println("Error is nil")
		return
	}

	errStr := err.Error()
	start := strings.Index(errStr, "{")
	if start == -1 {
		fmt.Println("Badly formatted error message")
		return
	}

	jsonPart := errStr[start:]

	var data map[string]any
	if jsonErr := json.Unmarshal([]byte(jsonPart), &data); jsonErr != nil {
		fmt.Println("Error parsing error message")
		return
	}

	errorMessage, ok := data["message"].(string)
	if !ok || errorMessage == "" {
		fmt.Println("Could not find error message")
		return
	}

	fmt.Println(errorMessage)
}
