package main

import (
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	lastTime := time.Now()
	black := 0
	currFrame := 0

	for {
		if time.Until(lastTime) < time.Duration(-30)*time.Millisecond {
			currFrame++
			err := clearScreen()
			if err != nil {
				panic("Clear command only available on linux.")
			}

			writeBar(os.Stdout, black)

			if black > 5 {
				black = 0
			} else {
				black++
			}
			lastTime = time.Now()
		}
	}
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func writeBar(w io.Writer, black int) {
	white_square := []byte("█")
	black_square := []byte(" ")
	new_line := []byte("\n")

	for i := black; i > 0; i-- {
		w.Write(white_square)
	}
	w.Write(black_square)
	for i := 6 - black; i > 0; i-- {
		w.Write(white_square)
	}
	w.Write(new_line)
}
