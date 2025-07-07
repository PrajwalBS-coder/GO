package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a int //here data type will be int and initial value will be 0 singned integer can store -ve value
	fmt.Println(a)
	var b uint //here data type will be unsigned int and initial value will be 0 cannnot store -ve values
	fmt.Println(b)//we can store 8,16,32,64 bit values
	var c byte //here data type will be byte and initial value will be 0
	fmt.Println(c,reflect.TypeOf(c))//we can store 8 bit values
	c=100
	fmt.Println(c,reflect.TypeOf(c))
}