//interfaces in go 
//using struct

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}


//single method interface

package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}


//interfaces in interfaces ***** do some research*****
// type conversions in interfaces 
// do some reseaarch on empty interfaces 


//type switches 
//interface switching
//interfaces in go
//interface switching

package main

import (
	"fmt"
)

func main() {
	var i interface{} = "weuwe"
	switch i.(type) { //this is called a type switch
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I dont know what i is")
	}
}
