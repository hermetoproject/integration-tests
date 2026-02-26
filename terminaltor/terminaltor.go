package terminaltor

import (
	"fmt"
	"os"

	"github.com/moby/term"
)

func Terminalte() (string, error) {
	fd := os.Stdin.Fd()
	if !term.IsTerminal(fd) {
		return "Not a terminal.", nil
	}

	ws, err := term.GetWinsize(fd)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Terminal with width=%d, height=%d", ws.Width, ws.Height), nil
}
