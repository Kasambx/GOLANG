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

//slices 
func main() {
	a:= [2]int{1,2,3}
	b:= a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n",cap(a))
	
}

//another way of creating a slice

func main() {
	a:= []int{1,2,3,4,5,6,7,8,9,10}
	b:= a[:] //slice of all elements
	c :=a[3:] //slice from 4th element to end
	d:=a[6:] //slice first 6 elements
	e:=a[3:6] //slice the 4th, 5th and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)	
}


//creating a slice using the make function
func main() {
	z:= make([]int, 3 , 100)
	fmt.Println(z)
	fmt.Printf("Length: %v\n", len(z))
	fmt.Printf("Capacity: %v\n",cap(z))
	
}



//slices size can be modified look on the code below 
func main(){
	a:= []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a,1)
	fmt.Println(a)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a,2,3,4,5,6)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, []int{2,3,4,5,6}...) // say if you want to concatenate arrays use "..." after the second string for the procedure to be successful 
}

func main() {
	a := []int{2,3,4,5,6}
	b := a[1: ]  //pops the first elemnent of the slice -- its called the shift operation 
	fmt.Println(b )
	c:= a[:len(a)-1] //removes the last element of the slice 
	fmt.Println(c)
	b := append(a[:2],a[3:]...)//removing the middle element by concatenating two slices   

}









