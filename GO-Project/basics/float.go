package main

import ("fmt"
"reflect")

func main(){
	var a float32 //here data type will be float 
	fmt.Println(a,reflect.TypeOf(a))
	b:=0.2
	fmt.Println(b,reflect.TypeOf(b))
}