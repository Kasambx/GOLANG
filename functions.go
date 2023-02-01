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


//functions with loops

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}


// function that prints out the values of other functions
package main

import (
	"fmt"
)

func main() {
	sayGreeting("Hello", " Kasamba!")

}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}



