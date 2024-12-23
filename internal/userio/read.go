package userio

import (
	"bufio"
	"log"
	"os"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadInput() (string, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(line), nil
}
