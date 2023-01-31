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


// creating pointers to objects 
package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}



// note the same also happens in maps 
package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

