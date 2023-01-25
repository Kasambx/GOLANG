// simple is statement
package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}
}

//using if and pop
package main

import "fmt"

func main() {
	statePopulations := make(map[string]int) //one way to make a map
	statePopulations = map[string]int{       //the second way to make a map
		"California":   39260017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	if pop, ok := statePopulations["New York"]; ok { //shows whether the element is present and its value
		fmt.Println(pop)
	}

}

//another example
package main

import "fmt"

func main() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too Low")

	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it !")
	}
	fmt.Println(number<=guess, number>= guess, number!=guess)  //returns boolean expressions for the variable in the func mamin statement
}


package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 370
	if guess < 1 || guess > 100 {		// || id the or operator command is executed if one of the requiements is fulfiled
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {		// && its the and operator both conditions have to be met for the command to be executed    // you can also replace this line with else{} operator 
		if guess < number {
			fmt.Println("Too Low")

		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it !")
		}
		fmt.Println(number <= guess, number >= guess, number != guess) //returns boolean expressions for the variable in the func mamin statement
	}
}



				//SWITCH STATEMENTS 
package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("not one ot two")
	}
}


// when dealing with more values 
package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1, 2, 3:
		fmt.Println("One, two or three")
	case 4, 5, 6:
		fmt.Println("four, five or six")
	default:
		fmt.Println("not one ot two")
	}
}

// another example 
package main

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
	case 1, 2, 3:
		fmt.Println("One, two or three")
	case 4, 5, 6:
		fmt.Println("four, five or six")
	default:
		fmt.Println("not one ot two")
	}
}



//tagless switches

package main

import (
	"fmt"
)

func main() {
	i := 8
	switch {
	case i <= 10:
		fmt.Println("Less than 10")
	case i <= 20:
		fmt.Println("Less than 20")
	default:
		fmt.Println("Beyond or under capacity ")
	}
}



//type switch 
// it finds the type of the variable 
package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}



package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1.09
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break 				// exits out of the switch statements early otice only the first line is printed 
		fmt.Println("This is true ")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}






