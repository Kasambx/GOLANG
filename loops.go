
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {    //first statement is the innititalizer, second the booloean producer and lastly the transfromation
		fmt.Println(i)
	}
}



// alternative for the code above 
package main

import "fmt"

func main() {
	i:= 0
	for ; i < 5; i++ {    //first statement is the innititalizer, second the booloean producer and lastly the transfromation
		fmt.Println(i)
	}
}





package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)

		if i%2 == 0 {     //check whether i is a multiple of 2
			i /= 2          // if so its divided by two
		} else {
			i = 2*i + 1   //if  not its multiplied by two and one is added 
		}
	}
}



//use of the break statement in looping
package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}


// nested loops 
package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}
}



// collections

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}


// looping on maps 

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
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

}

// looping on strings 
// n/b it will produce unicode representations of the string 
package main

import "fmt"

func main() {
	s := "Hello go !"
	for k, y := range s {
		fmt.Println(k, y)
		fmt.Println(k, string(y))
	}
}




