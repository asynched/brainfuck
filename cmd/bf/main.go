package main

import (
	"io"
	"os"

	"github.com/asynched/brainfuck/internal"
)

func main() {
	var input io.Reader

	// When an argument is provided, read from a file
	// instead of using stdin.
	if len(os.Args) > 1 {
		file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		input = file
	} else {
		input = os.Stdin
	}

	contents, err := io.ReadAll(input)

	if err != nil {
		panic(err)
	}

	mem := internal.NewMemory()
	parser := internal.NewParser(contents)

	for _, cmd := range parser.Parse() {
		cmd.Run(mem)
	}
}
