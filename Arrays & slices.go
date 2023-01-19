// You can edit this code!
package main

import (
	"fmt"
)

func main() {
	grade1 := 97
	grade2 := 93
	grade3 := 92
	fmt.Printf("Grades:%v,%v,%v", grade1, grade2, grade3)
}


//now for the array 

func main() {
  grades :=[...]int{97,85,93 }
  fmt.Printf("Grades: %v" , grades)
}

func main () {
	var students [3]string
	fmt.PrinfF("Students:%v", students)
}

func main() {
	var students [3]string
	fmt.Printf("Students: %v", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students: %v\n", students[2])
	fmt.Printf("No of Students: %v\n", len(students))
}



// making an array of an array
func main() {
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
}

func main() {
	var identifyMatrix [3][3]int
	identifyMatrix[0] = [3]int{1,0,0}
	identifyMatrix[1] = [3]int{0,1,0}
	identifyMatrix[2] = [3]int{0,0,1}
	fmt.Println(identifyMatrix)
}


func main() {
	a:= [...]int{1,2,3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}







