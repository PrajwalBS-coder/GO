 package main

import f "fmt"
// import r "reflect"

func main(){
	var a int = 10;
	var b int = 20;
	if(a > b) {
		f.Println("A is greater than B")
	} else {
		f.Println("B is greater than A")
	}
	if(a > b) {
		f.Println("A is greater than B")
	} 
	else {//it'l throw error if we use else in the same line as closing bracket
		f.Println("B is greater than A")
	}


}