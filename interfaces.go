//interfaces in go 


package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

