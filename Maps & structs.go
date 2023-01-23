package main

import "fmt"

func main() {
	      statePopulations :=make(map[string]int)  //one way to make a map
	      statePopulations = map[string]int{ 	//the second way to make a map 
	      	      "California":   39260017,
	      	      "Texas":        27862596,
	      	      "Florida":      20612439,
	      	      "New York":     19745289,
	      	      "Pennsylvania": 12802503,     
	      	      "Illinois":     12801539,     
	      	      "Ohio":         11614373,
	      }
	      statePopulations["Georgia"] = 10310371 //adding elements to a map
	      delete(statePopulations, "Georgia") //deleting elements
	      fmt.Println(statePopulations)
	      fmt.Println(statePopulations["New York"]) // getting a single value of a key      
  	      pop, ok := statePopulations["Georgia"]  //shows whether the element is present and its value
	      fmt.Println(pop, ok)
              fmt.Println(len(statePopulations))
	
//passing out the map as another variable and checking what happns if you delete one element
	      sp := statePopulations				
	      delete(sp, "Ohio")			
	      fmt.Println(sp)
	      fmt.Println(statePopulations)
}      


