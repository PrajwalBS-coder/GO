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
	b[0]=100//Changing value sof perticular index
	f.Println(b)
	z:=[5]float64{}//for fixed size array if we print it will show empty values size times eg[5] values[0,0,0,0,0]
	f.Println(z)
	x:=[5]float32{1}//for fixed size if we have not provided enough values it will add empty values[1,0,0,0,0]
	f.Println(x)

	c:=[3]int{1,2,3} //for fixed size array if we give more value it will show error .\array.go:12:18: index 3 is out of bounds (>= 3)
	f.Println(c,r.TypeOf(c))
}