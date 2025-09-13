package main

import (
	f "fmt"
	r "reflect"
)

func main(){
	//simple for loop
	for i := 0; i < 10; i++ { //here data type will be int here created variable using var keyword
		f.Println(i, r.TypeOf(i))
	}


	//for loop with break
	for o:=1; o<=10;o++{
		if o%3==0{
			break//Here it will break the loop if the condition is true
		}
		f.Println(o, r.TypeOf(o))
	}


	//for loop with continue
	for o:=1; o<=10;o++{
		if o%3==0{
			continue//Here it will continue the loop if the condition is true
		}
		f.Println(o, r.TypeOf(o))
	}


	//Nested loop
	for i := 0; i < 1; i++ { //here data type will be int here created variable using var keyword
		for j := 0; j < 10; j++ { //here data type will be int here created variable using var keyword
			f.Println(i, r.TypeOf(i),j, r.TypeOf(j))
		}
	}


	//for loop with range
	p:=[3]string{"Apple","Mango","Bannana"} //here data type will be [3]string here created variable using var keyword
	for idx,val :=range(p){
		f.Println(idx,val)
	}
}