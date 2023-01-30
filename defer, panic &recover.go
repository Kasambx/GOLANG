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



// exxample 2 
package main

import (
	"fmt"
)

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
