package main

import ("fmt"
		"reflect")

func main(){
	var a int;
	fmt.Println("Hello")
	fmt.Println("Type of variable",reflect.TypeOf(a))
}