package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func processStream(from io.Reader, to io.Writer) error {
	buffer, err := ioutil.ReadAll(from)
	if err != nil {
		return err
	}
	to.Write(blackfriday.Run(buffer))
	return nil
}

func main() {
	from := bufio.NewReader(os.Stdin)
	to := os.Stdout
	err := processStream(from, to)
	if err != nil {
		panic(err)
	}
}
