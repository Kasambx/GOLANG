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





