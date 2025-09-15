package main

import (
	f "fmt"
	r "reflect"
)

// Naming Rules for Go Functions
// A function name must start with a letter
// A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
// Function names are case-sensitive
// A function name cannot contain spaces
// If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

func add(a int, b int) int { //here data type will be int here created variable using var keyword	
	return a + b //here data type will be int here created variable using var keyword
}

func describe(name string, age int) { //here data type will be int here created variable using var keyword	
	f.Println(name, "is", age, "years old") //here data type will be int here created variable using var keyword
}
func main(){
	var a int = 10 //here data type will be int here created variable using var keyword
	f.Println(a, r.TypeOf(a))
	var b int = 20 //here data type will be int here created variable using var keyword
	f.Println(b, r.TypeOf(b))
	var sum int = add(a, b) //here data type will be int here created variable using var keyword
	f.Println(sum, r.TypeOf(sum))
	f.Println(add(100,-90))
	(describe("John", 25))


}