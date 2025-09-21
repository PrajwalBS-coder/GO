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
func addd(a, b int) int { //here data type will be int here created variable using var keyword	
	return a + b //here data type will be int here created variable using var keyword	
}

//named return
func des(a,b int)(result int) { //here data type will be int here created variable using var keyword	
	result = a + b //here data type will be int here created variable using var keyword
	return
}
func addmul(a,b int)(result int, mul int) { //here data type will be int here created variable using var keyword	
	result = a + b //here data type will be int here created variable using var keyword
	mul = a * b //here data type will be int here created variable using var keyword
	return
}

func main(){
	// var a int = 10 //here data type will be int here created variable using var keyword
	// f.Println(a, r.TypeOf(a))
	// var b int = 20 //here data type will be int here created variable using var keyword
	// f.Println(b, r.TypeOf(b))
	// var sum int = add(a, b) //here data type will be int here created variable using var keyword
	// f.Println(sum, r.TypeOf(sum))
	// f.Println(add(100,-90))
	// (describe("John", 25))
	// f.Println(addd(100,-90))
	// f.Println(des(100,-90))
	// ad,ml:=addmul(10,20)
	// f.Println(ad,ml)
	// testcount(1)
	// f.Println(factorial_recursion(5))
	f.Println(rec(4),r.TypeOf(rec(4)))
}
func testcount(x int) int {
  if x == 11 {
    return 0
  }
  f.Println(x)
  return testcount(x + 1)
}
//Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func rec(x int)int{//Factorial recursion
	if x==0{
		return 1
	}
	
	return x*rec(x-1)

}


