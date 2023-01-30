// defer function in Go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")			//defer mves this statement before the main function is returned
	fmt.Println("End")
}
