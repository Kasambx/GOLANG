//basic function in go 
package main

import "fmt"


func main () {
  fmt.Println("Hey I am the best")
}


// defining a function that takes parameters from another function
package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello Go!")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}




