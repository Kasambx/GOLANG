// go routines

package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()                      // the go is used to make a goroutine //n/b nothing will be displayed after printing
	time.Sleep(100 * time.Millisecond) // using time sleep to put some delay so that the print statement can be displayed
}

func sayHello() {
	fmt.Println("Hello")
}



// go routines
// you can also use a wait group instead of sleep
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // you can also use a wait group instead of sleep

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}




//GOMAXPROCS -- shows the number of threads used
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) //shows the number of threads -1 shows the no of threads that are used
}
