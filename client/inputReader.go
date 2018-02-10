package client

import (
	"bufio"
	"fmt"
)

const prompt = "[Message]: "

type InputReader struct {
	scanner *bufio.Scanner
}

func NewInputReader(scanner *bufio.Scanner) *InputReader {
	return &InputReader{scanner: scanner}
}

func (reader *InputReader) run(channel chan<- string) {
	fmt.Print(prompt)
	for reader.scanner.Scan() {
		channel <- reader.scanner.Text()
		fmt.Print(prompt)
	}
}