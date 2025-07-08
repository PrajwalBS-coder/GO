package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a string //here data type will be string and initial value will be empty string
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}