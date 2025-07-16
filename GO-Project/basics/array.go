package main

import (
	f "fmt" 
	r "reflect"
)
func main(){
	var array=[]int{1,2}//here data type will be []int here created variable using var keyword
	f.Println(array,r.TypeOf(array))
	a:=[]int{2,3} //here data type will be []int here created variable using := in oth cases array size is not fixed
	f.Println(a,r.TypeOf(a))
	b:=[3]int{1,2,3} //for fixed size array if we give more value it will show error .\array.go:12:18: index 3 is out of bounds (>= 3)
	f.Println(b,r.TypeOf(b))
	b[0]=100
	f.Println(b)
}