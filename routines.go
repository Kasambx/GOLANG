// go routines

package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()                      // the go is used to make a goroutine //n/b nothing will be displayed after printing
	time.Sleep(100 * time.Millisecond) // using time sleep to put some delay so that the print statement can be displayed
}

func sayHello() {
	fmt.Println("Hello")
}
