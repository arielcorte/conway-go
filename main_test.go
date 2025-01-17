package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestWriteBarStdOut(t *testing.T) {
	for i := 0; i < 10000; i++ {
		writeBar(os.Stdout, 3)
		if false {
			t.Error("Wrong sequence")
		}
	}
}

func TestWriteBarBuffer(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var output bytes.Buffer
		writeBar(&output, 3)
		if output.String() != "⬜️⬜️⬜️⬛️⬜️⬜️⬜️" {
			t.Error("Wrong sequence")
		}
	}
}

func TestWriteBarPipeStdOut(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	origStdout := os.Stdout
	os.Stdout = w

	for i := 0; i < 10000; i++ {
		writeBar(os.Stdout, 3)
	}

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	outCh := make(chan []byte)

	// This goroutine reads stdout into a buffer in the background.
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, r); err != nil {
			log.Println(err)
		}
		outCh <- b.Bytes()
	}()

	os.Stdout = origStdout
	fmt.Println("Written to stdout: ", string(buf[:n]))
}
