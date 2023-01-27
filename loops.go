
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {    //first statement is the innititalizer, second the booloean producer and lastly the transfromation
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
