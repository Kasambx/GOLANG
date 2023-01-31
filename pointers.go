//pointers

package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a     // * points b to address of a
	fmt.Println(&a, *b) // * pulls the value in the address of B
	a = 30				//value of b changes after changing value of a
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)		//value of a changes after changing value of b
}


