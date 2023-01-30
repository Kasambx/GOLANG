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


// the following program prints out panic which shows that program can no longer continue to execute 
// panic shows that the program cant continue to execute 
package main

import (
	"fmt"
)

func main() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}


//another example 
// the following program prints out panic which shows that program can no longer continue to execute
// panic shows that the program cant continue to execute
package main

import (
	"fmt"
)

func main() {
	fmt.Println("I'll try")
	panic("that's not enough ")
	fmt.Println("I'll do my best")
}


// a simple webserver using go  lang andthe panic function 
package main

import (
	"fmt"
)

func main() {
	http.handleFUnc("/", func(w http.ResponseWriter, r *http.Requsest) {
		w.Write([]byte("Hello Go !"))
	})
	err := http.ListenAndServer(":8080", nil)
	if err != nil {
		panic(err.Error())
	
	}
}
