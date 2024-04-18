package main

import (
	"io"
	"os"

	"golang.org/x/term"
)

func main() {
	// fd := int(os.Stdout.Fd())

	// term.MakeRaw(fd)

	term := term.NewTerminal(io.ReadWriter(os.Stdout), ">")

	term.Write([]byte("1\n2\n"))
}
