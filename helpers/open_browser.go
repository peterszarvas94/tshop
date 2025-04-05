package helpers

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "rundll32"
		args = append(args, "url.dll,FileProtocolHandler")
	default:
		cmd = "xdg-open"
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
