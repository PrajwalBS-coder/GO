package main

import (
	f "fmt"
	r "reflect"
)

type person struct{ //here data type will be person here created variable using var keyword
	name string //here data type will be string here created variable using var keyword
	age int //here data type will be int here created variable using var keyword
}

func main() {
	var p person //here data type will be person here created variable using var keyword
	p.name = "John" //here data type will be string here created variable using var keyword
	p.age = 25 //here data type will be int here created variable using var keyword
	f.Println(p, r.TypeOf(p)) //here data type will be person here created variable using var keyword
	print(p)
}

func print(p person){
	f.Println(p.age, r.TypeOf(p.age))
	f.Println(p.name, r.TypeOf(p.name))
}