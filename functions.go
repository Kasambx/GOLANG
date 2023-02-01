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


//variadic parameters

package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(greeting, name)

}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"		//if i change the name using pointers it will be modified it cannot change normally for it is declared  in another function
	fmt.Println(*name)
}



//variadric parameters 

package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5)		//made a generic sum function  
}

func sum(values ...int) {  //tells the go runtime to take all of the arguments into a slice caleed values

	fmt.Println(values)	//prints out all of the value variables
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}


// use of return statement
package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5) //made a generic sum function
}

func sum(values ...int) int { //tells the go runtime to take all of the arguments into a slice caleed values

	fmt.Println(values) //prints out all of the value variables
	result := 0
	for _, v := range values {
		result += v
	}
	return result		// returns the value of the variable result
}


// returning a pointer as a local variable
package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5) //made a generic sum function
}

func sum(values ...int) *int { //tells the go runtime to take all of the arguments into a slice caleed values

	fmt.Println(values) //prints out all of the value variables
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // returns the value of the variable result
}


//returning multiple return values form a function
package main

import (
	"fmt"
)

func main() {
	d := divide(5.0, 3.0)
	fmt.Println(d)
}

func divide(a, b float64) float64 {
	return a / b
}
